FROM golang:1.11

WORKDIR /go/src/github.com/newfolder31/yurko
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["yurko"]

#build docker image
#docker build -t yurko-image .

#run docker container
#docker run --publish 8081:8081 -it --rm --name yurko-container yurko-image