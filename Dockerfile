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

FROM microsoft/dotnet AS resource
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource AS tests
COPY --from=builder /tests /go-tests
WORKDIR /go-tests
RUN set -e; for test in /go-tests/*.test; do \
		$test; \
	done

FROM resource