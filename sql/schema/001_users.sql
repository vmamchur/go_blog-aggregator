-- +goose Up
CREATE TABLE users(
    id uuid NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name varchar(255) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
