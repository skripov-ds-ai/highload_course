-- +goose Up
-- +goose StatementBegin
create index if not exists public.users.second_first_name_btree_idx
    on public.users (second_name varchar_pattern_ops, first_name varchar_pattern_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists public.users.second_first_name_btree_idx;
-- +goose StatementEnd
