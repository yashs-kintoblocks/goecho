# goecho

Just a simple Go server that echoes the data sent to it.

## Installation

You're going to need `dep` for this project. Clone the repository and run `dep ensure`. If you don't have your Go path set up, keep reading.

## Running

You can run this project solely through Docker. The included `Dockerfile` is set up to pull down the dependencies, build the project and then create an image out of just the executable.

Run `docker build -t goecho .` to get an image built and then run `docker run -it --name <some_name> -p 8080:80 goecho` to get the server started. The container itself exposes port 80.

## API

### `GET /`

This is the standard hello world response just to check that the server is running.

### `GET /get`

This is the first of the two echo endpoints on the server. This will echo back any query parameters sent to the server in the response. For example,

For the query

```shell
curl http://localhost:8080/get?foo=bar&baz=bat
```

The server will return

```json
{
    "status": "ok",
    "queryParams": {
        "foo": "bar",
        "baz": "bat"
    }
}
```

### `POST /post`

The second endpoint echoes the body of the POST request back in the response. For example,

For the query

```shell
curl -X POST -d '{ "foo": "bar", "baz": "bat" }' http://localhost:8080/post
```

The server will return

```json
{
    "status": "ok",
    "body": {
        "foo": "bar",
        "baz": "bat"
    }
}
```

For this endpoint, the server only supports strings as values, anything else will lead to unintended behavior.

## Resources

* [Gorilla Mux reference](http://www.gorillatoolkit.org/pkg/mux)
* [Mux Logging](http://www.gorillatoolkit.org/pkg/handlers#LoggingHandler)
* [Gorilla Handlers](https://github.com/gorilla/handlers)
* [RESTful API in Mux](https://thenewstack.io/make-a-restful-json-api-go/)
* [KintoHub Go example](https://github.com/kintohub/go-examples)
