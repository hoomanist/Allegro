CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL, 
   creationdate timestamp NOT NULL,
   password VARCHAR (90) NOT NULL
);