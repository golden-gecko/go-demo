FROM rabbitmq:3.12.12-management-alpine

RUN rabbitmq-plugins enable --offline \
    rabbitmq_prometheus \
    rabbitmq_shovel \
    rabbitmq_shovel_management

RUN echo secret-cookie > /var/lib/rabbitmq/.erlang.cookie

CMD ["rabbitmq-server", "start"]
