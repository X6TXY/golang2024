FROM golang:1.19


RUN apt-get update && apt-get install -y wget

# Install dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /finalproject

COPY . .

RUN go mod download


RUN go build -o ./finalproject/cmd/main ./cmd

EXPOSE 8080


CMD ["dockerize", "-wait", "tcp://postgres:5432", "-timeout", "60s", "./finalproject/cmd/main"]
