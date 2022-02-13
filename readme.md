# gRPC

## Definisi
gRPC merupakan metode pendistribusian data melalui sebuah protofile atau *stub*.

## Kenapa Harus Belajar gRPC?
- Lebih simple dibanding REST dari segi banyaknya HTTP Method
- Performa lebih baik dibanding REST (katanya)
- Terdapat file proto yang dapat menjadi API Contract antara client dan server

## Kekurangan gRPC
- harus ada file Proto
- payload yang lebih banyak

di REST:

```json
{
  "status": "success",
  "statusCode": "200",
  "message": "Success Creating new Book called Dunia Shopie with book id 53e8dad6-4600-478d-8c14-7b85d483d0a4"
}
```
di gRPC:
```json
{
  "status": "success",
  "_status": "status",
  "statusCode": "200",
  "_statusCode": "statusCode",
  "message": "Success Creating new Book called Dunia Shopie with book id 53e8dad6-4600-478d-8c14-7b85d483d0a4",
  "_message": "message"
}
```

## Proto File
- Merupakan file yang mendefinisikan service apa saja yang akan dibuat.
- Setiap service memiliki *message* parameter dan *message* response.
- Dalam membuat protofile, perlu mendefinisikan proto versi berapa dan masuk ke package apa.
- Dalam membuat message, terdapat 3 tipe variable. ***required, optional, repeated***. Namun, lebih baik menghindari penggunaan ***required*** karena ketika suatu waktu kita mengubah required ke optional, dari sisi client akan membaca bahwa parameter tersebut tetap required. Sehingga, akan menimbulkan error. Lebih baik, validasi required atau tidak dihandle di sisi aplikasi.
- angka dalam *message* merupakan urutan variable.

## Cara Implementasi
1. Buat "interface" service di dalam protofile.
2. tentukan struktur data parameter dan returnnya.
3. compile proto kita menjadi file `pb` yang berfungsi sebagai stub server kita. Kita dapat menggunakan command:

```cmd 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
modules/books/proto/books.proto //lokasi file proto
```
4. Implementasikan service yang sudah kita buat di file proto tadi. Contoh implementasinya bisa dilihat di `modules/books/repositories.go`
5. Buat server dan masukkan service yang sudah diimplementasikan tadi ke server agar terbaca.
6. panggil service yang sudah kita buat dengan gRPC-client atau dengan menggunakan bantuan aplikasi `bloomRPC`.