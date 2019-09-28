#!/bin/sh

sed -n 's/[[:space:]]*F(\([a-zA-Z0-9_]*\)[[:space:]]*,[[:space:]]*\"[a-zA-Z0-9,=_[:space:]\"\\)]*/hi3516av200\/ko\/\1.o/p' ${CHIPFAMILY}/${CHIPFAMILY}_ko.h | tr '\n' ' ' #| sed 's/.$//'
