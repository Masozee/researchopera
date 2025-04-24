# syntax=docker/dockerfile:1
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o backend .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /app/backend /backend
EXPOSE 8080
CMD ["/backend"]
