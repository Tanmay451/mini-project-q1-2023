CREATE TABLE forum_chat_room (
  id serial PRIMARY KEY,
  name text NOT NULL,
  description text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);