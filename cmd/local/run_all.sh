# This script will run the webserver and backend service on the local machine.

# Set environment variables(optional)
export USTART_AUTH_PORT="5001"
export USTART_PROF_PORT="5002"
export USTART_PROJ_PORT="5003"

export BACKEND_PORT="5000"


# Then, we run the services
# you should build and run binaries when in production
go run -o microservices_grpc/auth/auth microservices_grpc/auth/auth.go
go run -o microservices_grpc/profile/prof microservices_grpc/profile/profile.go
go run -o microservices_grpc/project/proj microservices_grpc/project/project.go
go run -o backend backend.go &>backend_log.txt &disown

# Next, we execute them
# ./cmd/backend/backend &>backend_log.txt &disown

echo $?

echo "Backend service called.  Follow logs with \`tail -f backend_log.txt\`"

# ./cmd/web/web &>webserver_log.txt &disown
# echo "Webserver service called.  Follow logs with \`tail -f webserver_log.txt\`"

echo "Services spawned.  Exiting..."

