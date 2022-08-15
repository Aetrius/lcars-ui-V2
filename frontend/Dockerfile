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

RUN go get github.com/joho/godotenv
RUN go get github.com/mackerelio/go-osstat/cpu
RUN go get github.com/labstack/echo/v4
RUN go get golang.org/x/net
#RUN go get github.com/cosmtrek/air@latest

COPY *.go ./
COPY index.html ./
COPY cmd/ cmd/
COPY controllers/ controllers/
COPY routes/ routes/
COPY views/ views/
COPY stream.go stream.go
RUN go mod vendor

RUN ls
RUN go build -o /lcars
EXPOSE 80
CMD ["/lcars"] --v