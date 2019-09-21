FROM node:latest

WORKDIR /node
ADD dist/node /node
COPY data/ /node/data
RUN npm install

EXPOSE 3000
CMD ["node", "main.js"]

