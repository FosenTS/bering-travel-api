create table pointer(
    name varchar(25) not null,
    address varchar(100) not null,
    rating int not null
);

create table user_visit(
    user_id int not null,
    id int not null,
    pointer_id int not null,
    rating int not null ,
    comment varchar(255)
);