FROM golang:1.14.7-alpine3.12 as build

WORKDIR /go/src/ebook/ebook-ebook
COPY ./ /go/src/ebook/ebook-ebook
RUN  apk --no-cache add git \
  && go env -w  GOPROXY=https://goproxy.cn \
  && go get -d \
  && CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o ebook-ebook main.go \
  && chmod +x ebook-ebook


FROM alpine:3.12

ENV port=50801
COPY --from=build /go/src/ebook/ebook-ebook/ebook-ebook /bin/
EXPOSE ${port}
ENTRYPOINT [ "/bin/ebook-ebook" ]
