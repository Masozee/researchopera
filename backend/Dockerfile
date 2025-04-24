# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /app/backend /backend
EXPOSE 8080
USER nonroot:nonroot
CMD ["/backend"]

