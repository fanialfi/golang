package main // nama package

import "fmt"

func main()  {
  fmt.Println("Hello World")
}

/*
 untuk menginisialisasi project baru dengan menggunakan go modules gunakan command `go mod init <nama-project>`
 lalu untuk mengeksekusi file program dengan cara menggunakan command `go run <nama-file> [<nama-file>]`
 tapi command `go run` hanya bisa digunakan pada file dengan nama package main saja
 lalu untuk mengkompilasi file program menjadi atau menghasilkan file executable, maka dengan menggunakan command `go build`
*/
