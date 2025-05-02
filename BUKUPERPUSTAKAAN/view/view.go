package view

import (
	"BUKUPERPUSTAKAAN/model"
	"BUKUPERPUSTAKAAN/node"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Insert() {
	var id int
	var nama, notelp, email, tglPinjam, tglKembali, judulBuku, penulisBuku string
	var idGenre int

	reader := bufio.NewReader(os.Stdin)

	// Input data Peminjam
	fmt.Print("Masukkan ID Peminjam: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Peminjam: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Nomor Telepon: ")
	notelp, _ = reader.ReadString('\n')
	notelp = strings.TrimSpace(notelp)

	fmt.Print("Masukkan Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Input data Buku
	fmt.Print("Masukkan Judul Buku: ")
	judulBuku, _ = reader.ReadString('\n')
	judulBuku = strings.TrimSpace(judulBuku)

	fmt.Print("Masukkan Penulis Buku: ")
	penulisBuku, _ = reader.ReadString('\n')
	penulisBuku = strings.TrimSpace(penulisBuku)

	fmt.Print("Masukkan Tanggal Pinjam (dd-mm-yyyy): ")
	tglPinjam, _ = reader.ReadString('\n')
	tglPinjam = strings.TrimSpace(tglPinjam)

	fmt.Print("Masukkan Tanggal Kembali (dd-mm-yyyy): ")
	tglKembali, _ = reader.ReadString('\n')
	tglKembali = strings.TrimSpace(tglKembali)

	fmt.Print("Masukkan ID Genre Buku: ")
	idGenreStr, _ := reader.ReadString('\n')
	idGenre, _ = strconv.Atoi(strings.TrimSpace(idGenreStr))

	// Simpan data peminjam dan buku yang dipinjam
	peminjam := node.Peminjam{
		ID:             id,
		Nama:           nama,
		NoTelp:         notelp,
		Email:          email,
		TanggalPinjam:  tglPinjam,
		TanggalKembali: tglKembali,
		BukuDipinjam: node.Buku{
			Judul:   judulBuku,
			Penulis: penulisBuku,
			IdGenre: idGenre,
		},
	}

	// Menambahkan peminjam ke dalam daftar peminjam
	model.CreatePeminjam(peminjam)
	fmt.Println("== Data Peminjam berhasil ditambahkan ==")
	fmt.Println()
}

func Views() {
	fmt.Println("== Daftar Peminjaman ==")
	for i, pj := range model.ReadPeminjam() {
		// Ambil genre buku berdasarkan ID
		genreBuku := model.GetNamaGenre(pj.BukuDipinjam.IdGenre)

		// Menampilkan informasi peminjam dan buku
		fmt.Println("--- Peminjaman ke -", i+1, " ---")
		fmt.Println("ID Peminjam\t: ", pj.ID)
		fmt.Println("Nama Peminjam\t: ", pj.Nama)
		fmt.Println("No Telepon\t: ", pj.NoTelp)
		fmt.Println("Email\t\t: ", pj.Email)
		fmt.Println("Judul Buku\t: ", pj.BukuDipinjam.Judul)
		fmt.Println("Penulis Buku\t: ", pj.BukuDipinjam.Penulis)
		fmt.Println("Tanggal Pinjam\t: ", pj.TanggalPinjam)
		fmt.Println("Tanggal Kembali\t: ", pj.TanggalKembali)
		fmt.Println("Genre\t\t: ", genreBuku)
		fmt.Println()
	}
}


func Update() {
	var id, idGenre int
	var nama, notelp, email, tglPinjam, tglKembali, judulBuku, penulisBuku string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Peminjam yang akan diupdate: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Peminjam: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan No Telepon: ")
	notelp, _ = reader.ReadString('\n')
	notelp = strings.TrimSpace(notelp)

	fmt.Print("Masukkan Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan Judul Buku: ")
	judulBuku, _ = reader.ReadString('\n')
	judulBuku = strings.TrimSpace(judulBuku)

	fmt.Print("Masukkan Penulis Buku: ")
	penulisBuku, _ = reader.ReadString('\n')
	penulisBuku = strings.TrimSpace(penulisBuku)

	fmt.Print("Masukkan Tanggal Pinjam (dd-mm-yyyy): ")
	tglPinjam, _ = reader.ReadString('\n')
	tglPinjam = strings.TrimSpace(tglPinjam)

	fmt.Print("Masukkan Tanggal Kembali (dd-mm-yyyy): ")
	tglKembali, _ = reader.ReadString('\n')
	tglKembali = strings.TrimSpace(tglKembali)

	fmt.Print("Masukkan ID Genre Buku: ")
	idGenreStr, _ := reader.ReadString('\n')
	idGenre, _ = strconv.Atoi(strings.TrimSpace(idGenreStr))

	// Ambil nama genre dari daftar buku
	namaGenre := model.GetNamaGenre(idGenre)

	peminjam := node.Peminjam{
		ID:             id,
		Nama:           nama,
		NoTelp:         notelp,
		Email:          email,
		TanggalPinjam:  tglPinjam,
		TanggalKembali: tglKembali,
		BukuDipinjam: node.Buku{
			Judul:     judulBuku,
			Penulis:   penulisBuku,
			IdGenre:   idGenre,
			NamaGenre: namaGenre,
		},
	}

	cek := model.UpdatePeminjam(peminjam, id)
	if cek {
		fmt.Println("== Data Peminjam berhasil diupdate ==")
	} else {
		fmt.Println("Peminjam gagal diupdate")
	}
	fmt.Println()
}


func Delete() {
	reader := bufio.NewReader(os.Stdin)
	var id int

	fmt.Print("Masukkan ID Peminjam yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Input ID tidak valid!")
		return
	}

	cek := model.DeletePeminjam(id)
	if cek {
		fmt.Println("== Data Peminjam berhasil dihapus ==")
	} else {
		fmt.Println("Peminjam gagal dihapus")
	}
	fmt.Println()
}
