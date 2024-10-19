-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files_metadata (
    UUID VARCHAR UNIQUE NOT NULL,
    filename VARCHAR(255) NOT NULL,
    extension VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files_metadata;
-- +goose StatementEnd
