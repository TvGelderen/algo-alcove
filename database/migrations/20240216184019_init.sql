-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(64) UNIQUE NOT NULL,
    username VARCHAR(32) UNIQUE NOT NULL,
    role VARCHAR(32) UNIQUE NOT NULL,
    password_hash BYTEA,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE algorithms (
    id SERIAL PRIMARY KEY,
    text_id VARCHAR(128) UNIQUE NOT NULL,
    name VARCHAR(128) UNIQUE NOT NULL,
    type SMALLINT NOT NULL,
    position SMALLINT NOT NULL,
    explanation TEXT NOT NULL,
    has_code BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE algorithm_code (
    id SERIAL PRIMARY KEY,
    algorithm_id INT NOT NULL,
    language VARCHAR(32) NOT NULL,
    code TEXT NOT NULL,
    CONSTRAINT fk_algorithm 
        FOREIGN KEY(algorithm_id)
        REFERENCES algorithms(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users;
DROP TABLE algorithm_code;
DROP TABLE algorithms;
