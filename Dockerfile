FROM golang:1.7.5

COPY . /go/src/dockercon

RUN go get -d -v dockercon 
RUN go install -v dockercon

EXPOSE 80

CMD ["dockercon"]
