# Mentor Me

This is the front-end for the Mentor Me project, this time using Beego

## Setup
1. [Install Go](https://golang.org/doc/install)
    * You may need to add $GOPATH/bin to your $PATH
2. [Install Beego](https://beego.me/docs/install/)
3. [Install dep](https://github.com/golang/dep) - `go get -u github.com/golang/dep/...`
4. Run `dep ensure` to resolve dependencies
3. Run `bee run -downdoc=true -gendoc=true` to generate documentation and run the server

## Front-end check
To play with static front-end assets, see `localhost:8080/static/example.html` once you have the server running. 