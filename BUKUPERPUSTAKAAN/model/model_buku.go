package model

import "BUKUPERPUSTAKAAN/node"

var DaftarBuku node.ListBuku

// CREATE
func CreateBuku(bk node.Buku) bool {
	tempLL := node.ListBuku{
		Data: bk,
		Link: nil,
	}
	if DaftarBuku.Link == nil {
		DaftarBuku.Link = &tempLL
		return true
	} else {
		temp := &DaftarBuku
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

// READ
func ReadBuku() []node.Buku {
	daftarBuku := []node.Buku{}
	temp := &DaftarBuku
	for temp.Link != nil {
		daftarBuku = append(daftarBuku, temp.Link.Data)
		temp = temp.Link
	}
	return daftarBuku
}

// UPDATE
func UpdateBuku(bk node.Buku, id int) bool {
	temp := DaftarBuku.Link
	for temp != nil {
		if temp.Data.IdGenre == id {
			temp.Data = bk
			return true
		}
		temp = temp.Link
	}
	return false
}

// DELETE
func DeleteBuku(id int) bool {
	temp := &DaftarBuku
	for temp.Link != nil {
		if temp.Link.Data.IdGenre == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

// SEARCH
func SearchBuku(id int) bool {
	temp := &DaftarBuku
	for temp.Link != nil {
		if temp.Link.Data.IdGenre == id {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetNamaGenre(id int) string {
	temp := &DaftarBuku
	for temp.Link != nil {
		if temp.Link.Data.IdGenre == id {
			return temp.Link.Data.NamaGenre
		}
		temp = temp.Link
	}
	return ""
}
