-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

CREATE TABLE categories (
    id uuid DEFAULT uuid_generate_v4(),
    name character varying(100) NULL,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE categories;
-- +goose StatementEnd
