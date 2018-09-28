FROM node

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

CMD npm run build && http-server dist
