FROM golang

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o /kvstore .

EXPOSE 8080
EXPOSE 3000

CMD ["/kvstore"]