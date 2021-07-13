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

# Copy and download dependency using go mod
#COPY go.mod .
#COPY go.sum .
#RUN go mod download

# Copy the build image into the container
# need to build like this:
# $ CGO_ENABLED=0 go build -o printapi
ADD printapi .
ADD localhost.crt .
ADD localhost.key .

# Build the application
#RUN go build -o server .

# Export necessary port
EXPOSE 8050

# Command to run when starting the container
CMD ["/printapi/printapi"]