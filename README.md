## Requirements

We ask you to create a simple web server that should do the following:

* Runs locally on port 8000 and accepts `GET` requests at the index URL `/`
* It checks that the request has a query parameter called `favoriteTree` with a valid value
* For a successful request, returns a properly encoded HTML document with the following content:

If `favoriteTree` was specified (e.g. a call like `127.0.0.1:8000/?favoriteTree=baobab`):

```
It's nice to know that your favorite tree is a <value of "favoriteTree" from the url> 
```

if not specified (e.g. a call like `127.0.0.1:8000/`):

```
Please tell me your favorite tree
```

## Start Server

```sh
go run main.go
```

## Call the service with CURL

```sh
curl localhost:8000/?favoriteTree=baobab
```


## Improvements

- [ ] Run in docker container
- [ ] Write tests
- [ ] Use Go modules
- [ ] Restructure the folder (for example a separate file for handlers)