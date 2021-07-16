FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    ZVT_URL=pt:20007 \
    ZVT_LOGFILEPATH=/var/log/zvt\
    ZVT_DUMPFILEPATH=/var/log/zvt/dump

#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh
# Move to working directory /build
WORKDIR /printapi

ADD printapi .
ADD localhost.crt .
ADD localhost.key .

EXPOSE 8050

CMD ["/printapi/printapi"]