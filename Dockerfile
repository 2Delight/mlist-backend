FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN make build

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/mlist-backend ./mlist-backend
COPY --from=builder /app/migrations ./migrations
ENTRYPOINT ["./mlist-backend"]
