#!/bin/sh

#also run `mvn install -Dcheckstyle.skip -DskipTests` in K Framework

./kompile-clear.sh

go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/node/kompile
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/standalone/kompile

go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/elrond/node/kompile
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/kompile

go build ./...
go test ./...