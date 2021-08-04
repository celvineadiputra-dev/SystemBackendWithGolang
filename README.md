# Selamat Datang System Backend With Golang

### Silahkan download golang
<a href="https://golang.org/">Download Golang</a>
<br/>
silahkan sesuaikan dengan OS masing masing dan install seperti biasa, jika sudah cek di terminal kamu
```
go version
```
jika tidak ada error maka selamat golang sudah terinstall di lokal mesin kamu

### Cara menjalankan Project

1. Clone project ini
```
git clone URL
```

2. Silahkan import terlebih dahulu file database (SQL) yang berada di folder SQL

3. Silahkan buka file main.go lalu, cari
```
dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
```

4. Setelah di temukan silahkan ubah datanya, sesuai dengan lokal mesin anda, formatnya
```
dsn := "(USER DATABASE):(PASSWORD DATABASE)@tcp((URL ATAU HOST MYSQL):(PORT MYSQL))/(NAMA DATABASE)?charset=utf8mb4&parseTime=True&loc=Local"
```

5. Bukan terminal anda lalu arahkan ke dalam folder project
```
go mod download
```

Atau manual install (optional) jika langkah 5 tidak berhasil ✨
```
go get -u github.com/gin-gonic/gin
```
```
go get -u gorm.io/driver/mysql
```
```
go get github.com/Sirupsen/logrus
```
```
go get github.com/speps/go-hashids/v2
```

6. Masih di terminal dan di dalam folder project
```
go run main.go
```

7. Silahkan lalukan import Collection API ke APIDEVELOPMENT yang kamu cintai💖


# Salam dari saya pengguna baru golang 🐱‍👤
