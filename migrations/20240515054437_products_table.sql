-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE products (
    id uuid DEFAULT uuid_generate_v4(),
    category_id uuid NULL,
    name character varying(100) NULL,
    description text NULL,
    stock int DEFAULT 0,
    price int DEFAULT 0,
    image text NULL,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE products;
-- +goose StatementEnd
