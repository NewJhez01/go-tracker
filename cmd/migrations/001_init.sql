BEGIN;

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  email text UNIQUE NOT NULL,
  last_logged_in TIMESTAMPTZ NOT NULL
);

COMMIT;
