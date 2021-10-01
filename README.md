# OMDB in Go

This exercise implements a clean architecture using the concept of a hexagonal architecture to fetch the OMDB api using the HTTP JSON API and GRPC-Protobuf.


## Requirement

- go 1.15+
- mysql server

## Config

In this app using [viper](https://github.com/gunturaf/omdb-server#readme) for configuration environment variable use `.env` file

## Running App

`make run`

## Test

`make test`
