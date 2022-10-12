# syntax=docker/dockerfile:1

# Build environment
FROM node:18-alpine as build
WORKDIR /app
COPY . .
RUN yarn install --frozen-lockfile
RUN yarn run build

# Production environment
FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY --from=build /app/nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]