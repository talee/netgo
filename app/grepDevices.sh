#!/bin/bash
./netgo devices | grep --color=none -o "\d\+\.\d\.\d\.\d\+\|ttext\">[^<]\+<" | grep --color=none -o ">[^<]\+" | grep --color=auto -C1 ">[^\.:]\+$"
