# How to Generate Customized Go Documents for Private Repository
## Get Started in Local Machine

### godoc
the simple tool, provided by golang community, will help you get local customized go documents in a nutshell
```
godoc -http=:6060
```
after that, you will be able to review all go document under
```
open http://localhost:6060/pkg
```
Of course, the detail of documentation has to be done by you inline with code. This repository has some implementation in place, also here is [more godoc styling tricks](https://godoc.org/github.com/fluhus/godoc-tricks)

However, the documents generated will include all directories and subdirectories under `$GOPATH/src`, most of them may not be **necessary**

## Docker
### Dockerfile
```
# Use an official Golang runtime as a parent image
FROM golang:1.10.3

# env setup for reusability
ENV APP_NAME=godoc-example
ENV WORKING_DIR=$GOPATH/src/github.com/marvin5064/$APP_NAME

# Set the working directory to golang working space
WORKDIR $WORKING_DIR

# Copy the current directory contents into the container at current directory
ADD . .

CMD godoc -http=:6060
```
### Build & Run
```
docker build . \
    -t godoc-example/v1 # giving a tag for reference
docker run \
    -p 6060:6060 \ # (local):(docker) as we expose 6060 port in docker
    -it godoc-example/v1 # run for a built image by tag
```
for simplicity, the cmd has been stored in `Makefile`, you can simply trigger `make docker-run` to execute it

### Check it out:
after above steps, we can see all documentations about golang native library and current repository
```
open http://localhost:6060/pkg
open http://localhost:6060/pkg/github.com/marvin5064/godoc-example/
```
