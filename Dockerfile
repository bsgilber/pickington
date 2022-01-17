FROM golang:1.17 AS build
WORKDIR /app/src/
RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make test build

FROM alpine:3.14 AS run
WORKDIR /app
COPY --from=build /app/src/dist/lambda .
ENV ENVIRONMENT=""

ENTRYPOINT [ "/app/lambda" ]
