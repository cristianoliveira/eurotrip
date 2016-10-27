from ubuntu:16.04

RUN apt-get update && apt-get install -y \
    openjdk-8-jdk \
    git \
    maven \
    gradle \
    unzip \
 && rm -rf /var/lib/apt/lists/*

COPY example /data/example
COPY get_repo.sh /scripts/
COPY build.sh /scripts/
COPY run.sh /scripts/
ARG path
ARG type
COPY files /files
WORKDIR /data/test
RUN bash /scripts/get_repo.sh $type $path
RUN bash /scripts/build.sh
CMD ["bash", "/scripts/run.sh"]
