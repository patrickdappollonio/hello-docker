# "Hello, Docker!" container

[![Build Status](https://travis-ci.org/patrickdappollonio/hello-docker.svg?branch=master)](https://travis-ci.org/patrickdappollonio/docker-http-server)
[![Docker Automated buil](https://img.shields.io/docker/automated/patrickdappollonio/hello-docker.svg)](https://hub.docker.com/r/patrickdappollonio/hello-docker/)

This Docker container is just a `Hello, world!` example to practice with containers in general.
This container is useful to demonstrate routing capabilities, since, by default, it prints
the Hostname instead of the word `world`, so when ran in a container, it'll be the Container ID.

When load-balanced, the ID will jump between the different containers where the request is being
offloaded.

You can change two settings: the port the container will be listening at, which is managed by the
`$PORT` environment variable, and the name you want to be printed to the screen, in this case, managed
by the `$NAME` environment variable.

## Container published to the docker registry

The docker container [is published in the public Docker Registry](https://hub.docker.com/r/patrickdappollonio/hello-docker/)
under `patrickdappollonio/hello-docker`, you can pull it by executing:

```bash
docker pull patrickdappollonio/hello-docker
```

Then, to actually run it and print the Hostname, you'll do:

```bash
docker run -p 5000:8000 patrickdappollonio/hello-docker
```

Then just access `127.0.0.1:5000` in your browser to see the `Hello, $NAME!` message.
