FROM golang:1.14.7-alpine3.12 as build

WORKDIR /go/src/ebook/ebook-privilege
COPY ./ /go/src/ebook/ebook-privilege
RUN    apk --no-cache add git \
  && go env -w  GOPROXY=https://goproxy.cn \
  && go get -d \
  && CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o ebook-user main.go \
  && chmod +x ebook-privilege


FROM alpine:3.12
ENV port=50802
COPY --from=build /go/src/ebook/ebook-privilege/ebook-privilege /usr/bin/
EXPOSE ${port}
ENTRYPOINT [ "/bin/ebook-user" ]
