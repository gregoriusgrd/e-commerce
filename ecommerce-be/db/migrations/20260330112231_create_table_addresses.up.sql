create table if not exists addresses(
    id             serial,
    user_id        integer          not null,
    province_id    integer          not null,
    city_id        integer          not null,
    address_line   varchar(100)     not null,
    postal_code    varchar(100)     not null,
    receiver_name  varchar (100)    not null,
    phone_number   varchar(100)     not null,
    is_default     boolean          not null,
    created_at     TIMESTAMPTZ      not null,
    updated_at     TIMESTAMPTZ      not null,

    primary key (id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id),
    CONSTRAINT fk_province
        FOREIGN KEY (province_id)
        REFERENCES provinces(id),
    CONSTRAINT fk_city
        FOREIGN KEY (city_id)
        REFERENCES cities(id)
);