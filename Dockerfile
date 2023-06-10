# backend builder
FROM golang:1.20 AS backend-builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/rates/main.go

# frontend builder
FROM node:20-alpine3.17 AS frontend-builder
RUN mkdir /front
ADD ./frontend /front
WORKDIR /front
RUN npm install
RUN npm run build

# production
FROM alpine:latest AS production
COPY --from=backend-builder /app/app .
COPY --from=frontend-builder /front/dist/ /frontend
CMD ["./app"]