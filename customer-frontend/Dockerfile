FROM node:12-alpine AS build

COPY . /app

RUN cd /app && \
    npm ci install && \
    npm run build

FROM alpine:edge

RUN apk update && \
    apk add nginx && \
    rm /etc/nginx/conf.d/default.conf

COPY --from=build /app/dist /app
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

CMD [ "nginx", "-g", "pid /tmp/nginx.pid; daemon off;" ]
