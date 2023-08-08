#! /bin/bash

head=$(git ls-remote git@github.com:Horryportier/go-jisho.git HEAD)
words=($head)

repo="github.com/Horryportier/go-jisho"


echo -e "$(go get "$repo@${words[0]}")"
