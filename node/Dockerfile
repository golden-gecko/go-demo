FROM node:21.6.0-alpine3.18

WORKDIR /app

RUN npm install express

COPY . .

EXPOSE 2000

CMD ["node", "server.js"]
