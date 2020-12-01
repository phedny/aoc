#! /usr/bin/awk -f

NF {n += int($1 / 3) - 2}; END{print n}
