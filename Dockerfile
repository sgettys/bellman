FROM golang:1.16-alpine as builder

ARG REVISION

RUN mkdir -p /bellman/

WORKDIR /bellman

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
	-X gitlab.com/sgettys/bellman/pkg/version.REVISION=${REVISION}" \
	-a -o bin/bellman cmd/bellman/*

FROM alpine:3.14

LABEL maintainer="sgettys"

RUN addgroup -S app \
	&& adduser -S -G app app \
	&& apk --no-cache add \
	ca-certificates curl netcat-openbsd jq

WORKDIR /home/app

COPY --from=builder /bellman/bin/bellman .

RUN chown -R app:app ./

USER app

CMD ["./bellman"]
