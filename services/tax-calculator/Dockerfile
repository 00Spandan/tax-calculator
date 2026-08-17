FROM golang:1.23-alpine AS build
WORKDIR /workspace

COPY gen/go ./gen/go
COPY services/tax-calculator ./services/tax-calculator

WORKDIR /workspace/services/tax-calculator
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/tax-calculator ./cmd/server

FROM gcr.io/distroless/static-debian12
ENV PORT=8080
EXPOSE 8080
COPY --from=build /out/tax-calculator /tax-calculator
ENTRYPOINT ["/tax-calculator"]
