FROM golang:1.10-alpine



ADD . /go/src/servercorssl
RUN go install servercorssl

#
#
FROM alpine:latest
RUN mkdir -p /etc/certs
COPY fullchain.pem /etc/certs/fullchain.pem
COPY key.pem /etc/certs/key.pem
RUN apk	update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=0 /go/bin/servercorssl .
ENV PORT 8080
CMD ["./servercorssl"]

# To run this
#  docker build -t us.gcr.io/mchirico/gosslcorsserver:latest .

#  docker push us.gcr.io/mchirico/gosslcorsserver:latest

# kubectl run hello-server --image us.gcr.io/mchirico/gosslserver:latest --port 8080
# kubectl delete service hello-server

# // Not working ..
# kubectl rolling-update myapp --image=us.gcr.io/mchirico/hello:latest --image-pull-policy Always


