# BNMO-Backend
## Description
Merupakan bagian backend dari web application BNMO. Backend menggunakan bahasa pemrograman Go dengan framework Gin Gonic. 

## Daftar Isi
* [Technology Stack](#technology-stack)
* [Design Patterns](#design-patterns)
* [Requirements](#requirements)
* [Setup](#setup)
* [Endpoints](#endpoints)
* [Created By](#created-by)

## Technology Stack
* Bahasa Go dengan framework Gin Gonic
* MariaDB database (terdapat 3 buah tabel, yaitu tabel Account, tabel Requests, dan tabel Transfers)
* Redis cache (data di dalam redis di set untuk hilang setelah 24 jam)
* GORM (Go Object Relational Mapper)
* Docker (menyimpan database MariaDB dengan versi paling terbaru, Redis cache dengan versi paling terbaru, dan bahasa Go dengan versi paling terbaru)

## Design Patterns
Terdapat 2 design pattern yang digunakan dalam proses pembuatan backend dari web application BNMO ini, antara lain:
* **Facade**: digunakan untuk routing, ketika API backend dipanggil oleh frontend, maka facade akan menerima panggilan tersebut dan menjalankan proses lain setelahnya (seperti mengakses database, mengakses redis cache, atau melakukan kalkulasi).
* **Singleton**: digunakan untuk objek database dan redis cache. Objek database dibentuk ketika program utama dijalankan dan akan terus menerus dipanggil apabila data yang tersimpan dibutuhkan atau akan diganti. Objek redis akan dibentuk ketika bagian frontend memanggil API yang membutuhkan konversi mata uang, dan setelah pembentukannya, maka objek redis ini akan terus menerus dipanggil untuk mengambil data di dalamnya.

## Requirements
- [Docker](https://docs.docker.com/desktop/install/windows-install/)

## Setup
1. Download file ZIP dari repository ini atau clone repository ke dalam komputer Anda
2. Pastikan Docker **sudah di-install** dan dapat dijalankan
3. Buka command prompt atau powershell dan pindahkan direktori ke folder tempat repository ini berada
4. Jalankan command berikut untuk melakukan inisialisasi Dockerfile
```
docker build -t bnmo/redis -f Dockerfile .
```
5. Setelah build Docker selesai, jalankan command berikut untuk melakukan inisialisasi service yang dibutuhkan di dalam container
```
docker-compose up -d
```
6. Setelah selesai, maka untuk menghubungkan program dengan database MariaDB dan redis cache, maka command berikut **perlu** dijalankan kembali
```
docker build -t bnmo/redis -f Dockerfile .
```
7. Apabila *Windows Defender Firewall* muncul, maka tekan `Allow Access` 
8. Bagian backend telah berhasil dijalankan beserta database dan redis cache
9. Apabila Anda ingin melihat logs dari program utama, maka bisa menjalankan command 
```
docker logs -f bnmo
```
10. Untuk menghentikan docker container, maka jalankan command
```
docker-compose stop
```

## Endpoints
Terdapat 18 total endpoints yang digunakan dalam web application ini, dibagi kedalam 5 endpoints yang tidak membutuhkan autentikasi dan 13 endpoints yang membutuhkan autentikasi untuk berjalan. Seluruh endpoints dapat dilihat di file `router.go` pada folder routes.
### Tidak membutuhkan autentikasi
* **/api/register (Method: POST)**
  
  Endpoint ini dipanggil ketika user ingin melakukan registrasi. Data user akan dicek dan dimasukkan ke dalam database. Password yang dimasukkan oleh user akan di-hash menggunakan fungsi bcrypt. Apabila terdapat kesalahan input dari user, maka endpoint akan mengembalikan status beserta pesannya.
* **/api/upload-image (Method: POST)**
  
  Endpoint ini dipanggil ketika user melakukan upload gambar KTP mereka. Gambar tersebut akan disimpan ke dalam folder images dan mengembalikan URL.
* **/api/login (Method: POST)**
  
  Endpoint ini dipanggil ketika user melakukan login. Data user akan dicek dengan data yang terdapat di dalam database. Apabila terdapat kesalahan input dari user, maka endpoint akan mengembalikan status beserta pesannya. Akun yang belum diverifikasi oleh admin tidak akan bisa login.
* **/api/logout (Method: POST)**

  Endpoint ini dipanggil ketika user melakukan logout. Cookie yang terdapat di browser akan diganti dengan cookie baru yang memiliki durasi sangat pendek, sehingga akan hilang.
* **/api/uploads (Method: Static)**

  Endpoint ini dipanggil ketika gambar yang dimasukkan oleh user diperlukan oleh bagian frontend. 

### Membutuhkan autentikasi
#### Sisi customer
* **/api/profile/:id (Method: GET)**
  
  Endpoint ini dipanggil ketika user ingin melihat profile mereka. Data profile akan dikirimkan sesuai dengan id yang diberikan.
* **/api/customerrequest/add (Method: POST)**
  
  Endpoint ini dipanggil ketika user mengirimkan request untuk menambah balance mereka. Angka yang diminta beserta mata uang yang dipakai akan dikalkulasi dan menghasilkan nilai yang sudah dikonversi ke IDR. Request ini akan disimpan ke dalam database.
* **/api/customerrequest/subtract (Method: POST)**
  
  Endpoint ini dipanggil ketika user mengirimkan request untuk mengurangi balance mereka. Angka yang diminta beserta mata uang yang dipakai akan dikalkulasi dan menghasilkan nilai yang sudah dikonversi ke IDR. Request ini akan disimpan ke dalam database.
* **/api/displayaccounts (Method: GET)**

  Endpoint ini dipanggil ketika user berada di halaman transfer. Endpoint ini akan mengembalikan data semua akun yang berada di dalam database, dan digunakan untuk memberikan pilihan destinasi transfer.
* **/api/transfer (Method: POST)**

  Endpoint ini dipanggil ketika user melakukan transfer dari akun mereka ke akun destinasi pilihan. Jumlah transaksi beserta mata uang yang dipilih akan dikalkulasi dan menghasilkan nilai yang sudah dikonversi ke IDR. Apabila balance dari user tidak mencukupi, maka transfer akan gagal. Data transfer ini akan disimpan ke dalam database.
* **/api/requesthistory (Method: GET)**
  
  Endpoint ini dipanggil ketika user membuka halaman request history. Endpoint akan mengembalikan seluruh request history sesuai dengan ID user yang sedang membukanya. Pagination dilakukan oleh backend dengan cara mengirimkan hanya 5 data setiap pemanggilan beserta metadatanya yang berisi jumlah halaman maksimal dan total history yang ada.
* **/api/transferhistory (Method: GET)**
  
  Endpoint ini dipanggil ketika user membuka halaman transfer history. Endpoint akan mengembalikan seluruh transfer history sesuai dengan ID user yang sedang membukanya. Pagination dilakukan oleh backend dengan cara mengirimkan hanya 5 data setiap pemanggilan beserta metadatanya yang berisi jumlah halaman maksimal dan total history yang ada.
* **/api/updateimage (Method: POST)**
  
  Endpoint ini dipanggil ketika user melakukan perubahan gambar KTP mereka. Endpoint akan mengganti data gambar di database, menghapus gambar yang terdapat di folder images, dan menggantinya dengan gambar baru.

#### Sisi Admin
* **/api/displayrequest (Method: GET)**

  Endpoint ini dipanggil ketika admin membuka halaman validate request. Endpoint akan mengembalikan data berisi request yang masih berstatus pending. Pagination juga dilakukan oleh backend dengan cara yang sama seperti request dan transfer history.
* **/api/validaterequest (Method: POST)**

  Endpoint ini dipanggil ketika admin memilih untuk **accept** atau **reject** request dari user. Status request di dalam database akan berubah sesuai dengan pilihan admin. Apabila status accept, maka pengubahan nilai balance user akan dilakukan. Apabila status accept tetapi user tidak memiliki balance yang cukup (dalam kasus subtract balance), maka status akan menjadi reject.
* **/api/displaypending (Method: GET)**
  
  Endpoint ini dipanggil ketika admin membuka halaman customer verification. Endpoint akan mengembalikan data berisi akun yang masih berstatus pending. Pagination juga dilakukan oleh backend dengan cara yang sama seperti request dan transfer history.
* **/api/validateaccount (Method: POST)**
  
  Endpoint ini dipanggil ketika admin memilih untuk **accept** atau **reject** akun yang ada. Status request di dalam database akan berubah sesuai dengan pilihan admin. Apabila status accept, maka user akan diperbolehkan untuk login ke dalam aplikasi. Apabila status reject, maka user akan ditolak untuk login ke dalam aplikasi.
* **/api/customerdata (Method: GET)**
  
  Endpoint ini dipanggil ketika admin membuka halaman customer data. Endpoint akan mengembalikan data berisi seluruh akun yang melakukan registrasi ke dalam aplikasi dan sudah di verifikasi oleh admin. Tidak terdapat pagination pada endpoint ini.

## Created By
Nama                      | NIM
----                      | ---
Yakobus Iryanto Prasethio | 13520104