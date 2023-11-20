/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

create table product
(
    id serial PRIMARY KEY,
    name varchar(255),
    description    varchar(255),
    image_link     varchar(255),
    price       float,
    rating int
);

alter table users
    owner to postgres;