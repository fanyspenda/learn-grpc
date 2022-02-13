package books

import (
	context "context"
	"fmt"

	proto_books "learn-grpc/modules/books/proto-books"
)

type serverBook struct {
	//ini perlu diimplementasikan karena akan meneruskan repo sebagai 'BookServer'
	//yang diperlukan ketika menginstansiasi gRPC server
	proto_books.UnimplementedBookServer
}

func NewBookServer() *serverBook {
	s := &serverBook{}
	return s
}

func (r *serverBook) CreateBook(ctx context.Context, book *proto_books.CreateBookRequest) (*proto_books.BookResponse, error) {

	// Logic diterapkan di sini
	// Management data di database juga bisa diimplementasikan di sini
	status, message := "success", fmt.Sprintf("Success Creating new Book called %s with book id %s", book.GetName(), book.GetBookId())
	var statusCode int64 = 200

	return &proto_books.BookResponse{
		Status:     &status,
		StatusCode: &statusCode,
		Message:    &message,
	}, nil
}
