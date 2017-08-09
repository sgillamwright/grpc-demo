# GRPC Demo

This is a basic GRPC demo. It has a node server, with 3 clients talking with it (GO, PHP, NODE).

It contains:
- An example of SSL authentication (node client only)
- An example of a client interceptor (go client only)
- A streaming example (node client is only current implementation)

## Setup
- Clone the repo
- move the `clients/go` section if needed to be within your GO path.
- run `npm install`

### Server
To run the server, simply run `npm run server` from the command line

### Clients
To run the clients follow the steps below:

#### PHP
 - run `composer install` in the `clients/php` foler
 - run `npm run client:php` from command line

#### NODE
 - run `npm run client:node` from the command line

#### GO
 - If you cloned the repo within your GO path, you can simply run `npm run client:go` from the command line
 - Otherwise run `go run main.go` from the folder your go client is located
