FROM alpine:3.4
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
ADD micro /
ENTRYPOINT [ "/micro" ]