FROM golang:alpine
COPY . /go/src/DHTWeather
WORKDIR /go/src/DHTWeather
RUN apk add --no-cache git \
    && go get github.com/go-sql-driver/mysql \
    && apk del git
RUN go install DHTWeather
CMD ["DHTWeather"]