FROM node:13.12.0-alpine3.11 as dev

ADD package.json /tmp/package.json
RUN cd /tmp && npm install
RUN mkdir /bff && cp -a /tmp/node_modules /bff

WORKDIR /bff
COPY . .
CMD ["npm", "run", "start"]
# CMD /bin/bash

# FROM dev as build
# RUN npm run build

# FROM nginx:1.14-alpine as nginx
# COPY --from=build /frontend/dist /usr/share/nginx/html/
