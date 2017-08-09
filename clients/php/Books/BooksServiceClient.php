<?php
// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Syntax the proto file uses
namespace Books {

  // My Books Service
  class BooksServiceClient extends \Grpc\BaseStub {

     // @param string $hostname hostname
     // @param array $opts channel options
     // @param \Grpc\Channel $channel (optional) re-use channel object
    public function __construct($hostname, $opts, $channel = null) {
      parent::__construct($hostname, $opts, $channel);
    }

     // Get All Books
     // @param \Books\EmptyMessage $argument input argument
     // @param array $metadata metadata
     // @param array $options call options
    public function getAllBooks(\Books\EmptyMessage $argument,
      $metadata = [], $options = []) {
      return $this->_simpleRequest('/Books.BooksService/getAllBooks',
      $argument,
      ['\Books\BookListResponse', 'decode'],
      $metadata, $options);
    }

     // Stream of books
     // @param \Books\EmptyMessage $argument input argument
     // @param array $metadata metadata
     // @param array $options call options
    public function getAllBooksStream(\Books\EmptyMessage $argument,
      $metadata = [], $options = []) {
      return $this->_serverStreamRequest('/Books.BooksService/getAllBooksStream',
      $argument,
      ['\Books\Book', 'decode'],
      $metadata, $options);
    }

     // Get Single Book
     // @param \Books\BookIDRequest $argument input argument
     // @param array $metadata metadata
     // @param array $options call options
    public function getOneBook(\Books\BookIDRequest $argument,
      $metadata = [], $options = []) {
      return $this->_simpleRequest('/Books.BooksService/getOneBook',
      $argument,
      ['\Books\Book', 'decode'],
      $metadata, $options);
    }

     // Add Book
     // @param \Books\Book $argument input argument
     // @param array $metadata metadata
     // @param array $options call options
    public function addBook(\Books\Book $argument,
      $metadata = [], $options = []) {
      return $this->_simpleRequest('/Books.BooksService/addBook',
      $argument,
      ['\Books\Book', 'decode'],
      $metadata, $options);
    }

     // Remove Book
     // @param \Books\BookIDRequest $argument input argument
     // @param array $metadata metadata
     // @param array $options call options
    public function removeBook(\Books\BookIDRequest $argument,
      $metadata = [], $options = []) {
      return $this->_simpleRequest('/Books.BooksService/removeBook',
      $argument,
      ['\Books\EmptyMessage', 'decode'],
      $metadata, $options);
    }

  }

}
