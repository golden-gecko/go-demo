FROM node:21.5.0-alpine3.18

WORKDIR /usr/local/app

RUN npm install express

COPY . .

EXPOSE 2000

ENTRYPOINT ["node"]

CMD ["server.js"]
