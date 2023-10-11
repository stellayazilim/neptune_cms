-- Active: 1697054873948@@127.0.0.1@5432@neptune
-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Accounts (
    id UUID NOT NULL DEFAULT UUID_GENERATE_V4() PRIMARY KEY,
    email VARCHAR(64) NOT NULL,
    password BYTEA NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS Accounts;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
