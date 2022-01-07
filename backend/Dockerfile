FROM golang:1.16.3

RUN mkdir -p /backend
COPY . /backend

WORKDIR /backend
RUN go get

# Docker Compose Wait Tool
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# Hot Reload
RUN [ "go", "get", "github.com/githubnemo/CompileDaemon" ]
CMD /wait && CompileDaemon -log-prefix=false -build="go build" -command="./main"