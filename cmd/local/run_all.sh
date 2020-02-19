# This script will run the webserver and backend service on the local machine.

# Set environment variables(optional)
export USTART_AUTH_PORT="5001"
export USTART_PROF_PORT="5002"
export USTART_PROJ_PORT="5003"

export BACKEND_PORT="5000"


# Then, we run the services
# you should build and run binaries with log files when in production
cd microservices_grpc/auth/
screen -d -m -S auth go run microservices_grpc/auth/run.go
cd ../profile/
screen -d -m -S profile go run microservices_grpc/profile/run.go
cd ../project/
screen -d -m -S project go run microservices_grpc/project/run.go
cd ../..
screen -d -m -S backend go run backend backend.go

# Next, we execute them
# ./cmd/backend/backend &>backend_log.txt &disown

echo $?

echo "Backend service called.  Follow logs with \`screen -ls\`"
echo "Then select service to follow with \`screen -r *screen name*\`"

# ./cmd/web/web &>webserver_log.txt &disown
# echo "Webserver service called.  Follow logs with \`tail -f webserver_log.txt\`"

echo "Services spawned.  Exiting..."

