#!/usr/bin/env bash

# This whole script will run under this bus_route_challenge folder
# irrespective of where it was invoked from
cd $(dirname $0)

dev_test() {
  # Do what you need to run your unit tests, e.g. mvn test
  true
}

dev_build() {
  # Do what you need to package your app, e.g. mvn package
  true
}

dev_run() {
  # Do what you need to run your app in the foreground
  # e.g. java -jar target/magic.jar $*
  sleep 600
}

builtin_help() {
  cat <<EOF
Usage:
  $0 <command> <args>
Developer-defined commands:
  test             : run your unit tests
  build            : builds and packages your app
  run <data-file>  : starts your app in the foreground
Built-in commands:
  smoke            : runs our integration test suite on localhost
  docker-build     : packages your app into a docker image
  docker-run       : runs your app using a docker image
  docker-smoke     : runs same smoke tests inside a docker container
EOF
}

builtin_smoke() {
  if builtin_smoke_run $*; then
    echo "Tests Passed"
    exit 0
  else
    echo "Tests Failed"
    exit 1
  fi
}

builtin_smoke_run() {
  baseUrl="http://localhost:8088"
  echo "Running smoke tests on $baseUrl..." && \
    (curl -fsS "$baseUrl/api/direct?dep_sid=3&arr_sid=4" | grep -E 'true|false') && \
    (curl -fsS "$baseUrl/api/direct?dep_sid=0&arr_sid=1" | grep -E 'true|false')
}

docker_build() {
  docker build -t goeuro:devtest .
}

docker_run() {
  docker run --rm -it goeuro:devtest
}

docker_smoke() {
  containerId=$(docker run -d goeuro:devtest)
  echo "Waiting 5 seconds for service to start..."
  sleep 5
  docker exec $containerId /src/service.sh smoke
  retval=$?
  docker rm -f $containerId
  exit $retval
}

cmd="$1"
cmd=${cmd:-"help"}
shift
args="$*"

case $cmd in
  test)
    dev_test
    ;;
  build)
    dev_build
    ;;
  run)
    dev_run $args
    ;;
  smoke)
    builtin_smoke
    ;;
  docker-build)
    docker_build
    ;;
  docker-run)
    docker_build
    docker_run
    ;;
  docker-smoke)
    docker_build
    docker_smoke
    ;;
  help)
    builtin_help
    ;;
  *)
    echo "Unknown command: $cmd"
    builtin_help
    exit 1
esac
