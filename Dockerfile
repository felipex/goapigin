From golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go mod init github.com/felipex/apigin
RUN go mod tidy
RUN go get github.com/gin-gonic/gin
RUN go build -o main main.go

CMD ["./main"]