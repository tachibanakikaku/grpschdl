#!/bin/sh

VERSION=alpha-001

if [ $# = 0 ]; then
    echo "Usage: $0 <project> [<project> ...]"
fi

gcloud config configurations activate default

while [ $# -gt 0 ]; do
    read -p "Deploy to $1 (Y/n)? " ans
    if [ x"$ans" = x"" ]; then
        ans="Y"
    fi
    if [ x"$ans" = x"Y" ]; then
        # app.yaml cron.yaml index.yaml queue.yaml
        gcloud app deploy *.yaml --project=$1 -v $VERSION

    elif [ x"$ans" = x"n" ]; then
        echo "Input is n, so skip app [$1]"
    else
        echo "Invalid input. Skip app [$1]"
    fi
    shift;
done
