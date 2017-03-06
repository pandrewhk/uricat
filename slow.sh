#!/bin/sh

jot 100|while read a;do
    echo '{ "Url": "http://httpbin.org/get?'$a'" }'
    sleep 1
done
