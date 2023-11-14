#!/bin/sh
ori=$(pwd)

cd /app

go work init

go work use -r ./

cd ${ori}
/go/bin/air
