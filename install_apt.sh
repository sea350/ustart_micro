apt install git golang tree -y

export GO111MODULE=on
go get google.golang.org/grpc@v1.28.1

PROTOC_VERSION=3.11.4# change as needed
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip
unzip protoc-$PROTOC_VERSION-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"
rm -f protoc-$PROTOC_VERSION-linux-x86_64.zip


export GOPATH=$HOME/go # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"
export PATH=$PATH:/usr/local/go/bin

go get github.com/olivere/elastic
env GIT_TERMINAL_PROMPT=1 go get github.com/sea350/ustart_micro