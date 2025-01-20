-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS receipts;

create table receipts (
    id serial primary key,
    date date not null,
    time time not null,
    amount numeric(10, 2) not null,
    fiscal_number bigint not null,
    fiscal_document integer not null,
    fiscal_sign bigint not null,
    created_at timestamp with time zone default current_timestamp
);

insert into receipts (date, time, amount, fiscal_number, fiscal_document, fiscal_sign) values
    ('2024-12-19', '14:20:00', 800.00, '7384440800122779', '121', '427379931'),
    ('2024-12-20', '15:30:00', 1500.50, '7384440800122780', '122', '427379932'),
    ('2024-12-21', '16:40:00', 1200.75, '7384440800122781', '123', '427379933')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table receipts;
-- +goose StatementEnd

