#!/bin/sh

#TODO check $1 PROG device should exist

case "$2" in
  check)
    flashrom -p serprog:dev=/dev/ttyPROG$1:2000000 -V
    ;;
  read)
    #flashrom -p serprog:dev=/dev/ttyPROG$1:2000000 -r 
    ;;
  *)
    echo "Usage: $0 device {check|read} [options...]"
    echo "       $0 14 check"
    exit 1
esac

exit $?

