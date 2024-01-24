-- +goose Up
-- +goose StatementBegin
CREATE TABLE hello (
                          id SERIAL PRIMARY KEY,
                          user_id INT UNIQUE REFERENCES users (id) ON DELETE CASCADE,
                          token_hash TEXT UNIQUE NOT NULL --the unique constrain does not let the user have more than one session at a time
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE hello;
-- +goose StatementEnd
