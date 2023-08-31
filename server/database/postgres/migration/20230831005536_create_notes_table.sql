-- migrate:up
CREATE TYPE status AS ENUM ('normal', 'favorite', 'archived', 'trashed');
CREATE TABLE notes (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INT NOT NULL,
  space_id INT NOT NULL,

  title VARCHAR(50) NOT NULL,
  body TEXT,
  status status DEFAULT 'normal'::status,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_note_user FOREIGN KEY(user_id) REFERENCES users(id) 
    ON DELETE CASCADE,
  CONSTRAINT fk_note_space FOREIGN KEY(space_id) REFERENCES spaces(id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE notes;
