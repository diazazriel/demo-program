package main

import (
	"BUKUPERPUSTAKAAN/model"
	"BUKUPERPUSTAKAAN/node"
	"BUKUPERPUSTAKAAN/view"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func menuUtama() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner((reader))

	for {
		reader.Discard(reader.Buffered())

		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Data Peminjam")
		fmt.Println("2. Tampilkan Data Peminjam")
		fmt.Println("3. Update Data Peminjam")
		fmt.Println("4. Hapus Data Peminjam")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		if scanner.Scan() {
			input := scanner.Text()
			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Tidak Valid, silakan masukkan angka")
				continue
			}
			switch choice {
			case 1:
				view.Insert()
			case 2:
				view.Views()
			case 3:
				view.Update()
			case 4:
				view.Delete()
			case 5:
				fmt.Println("Terima Kasih.")
				return
			default:
				fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			}
		}
	}
}

func main() {
	// Tambahkan data genre buku dengan deskripsi di awal
	model.CreateBuku(node.Buku{IdGenre: 1, NamaGenre: "Novel", DeskripsiGenre: "Kisah Inspiratif"})
	model.CreateBuku(node.Buku{IdGenre: 2, NamaGenre: "Self Improvement", DeskripsiGenre: "Mengubah kebiasaan kecil untuk hidup yang lebih baik"})
	model.CreateBuku(node.Buku{IdGenre: 3, NamaGenre: "Romansa", DeskripsiGenre: "Kisah cinta penuh emosi dan perjalanan hati"})
	model.CreateBuku(node.Buku{IdGenre: 4, NamaGenre: "Fantasi", DeskripsiGenre: "Petualangan di dunia lain dengan berbagai makhluk fantastis"})

	// Tampilkan daftar buku (IdGenre, Genre, dan Deskripsi Genre)
	fmt.Println("Daftar Genre Buku:")
	for _, buku := range model.ReadBuku() {
		fmt.Printf("IdGenre\t\t: %d\n", buku.IdGenre)
		fmt.Printf("Genre\t\t: %s\n", buku.NamaGenre)
		fmt.Printf("Deskripsi genre: %s\n\n", buku.DeskripsiGenre)
	}

	// Jalankan menu utama
	menuUtama()
}
