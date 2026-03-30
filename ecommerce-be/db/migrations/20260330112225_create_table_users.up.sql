create type role as ENUM('CUSTOMER', 'ADMIN');


create table if not exists users(
    id              serial,
    name            varchar(100)    not null,
    email           varchar(100)    not null unique,
    password        varchar(100)    not null,
    profile_picture varchar(100)    not null,
    role            role            not null default 'CUSTOMER',
    created_at      TIMESTAMPTZ     not null,
    updated_at      TIMESTAMPTZ     not null,
    primary key (id)
);