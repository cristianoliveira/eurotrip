FROM ubuntu:16.04

RUN apt-get -qq -y update && \
    apt-get -qq -y install openjdk-8-jdk maven gradle curl && \
    rm -rf /var/cache/apt /var/lib/apt/lists/*

WORKDIR /src
ADD . .
RUN ./service.sh test && \
    ./service.sh build

CMD ["/src/service.sh", "run", "tests/example"]
