#!/bin/bash
for file in `find . -type d -not -ipath "*git*" -print`; do
    url=`echo $file | sed 's#^.#https://godoc.org/github.com/gf0842wf/unioffice#'`
    echo $url
    curl -s $url -o /dev/null
    sleep 10
done
