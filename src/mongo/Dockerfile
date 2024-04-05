FROM mongo:4.4.28

COPY mongod.conf mongo.key /etc/mongo/

RUN chown mongodb /etc/mongo/mongo.key
RUN chmod 400 /etc/mongo/mongo.key
