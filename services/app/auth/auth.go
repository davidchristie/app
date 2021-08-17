//go:generate mockgen -destination ../mocks/auth.go -package mocks github.com/davidchristie/app/services/app/auth Auth

package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/davidchristie/app/services/app/auth/github"
	"github.com/davidchristie/app/services/app/auth/google"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/entities"
	"github.com/davidchristie/app/services/app/repositories"
	"github.com/davidchristie/app/services/app/utilities"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Auth interface {
	Authorize(providerID string) (*AuthorizeResult, error)
	Callback(ctx context.Context, providerID, state, code string) (*CallbackResult, error)
	Session(ctx context.Context, token string) (*Session, error)
}

type AuthorizeResult struct {
	Redirect string
}

type CallbackResult struct {
	SessionToken     string
	SessionExpiresAt time.Time
	Redirect         string
}

type Profile struct {
	ID           string
	PrimaryEmail string
	FullName     string
	AvatarURL    string
}

type Provider struct {
	Config       *oauth2.Config
	State        string
	FetchProfile func(client *http.Client) (*Profile, error)
}

type Session struct {
	User *User
}

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	AvatarURL string
}

type auth struct {
	providers         map[string]*Provider
	userRepository    repositories.UserRepository
	accountRepository repositories.AccountRepository
	sessionRepository repositories.SessionRepository
}

func NewAuth(
	config *config.Config,
	userRepository repositories.UserRepository,
	accountRepository repositories.AccountRepository,
	sessionRepository repositories.SessionRepository,
) Auth {
	providers := map[string]*Provider{
		"github": {
			Config: &oauth2.Config{
				ClientID:     config.GitHubClientID,
				ClientSecret: config.GitHubClientSecret,
				Endpoint: oauth2.Endpoint{
					AuthURL:  config.GitHubAuthURL,
					TokenURL: config.GitHubTokenURL,
				},
				RedirectURL: config.GitHubRedirectURL,
				Scopes:      []string{"user:email"},
			},
			State: utilities.MustGenerateSecureToken(32),
			FetchProfile: func(client *http.Client) (*Profile, error) {
				user, err := github.FetchUser(client, config.GitHubUserURL)
				if err != nil {
					return nil, err
				}
				profile := Profile{
					ID:           fmt.Sprint(user.ID),
					PrimaryEmail: user.Email,
					FullName:     user.Name,
					AvatarURL:    user.AvatarURL,
				}
				if profile.PrimaryEmail == "" {
					primaryEmail, err := github.FetchPrimaryEmail(client, config.GitHubEmailsURL)
					if err != nil {
						return nil, err
					}
					profile.PrimaryEmail = primaryEmail
				}
				return &profile, nil
			},
		},
		"google": {
			Config: &oauth2.Config{
				ClientID:     config.GoogleClientID,
				ClientSecret: config.GoogleClientSecret,
				Endpoint: oauth2.Endpoint{
					AuthURL:  config.GoogleAuthURL,
					TokenURL: config.GoogleTokenURL,
				},
				RedirectURL: config.GoogleRedirectURL,
				Scopes:      []string{"email", "profile"},
			},
			State: utilities.MustGenerateSecureToken(32),
			FetchProfile: func(client *http.Client) (*Profile, error) {
				user, err := google.FetchUser(client, config.GoogleUserURL)
				if err != nil {
					return nil, err
				}
				return &Profile{
					ID:           user.Sub,
					PrimaryEmail: user.Email,
					FullName:     user.Name,
					AvatarURL:    user.Picture,
				}, nil
			},
		},
	}
	return &auth{
		providers:         providers,
		userRepository:    userRepository,
		accountRepository: accountRepository,
		sessionRepository: sessionRepository,
	}
}

func (a *auth) Authorize(providerID string) (*AuthorizeResult, error) {
	provider, ok := a.providers[providerID]
	if !ok {
		return nil, fmt.Errorf("Unsupported provider ID: %s", providerID)
	}
	return &AuthorizeResult{
		Redirect: provider.Config.AuthCodeURL(provider.State),
	}, nil
}

func (a *auth) Callback(ctx context.Context, providerID, state, code string) (*CallbackResult, error) {
	provider, ok := a.providers[providerID]
	if !ok {
		return nil, fmt.Errorf("Unsupported provider ID: %s", providerID)
	}
	if state != provider.State {
		return nil, fmt.Errorf("Invalid oauth2 state: %s", state)
	}
	token, err := provider.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	client := provider.Config.Client(ctx, token)
	profile, err := provider.FetchProfile(client)
	if err != nil {
		return nil, err
	}
	user, err := a.findOrCreateUserForProfile(ctx, "oauth", providerID, profile)
	session := &entities.Session{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		ExpiresAt:    time.Now().UTC().Add(7 * 24 * time.Hour),
		SessionToken: utilities.MustGenerateSecureToken(32),
		UserID:       user.ID,
	}
	err = a.sessionRepository.Insert(ctx, session)
	if err != nil {
		return nil, err
	}
	return &CallbackResult{
		SessionToken:     session.SessionToken,
		SessionExpiresAt: session.ExpiresAt,
		Redirect:         "/",
	}, nil
}

func (a *auth) Session(ctx context.Context, token string) (*Session, error) {
	if token == "" {
		return &Session{
			User: nil,
		}, nil
	}
	session, err := a.sessionRepository.FindBySessionToken(ctx, token)
	if err == repositories.ErrRecordNotFound {
		return &Session{
			User: nil,
		}, nil
	} else if err != nil {
		return nil, err
	}
	if session.ExpiresAt.Before(time.Now()) {
		return &Session{
			User: nil,
		}, nil
	}
	user, err := a.userRepository.FindByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}
	return &Session{
		User: &User{
			ID:        user.ID,
			Name:      user.FullName,
			Email:     user.PrimaryEmail,
			AvatarURL: user.AvatarURL,
		},
	}, nil
}

func (a *auth) findOrCreateUserForProfile(
	ctx context.Context,
	providerType string,
	providerID string,
	profile *Profile,
) (*entities.User, error) {
	existingAccount, err := a.accountRepository.FindByProvider(ctx, providerType, providerID, profile.ID)
	if err == repositories.ErrRecordNotFound {
		newUser := &entities.User{
			ID:           uuid.New(),
			CreatedAt:    time.Now().UTC(),
			UpdatedAt:    time.Now().UTC(),
			PrimaryEmail: profile.PrimaryEmail,
			FullName:     profile.FullName,
			AvatarURL:    profile.AvatarURL,
		}
		if err = a.userRepository.Insert(ctx, newUser); err != nil {
			return nil, err
		}
		newAccount := &entities.Account{
			ID:                uuid.New(),
			CreatedAt:         time.Now().UTC(),
			UpdatedAt:         time.Now().UTC(),
			ProviderType:      providerType,
			ProviderID:        providerID,
			ProviderAccountID: profile.ID,
			UserID:            newUser.ID,
		}
		if err = a.accountRepository.Insert(ctx, newAccount); err != nil {
			return nil, err
		}
		return newUser, nil
	} else if err != nil {
		return nil, err
	}
	existingUser, err := a.userRepository.FindByID(ctx, existingAccount.UserID)
	if err != nil {
		return nil, err
	}
	return existingUser, nil
}
