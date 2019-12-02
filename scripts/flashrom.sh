#!/bin/sh

#TODO check $1 PROG device should exist

if [ -c "/dev/ttyPROG$1" ]; then
    echo "Device /dev/ttyPROG$1 exist"
else
    echo "Device /dev/ttyPROG$1 does not exist"
    exit 1
fi

check () {
    #echo $1
    CHIP=$(flashrom -p serprog:dev=/dev/ttyPROG$1:2000000 -V 2>&1 | \
        grep -oP '^Found [a-zA-z].* flash chip "([a-zA-z0-9].*)"' | \
        head -1 | \
        cut -d " " -f5 | \
        tr -d '"')
    echo ${CHIP}
}

case "$2" in
  check)
    CHIP=$(check $1)
    if [ -z "$CHIP" ]
    then
        echo "Memory chip not found"
        exit 1
    else
        echo "Found memory chip ${CHIP}"
    fi
    ;;
  read)
    if [ -z "$3" ]
    then
        echo "Output file not selected"
        exit 1
    fi
    if [ -f "$3" ]; then
        echo "File $3 already exist"
        exit 1
    fi
    CHIP=$(check $1)
    flashrom -p serprog:dev=/dev/ttyPROG$1:2000000 -c ${CHIP} -r $3
    ;;
  write)
    if [ -z "$3" ]
    then
        echo "Input file not selected"
        exit 1
    fi
    CHIP=$(check $1)
    flashrom -p serprog:dev=/dev/ttyPROG$1:2000000 -c ${CHIP} -w $3
    ;;
  write2)
    #TODO
    exit 1
    ;;
  *)
    echo "Usage: $0 device {check|read|write} [options...]"
    echo "       $0 device check"
    echo "       $0 device read outfile"
    echo "       $0 device write infile"
    echo "       $0 device write2 infile offset"
    exit 1
esac

exit $?

