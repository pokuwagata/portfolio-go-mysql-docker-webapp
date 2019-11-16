#!/bin/bash

sed 's/\([\t ]*\)\(.*\)/\1"\2",/' $1 | pbcopy