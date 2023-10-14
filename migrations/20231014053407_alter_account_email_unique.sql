-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS Accounts ADD CONSTRAINT unique_email UNIQUE (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS Accounts DROP CONSTRAINT unique_email;
-- +goose StatementEnd
