
package model

import "BUKUPERPUSTAKAAN/node"

var DaftarPeminjam node.ListPeminjam

// CREATE
func CreatePeminjam(pj node.Peminjam) bool {
	tempLL := node.ListPeminjam{
		Data: pj,
		Link: nil,
	}
	if DaftarPeminjam.Link == nil {
		DaftarPeminjam.Link = &tempLL
		return true
	} else {
		temp := &DaftarPeminjam
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

// READ
func ReadPeminjam() []node.Peminjam {
	daftarPeminjam := []node.Peminjam{}
	temp := &DaftarPeminjam
	for temp.Link != nil {
		daftarPeminjam = append(daftarPeminjam, temp.Link.Data)
		temp = temp.Link
	}
	return daftarPeminjam
}

// UPDATE
func UpdatePeminjam(pj node.Peminjam, id int) bool {
	temp := DaftarPeminjam.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Data = pj
			return true
		}
		temp = temp.Link
	}
	return false
}

// DELETE
func DeletePeminjam(id int) bool {
	temp := &DaftarPeminjam
	for temp.Link != nil {
		if temp.Link.Data.ID == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}
