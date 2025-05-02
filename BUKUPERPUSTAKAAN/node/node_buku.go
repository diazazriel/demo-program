package node

type Buku struct {
	Judul         string
	Penulis       string
	IdGenre       int
	NamaGenre     string
	DeskripsiGenre string
}

type ListBuku struct {
	Data Buku
	Link *ListBuku
}
