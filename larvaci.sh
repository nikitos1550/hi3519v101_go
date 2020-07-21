#!/bin/bash
# Do here all neeeded for CI tests

make NO_USER_MAKEFILE=Y prepare >&2
cd tests && authbind --deep python -m ci
