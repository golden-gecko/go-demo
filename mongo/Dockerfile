FROM mongo:4.4.6

COPY mongo.key /etc/mongo/mongo.key
COPY mongod.conf /etc/mongo/mongod.conf

RUN chown mongodb /etc/mongo/mongo.key
RUN chmod 400 /etc/mongo/mongo.key
