#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

PACKAGE_DIRS=`go list -e ./... | egrep -v "binary_output_dir|.git"`

# Run a godoc server which we will scrape. Clobber the GOPATH to include
# only our dependencies.
godoc -http=:6060 &
DOC_PID=$!

# Wait for the server to init
while :
do
    curl -s "http://localhost:6060" > /dev/null
    if [ $? -eq 0 ] # exit code is 0 if we connected
    then
        break
    fi
done

wget -r -np -N -E -p -k http://localhost:6060/pkg

# Stop the godoc server
kill -9 $DOC_PID

mv localhost\:6060 /tmp