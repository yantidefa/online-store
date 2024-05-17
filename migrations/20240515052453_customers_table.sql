-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    name character varying(100) NULL,
    email character varying(100) NULL,
    password character varying(100) NULL,
    token text NULL,
    is_login bool DEFAULT false,
    gender character varying(20),
    address text NULL,
    phone character varying(15),
    role character varying(15),
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
