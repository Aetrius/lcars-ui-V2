FROM golang:1.19-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy
RUN go mod vendor

RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
#RUN go get github.com/cosmtrek/air@latest

COPY *.go ./
RUN go mod vendor

RUN ls
RUN go build -o /lcars
EXPOSE 8081
CMD ["/lcars"] --v