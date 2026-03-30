create table if not exists categories
(
    id          serial,
    name        varchar(100) not null,
    parent_id   integer      null,
    created_at  TIMESTAMPTZ  not null,
    updated_at  TIMESTAMPTZ  not null,
    primary key(id),
    CONSTRAINT fk_parent
        FOREIGN KEY (parent_id)
        REFERENCES categories(id)
        ON DELETE SET NULL
)