create table if not exists products
(
    id          serial,
    category_id integer      not null,
    name        varchar(100) not null,
    slug        varchar(100) not null,
    description text         not null,
    price       decimal      not null,
    stock       integer      not null,
    is_active   boolean      not null,
    created_at  TIMESTAMPTZ  not null,
    updated_at  TIMESTAMPTZ  not null,
    primary key(id),
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
        REFERENCES categories(id)
)