create table if not exists order_items
(
    id          serial,
    order_id    integer     not null,
    product_id  integer     not null,
    quantity    integer     not null,
    price       decimal     not null,
    created_at  TIMESTAMPTZ not null,
    updated_at  TIMESTAMPTZ not null,
    primary key(id),
    CONSTRAINT fk_order
        FOREIGN KEY(order_id)
        REFERENCES orders(id),
    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES products(id)
);