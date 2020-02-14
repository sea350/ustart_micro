# This script will run the webserver and backend service on the local machine.

# Set environment variables(optional)
export USTART_AUTH_PORT="5001"
export USTART_PROF_PORT="5002"
export USTART_PROJ_PORT="5003"

export BACKEND_PORT="5000"


# Then, we build the services
go build -o backend backend.go
go build -o microservices_grpc/auth/auth microservices_grpc/auth/auth.go
go build -o microservices_grpc/profile/prof microservices_grpc/profile/profile.go
go build -o microservices_grpc/project/proj microservices_grpc/project/project.go


# Next, we execute them
./cmd/backend/backend &>backend_log.txt &disown

echo $?

echo "Backend service called.  Follow logs with \`tail -f backend_log.txt\`"

# ./cmd/web/web &>webserver_log.txt &disown
# echo "Webserver service called.  Follow logs with \`tail -f webserver_log.txt\`"

echo "Services spawned.  Exiting..."

