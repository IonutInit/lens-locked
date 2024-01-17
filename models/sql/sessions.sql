CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE,
    token_hash TEXT UNIQUE NOT NULL --the unique constrain does not let the user have more than one session at a time
)