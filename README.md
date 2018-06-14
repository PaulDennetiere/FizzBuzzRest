# Fizz Buzz exercise

This is a Rest Server that can received parametrized query and will respond a json representing the list of output for fizzbuzz with those parameters.

It is written entirely in Go (aka Golang).

## Usage
____
### Build

To build this project you will have to get its dependencies. 
You can use [dep](https://github.com/golang/dep) to vendor those dependencies.
```sh
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

Once you have the dependencies you can build the project :
```sh
# While at the project's root folder:
go build .
```
____
### Run

This project need some configuration in order to run.
This configuration will set the **port** to listen to, the **read timeout** for the incoming connections (default is 5s), and the **graceful stop time** which is the time to wait for finishing connections before shutting down the server (default is 5s).

This configuration is based on environment variables, and a script is provided to set them. 
```sh
# To set you environment variables
source env.sh
```
Then only you can run the binary: 
```sh
./fizzbuzz_rest
```
____
### Usage

This expose a rest API, that you can query easily. This API will always return a JSON. It contains the data if everything is OK, an error describing the problem otherwise. This API will only accept 'GET' method, and will trigger errors if parameters doesn't respect internal rules.   

Example : 
```sh
curl 'http://localhost:8080/?int1=3&int2=5&string1=fizz&string2=buzz&limit=10'
# {"data":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz"]}

curl 'http://localhost:8080/?int1=3&int2=5&string1=fizz&string2=buzz&limit=10' -X 'POST'
# {"error":"Bad method"}
```
____
## Tests

Business code and handlers are tested. 

To run those tests you can: 
```sh
# While at project's root folder
go test -p=1 -cover -race `go list ./... | grep -v /vendor/`
```
This command will run every tests, print code coverage, and test if race condition can be detected.