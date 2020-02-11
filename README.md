# UStart mono

Re factored beta2 for ustart. Monolithic design stores frontend and backend in one hub.
NOTE: This repo is in the middle of refactoring.  Delete this sentence when it is stable.

## Installation

Make sure you have golang, protoc (and all associated dependancies), olivere/elastic for elastic7 installed.

Before running anything spin up `gen-proto.sh` you should then find auto generated go files in all `*pb` folders. After that you should be good to go.

## How to configure the servers

Iunno, but I'll figure it out. Sit tight!

## How to run the servers

Every executable in this project lives in the `/cmd` folder.

The implementations of those executables live elsewhere – look at the import statements in `/cmd` to figure out where.

In general, you should always be able to execute `go run /cmd/svc/svc.go`
