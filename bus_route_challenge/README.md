# Bus Route Challenge

### Problem

We are adding a new bus provider to our system. In order to implement a very
specific requirement of this bus provider our system needs to be able to filter
direct connections. We have access to a weekly updated list of bus routes
in form of a **bus route data file**. As this provider has a lot of long bus
routes, we need to come up with a proper service to quickly answer if two given
stations are connected by a bus route.


### Task

The bus route data file provided by the bus provider contains a list of bus
routes. These routes consist of an unique identifier and a list of stations
(also just unique identifiers). A bus route **connects** its list of stations.

Your task is to implement a micro service which is able to answer whether there
is a bus route providing a direct connection between two given stations. *Note:
The station identifiers given in a query may not be part of any bus route!*


### Bus Route Data

The first line of the data gives you the number **N** of bus routes, followed by
**N** bus routes. For each bus route there will be **one** line containing a
space separated list of integers. This list contains at least three integers. The
**first** integer represents the bus **route id**. The bus route id is unique
among all other bus route ids in the input. The remaining integers in the list
represent a list of **station ids**. A station id may occur in multiple bus
routes, but can never occur twice within the same bus route.

You may assume 100,000 as upper limit for the number of bus routes, 1,000,000 as
upper limit for the number of stations, and 1,000 as upper limit for the number
of station of one bus route. Therefore, your internal data structure should
still fit into memory on a suitable machine.

*Note: The bus route data file will be a local file and your service will get
the path to file as the first command line argument. Your service will get
restarted if the file or its location changes.*


### REST API

Your micro service has to implement a REST-API supporting a single URL and only
GET requests. It has to serve
`http://localhost:8088/api/direct?dep_sid={}&arr_sid={}`. The parameters
`dep_sid` (departure) and `arr_sid` (arrival) are two station ids (sid)
represented by 32 bit integers.

The response has to be a single JSON Object:

```
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "dep_sid": {
      "type": "integer"
    },
    "arr_sid": {
      "type": "integer"
    },
    "direct_bus_route": {
      "type": "boolean"
    }
  },
  "required": [
    "dep_sid",
    "arr_sid",
    "direct_bus_route"
  ]
}
```

The `direct_bus_route` field has to be set to `true` if there exists a bus route
in the input data that connects the stations represented by `dep_sid` and
`arr_sid`. Otherwise `direct_bus_route` must be set to `false`.




### Example Data

Bus Routes Data File:
```
3
0 0 1 2 3 4
1 3 1 6 5
2 0 6 4
```

Query:
```
http://localhost:8088/api/direct?dep_sid=3&arr_sid=6
```

Response:
```
{
    "dep_sid": 3,
    "arr_sid": 6,
    "direct_bus_route": true
}
```

### Implementation

Please implement your solution in Java, preferably Java 8. We expect you to
demonstrate best practices for general software development. Feel free to use
helpful open source libraries if applicable. We will evaluate your source code
as well as the functionality and compliance of the application.


### Packaging

Done with the fun part, the implementation? We have a few more requirements
for you that might sound boring but help us a lot in testing your code. Before you
send us your solution, have some build scripts for it.

Because we are a friendly bunch of developers, we share a (simplified) script.
Please open `service.sh` and fill out these 2 minimum steps:

```
dev_build() {
  # Do what you need to package your app, e.g. mvn package
}

dev_run() {
  # Do what you need to run your app, e.g. java -jar $DIR/target/magic.jar $*
}
```

Once you have filled out those 2 steps, you can do a bunch of things.
Run `./service.sh` to see usage. For a starter:

```
./service.sh           : prints usage
./service.sh dev_build : builds your app
./service.sh dev_run   : runs your app
```

### Smoke Tests

*Note: This smoke test only checks for compliance, not for correctness!*

```
./service.sh dev_smoke : runs our smoke tests on your machine
```

If you have `docker` installed on your machine, you can also run all the
scripts and tests inside a docker container:

```
./service.sh docker_build : packages your app into a docker image
./service.sh docker_run   : runs your app using a docker image
./service.sh docker_smoke : runs same smoke tests inside a docker container
```

### Shipping

The preferred option is a link to a **git** repository. GitHub, GitLab,
Bitbucket or your self hosted git. As long as `git clone LINK` works we're
happy. If you are not able to share a git repository, e.g. for privacy reasons,
please share a download link for a **zip** file. This zip file should contain a
single folder that contains your project. The content of the zip file should
look like this:

```
project_folder
├── src
├── data/
├── build.gradle/pom.xml
├── Dockerfile
└── service.sh
```
