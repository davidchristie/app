CREATE TABLE sessions (
  id uuid PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  session_token VARCHAR NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);