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

COPY cmd/main.go cmd/main.go
#COPY cmd/ ./
COPY controllers/ controllers/
COPY routes/ ./
COPY views/ ./
COPY models/ models/
COPY models/models.go /models/models.go
#COPY stream.go ./
RUN go mod vendor

RUN ls
RUN go build .
RUN ls
EXPOSE 8082
CMD ["go", "run" "cmd/main.go"] --v