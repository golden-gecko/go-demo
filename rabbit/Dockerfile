FROM rabbitmq:3.12.12-management-alpine

WORKDIR /root

RUN rabbitmq-plugins enable --offline \
    rabbitmq_prometheus \
    rabbitmq_shovel \
    rabbitmq_shovel_management

COPY entrypoint.sh .

CMD ["/root/entrypoint.sh"]
