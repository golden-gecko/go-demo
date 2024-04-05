FROM postgis/postgis:16-3.4

RUN apt update \
    && apt install -y build-essential cmake git make gcc libboost-graph-dev libpq-dev postgresql-server-dev-16

RUN git clone --depth 1 --branch v3.6.1 https://github.com/pgRouting/pgrouting.git \
    && cd pgrouting \
    && mkdir build \
    && cd build \
    && cmake .. \
    && make \
    && make install

COPY init.sql /docker-entrypoint-initdb.d/
