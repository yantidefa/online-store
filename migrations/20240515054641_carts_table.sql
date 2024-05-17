-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE carts (
    id uuid DEFAULT uuid_generate_v4(),
    product_id uuid NULL,
    user_id character varying(100) NULL,
    qty int DEFAULT 0,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE carts;
-- +goose StatementEnd
