create table pointer(
    id bigint generated always as identity primary key,
    name text not null,
    description text not null,
    latitude float not null,
    longitude float not null,
    rating int not null
);

create table user_visit(
    id bigint generated always as identity primary key,
    user_id BIGINT not null,
    pointer_id BIGINT REFERENCES pointer(id) not null,
    rating int not null,
    comment text not null,
    user_activity bool not null
);

create table user_coins(
    user_id BIGINT not null,
    coin_count INT not null,
 );