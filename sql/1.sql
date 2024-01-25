CREATE DATABASE cinema_labora;

CREATE TABLE IF NOT EXISTS theater (
  id SERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  capacity INT NOT NULL,
  last_row VARCHAR(2) NOT NULL,
  last_column INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS screening (
  id SERIAL PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  movie_id INT NOT NULL,
  theater_id INT REFERENCES theater(id),
  available_seats INT NOT NULL,
  taken_seats VARCHAR(5) [],
  showtime TIMESTAMP NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  language VARCHAR(10),
  views_count INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(30),
  last_name VARCHAR(30),
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL,
  gender CHAR,
  picture_url VARCHAR(256) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS ticket (
  id SERIAL PRIMARY KEY,
  pickup_id VARCHAR(10),
  user_id INT REFERENCES "user"(id),
  screening_id INT REFERENCES screening(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comment (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES "user"(id),
  movie_id INT,
  content TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
