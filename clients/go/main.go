package main

import (
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "goGrpc/books"
)

const (
  address = "localhost:50054"
)

// Example of an interceptor on the client
func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
  log.Printf("Method Called: %s", method)
  return invoker(ctx, method, req, reply, cc, opts...)
}

func main() {
  var opts []grpc.DialOption
  opts = append(opts, grpc.WithInsecure())
  opts = append(opts, grpc.WithUnaryInterceptor(clientInterceptor))

  // Set up a connection to the server.
  conn, err := grpc.Dial(address, opts...)

  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }

  // Close my connection after run
  defer conn.Close()
  c := pb.NewBooksServiceClient(conn)

  // Contact the server and print out its response.
  r, err := c.GetAllBooks(context.Background(), &pb.EmptyMessage{})
  // r, err := c.GetOneBook(context.Background(), &pb.BookIDRequest{Id: "2"})

  if err != nil {
    log.Fatalf("could not get books: %v", err)
  }

  log.Printf("Books: %s", r)
}
