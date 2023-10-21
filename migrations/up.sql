create table pointer(
    id bigint generated always as identity primary key,
    name text not null,
    address text not null,
    rating int not null
);

create table user_visit(
    id bigint generated always as identity primary key,
    user_id BIGINT REFERENCES pointer(id) not null,
    pointer_id int not null,
    rating int not null,
    comment text
);
