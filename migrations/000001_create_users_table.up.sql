CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR,
    password VARCHAR,
    email VARCHAR,
    age INT,
    first_name VARCHAR,
    last_name VARCHAR
);