CREATE TABLE IF NOT EXISTS restaurants(
    id SERIAL PRIMARY KEY,
    name text NOT NULL, 
    city bigserial NOT NULL REFERENCES cities ON DELETE CASCADE,
    street bigserial NOT NULL REFERENCES streets ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS menus(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    restaurants bigserial NOT NULL REFERENCES restaurants
);
