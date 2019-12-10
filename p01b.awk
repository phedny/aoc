#! /usr/bin/awk -f

NF {n = $1; while (n > 3) { n = int(n / 3 - 2); nn += n }}; END{print nn}
