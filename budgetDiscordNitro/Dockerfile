FROM golang:latest as builder

ARG GIT_COMMIT
ARG GIT_COMMIT="unversioned"
ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=$GIT_COMMIT"

EXPOSE 1337

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/src/app/budgetDiscordNitro /app/

EXPOSE 1337
