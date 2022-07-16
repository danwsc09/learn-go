#!/bin/bash

export DB_USER=wonsang
export DB_NAME=testdb
export DB_HOST=127.0.0.1

go build .
./data-access