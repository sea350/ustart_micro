# UStart mono

Re factored beta2 for ustart. Monolithic design stores frontend and backend in one hub.
NOTE: This repo is in the middle of refactoring.  Delete this sentence when it is stable.

## How to configure the servers

Iunno, but I'll figure it out. Sit tight!

## How to run the servers

Every executable in this project lives in the `/cmd` folder.

The implementations of those executables live elsewhere – look at the import statements in `/cmd` to figure out where.

In general, you should always be able to execute `go run /cmd/svc/svc.go`
