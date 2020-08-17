FROM golang:1.14.7-alpine3.12 as build

WORKDIR /go/src/ebook/ebook-user
COPY ./ /go/src/ebook/ebook-user
RUN  apk --no-cache add git \
  && go env -w  GOPROXY=https://goproxy.cn \
  && go get -d \
  && CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o ebook-user main.go \
  && chmod +x ebook-user


FROM alpine:3.12

ENV port=50801
COPY --from=build /go/src/ebook/ebook-user/ebook-user /bin/
EXPOSE ${port}
ENTRYPOINT [ "/bin/ebook-user" ]
