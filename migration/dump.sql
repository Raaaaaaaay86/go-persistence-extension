BEGIN;

CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username VARCHAR(200) UNIQUE NOT NULL,
  email VARCHAR(100) NOT NULL,
  age INT NOT NULL,
  birthday DATE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP WITH TIME ZONE
);

INSERT INTO users (username, email, age, birthday) VALUES ('user1', 'test@mail.com', 20, '2000-03-03');
INSERT INTO users (username, email, age, birthday) VALUES ('user2', 'test@mail.com', 21, '2000-04-04');
INSERT INTO users (username, email, age, birthday) VALUES ('user3', 'test@mail.com', 20, '2000-05-05');
INSERT INTO users (username, email, age, birthday) VALUES ('user4', 'test@mail.com', 23, '2000-06-06');
INSERT INTO users (username, email, age, birthday) VALUES ('user5', 'test@mail.com', 24, '2000-07-07');

COMMIT;