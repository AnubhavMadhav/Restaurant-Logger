FROM golang:1.17-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o topmenu .

FROM alpine:latest
COPY --from=build /app/topmenu /usr/local/bin/topmenu
CMD ["topmenu"]