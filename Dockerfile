FROM golang:1.17 as build-env 

# Set the working directory to /app
WORKDIR /app

COPY . /app

RUN go mod download 

ENV NODE_ENV=development

RUN go build -o main .

EXPOSE 3002

CMD  ["./main"]