-- migrate:up
CREATE TABLE refresh_tokens (
  token TEXT NOT NULL PRIMARY KEY
);

-- migrate:down
DROP TABLE refresh_tokens;
