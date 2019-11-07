FROM node:8.11 as dev

ADD package.json /tmp/package.json
RUN cd /tmp && npm install
RUN mkdir /frontend && cp -a /tmp/node_modules /frontend

WORKDIR /frontend
COPY package*.json ./
RUN npm install
COPY . ./

FROM dev as build
RUN npm run build

FROM nginx:1.14-alpine
COPY --from=build /frontend/dist /usr/share/nginx/html/
RUN mv /etc/nginx/nginx.conf.default /etc/nginx/nginx.conf.default.disabled
