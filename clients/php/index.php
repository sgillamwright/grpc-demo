<?php

require dirname(__FILE__).'/vendor/autoload.php';

@include_once dirname(__FILE__).'/GPBMetadata/Books.php';

@include_once dirname(__FILE__).'/Books/BookIDRequest.php';
@include_once dirname(__FILE__).'/Books/BookListResponse.php';
@include_once dirname(__FILE__).'/Books/Book_Published.php';
@include_once dirname(__FILE__).'/Books/EmptyMessage.php';
@include_once dirname(__FILE__).'/Books/Genre.php';
@include_once dirname(__FILE__).'/Books/Author.php';
@include_once dirname(__FILE__).'/Books/Book.php';

@include_once dirname(__FILE__).'/Books/BooksServiceClient.php';

$client = new Books\BooksServiceClient('127.0.0.1:50054', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

function getAllBooks () {
  echo "Running Get All Books...\n\n";
  global $client;

  $emptyMessage = new Books\EmptyMessage();

  list($response, $status) = $client->getAllBooks($emptyMessage)->wait();

  foreach ($response->getBooks() as $key => $book) {
    echo "Book: \n";
    echo $book->getName() . "\n";
    echo $book->getAuthor()->getName() . "\n";

    foreach ($book->getGenres() as $key => $genre) {
      echo $genre->getName() . "\n";
    }

    echo "\n";
  }
}

getAllBooks();
