FROM golang:1.14.7-alpine
WORKDIR /src
COPY . .
CMD ["go", "run", "main.go"]
