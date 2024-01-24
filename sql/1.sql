CREATE DATABASE cinema_labora;

CREATE TABLE IF NOT EXISTS movie (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL
  adult BOOLEAN,
  genres VARCHAR(50)[],
  release_date DATE NOT NULL,
  poster_path VARCHAR(100),
  imdb_id VARCHAR(50),
  original_language VARCHAR(10),
  original_title VARCHAR(100),
  overview VARCHAR(300),
  popularity DECIMAL(10,2),
  runtime DECIMAL(10,2),
  status VARCHAR(50),
  tagline VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS theater (
  id SERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  capacity INT NOT NULL,
  last_row VARCHAR(2),
  last_column INT,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS screening (
  id SERIAL PRIMARY KEY,
  name VARCHAR(20),
  movie_id INT
  theater_id INT REFERENCES theater(id),
  available_seats INT,
  taken_seats VARCHAR(5)[],
  showtime TIMESTAMP,
  language VARCHAR(10),
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ticket (
  id SERIAL PRIMARY KEY,
  pickup_id VARCHAR(10),
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(30) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL,
  gender CHAR,
  ticket_ids INT[] REFERENCES ticket(id),
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);