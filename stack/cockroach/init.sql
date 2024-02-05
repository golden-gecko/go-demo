CREATE DATABASE go_demo;

CREATE USER go_user;

GRANT ALL ON DATABASE go_demo TO go_user WITH GRANT OPTION;

CREATE TABLE go_demo.public.users (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id ASC)
);

CREATE TABLE go_demo.public.vehicles (
    plate VARCHAR NOT NULL,
    brand VARCHAR NOT NULL,
    model VARCHAR NOT NULL,
    year INT8 NOT NULL,
    category VARCHAR NOT NULL,
    CONSTRAINT vehicles_pkey PRIMARY KEY (plate ASC)
);

GRANT ALL ON TABLE go_demo.public.* TO go_user WITH GRANT OPTION;
