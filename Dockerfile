FROM golang:1.19 as build

WORKDIR /go/src/github.com/daniil49926/FastTubeAnalytics/app

COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/starter cmd/api/main.go

FROM alpine:3.11.13

COPY --from=build /go/src/github.com/daniil49926/FastTubeAnalytics/app/configs/prod/api.toml ./configs/prod/api.toml
COPY --from=build /go/bin/starter .

EXPOSE 8350

ENTRYPOINT [ "./starter" ]