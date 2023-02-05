FROM golang:1.19-alpine AS build

WORKDIR /build-dir

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/ .

RUN CGO_ENABLED=0 go build -o ./out/rewards_checker .

FROM alpine:3.9.6

COPY --from=build /build-dir/out/rewards_checker /app/rewards_checker

CMD ["/app/rewards_checker"]