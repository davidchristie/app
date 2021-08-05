CREATE TABLE accounts (
  id uuid PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  provider_type VARCHAR NOT NULL,
  provider_id VARCHAR NOT NULL,
  provider_account_id VARCHAR NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  UNIQUE(provider_type, provider_id, provider_account_id)
);