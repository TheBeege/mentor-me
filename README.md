# Mentor Me
Mentor-Me is a simple mentor matchmaking app. Users can sign up as a mentor and/or mentee. Mentors can list a skill they'd like to mentor in along with their level of expertise in said skill. Mentees can search for mentors by skill and level. A mentee can submit a request for mentorship upon finding a desired mentor. The mentor can then accept or reject the request. Upon acceptance, the mentor and mentee will receive each other's contact information.

This is the front-end for the Mentor Me project, this time using Beego. You can find the [API here](https://github.com/TheBeege/mentor-me-api).

## Setup
1. [Install Go](https://golang.org/doc/install)
    * You may need to add $GOPATH/bin to your $PATH
2. [Install Beego](https://beego.me/docs/install/)
3. [Install dep](https://github.com/golang/dep) - `go get -u github.com/golang/dep/...`
4. Run `dep ensure` to resolve dependencies
3. Run `bee run -downdoc=true -gendoc=true` to generate documentation and run the server

## Front-end check
To play with static front-end assets, see `localhost:8080/static/example.html` once you have the server running. 