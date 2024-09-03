-- +goose Up
-- +goose StatementBegin

create table if not exists public.users
(
    id            uuid        not null
        default gen_random_uuid()
        constraint user_pk
            primary key,
    first_name    varchar(50) not null,
    second_name   varchar(50) not null,
    birthdate     date        not null,
    biography     text,
    city          text,
    password_hash text        not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists public.users;

-- +goose StatementEnd
