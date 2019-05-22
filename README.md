# UStart mono

Re factored beta2 for ustart. Monolithic design stores frontend and backend in one hub.
NOTE: This repo is in the middle of refactoring.  Delete this sentence when it is stable.

## How to run the servers

Every executable in this project lives in the `/cmd` folder.

The implementations of those executables live elsewhere – look at the import statements in `/cmd` to figure out where.

In general, you should always be able to execute `go run /cmd/svc/svc.go`

## appserver

Manages the U·START mobile applications

## backend

Contains backend services

## webserver

Contains web services
