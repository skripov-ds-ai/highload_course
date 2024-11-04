-- +goose Up
-- +goose StatementBegin
create index if not exists second_first_name_btree_idx
on public.users
    (second_name varchar_pattern_ops, first_name varchar_pattern_ops)
include
    (id, birthdate, biography, city, gender);

create index if not exists first_name_btree_idx
on public.users
    (first_name varchar_pattern_ops)
include
    (id, birthdate, biography, city, gender)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists public.second_first_name_btree_idx;
drop index if exists public.first_name_btree_idx;
-- +goose StatementEnd
