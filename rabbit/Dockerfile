FROM rabbitmq:3.12.12-management-alpine

RUN rabbitmq-plugins enable --offline \
    rabbitmq_prometheus \
    rabbitmq_shovel \
    rabbitmq_shovel_management
