FROM golang:1.14-alpine AS go
ENV APP_DIR /var/www/
ENV BUILD_DIR /home/

COPY ./ ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN set -ex && \
    go build ./
        
RUN set -ex && \
    mkdir ${APP_DIR} && \
    cp ./invoke-webhook ${APP_DIR}
FROM alpine:latest
ENV APP_DIR /var/www/

COPY --from=go ${APP_DIR} ${APP_DIR}
COPY ./start.sh /usr/local/bin/
WORKDIR ${APP_DIR}
RUN set -ex && \
    chmod u+x ./invoke-webhook && \
    chmod u+x /usr/local/bin/start.sh

CMD ["start.sh"]