CREATE TYPE role AS ENUM('USER', 'MODERATOR', 'ADMIN');

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    phone text NOT NULL,
    email text NOT NULL,
    role role NOT NULL DEFAULT 'USER',
    username text UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);