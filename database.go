package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Koneksi ke database
func ConnectDatabase() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:8111)/jual_barang")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Koneksi ke database berhasil!")
}

// Tambahkan barang ke database
func AddBarangToDatabase(barang Barang) {
	query := "INSERT INTO barang (nama, jumlah, harga) VALUES (?, ?, ?)"
	_, err := db.Exec(query, barang.Nama, barang.Jumlah, barang.Harga)
	if err != nil {
		fmt.Println("Gagal menambahkan barang ke database:", err)
		return
	}
	fmt.Println("Barang berhasil disimpan ke database!")
}

// Ambil semua barang dari database
func GetAllBarangFromDatabase() []Barang {
	query := "SELECT id, nama, jumlah, harga FROM barang"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Gagal mengambil barang dari database:", err)
		return nil
	}
	defer rows.Close()

	var barangList []Barang
	for rows.Next() {
		var barang Barang
		err := rows.Scan(&barang.ID, &barang.Nama, &barang.Jumlah, &barang.Harga)
		if err != nil {
			fmt.Println("Gagal membaca data:", err)
			continue
		}
		barangList = append(barangList, barang)
	}
	return barangList
}
