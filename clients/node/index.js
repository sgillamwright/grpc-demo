const grpc = require('grpc');
const path = require('path');
const fs = require('fs');

// Set my proto path
const PROTO_PATH = path.join(__dirname, '../../proto/books.proto');
const Client = grpc.load(PROTO_PATH).Books;


let credentials = grpc.credentials.createInsecure();

// Example of SSL authentication
// let credentials = grpc.credentials.createSsl(
//   fs.readFileSync(path.join(__dirname + '/../../ssl/ca.crt')),
//   fs.readFileSync(path.join(__dirname + '/../../ssl/client.key')),
//   fs.readFileSync(path.join(__dirname + '/../../ssl/client.crt'))
// );

const BooksClient = new Client.BooksService('localhost:50054', credentials);

function getAllBooks () {
  console.log('GetAllBooks');

  BooksClient.getAllBooks({}, function (err, response) {
    if (err) {
      console.log('error', err);
      return;
    }

    console.log('response', response);
    return;
  });
};

function getAllBooksStream (calback) {
  var response = BooksClient.getAllBooksStream();

  response.on('data', (book) => {
    console.log('Got Book', book);
  });

  response.on('end', (callback) => {
    console.log('Stream ended');
  });
};

function getOneBook (id) {
  console.log('Get One Book');

  BooksClient.getOneBook(id, function (err, response) {
    if (err) {
      console.log('error', err);
      return err;
    }

    console.log('response', response);
    return;
  });
};

function addBook (book) {
  console.log('Add book');

  BooksClient.addBook(book, function (err, response) {
    if (err) {
      console.log('error', err);
      return;
    }

    console.log('response', response);
    return;
  });
};

function removeBook (id) {
  console.log('RemoveBook');

  BooksClient.removeBook(id, function (err, response) {
    if (err) {
      console.log('error', err);
      return;
    }

    console.log('response', response);
    return;
  });
};

removeBook('2');
getOneBook('2');
getAllBooks();

getAllBooksStream();

addBook({
  name: 'My First Book',
  genres: [
    {
      id: '1',
      name: 'Horror'
    },
    {
      id: '2',
      name: 'Thriller'
    }
  ],
  published: 'UNPUBLISHED',
  author: {
    id: '1',
    name: 'Author Name',
    website: 'My Fake Website',
    description: 'The worst'
  }
});

addBook({
  name: 'My Second Book',
  genres: [
    {
      id: '1',
      name: 'Horror'
    },
    {
      id: '2',
      name: 'Thriller'
    }
  ],
  published: 'UNPUBLISHED',
  author: {
    id: '1',
    name: 'Author Name',
    website: 'My Fake Website',
    description: 'The worst'
  }
});
