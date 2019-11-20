FROM alpine:edge

COPY bin/koala-pos /app/koala-pos
COPY docker/config-prod.toml /app/config.toml

WORKDIR /app

CMD [ "/app/koala-pos", "-c", "/app/config.toml", "-d" ]
