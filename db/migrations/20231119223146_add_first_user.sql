-- +goose Up
-- +goose StatementBegin

insert into public.users
    (id, first_name, second_name, birthdate, password_hash)
values
    ('35df14eb-a4b8-4121-a85c-2bffabfde9d9', 'First', 'Second', now()::date, '$2a$14$maBIpxNvYquBO2YRt4GcK.vJ6ehR3MYHDm2Zyi6AaLcByRHy75fMO')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

delete from public.users
where first_name = 'First'
;
-- +goose StatementEnd
