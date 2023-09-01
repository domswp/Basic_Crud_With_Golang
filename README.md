# Basic Crud With Golang
Aplikasi ini adalah contoh sederhana dari operasi CRUD (Create, Read, Update, Delete) untuk mengelola daftar kontak. Aplikasi ini dibuat menggunakan bahasa pemrograman Go.

## Cara Menggunakan

1. Pastikan Anda memiliki Go terinstal di komputer Anda. Jika belum, Anda dapat mengunduh dan menginstalnya dari [situs web resmi Go](https://golang.org/dl/).

2. Clone repositori ini ke komputer Anda atau unduh kode sumbernya sebagai ZIP.

3. Buka terminal dan navigasi ke direktori tempat Anda menyimpan kode sumber.

4. Jalankan aplikasi dengan perintah berikut: ``go run basic_crud.go``

5. Anda akan melihat menu dengan opsi berikut:
- Tambah Kontak
- Tampilkan Kontak
- Edit Kontak
- Hapus Kontak
- Keluar

6. Pilih opsi yang sesuai dengan apa yang ingin Anda lakukan.

## Struktur Kode

- `basic.crud.go`: File utama yang berisi kode aplikasi.
- `Contact`: Struktur data untuk menyimpan informasi kontak.
- Fungsi-fungsi seperti `addContact()`, `displayContacts()`, `editContact()`, dan `deleteContact()` mengelola operasi CRUD.

## Catatan

- Aplikasi ini hanya menyimpan data kontak dalam memori. Data akan hilang setelah aplikasi ditutup. Untuk implementasi yang lebih permanen, Anda perlu mempertimbangkan penggunaan database atau penyimpanan berkas.

- Ini adalah contoh sederhana dan bukan aplikasi produksi sebenarnya. Anda dapat mengembangkan lebih lanjut sesuai kebutuhan Anda.
