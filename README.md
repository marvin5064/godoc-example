# How to Generate Customized Go Documents for Private Repository

## Get Started in Local Machine

### Godoc

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

### Check it out

after above steps, we can see all documentations about golang native library and current repository
```
open http://localhost:6060/pkg
open http://localhost:6060/pkg/github.com/marvin5064/godoc-example/
```

## Extensive
Let s further exclude the golang native library from the documentation

After some research, `godoc` no yet support specific file documentation, there are some tricks for it
1. Generate static web files for the repository you are working with, here we using docker with a simple script to generate it
```
#!/bin/bash
CURRENT_PKG=`go list -e`

# run a godoc server
godoc -http=:6060 & DOC_PID=$!

# Wait for the server to start
until curl -sSf "http://localhost:6060/pkg/$CURRENT_PKG/" > /dev/null
do
    sleep 1
done
sleep 1

# recursive fetch entire web including CSS & JS
# turn off robots check, otherwise might get blocked with details in `robots.txt` file
# only get the directories we are looking for
wget -r -p \
    -e robots=off \
    --include-directories="/lib/godoc,/pkg/$CURRENT_PKG,/src/$CURRENT_PKG" \
    --exclude-directories="/pkg/$CURRENT_PKG/vendor,/src/$CURRENT_PKG/vendor" \
    "http://localhost:6060/pkg/$CURRENT_PKG/"

# Stop the godoc server
kill -9 $DOC_PID

# all file will be generated into `localhost:6060` folder, hence we move them out from docker to local machine
mv localhost:6060 /tmp
```
you can simply achieve this by `Makefile`
```
make docker-gen-docs
```
2. Serve it out by nodejs http server package as a static web page
```
make docker-gen-docs-server
```
3. Check it out!
```
open http://localhost:8080/pkg/github.com/marvin5064/godoc-example/
```

### Reference
- https://github.com/golang/go/issues/2381
- https://groups.google.com/forum/#!topic/golang-nuts/rzdX87PojAo
- https://forum.golangbridge.org/t/running-go-doc-only-for-my-code/2832

### PS - Improvements

- use NGINX as a replacement of normal node server 
    - NGINX is native for static website serving
    - Leverage http rules out of application/server side
- use docker-compose
    - hence we can run on docs generation & http server on same container
    - if you run on a CD/CI process, different container do different job will be better
- *Extensive* version has too many drawbacks
    - wait for golang community to make improvement on the godoc
    - documents are generated by crawler, it can by bugger at certain point
    - lack of their party documentation
        - if the repo itself has third party dependencies (vendor). For the approach with docker, I would suggest move all files in `vendor` to `$GOPATH/src`, and generate the docs right after.

[Chinese Version](https://marvin-code.blogspot.com/2018/07/golang-godoc.html)