#!/bin/sh

sed -n 's/[[:space:]]*F(\([a-zA-Z0-9_]*\)[[:space:]]*,[[:space:]]*\"[a-zA-Z0-9,=_[:space:]\"\\)]*/hi3516av200\/ko\/\1.o/p' hi3516av200/hi3516av200_ko.h | tr '\n' ' ' #| sed 's/.$//'
