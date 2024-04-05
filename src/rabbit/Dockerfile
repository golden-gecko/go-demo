FROM rabbitmq:3.13.0-management-alpine

RUN rabbitmq-plugins enable --offline \
    rabbitmq_prometheus \
    rabbitmq_shovel \
    rabbitmq_shovel_management

WORKDIR /app

COPY entrypoint.sh .

ENTRYPOINT ["sh", "/app/entrypoint.sh"]
