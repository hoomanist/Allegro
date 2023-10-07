CREATE TABLE IF NOT EXISTS composers(
   composer_id serial PRIMARY KEY,
   composer_name VARCHAR (50) UNIQUE NOT NULL,
   photo VARCHAR (200) UNIQUE NOT NULL,
   description VARCHAR (300) UNIQUE NOT NULL,
   birth SMALLINT,
   death SMALLINT
);
