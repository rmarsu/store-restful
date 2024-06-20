CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE products
(
    id serial not null unique,
    image varchar(255) not null,
    name varchar(255) not null,
    price integer not null

);
