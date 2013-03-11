#!/bin/bash
netgo devices 2>/dev/null | grep --color=none -o "\d\+\.\d\.\d\.\d\+\|ttext\">[^<]\+<" | grep --color=none -o ">[^<]\+" | grep --color=auto -C1 ">[^\.:]\+$" || {
	echo "netgo doesn't exist in PATH"
}
