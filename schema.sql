drop table if exists addresses;
drop table if exists users;

create table if not exists users(
    id bigserial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    birth_date timestamp not null
);

create table if not exists addresses(
    id bigserial primary key,
    street varchar(255) not null,
    number int not null,
    user_id int not null,
    city varchar(255),
    constraint fk_users
        foreign key(user_id)
            references users(id) ON DELETE CASCADE
);

CREATE INDEX ON addresses (user_id);