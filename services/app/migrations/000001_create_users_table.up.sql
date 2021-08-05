CREATE TABLE users (
  id uuid PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  primary_email VARCHAR NOT NULL UNIQUE,
  full_name VARCHAR NOT NULL,
  avatar_url VARCHAR
);