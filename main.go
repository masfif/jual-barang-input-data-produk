package main

import "fmt"

func main() {
	// Koneksi ke database
	ConnectDatabase()

	// Membuat Stack
	stack := &Stack{}

	// Menu aplikasi
	for {
		fmt.Println("\n=== MENU JUAL BARANG ===")
		fmt.Println("1. Tambah Barang (Push)")
		fmt.Println("2. Hapus Barang Terakhir (Pop)")
		fmt.Println("3. Lihat Barang Terakhir (Peek)")
		fmt.Println("4. Tampilkan Semua Barang (Stack)")
		fmt.Println("5. Simpan Barang ke Database")
		fmt.Println("6. Tampilkan Barang dari Database")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id, jumlah int
			var nama string
			var harga float64

			fmt.Print("Masukkan Nama Barang: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan Jumlah Barang: ")
			fmt.Scan(&jumlah)
			fmt.Print("Masukkan Harga Barang: ")
			fmt.Scan(&harga)

			stack.Push(Barang{
				ID:     id,
				Nama:   nama,
				Jumlah: jumlah,
				Harga:  harga,
			})

		case 2:
			stack.Pop()

		case 3:
			barang := stack.Peek()
			if barang.Nama != "" {
				fmt.Printf("Barang terakhir di Stack: ID=%d, Nama=%s, Jumlah=%d, Harga=%.2f\n",
					barang.ID, barang.Nama, barang.Jumlah, barang.Harga)
			}

		case 4:
			stack.Display()

		case 5:
			for _, barang := range stack.items {
				AddBarangToDatabase(barang)
			}

		case 6:
			barangList := GetAllBarangFromDatabase()
			fmt.Println("Daftar barang dari database:")
			for i, barang := range barangList {
				fmt.Printf("%d. ID=%d, Nama=%s, Jumlah=%d, Harga=%.2f\n",
					i+1, barang.ID, barang.Nama, barang.Jumlah, barang.Harga)
			}

		case 7:
			fmt.Println("Terima kasih! Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
