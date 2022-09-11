FROM golang:1.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]

## Base image for building the go project
#FROM golang:1.14-alpine AS build
#
## Updates the repository and installs git
#RUN apk update && apk upgrade && \
#    apk add --no-cache git
#
#WORKDIR /tmp/app
##ENV APP_HOME /go/src/tamircity
##RUN mkdir -p "$APP_HOME"
##WORKDIR "$APP_HOME"
#
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#
#COPY . .
#
## Builds the current project to a binary file called api
## The location of the binary file is /tmp/app/out/api
#RUN GOOS=linux go build -o ./out/api .
#
## The project has been successfully built and we will use a
## lightweight alpine image to run the server
#FROM alpine:latest
#
## Adds CA Certificates to the image
#RUN apk add ca-certificates
#
## Copies the binary file from the BUILD container to /app folder
#COPY --from=build /tmp/app/out/api /app/api
#
## Switches working directory to /app
#WORKDIR "/app"
#
## Exposes the 5000 port from the container
#EXPOSE 5000
#
## Runs the binary once the container starts
#CMD ["./api"]
#
## Later these in terminal type docker build -t tamircity-api .
## Then docker images
## Then docker run --name api --rm -p 5000:5000 tamircity-api