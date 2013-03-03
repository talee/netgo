#!/bin/bash
BIN=../../../../bin
if [ -f $BIN/netgo ];
then
	cd $BIN && ./netgo devices | grep --color=none -o "\d\+\.\d\.\d\.\d\+\|ttext\">[^<]\+<" | grep --color=none -o ">[^<]\+" | grep --color=auto -C1 ">[^\.:]\+$"
else
	echo "$BIN/netgo doesn't exist"
fi
