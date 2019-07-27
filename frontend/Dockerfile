FROM node:8.11 as dev
WORKDIR /frontend
COPY . .
RUN npm install
CMD ["npm", "run", "start"]

FROM dev as build
RUN npm run build

FROM nginx:1.14-alpine as nginx
# COPY --from=build /frontend/dist /usr/share/nginx/html/
