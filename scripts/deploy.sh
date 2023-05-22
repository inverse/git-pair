#!/bin/bash

cd dist

for file in *.rpm *.deb; do
    fury push $file --api-token=$FURY_TOKEN
done
