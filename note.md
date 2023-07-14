untuk menginisialisasi project baru dengan menggunakan go modules gunakan command `go mod init <nama-project>`
lalu untuk mengeksekusi file program dengan cara menggunakan command `go run <nama-file> [<nama-file>]`
tapi command `go run` hanya bisa digunakan pada file dengan nama package main saja
lalu untuk mengkompilasi file program menjadi atau menghasilkan file executable, maka dengan menggunakan command `go build`

- `go get <source>`

digunakan untuk mendownload package, package yang sudah terunduh disimpan didalam temporary folder yang ter-link dengan project folder dimana command `go get` di jalankan. Command `go get` harus dijalankan di dalam folder project, kalau tidak maka akan diunduh ke `GOPATH`.

- `go mod tidy`

digunakan untuk memvalidasi dependensi, jika ada dependensi yang belum terdownload, maka secara otomatis akan di download.
