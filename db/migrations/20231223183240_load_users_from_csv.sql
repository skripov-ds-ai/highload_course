-- +goose Up
-- +goose StatementBegin
copy public.users
    (second_name, first_name, birthdate, city, password_hash)
from program
    'awk -F, ''{split($1, a, " "); printf a[1] "," a[2] "," strftime("%Y-%m-%d", systime() - $2 * 3600 * 24 * 365.25) "," $3 ",example" "\n"}'' ''/data/people_shuffled.csv'''
with delimiter ','
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from public.users
where id != '35df14eb-a4b8-4121-a85c-2bffabfde9d9'
;

-- +goose StatementEnd
