FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /demo-login-server

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /demo-login-server /demo-login-server
COPY ./keys/app.rsa .
COPY ./keys/app.rsa.pub .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/demo-login-server"]