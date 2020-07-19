#!/bin/bash
# Do here all neeeded for CI tests

make prepare >&2
cd tests && authbind --deep python -m ci
