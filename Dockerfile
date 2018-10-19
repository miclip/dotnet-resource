FROM golang:alpine as builder
COPY . /go/src/github.com/miclip/dotnet-resource
ENV CGO_ENABLED 0
RUN go build -o /assets/in github.com/miclip/dotnet-resource/cmd/in
RUN go build -o /assets/out github.com/miclip/dotnet-resource/cmd/out
RUN go build -o /assets/check github.com/miclip/dotnet-resource/cmd/check
WORKDIR /go/src/github.com/miclip/dotnet-resource
RUN set -e; for pkg in $(go list ./...); do \
		go test -o "/tests/$(basename $pkg).test" -c $pkg; \
	done

FROM microsoft/dotnet:2.2-sdk-alpine AS resource
COPY --from=builder assets/ /opt/resource/
RUN mkdir -p /etc/BUILDS/ && \
    printf "Build of miclip/dotnet-resource, date: %s\n"  `date -u +"%Y-%m-%dT%H:%M:%SZ"` > /etc/BUILDS/alpine-golang && \
    apk add curl && \
    curl https://storage.googleapis.com/golang/go1.11.linux-amd64.tar.gz | tar xzf - -C / && \
    mv /go /goroot && \
    apk del curl && \
    rm -rf /var/cache/apk/* && \
    chmod +x /opt/resource/*

ENV GOROOT=/goroot \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PATH=${PATH}:/goroot/bin:/gopath/bin
    

FROM resource AS tests
COPY --from=builder /tests /go-tests
WORKDIR /go-tests
RUN set -e; for test in /go-tests/*.test; do \
		$test; \
	done

FROM resource