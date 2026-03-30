create table if not exists cities(
    id          serial,
    province_id integer         not null,
    name        varchar(100)    not null,
    primary key (id),
    CONSTRAINT fk_province
        FOREIGN KEY (province_id)
        REFERENCES provinces(id)
)