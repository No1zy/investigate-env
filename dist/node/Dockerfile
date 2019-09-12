FROM node:latest

WORKDIR /node
ADD dist/node /node
COPY data/ /node/data

CMD ["node", "main.js"]

