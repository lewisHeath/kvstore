FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /kvstore .

FROM scratch

COPY --from=build /kvstore /kvstore

CMD ["/kvstore"]