package books

import (
	context "context"
	"fmt"

	proto_books "learn-grpc/modules/books/proto-books"

	"database/sql"
)

type serverBook struct {
	// //ini perlu diimplementasikan karena akan meneruskan repo sebagai 'BookServer'
	// //yang diperlukan ketika menginstansiasi gRPC server
	proto_books.UnimplementedBookServer
	dbConfig *sql.DB
}

func NewBookServer(db *sql.DB) proto_books.BookServer {
	return &serverBook{
		dbConfig: db,
	}

}

func (r *serverBook) CreateBook(ctx context.Context, book *proto_books.CreateBookRequest) (*proto_books.BookResponse, error) {

	// Logic diterapkan di sini
	// Management data di database juga bisa diimplementasikan di sini
	// bisa dipisah menjadi service - repository seperti pada clean architecture
	status, message := "success", fmt.Sprintf("Success Creating new Book called %s with book id %s", book.GetName(), book.GetBookId())
	var statusCode int64 = 200

	return &proto_books.BookResponse{
		Status:     &status,
		StatusCode: &statusCode,
		Message:    &message,
		Data: &proto_books.BookItems{
			Data: []*proto_books.BookItem{
				{
					Id:     &statusCode,
					Name:   book.Name,
					BookId: book.BookId,
				},
			},
		},
	}, nil
}

func (r *serverBook) DeleteBook(ctx context.Context, book *proto_books.DeleteBookRequest) (*proto_books.BookResponse, error) {
	fmt.Println("bookId", book.BookId)
	status, message := "success", fmt.Sprintf("Success deleting book with ID: %d", *book.BookId)
	var statusCode int64 = 200
	return &proto_books.BookResponse{
		Status:     &status,
		StatusCode: &statusCode,
		Message:    &message,
	}, nil
}

func (r *serverBook) ListBooks(ctx context.Context, _ *proto_books.EmptyRequest) (*proto_books.BookResponse, error) {
	status, message := "success", fmt.Sprintf("Success getting books data")
	var statusCode int64 = 200
	bookData := []struct {
		id     int64
		title  string
		bookId string
	}{
		{
			id:     1,
			title:  "book 1",
			bookId: "bookID 1",
		},
		{
			id:     2,
			title:  "book 2",
			bookId: "bookID 2",
		},
		{
			id:     3,
			title:  "book 3",
			bookId: "bookID 3",
		},
	}

	return &proto_books.BookResponse{
		Status:     &status,
		StatusCode: &statusCode,
		Message:    &message,
		Data: &proto_books.BookItems{
			Data: []*proto_books.BookItem{
				{
					Id:     &bookData[0].id,
					Name:   &bookData[0].title,
					BookId: &bookData[0].bookId,
				},
				{
					Id:     &bookData[1].id,
					Name:   &bookData[1].title,
					BookId: &bookData[1].bookId,
				},
				{
					Id:     &bookData[2].id,
					Name:   &bookData[2].title,
					BookId: &bookData[2].bookId,
				},
			},
		},
	}, nil
}

func (r *serverBook) UpdateBook(ctx context.Context, book *proto_books.UpdateBookRequest) (*proto_books.BookResponse, error) {
	fmt.Println("bookId", *book.BookData.BookId)
	status, message := "success", fmt.Sprintf("Success updating book with bookID: %s", book.BookData.GetBookId())
	var statusCode int64 = 200

	return &proto_books.BookResponse{
		Status:     &status,
		StatusCode: &statusCode,
		Message:    &message,
		Data: &proto_books.BookItems{
			Data: []*proto_books.BookItem{
				{
					Id:     book.BookData.Id,
					Name:   book.BookData.Name,
					BookId: book.BookData.BookId,
				},
			},
		},
	}, nil
}
