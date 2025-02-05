DO $$ 
BEGIN
  IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'users_db') THEN
    CREATE DATABASE users_db;
  END IF;
END $$;

\c users_db;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS relations (
  id SERIAL PRIMARY KEY,
  user_1_id INT REFERENCES users(id) ON DELETE CASCADE,
  user_2_id INT REFERENCES users(id) ON DELETE CASCADE,
  status VARCHAR(20) CHECK (status IN ('pending', 'accepted', 'blocked')),
  status_creator INT, 
  lesser_user INT GENERATED ALWAYS AS (LEAST(user_1_id, user_2_id)) STORED,
  greater_user INT GENERATED ALWAYS AS (GREATEST(user_1_id, user_2_id)) STORED,
  CONSTRAINT dialogs_users_pair_unique UNIQUE (lesser_user, greater_user),
  CONSTRAINT status_creator_check CHECK (status_creator IN (user_1_id, user_2_id))
)
