FROM golang:1.10 AS builder

# need dep, want dep, can't not have dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# copy, compile
WORKDIR $GOPATH/src/github.com/yashdalfthegray/goecho
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 80
ENTRYPOINT ["./app"]
