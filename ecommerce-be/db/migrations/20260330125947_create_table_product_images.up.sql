create table if not exists product_images
(
    id         serial,
    product_id integer      not null,
    url        varchar(100) not null,
    is_primary boolean      not null,
    created_at TIMESTAMPTZ  not null,
    updated_at TIMESTAMPTZ  not null,
    primary key(id),
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
);