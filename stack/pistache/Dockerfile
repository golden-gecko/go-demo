FROM gcc:13.2.0-bookworm

WORKDIR /app

RUN apt update \
    && apt install -y meson

RUN git clone https://github.com/pistacheio/pistache.git \
    && cd pistache \
    && meson setup build \
    && meson install -C build \
    && ldconfig

WORKDIR /app/src

COPY . .

RUN gcc main.cpp -lstdc++ -lpistache -o pistache

EXPOSE 9000

CMD ["./pistache"]
