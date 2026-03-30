create table if not exists payments
(
    id              serial,
    order_id        integer      not null,
    payment_method  integer      not null,
    payment_status  varchar(100) not null,
    amount          decimal      not null,
    payment_proof   varchar(100) not null,
    paid_at         TIMESTAMPTZ  not null,
    expired_at      TIMESTAMPTZ  not null,
    created_at      TIMESTAMPTZ  not null,
    updated_at      TIMESTAMPTZ  not null,

    primary key (id),
    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
)