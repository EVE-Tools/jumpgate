FROM alpine:3.7

#
# Copy release to container and set command
#

# Add faster mirror and upgrade packages in base image, load ca certs, otherwise no TLS for us
RUN printf "http://mirror.leaseweb.com/alpine/v3.7/main\nhttp://mirror.leaseweb.com/alpine/v3.7/community" > etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

# Copy build
COPY jumpgate jumpgate

ENV LISTEN_HOST ":8000"
EXPOSE 8000

CMD ["/jumpgate"]