# go-api-test-service
This library allows you to launch different REST, GraphQL and GraphQL subgraphs to help test your APIs.

The available servcies are:
- rest
- users-rest
- posts-rest
- comments-rest
- users-subgraph
- posts-subgraph
- comments-subgraph
- notificaitons-subgraph (federated subscription)

The same docker file will allow you to run the different services. You can launch the service you want by setting the entrypoint to the docker container to be `entrypoint: ./$SERVICE_NAME/server`.

### Data
The data used in the users, posts, and comments is static data from the [jsonplaceholder](https://github.com/typicode/jsonplaceholder) library. 

The notification data is generated using the `loremipsum` library.

### How to use
You can use the `docker-compose.yml` to stand up and test these services. 

Run `docker-compose up` from the root of the repo. You should be able to access the services on the following ports:
- rest: `5000`
- users-rest: `5001`
- posts-rest: `5002`
- comments-rest: `50003`
- users-subgraph: `4001`
- posts-subgraph: `4002`
- comments-subgraph: `4003`
- notifications-subgraph: `4004`

### Containers
Containers are available on DockerHub under [zalbiraw/go-api-test-service](https://hub.docker.com/r/zalbiraw/go-api-test-service).

