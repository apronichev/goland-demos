-- Seed data for the "postgres" service in docker-compose.yaml.
-- Mounted into the container at /docker-entrypoint-initdb.d and executed
-- automatically the first time the database is initialized.

CREATE TABLE IF NOT EXISTS users (
    id    SERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

INSERT INTO users (name, email) VALUES
    ('Ada Lovelace',   'ada@example.com'),
    ('Alan Turing',    'alan@example.com'),
    ('Grace Hopper',   'grace@example.com'),
    ('Rob Pike',       'rob@example.com')
ON CONFLICT (email) DO NOTHING;
