FROM golang:1.7.5

COPY . /go/src/dockercon

RUN go get -d -v dockercon 
RUN go install -v dockercon

EXPOSE 80

HEALTHCHECK --interval=2s --timeout=3s --retries=3 \
  CMD curl -f http://localhost:80/healthcheck

CMD ["dockercon"]
