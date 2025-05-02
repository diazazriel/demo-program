package node

type Peminjam struct {
	ID             int
	Nama           string
	NoTelp         string
	Email          string
	TanggalPinjam  string
	TanggalKembali string
	BukuDipinjam   Buku  // Menambahkan relasi ke Buku yang dipinjam
}

type ListPeminjam struct {
	Data Peminjam
	Link *ListPeminjam
}
