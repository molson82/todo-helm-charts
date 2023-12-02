FROM golang:1.21.4-alpine3.18 AS build
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o bin/todo -v .

FROM alpine:3.18
EXPOSE 8880
COPY --from=build /app/bin/todo .
CMD [ "/todo"]
