#!/bin/bash
CURRENT_PKG=`go list -e`

# Run a godoc server which we will scrape. Clobber the GOPATH to include
# only our dependencies.
godoc -http=:6060 & DOC_PID=$!

# Wait for the server to start
until curl -sSf "http://localhost:6060/pkg/$CURRENT_PKG/" > /dev/null
do
    echo "waiting for godoc server ..."
    sleep 1
done
sleep 1

echo "generating documents HTML/CSS/JS files for package ${CURRENT_PKG} ..."
# Scrape the pkg directory for the API docs. Scrap lib for the CSS/JS. Ignore everything else.
# The output is dumped to the directory "localhost:6060".
wget -r -p \
    -e robots=off \
    --include-directories="/lib/godoc,/pkg/$CURRENT_PKG,/src/$CURRENT_PKG" \
    --exclude-directories="/pkg/$CURRENT_PKG/vendor,/src/$CURRENT_PKG/vendor" \
    "http://localhost:6060/pkg/$CURRENT_PKG/"

# Stop the godoc server
kill -9 $DOC_PID

mv localhost:6060 /tmp