const grpc = require('grpc');
const path = require('path');
const fs = require('fs');

// Set my proto paths
const BOOKS_PATH = path.join(__dirname + '/../proto/books.proto');

// Load my controllers
const BooksController = require(__dirname + '/controllers/BooksController');

// Load my proto files
const Books = grpc.load(BOOKS_PATH).Books;

// Create my server
const server = new grpc.Server();

// Add my proto services to the server
server.addService(Books.BooksService.service, new BooksController);

const creds = grpc.ServerCredentials.createInsecure();

// This is an example of SSL authentication. Only the node client is set up to use this
// const creds = grpc.ServerCredentials.createSsl(fs.readFileSync(path.join(__dirname + '/../ssl/ca.crt')), [{
//     cert_chain: fs.readFileSync(path.join(__dirname + '/../ssl/server.crt')),
//     private_key: fs.readFileSync(path.join(__dirname + '/../ssl/server.key'))
// }], true);

// Bind my server with insecure credentials
server.bind('0.0.0.0:50054', creds);

// Start my server
server.start();
