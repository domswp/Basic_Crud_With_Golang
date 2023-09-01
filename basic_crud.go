package main

import (
	"fmt"
	"os"
)

type Contact struct {
	Name    string
	Address string
	Phone   string
	Email   string
}

var contacts []Contact

func main() {
	for {
		fmt.Println("Pilih Menu:")
		fmt.Println("1. Tambah Kontak")
		fmt.Println("2. Tampilkan Kontak")
		fmt.Println("3. Edit Kontak")
		fmt.Println("4. Hapus Kontak")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan Menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addContact()
		case 2:
			displayContacts()
		case 3:
			editContact()
		case 4:
			deleteContact()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak terdaftar.")
		}
	}
}

func addContact() {
	var contact Contact
	fmt.Println("Tambah Kontak")
	fmt.Print("Nama: ")
	fmt.Scanln(&contact.Name)
	fmt.Print("Alamat: ")
	fmt.Scanln(&contact.Address)
	fmt.Print("Nomor Telepon: ")
	fmt.Scanln(&contact.Phone)
	fmt.Print("Email: ")
	fmt.Scanln(&contact.Email)

	contacts = append(contacts, contact)
	fmt.Println("Kontak berhasil ditambahkan.")
}

func displayContacts() {
	fmt.Println("Daftar Kontak:")
	for i, contact := range contacts {
		fmt.Printf("%d. Nama: %s, Alamat: %s, Nomor Telepon: %s, Email: %s\n", i+1, contact.Name, contact.Address, contact.Phone, contact.Email)
	}
}

func editContact() {
	fmt.Print("Masukkan kontak yang ingin diedit: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(contacts) {
		fmt.Println("Kontak tidak valid.")
		return
	}

	var contact Contact
	fmt.Print("Nama: ")
	fmt.Scanln(&contact.Name)
	fmt.Print("Alamat: ")
	fmt.Scanln(&contact.Address)
	fmt.Print("Nomor Telepon: ")
	fmt.Scanln(&contact.Phone)
	fmt.Print("Email: ")
	fmt.Scanln(&contact.Email)

	contacts[index-1] = contact
	fmt.Println("Kontak berhasil diedit.")
}

func deleteContact() {
	fmt.Print("Masukkan kontak yang ingin dihapus: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(contacts) {
		fmt.Println("Kontak tidak valid.")
		return
	}

	contacts = append(contacts[:index-1], contacts[index:]...)
	fmt.Println("Kontak berhasil dihapus.")
}

