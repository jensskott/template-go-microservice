# Template for go microservice

General boilerplate for creating a microservice in golang.
Contains basic files for building and testing the app and  
some standard setup in go for gin framework and how to model  
the code.


## Dockerfile and compose

Builds and tests the go app in the dockerfile then creates an  
artifact docker container with the specific built binary that  
it will execute on run. 
Compose file is to build and run locally with docker.

## Makefile

Different targets

test: test the code  
build: tests the code then builds binary  

test-docker: special test stage for docker to integrate with sonarqube  
build-docker: runs the docker build and creates the artifact container

## Code structure

TODO

## Usage

Prepare environment: 

* Setup golang [golang.org](https://golang.org)
* Install dep [Instructions](https://github.com/golang/dep)
* Clone this repo into `$GOPATH/src/github.com/getndazn`
* Rename the directory to your app `mv template-go-microservice my-cool-app`
* Reinit git, `rm -rf .git && git init`
* Get dependencies `dep ensure`

