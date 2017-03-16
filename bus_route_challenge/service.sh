#!/usr/bin/env bash
DIR=$(cd `dirname $0`; pwd)

dev_test() {
  # Do what you need to run your unit tests
  # e.g. mvn test
  true
}

dev_build() {
  # Do what you need to package your app
  # e.g. mvn package
  true
}

dev_run() {
  # Do what you need to run your app
  # e.g. java -jar $DIR/target/magic.jar $*
  sleep 600
}

builtin_help() {
  cat <<EOF
Usage:
  $0 <command> <args>
Developer-defined commands:
  test            : run your unit tests
  build           : builds and packages your app
  run <data-file> : starts your app in the foreground
Built-in commands:
  e2e <base-url>  : runs our integration test suite
  docker-build    : packages your app into a docker image
  docker-run      : runs your app using a docker image
EOF
}

builtin_e2e() {
  if builtin_e2e_run; then
    echo "Tests Passed"
    exit 0
  else
    echo "Tests Failed"
    exit 1
  fi
}

builtin_e2e_run() {
  baseUrl=$1
  baseUrl=${baseUrl:-"http://localhost:8088"}
  echo "Running e2e tests on $baseUrl..."
  false
  curl -fsS "$baseUrl/api/direct?dep_sid=3&arr_sid=4" | grep -E 'true|false'
  curl -fsS "$baseUrl/api/direct?dep_sid=0&arr_sid=1" | grep -E 'true|false'
}

docker_build() {
  docker build -t goeuro:devtest $DIR
}

docker_run() {
  docker run --rm -it goeuro:devtest
}

docker_e2e() {
  containerId=$(docker run -d goeuro:devtest)
  echo "Waiting 5 seconds for service to start..."
  sleep 5
  docker exec $containerId /src/service.sh e2e
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
  e2e)
    builtin_e2e $args
    ;;
  docker-build)
    docker_build
    ;;
  docker-run)
    docker_run
    ;;
  docker-e2e)
    docker_e2e
    ;;
  help)
    builtin_help
    ;;
  *)
    echo "Unknown command: $cmd"
    builtin_help
    exit 1
esac
