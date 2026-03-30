create table if not exists carts 
(
    id         serial,
    user_id    integer         not null,
    created_at TIMESTAMPTZ     not null,
    updated_at TIMESTAMPTZ     not null,
    primary key (id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
);