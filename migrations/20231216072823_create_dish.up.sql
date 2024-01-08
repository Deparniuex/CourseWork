CREATE TABLE IF NOT EXISTS dishes(
    id SERIAL PRIMARY KEY, 
    name text NOT NULL,
    weight bigint NOT NULL, 
    price bigint NOT NULL, 
    menu bigint NOT NULL REFERENCES menus ON DELETE CASCADE
)