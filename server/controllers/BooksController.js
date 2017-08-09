'use strict';

let grpc = require('grpc');
let db = [];

class Books {
  getAllBooks (call, callback) {
    console.log('Get All Books');
    return callback(null, { books: db });
  }

  getAllBooksStream (call, callback) {
    console.log('Get All Books Stream');
    db.forEach((book) => {
      call.write(book);
    });

    call.end(callback);
  };

  getOneBook (call, callback) {
    console.log('Get One Book');
    for (let i = 0; i < db.length; i++) {
      if (db[i].id === call.request.id) {
        return callback(null, db[i]);
      }
    }

    return callback({
      code: grpc.status.NOT_FOUND,
      message: 'Not Found'
    });
  }

  addBook (call, callback) {
    console.log('Add Book');
    let newBook = Object.assign({}, call.request);
    newBook.id = (db.length + 1) + '';
    db.push(newBook);
    return callback(null, newBook);
  }

  removeBook (call, callback) {
    console.log('Remove Book');
    for (let i = 0; i < db.length; i++) {
      if (db[i].id === call.request.id) {
        db.splice(i, 1);
        return callback(null);
      }
    }

    return callback({
      code: grpc.status.NOT_FOUND,
      message: 'Can not find book by id'
    });
  }
}

module.exports = Books;
