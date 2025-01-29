DO $$ 
BEGIN
  IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'messages_db') THEN
    CREATE DATABASE messages_db;
  END IF;
END $$;

\c messages_db;

CREATE TABLE IF NOT EXISTS dialogs (
  id SERIAL PRIMARY KEY,
  user_1_id VARCHAR(50) UNIQUE NOT NULL,
  user_2_id VARCHAR(50) UNIQUE NOT NULL,
  last_message VARCHAR(50)
)

CREATE TABLE IF NOT EXISTS messages (
  message_id SERIAL PRIMARY KEY,
  dialog_id INT REFERENCES dialogs(id) ON DELETE CASCADE,
  sender_id INT UNIQUE NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
)