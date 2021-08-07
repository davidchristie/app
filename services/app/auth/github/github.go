package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Email struct {
	Email      string  `json:"email"`
	Verified   bool    `json:"verified"`
	Primary    bool    `json:"primary"`
	Visibility *string `json:"visibility"`
}

type User struct {
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

func FetchEmails(client *http.Client, url string) ([]Email, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code: %s", response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	emails := []Email{}
	err = json.Unmarshal(body, &emails)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func FetchPrimaryEmail(client *http.Client, url string) (string, error) {
	emails, err := FetchEmails(client, url)
	if err != nil {
		return "", fmt.Errorf("Error fetching primary GitHub email: %s", err)
	}
	for _, email := range emails {
		if email.Primary {
			return email.Email, nil
		}
	}
	return "", errors.New("Primary GitHub email not found")
}

func FetchUser(client *http.Client, url string) (*User, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code: %s", response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
