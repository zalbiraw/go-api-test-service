# go-api-test-service
This library allows you to launch different federated GraphQL subgraphs to help test your APIs.

The available servcies are:
- users
- posts
- comments
- notificaitons (federated subscription)

The same docker file will allow you to run the different services. You can launch the service you want by setting the entrypoint to the docker container to be `entrypoint: ./$SERVICE_NAME/server`.

### Data
The data used in the users, posts, and comments is static data from the [jsonplaceholder](https://github.com/typicode/jsonplaceholder) library. 

The notification data is generated using the `loremipsum` library.

### How to use
You can use the `docker-compose.yml` to stand up and test these services. 

Run `docker-compose up` from the root of the repo. You should be able to access the services on the following ports:
- Users: `4000`
- Posts: `4001`
- Comments: `4002`
- Notifications: `4003 `

### Containers
Containers are available on DockerHub under [zalbiraw/go-api-test-service](https://hub.docker.com/r/zalbiraw/go-api-test-service).

