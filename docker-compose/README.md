# Setup testing environment 

## How To

1. Launch testbed

    ```bash
    cd eventpublishplugin
    make start-dev-env
    ```

2. Setup proxy for build binary
    
    ````bash
    git clone https://github.com/trusted-cloud/trusted-cloud-proxy.git

    cd trusted-cloud-proxy
    make release-image
    make REPO_TOKEN=<TOEKN> proxy-up-pegasus-network
    ````
    **Note:** Replace your token with value of `<TOKEN>`.

3. Compile and Run

    3-1. Inside eep-server container.

    If eep-server container is enabled, you can login to the container and compile the code manually.

    * Use `go run` inside container. 
        ```bash
        docker exec -ti eep-server bash
        make GO_PROXY=$GOPROXY start
        ```

    * or use `go build` to create binary in container.
        ```bash
        docker exec -ti ep-server bash
        make GO_PROXY=$GOPROXY go-build
        ```

    **Note:** `GOPROXY` env variable is defined in `.env` file, and will be passed to the container.

    3-2. In DevContainer

    If you are using VSCode devcontainer, you should use the following configuration to launch the container. 
    Then you can run the plugin in the devcontainer.

    ```json
        "runArgs": [
            "--network=proxy-network",
            "--network=pegasus-cloud-network"
        ],
        "mounts": [
            { 
                "source": "pegasus-cloud-eventpublishplugin", 
                "target": "/run", 
                "type": "volume" 
            }
        ]
    ```

4. Invoke gRPC service from gRPC UI in port 8080.
    
    **Note:** Launching gRPC UI will failed if sockets are not available. It will restart automatically and until the sockets are available. 

5. Verify PubSub from Redis Insight in port 5540

6. Tear down 

    ```bash
    # goproxy and VSCode connect to pegasus-cloud-network, shutdown these two first 
    
    # Close VSCode devcontainer

    # Close proxy
    cd trusted-cloud-proxy
    make proxy-down

    # Then tear down test environment
    make purge-dev-all
    ```

## Architecture


```
                      (volume mount on /run )
               +----------------------------------+
            +--+ pegasus-cloud-eventpublishplugin +-------+
            |  +-----------------------+----------+       |
            |                          |                  | (listen on 8080)    (listen on 5540)
  +---------+----+               +-----+-----+         +--+----------+         +--------------+
  |    VSCode    |               |           |         |             |         |              |
  |              |               | eep-server|         |   gRPC UI   |         | RedisInsight |
  | devContainer |               |           |         |             |         |              |
  +-------+--+---+               +-----+-----+         +------+------+         +-------+------+
          |  |                         |                      |                        |
          |  +---------------+         |                      |                        |
          |                  |         |                      |                        |
+---------+--------+     +---+---------+----------------------+------------------------+-----------------+
|                  |     |                                                                               |
|   proxy-network  |     |                               pegasus-cloud-network                           |
|                  |     |                                   172.40.0.0/16                               |
+---------+--------+     +---+---------+------------------+-------------------+-------------------+------+
          |                  |         |                  |                   |                   |
          |  +---------------+         |                  |                   |                   |
          |  |                         |                  |                   |                   |
     +----+--+-+                 +-----+----+      +------+------+    +-------+-------+    +------+-----+
     |         |                 |          |      |    Redis    |    | Event Publish |    |   Network  |
     | GoProxy |                 | RabbitMQ |      |      &      |    |    Plugin     |    |            |
     |         |                 |          |      |   Sentinel  |    |   (STOPPED)   |    |  (STOPPED) |
     +---------+                 +----------+      +-------------+    +---------------+    +------------+
   (listen on 8078)
```