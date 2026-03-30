create table if not exists cart_items
(
    id          serial,
    cart_id     integer     not null,
    product_id  integer     not null,
    quantity    integer     not null,
    created_at  TIMESTAMPTZ not null,
    updated_at  TIMESTAMPTZ not null,
    primary key(id),
    CONSTRAINT fk_cart
        FOREIGN KEY (cart_id)
        REFERENCES carts(id),
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
);