FROM node:20-alpine3.18 AS frontend-builder
WORKDIR /app

COPY ./frontend/package.json ./frontend/package-lock.json ./
RUN npm ci

COPY frontend .
COPY .env ../.env
RUN npm run build

FROM golang:1.21-alpine3.18 AS backend-builder
WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o twp

FROM alpine:3.18 AS runner
WORKDIR /app

COPY --from=frontend-builder /app/dist ./frontend/dist
COPY --from=backend-builder /app/twp .

ENTRYPOINT [ "/app/twp" ]
