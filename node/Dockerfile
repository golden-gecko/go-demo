FROM node:19.1.0-alpine3.16

WORKDIR /usr/local/app

RUN npm install express

COPY . .

EXPOSE 2000

ENTRYPOINT ["node"]

CMD ["server.js"]
