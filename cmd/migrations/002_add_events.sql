BEGIN;

CREATE TABLE events (
  id bigserial PRIMARY KEY,
  user_id BIGINT NOT NUll,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  headline text NOT NUll,
  description text,

  CONSTRAINT fk_event_users
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

COMMIT;
