CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    status VARCHAR(50)
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name text,
    surname text,
    phone text,
    email text,
    password text,
    token text,
    created_at timestamp,
    updated_at timestamp
)