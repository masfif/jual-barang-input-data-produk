package main

import "fmt"

type Stack struct {
	items []Barang
}

// Push
func (s *Stack) Push(barang Barang) {
	s.items = append(s.items, barang)
	fmt.Println("Barang berhasil ditambahkan ke Stack!")
}

// Pop
func (s *Stack) Pop() Barang {
	if len(s.items) == 0 {
		fmt.Println("Stack kosong, tidak ada barang yang dapat dihapus.")
		return Barang{}
	}
	removedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Printf("Barag %s berhasil dihapus dari Stack!\n", removedItem.Nama)
	return removedItem
}

// Peek
func (s *Stack) Peek() Barang {
	if len(s.items) == 0 {
		fmt.Println("Stack kosong, tidak ada barang yang dapat dilihat.")
		return Barang{}
	}
	return s.items[len(s.items)-1]
}

// Display
func (s *Stack) Display() {
	if len(s.items) == 0 {
		fmt.Println(("Stack kosong, tidak ada barang yang dapat ditampilkan."))
		return
	}
	fmt.Println("Daftar barang di Stack:")
	for i, barang := range s.items {
		fmt.Printf("%d. ID=%d, Nama=%s, Jumlah=%d, Harga=%.2f\n", i+1, barang.ID, barang.Nama, barang.Jumlah, barang.Harga)
	}
}
