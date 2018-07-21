# Use an official Golang runtime as a parent image
FROM golang:1.10.3

# env setup for reusebility
ENV APP_NAME=godoc-example
ENV WORKING_DIR=$GOPATH/src/github.com/marvin5064/$APP_NAME

# Set the working directory to golang working space
WORKDIR $WORKING_DIR

# Copy the current directory contents into the container at current directory
ADD . .