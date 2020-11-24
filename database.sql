CREATE DATABASE test_wprv;
\c test_wprv
CREATE TABLE books (
  id SERIAL,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  autor VARCHAR(100) NOT NULL,
  date DATE NOT NULL,
  PRIMARY KEY (id)
);