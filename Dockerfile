FROM golang:1.22.6-alpine3.20 as build
ARG app_name

WORKDIR /${app_name}
COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd/server/
# ------------------------------------------------------------------------------------------------------------------------
FROM alpine:3.20
ARG app_name

WORKDIR /${app_name}
COPY --from=build /${app_name}/bin/main ./

CMD ["./main"]