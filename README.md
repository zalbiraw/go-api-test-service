# go-api-test-service
This library allows you to launch different REST, GraphQL, GraphQL subgraphs, and Kafka producers to help test your APIs.

The available services are:
- rest
- users-rest
- posts-rest
- comments-rest
- users-graph
- posts-graph
- comments-graph
- notifications-graph
- users-subgraph
- posts-subgraph
- comments-subgraph
- notifications-subgraph (federated subscription)
- notifications-kafka-producer

The same docker file will allow you to run the different services. You can launch the service you want by setting the entrypoint to the docker container to be `entrypoint: ./$SERVICE_NAME/server`.

### Data
The data used in the users, posts, and comments is static data from the [jsonplaceholder](https://github.com/typicode/jsonplaceholder) library. 

The notification data is generated using the `loremipsum` library.

### How to use
You can use the `docker-compose.yml` to stand up and test these services. 

Run `docker-compose up` from the root of the repo. You should be able to access the services on the following ports:
- rest: `3100`
- users-rest: `3101`
- posts-rest: `3102`
- comments-rest: `3103`
- users-graph: `4101`
- posts-graph: `4102`
- comments-graph: `4103`
- notifications-graph: `4104`
- users-subgraph: `4201`
- posts-subgraph: `4202`
- comments-subgraph: `4203`
- notifications-subgraph: `4204`

### Containers
Containers are available on DockerHub under [zalbiraw/go-api-test-service](https://hub.docker.com/r/zalbiraw/go-api-test-service).

