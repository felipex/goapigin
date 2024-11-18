FROM golang:1.20-alpine AS build

WORKDIR /go/src/app

COPY . .

RUN go mod init github.com/felipex/apigin
RUN go mod tidy
RUN go get github.com/gin-gonic/gin
#RUN go build -o main main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:edge

WORKDIR /go/src/app
COPY --from=build /go/src/app/main .

CMD ["./main"]

