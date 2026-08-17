FROM golang:1.23-alpine AS build
WORKDIR /workspace

COPY gen/go ./gen/go
COPY services/bff ./services/bff

WORKDIR /workspace/services/bff
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/bff ./cmd/server

FROM gcr.io/distroless/static-debian12
ENV PORT=8080
EXPOSE 8080
COPY --from=build /out/bff /bff
ENTRYPOINT ["/bff"]
