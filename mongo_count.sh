#!/bin/bash

while true; do
    s=$(mongo 127.0.0.1:27017 < mongo_count.js | tail -4 | grep -v bye)
    sum=$(echo $s | awk '{print $1+$2+$3}')
    echo $s $sum
    sleep 60
done
