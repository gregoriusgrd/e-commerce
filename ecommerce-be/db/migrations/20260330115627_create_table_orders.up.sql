create table if not exists orders
(
    id               serial,
    user_id          integer      not null,
    address_id       integer      not null,
    total_price      decimal      not null,
    order_status     varchar(100) not null,
    invoice_number   varchar(100) not null,
    created_at       TIMESTAMPTZ  not null,
    updated_at       TIMESTAMPTZ  not null,
    primary key (id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_address
        FOREIGN KEY (address_id)
        REFERENCES addresses(id)
);