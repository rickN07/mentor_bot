-- +goose Up
CREATE TABLE bot_users (
  id serial primary key,
  chat_id int,
  info jsonb,
  created_at timestamp default now(),
  mailing jsonb,
  sent_lead timestamp
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS bot_users;
