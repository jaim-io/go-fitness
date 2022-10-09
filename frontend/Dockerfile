# syntax=docker/dockerfile:1

FROM node:18-alpine

WORKDIR /app
COPY . .

RUN yarn install --frozen-lockfile
RUN yarn run build

ENV NODE_ENV production

EXPOSE 3000

CMD [ "npx", "serve", "build" ]