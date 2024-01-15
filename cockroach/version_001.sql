DROP TABLE go_demo.public.vehicles;

CREATE TABLE go_demo.public.vehicles (
    plate VARCHAR NOT NULL,
    brand VARCHAR NOT NULL,
    model VARCHAR NOT NULL,
    year INT8 NOT NULL,
    category VARCHAR NOT NULL,
    CONSTRAINT vehicles_pkey PRIMARY KEY (plate ASC)
);

GRANT ALL ON TABLE go_demo.public.* TO go_user WITH GRANT OPTION;
