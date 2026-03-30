create table if not exists reviews
(
    id          serial,
    user_id     integer     not null,
    product_id  integer     not null,
    rating      integer     not null,
    comment     text,
    created_at  TIMESTAMPTZ not null,
    updated_at  TIMESTAMPTZ not null,
    primary key(id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id),
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
);