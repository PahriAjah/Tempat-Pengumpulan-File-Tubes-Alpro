package main

import (
	"fmt"
	"strings"
)

const MAKS = 100

type BahanMakanan struct {
	Nama      string
	Jumlah    int
	Tanggal   string
	Digunakan bool
}

var stok [MAKS]BahanMakanan
var jumlahData int

func bacaString(pesan string) string {
	var input string
	fmt.Print(pesan)
	fmt.Scanln(&input)
	return input
}

func bacaInteger(pesan string) int {
	var input int
	fmt.Print(pesan)
	fmt.Scanln(&input)
	return input
}

func tambahBahan() {
	if jumlahData >= MAKS {
		fmt.Println("🚫 Stok penuh!")
		return
	}
	fmt.Println("\n➕ Tambah Bahan Baru")
	fmt.Println("----------------------")
	stok[jumlahData].Nama = bacaString("📌 Nama bahan: ")
	stok[jumlahData].Jumlah = bacaInteger("📦 Jumlah: ")
	stok[jumlahData].Tanggal = bacaString("📅 Tanggal kedaluwarsa (YYYY-MM-DD): ")
	stok[jumlahData].Digunakan = false
	jumlahData++
	fmt.Println("✅ Bahan berhasil ditambahkan.")
}

func ubahBahan() {
	fmt.Println("\n✏️ Ubah Data Bahan")
	fmt.Println("----------------------")
	nama := bacaString("🔍 Masukkan nama bahan yang ingin diubah: ")
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("❌ Bahan tidak ditemukan.")
		return
	}
	stok[idx].Nama = bacaString("📌 Nama baru: ")
	stok[idx].Jumlah = bacaInteger("📦 Jumlah baru: ")
	stok[idx].Tanggal = bacaString("📅 Tanggal kedaluwarsa baru (YYYY-MM-DD): ")
	fmt.Println("✅ Data berhasil diubah.")
}

func hapusBahan() {
	fmt.Println("\n🗑️ Hapus Bahan")
	fmt.Println("----------------------")
	nama := bacaString("🔍 Masukkan nama bahan yang ingin dihapus: ")
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("❌ Bahan tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahData-1; i++ {
		stok[i] = stok[i+1]
	}
	jumlahData--
	fmt.Println("✅ Data berhasil dihapus.")
}

func sequentialSearch(nama string) int {
	nama = strings.ToLower(nama)
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(stok[i].Nama) == nama {
			return i
		}
	}
	return -1
}

func binarySearch(nama string) int {
	urutkanNama()
	kiri, kanan := 0, jumlahData-1
	nama = strings.ToLower(nama)
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(stok[tengah].Nama)
		if nama == namaTengah {
			return tengah
		} else if nama < namaTengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func urutkanNama() {
	for i := 0; i < jumlahData-1; i++ {
		for j := 0; j < jumlahData-i-1; j++ {
			if strings.ToLower(stok[j].Nama) > strings.ToLower(stok[j+1].Nama) {
				stok[j], stok[j+1] = stok[j+1], stok[j]
			}
		}
	}
}

func cariBahan() {
	fmt.Println("\n🔍 Pencarian Bahan")
	fmt.Println("----------------------")
	fmt.Println("1️⃣  Sequential Search")
	fmt.Println("2️⃣  Binary Search")
	pilih := bacaInteger("📌 Pilihan: ")

	nama := bacaString("🔍 Nama bahan yang dicari: ")
	var idx int
	if pilih == 1 {
		idx = sequentialSearch(nama)
	} else if pilih == 2 {
		idx = binarySearch(nama)
	} else {
		fmt.Println("❌ Metode tidak valid.")
		return
	}

	if idx == -1 {
		fmt.Println("❌ Bahan tidak ditemukan.")
	} else {
		fmt.Printf("✅ Ditemukan: %s (%d) - Kedaluwarsa: %s\n", stok[idx].Nama, stok[idx].Jumlah, stok[idx].Tanggal)
	}
}

func peringatanKedaluwarsa() {
	fmt.Println("\n⚠️  Bahan mendekati tanggal kedaluwarsa:")
	ada := false
	for i := 0; i < jumlahData; i++ {
		if stok[i].Tanggal != "" && strings.HasPrefix(stok[i].Tanggal, "2025-06") {
			fmt.Printf("⏰ %s (%d) - Kedaluwarsa: %s\n", stok[i].Nama, stok[i].Jumlah, stok[i].Tanggal)
			ada = true
		}
	}
	if !ada {
		fmt.Println("✅ Tidak ada bahan yang mendekati kedaluwarsa.")
	}
}

func urutkanJumlah() {
	for i := 0; i < jumlahData-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahData; j++ {
			if stok[j].Jumlah < stok[idxMin].Jumlah {
				idxMin = j
			}
		}
		stok[i], stok[idxMin] = stok[idxMin], stok[i]
	}
	fmt.Println("✅ Diurutkan berdasarkan jumlah (🔢 Selection Sort).")
}

func urutkanTanggal() {
	for i := 1; i < jumlahData; i++ {
		key := stok[i]
		j := i - 1
		for j >= 0 && stok[j].Tanggal > key.Tanggal {
			stok[j+1] = stok[j]
			j--
		}
		stok[j+1] = key
	}
	fmt.Println("✅ Diurutkan berdasarkan tanggal kedaluwarsa (📅 Insertion Sort).")
}

func laporanStok() {
	total := 0
	digunakan := 0
	for i := 0; i < jumlahData; i++ {
		total += stok[i].Jumlah
		if stok[i].Digunakan {
			digunakan += stok[i].Jumlah
		}
	}
	fmt.Println("\n📊 Laporan Stok Bahan")
	fmt.Println("----------------------")
	fmt.Printf("📦 Total bahan tersedia: %d\n", total)
	fmt.Printf("✅ Total bahan digunakan: %d\n", digunakan)
	fmt.Printf("🗂️  Jumlah entri: %d\n", jumlahData)
}

func tampilkanData() {
	fmt.Println("\n📋 Daftar Bahan Makanan")
	fmt.Println("----------------------")
	for i := 0; i < jumlahData; i++ {
		status := "❌ Belum Digunakan"
		if stok[i].Digunakan {
			status = "✅ Sudah Digunakan"
		}
		fmt.Printf("%d. %s (%d) - Kedaluwarsa: %s [%s]\n", i+1, stok[i].Nama, stok[i].Jumlah, stok[i].Tanggal, status)
	}
}

func tandaiDigunakan() {
	nama := bacaString("🔖 Nama bahan yang ingin ditandai digunakan: ")
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("❌ Bahan tidak ditemukan.")
		return
	}
	if stok[idx].Jumlah == 0 {
		fmt.Println("⚠️ Tidak ada stok tersisa untuk bahan ini.")
		return
	}

	jumlah := bacaInteger("🔢 Jumlah yang ingin ditandai digunakan: ")
	if jumlah <= 0 || jumlah > stok[idx].Jumlah {
		fmt.Println("❌ Jumlah tidak valid atau melebihi stok.")
		return
	}

	stok[idx].Jumlah -= jumlah

	// Tandai sebagai digunakan jika seluruh stok telah digunakan
	if stok[idx].Jumlah == 0 {
		stok[idx].Digunakan = true
	}

	fmt.Printf("✅ %d dari %s ditandai sebagai digunakan.\n", jumlah, stok[idx].Nama)
}

func menu() {
	for {
		fmt.Println("\n==============================")
		fmt.Println("🍽️  MANAJEMEN STOK BAHAN MAKANAN 🍽️")
		fmt.Println("==============================")
		fmt.Println("1️⃣  ➕ Tambah Bahan")
		fmt.Println("2️⃣  ✏️  Ubah Bahan")
		fmt.Println("3️⃣  🗑️  Hapus Bahan")
		fmt.Println("4️⃣  🔍 Cari Bahan")
		fmt.Println("5️⃣  📋 Tampilkan Semua Data")
		fmt.Println("6️⃣  🔢 Urutkan Jumlah")
		fmt.Println("7️⃣  📅 Urutkan Tanggal Kedaluwarsa")
		fmt.Println("8️⃣  🔖 Tandai Digunakan")
		fmt.Println("9️⃣  ⚠️  Peringatan Kedaluwarsa")
		fmt.Println("🔟  📊 Laporan Stok")
		fmt.Println("0️⃣  ❌ Keluar")
		fmt.Println("------------------------------")

		pilih := bacaInteger("📌 Pilih menu: ")
		switch pilih {
		case 1:
			tambahBahan()
		case 2:
			ubahBahan()
		case 3:
			hapusBahan()
		case 4:
			cariBahan()
		case 5:
			tampilkanData()
		case 6:
			urutkanJumlah()
			tampilkanData()
		case 7:
			urutkanTanggal()
			tampilkanData()
		case 8:
			tandaiDigunakan()
		case 9:
			peringatanKedaluwarsa()
		case 10:
			laporanStok()
		case 0:
			fmt.Println("👋 Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("❌ Menu tidak valid.")
		}
	}
}

func main() {
	menu()
}
