#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <host>"
    exit 1
fi

ping -c 4 "$1"
