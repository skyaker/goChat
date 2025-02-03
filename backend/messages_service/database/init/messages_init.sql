DO $$ 
BEGIN
  IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'messages_db') THEN
    CREATE DATABASE messages_db;
  END IF;
END $$;

CREATE TABLE IF NOT EXISTS dialogs (
  id SERIAL PRIMARY KEY,
  user_1_id INT NOT NULL,
  user_2_id INT NOT NULL,
  last_message VARCHAR(50),
  last_message_at TIMESTAMP DEFAULT NOW(),
  lesser_user INT GENERATED ALWAYS AS (LEAST(user_1_id, user_2_id)) STORED,
  greater_user INT GENERATED ALWAYS AS (GREATEST(user_1_id, user_2_id)) STORED,
  CONSTRAINT dialogs_users_pair_unique UNIQUE (lesser_user, greater_user)
);

CREATE TABLE IF NOT EXISTS messages (
  message_id SERIAL PRIMARY KEY,
  dialog_id INT REFERENCES dialogs(id) ON DELETE CASCADE,
  sender_id INT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);