yum install git golang tree -y

export PROTOC_ZIP=protoc-3.11.4-linux-x86_64.zip # change as needed
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/$PROTOC_ZIP
unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

export GIT_TAG=v1.2.0 # change as needed
export GOPATH=$HOME/go # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go

export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin

go get github.com/olivere/elastic
env GIT_TERMINAL_PROMPT=1 go get github.com/sea350/ustart_micro