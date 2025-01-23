drop table if exists receipts;

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
