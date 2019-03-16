#!/bin/bash 

location="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"


docker run --rm -v "`dirname $location`:/csvgo-reader" -e "docker=1" golang:1.12 bash -c 'DEBUG=1 bash /csvgo-reader/test/run.sh'
#'/bin/bash /app/tests/run.sh'
