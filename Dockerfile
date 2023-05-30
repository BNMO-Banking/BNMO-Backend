FROM golang:latest

WORKDIR /bnmo-backend
RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["CompileDaemon", "-command=./BNMO",  "-polling"]