package main

import (
	"fmt"
	"strings"
	"time"
)

type Barang struct {
	ID         string
	Nama       string
	Kategori   string
	Harga      float64
	Stok       int
	StockAlert int
}

type Transaksi struct {
	IDTransaksi string
	IDBarang    string
	Jenis       string
	Jumlah      int
	Tanggal     string
}

var inventaris [100]Barang
var transaksi [100]Transaksi
var totalBarang int
var totalTransaksi int
var kategori []string
var totalKategori int

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	msgInfo     = "[INFO] "
	msgError    = "[ERROR] "
	msgSuccess  = "[SUCCESS] "
)

func displayItems() {
	if totalBarang == 0 {
		fmt.Println(colorYellow + msgInfo + "Inventaris kosong!" + colorReset)
		return
	}

	fmt.Println("\nDaftar Barang:")
	fmt.Println("┌──────┬────────────┬────────────┬──────────┬──────┬─────────────┐")
	fmt.Println("│  ID  │    Nama    │  Kategori  │   Harga  │ Stok │   Status    │")
	fmt.Println("├──────┼────────────┼────────────┼──────────┼──────┼─────────────┤")
	for i := 0; i < totalBarang; i++ {
		status := ""
		if inventaris[i].Stok <= inventaris[i].StockAlert {
			status = colorRed + "STOCK ALERT!" + colorReset
		} else {
			status = colorGreen + "STOK AMAN" + colorReset
		}

		fmt.Printf("│ %-4s │ %-10s │ %-10s │ %8.2f │ %4d │ %-11s │\n",
			inventaris[i].ID,
			inventaris[i].Nama,
			inventaris[i].Kategori,
			inventaris[i].Harga,
			inventaris[i].Stok,
			status)
	}
	fmt.Println("└──────┴────────────┴────────────┴──────────┴──────┴─────────────┘")
}

func searchByID(id string) int {
	for i := 0; i < totalBarang; i++ {
		if inventaris[i].ID == id {
			return i
		}
	}
	return -1
}

func searchByName(nama string) int {
	for i := 0; i < totalBarang; i++ {
		if strings.ToLower(inventaris[i].Nama) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

func generateID() string {
	return fmt.Sprintf("BRG%03d", totalBarang+1)
}

func updateItem() {
	if totalBarang == 0 {
		fmt.Println(colorYellow + msgInfo + "Inventaris kosong!" + colorReset)
		return
	}

	fmt.Println("\nDaftar Barang:")
	fmt.Println("┌────┬──────┬────────────┐")
	fmt.Println("│ No │  ID  │    Nama    │")
	fmt.Println("├────┼──────┼────────────┤")
	for i := 0; i < totalBarang; i++ {
		fmt.Printf("│ %2d │ %-4s │ %-10s │\n", i+1, inventaris[i].ID, inventaris[i].Nama)
	}
	fmt.Println("└────┴──────┴────────────┘")

	var idBarang string
	fmt.Print("\nMasukkan ID barang: ")
	fmt.Scanln(&idBarang)

	idx := searchByID(idBarang)
	if idx == -1 {
		fmt.Println(colorRed + msgError + "Barang tidak ditemukan!" + colorReset)
		return
	}

	fmt.Print("Masukkan nama baru: ")
	fmt.Scanln(&inventaris[idx].Nama)

	fmt.Println("\nList Kategori Barang:")
	fmt.Println("ID | Kategori")
	fmt.Println("-------------")
	for i := 0; i < totalKategori; i++ {
		fmt.Printf("%d. %s\n", i+1, kategori[i])
	}

	var idKategori int
	fmt.Print("\nMasukkan ID kategori baru: ")
	fmt.Scanln(&idKategori)

	if idKategori > 0 && idKategori <= totalKategori {
		inventaris[idx].Kategori = kategori[idKategori-1]
	} else {
		fmt.Println("ID kategori tidak valid!")
		return
	}

	fmt.Print("Masukkan harga baru: ")
	fmt.Scanln(&inventaris[idx].Harga)

	fmt.Print("Masukkan Stock Alert baru (batas minimum stok): ")
	fmt.Scanln(&inventaris[idx].StockAlert)

	fmt.Println(colorGreen + msgSuccess + "Barang berhasil diupdate!" + colorReset)
}

func manajemenBarang() {
	for {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      MANAJEMEN BARANG          ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Tampilkan Barang            ║")
		fmt.Println("║ 2. Tambah Barang               ║")
		fmt.Println("║ 3. Edit Barang                 ║")
		fmt.Println("║ 4. Hapus Barang                ║")
		fmt.Println("║ 5. Kembali                     ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu (1-5): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			displayItems()
		case 2:
			addItem()
		case 3:
			updateItem()
		case 4:
			deleteItem()
		case 5:
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}

func main() {
	// Initialize default categories
	kategori = []string{"Elektronik", "Makanan", "Minuman", "Pakaian", "Lainnya"}
	totalKategori = len(kategori)

	showWelcome()

	for {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║         MENU INVENTORI         ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Tampilkan Daftar Barang     ║")
		fmt.Println("║ 2. Manajemen Barang            ║")
		fmt.Println("║ 3. Catat Transaksi             ║")
		fmt.Println("║ 4. Manajemen Kategori          ║")
		fmt.Println("║ 5. Generate Data Sample        ║")
		fmt.Println("║ 6. Keluar                      ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu (1-6): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			displayItems()
		case 2:
			manajemenBarang()
		case 3:
			recordTransaction()
		case 4:
			manajemenKategori()
		case 5:
			generateSampleData()
		case 6:
			showGoodbye()
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}

func deleteItem() {
	if totalBarang == 0 {
		fmt.Println(colorYellow + msgInfo + "Inventaris kosong!" + colorReset)
		return
	}

	fmt.Println("\nDaftar Barang:")
	fmt.Println("┌────┬──────┬────────────┐")
	fmt.Println("│ No │  ID  │    Nama    │")
	fmt.Println("├────┼──────┼────────────┤")
	for i := 0; i < totalBarang; i++ {
		fmt.Printf("│ %2d │ %-4s │ %-10s │\n", i+1, inventaris[i].ID, inventaris[i].Nama)
	}
	fmt.Println("└────┴──────┴────────────┘")

	var id string
	fmt.Print("\nMasukkan ID barang yang akan dihapus: ")
	fmt.Scanln(&id)

	idx := searchByID(id)
	if idx == -1 {
		fmt.Println(colorRed + msgError + "Barang tidak ditemukan!" + colorReset)
		return
	}

	// Geser semua elemen setelah idx satu posisi ke kiri
	for i := idx; i < totalBarang-1; i++ {
		inventaris[i] = inventaris[i+1]
	}
	totalBarang--
	fmt.Println(colorGreen + msgSuccess + "Barang berhasil dihapus!" + colorReset)
}

func recordTransaction() {
	fmt.Println("\n╔��═══════════════════════════════╗")
	fmt.Println("║        CATAT TRANSAKSI         ║")
	fmt.Println("╠════════════════════════════════╣")

	if totalTransaksi >= 100 {
		fmt.Println(colorRed + msgError + "Daftar transaksi penuh!" + colorReset)
		return
	}

	displayItems()
	var newTransaksi Transaksi

	// Generate ID transaksi (timestamp)
	newTransaksi.IDTransaksi = fmt.Sprintf("TRX%d", time.Now().Unix())

	fmt.Print("Masukkan ID Barang: ")
	fmt.Scanln(&newTransaksi.IDBarang)

	idx := searchByID(newTransaksi.IDBarang)
	if idx == -1 {
		fmt.Println(colorRed + msgError + "Barang tidak ditemukan!" + colorReset)
		return
	}

	fmt.Print("Jenis Transaksi (masuk/keluar): ")
	fmt.Scanln(&newTransaksi.Jenis)

	fmt.Print("Jumlah barang: ")
	fmt.Scanln(&newTransaksi.Jumlah)

	// Validasi stok untuk transaksi keluar
	if strings.ToLower(newTransaksi.Jenis) == "keluar" {
		if inventaris[idx].Stok < newTransaksi.Jumlah {
			fmt.Println(colorRed + msgError + "Stok tidak mencukupi!" + colorReset)
			return
		}
		inventaris[idx].Stok -= newTransaksi.Jumlah
	} else if strings.ToLower(newTransaksi.Jenis) == "masuk" {
		inventaris[idx].Stok += newTransaksi.Jumlah
	} else {
		fmt.Println(colorRed + msgError + "Jenis transaksi tidak valid!" + colorReset)
		return
	}

	// Set tanggal transaksi
	newTransaksi.Tanggal = time.Now().Format("2006-01-02")

	// Simpan transaksi
	transaksi[totalTransaksi] = newTransaksi
	totalTransaksi++

	fmt.Println(colorGreen + msgSuccess + "Transaksi berhasil dicatat!" + colorReset)
}

// Fungsi sorting
func selectionSortByPrice(arr []Barang, desc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if desc {
				if arr[j].Harga > arr[minIdx].Harga {
					minIdx = j
				}
			} else {
				if arr[j].Harga < arr[minIdx].Harga {
					minIdx = j
				}
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func insertionSortByNama(arr []Barang, desc bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		if desc {
			for j >= 0 && arr[j].Nama < key.Nama {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && arr[j].Nama > key.Nama {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

func tampilkanKategori() {
	if totalKategori == 0 {
		fmt.Println("Belum ada kategori!")
		return
	}

	fmt.Println("\nDaftar Kategori:")
	fmt.Println("┌────┬────────────┐")
	fmt.Println("│ ID │  Kategori  │")
	fmt.Println("├────┼────────────┤")
	for i := 0; i < totalKategori; i++ {
		fmt.Printf("│ %2d │ %-10s │\n", i+1, kategori[i])
	}
	fmt.Println("└────┴────────────┘")
}

func tambahKategori() {
	var namaKategori string
	fmt.Print("Masukkan nama kategori baru: ")
	fmt.Scanln(&namaKategori)

	// Cek apakah kategori sudah ada
	for _, kat := range kategori {
		if strings.ToLower(kat) == strings.ToLower(namaKategori) {
			fmt.Println("Kategori sudah ada!")
			return
		}
	}

	kategori = append(kategori, namaKategori)
	totalKategori++
	fmt.Println("Kategori berhasil ditambahkan!")
}

func editKategori() {
	if totalKategori == 0 {
		fmt.Println("Belum ada kategori!")
		return
	}

	tampilkanKategori()

	var id int
	fmt.Print("Masukkan ID kategori yang akan diedit: ")
	fmt.Scanln(&id)

	if id < 1 || id > totalKategori {
		fmt.Println("ID kategori tidak valid!")
		return
	}

	var namaBaruKategori string
	fmt.Print("Masukkan nama baru untuk kategori: ")
	fmt.Scanln(&namaBaruKategori)

	// Update kategori di semua barang yang menggunakan kategori ini
	oldKategori := kategori[id-1]
	for i := 0; i < totalBarang; i++ {
		if inventaris[i].Kategori == oldKategori {
			inventaris[i].Kategori = namaBaruKategori
		}
	}

	kategori[id-1] = namaBaruKategori
	fmt.Println("Kategori berhasil diupdate!")
}

func hapusKategori() {
	if totalKategori == 0 {
		fmt.Println("Belum ada kategori!")
		return
	}

	tampilkanKategori()

	var id int
	fmt.Print("Masukkan ID kategori yang akan dihapus: ")
	fmt.Scanln(&id)

	if id < 1 || id > totalKategori {
		fmt.Println("ID kategori tidak valid!")
		return
	}

	// Cek apakah kategori sedang digunakan
	kategoriYangAkanDihapus := kategori[id-1]
	for i := 0; i < totalBarang; i++ {
		if inventaris[i].Kategori == kategoriYangAkanDihapus {
			fmt.Printf("Kategori '%s' tidak dapat dihapus karena sedang digunakan oleh barang '%s'!\n",
				kategoriYangAkanDihapus, inventaris[i].Nama)
			return
		}
	}

	// Hapus kategori dengan menggeser elemen array
	for i := id - 1; i < totalKategori-1; i++ {
		kategori[i] = kategori[i+1]
	}
	totalKategori--
	kategori = kategori[:totalKategori]
	fmt.Println("Kategori berhasil dihapus!")
}

func manajemenKategori() {
	for {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      MANAJEMEN KATEGORI        ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Tampilkan Kategori          ║")
		fmt.Println("║ 2. Tambah Kategori             ║")
		fmt.Println("║ 3. Edit Kategori               ║")
		fmt.Println("║ 4. Hapus Kategori              ║")
		fmt.Println("║ 5. Kembali                     ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu (1-5): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanKategori()
		case 2:
			tambahKategori()
		case 3:
			editKategori()
		case 4:
			hapusKategori()
		case 5:
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}

func showWelcome() {
	fmt.Println(colorCyan + "\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                                              ║")
	fmt.Println("║          SISTEM INVENTORI BARANG             ║")
	fmt.Println("║                                              ║")
	fmt.Println("╚══════════════════════════════════════════════╝" + colorReset)
}

func showGoodbye() {
	fmt.Println(colorYellow + "\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                                              ║")
	fmt.Println("║       Terima kasih telah menggunakan         ║")
	fmt.Println("║           SISTEM INVENTORI BARANG            ║")
	fmt.Println("║                                              ║")
	fmt.Println("╚══════════════════════════════════════════════╝" + colorReset)
}

func addItem() {
	if totalBarang >= 100 {
		fmt.Println(colorRed + msgError + "Inventaris penuh!" + colorReset)
		return
	}

	var barang Barang
	barang.ID = generateID()

	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scanln(&barang.Nama)

	// Tampilkan daftar kategori yang tersedia
	fmt.Println("\nList Kategori Barang:")
	fmt.Println("ID | Kategori")
	fmt.Println("-------------")
	for i := 0; i < totalKategori; i++ {
		fmt.Printf("%d. %s\n", i+1, kategori[i])
	}

	var idKategori int
	fmt.Print("\nMasukkan ID kategori: ")
	fmt.Scanln(&idKategori)

	if idKategori > 0 && idKategori <= totalKategori {
		barang.Kategori = kategori[idKategori-1]
	} else {
		fmt.Println("ID kategori tidak valid!")
		return
	}

	fmt.Print("Masukkan Harga: ")
	fmt.Scanln(&barang.Harga)
	fmt.Print("Masukkan Stok: ")
	fmt.Scanln(&barang.Stok)
	fmt.Print("Masukkan Stock Alert (batas minimum stok): ")
	fmt.Scanln(&barang.StockAlert)

	inventaris[totalBarang] = barang
	totalBarang++
	fmt.Println(colorGreen + msgSuccess + "Barang berhasil ditambahkan!" + colorReset)
}

func generateSampleData() {
	for {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      GENERATE DATA SAMPLE      ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Generate Kategori           ║")
		fmt.Println("║ 2. Generate Barang             ║")
		fmt.Println("║ 3. Generate Transaksi          ║")
		fmt.Println("║ 4. Kembali                     ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu (1-4): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			generateKategori()
		case 2:
			generateBarang()
		case 3:
			generateTransaksi()
		case 4:
			return
		default:
			fmt.Println(colorRed + msgError + "Menu tidak valid!" + colorReset)
		}
	}
}

func generateKategori() {
	var jumlah int
	fmt.Print("Masukkan jumlah kategori yang ingin digenerate: ")
	fmt.Scanln(&jumlah)

	sampleKategori := []string{"Elektronik", "Makanan", "Minuman", "Pakaian", "Alat Tulis",
		"Perabotan", "Kosmetik", "Mainan", "Olahraga", "Otomotif"}

	for i := 0; i < jumlah && i < len(sampleKategori); i++ {
		if totalKategori >= len(kategori) {
			fmt.Println(colorRed + msgError + "Kapasitas kategori penuh!" + colorReset)
			return
		}
		kategori = append(kategori, sampleKategori[i])
		totalKategori++
	}
	fmt.Println(colorGreen + msgSuccess + "Berhasil generate " + fmt.Sprint(jumlah) + " kategori!" + colorReset)
}

func generateBarang() {
	var jumlah int
	fmt.Print("Masukkan jumlah barang yang ingin digenerate: ")
	fmt.Scanln(&jumlah)

	sampleNama := []string{"Laptop", "Buku", "Pensil", "Tas", "Sepatu", "Baju", "Celana",
		"Mouse", "Keyboard", "Monitor"}

	for i := 0; i < jumlah; i++ {
		if totalBarang >= 100 {
			fmt.Println(colorRed + msgError + "Kapasitas barang penuh!" + colorReset)
			return
		}

		var barang Barang
		barang.ID = generateID()
		barang.Nama = sampleNama[i%len(sampleNama)]
		barang.Kategori = kategori[i%totalKategori]
		barang.Harga = float64((i + 1) * 10000)
		barang.Stok = (i + 1) * 5
		barang.StockAlert = 5

		inventaris[totalBarang] = barang
		totalBarang++
	}
	fmt.Println(colorGreen + msgSuccess + "Berhasil generate " + fmt.Sprint(jumlah) + " barang!" + colorReset)
}

func generateTransaksi() {
	if totalBarang == 0 {
		fmt.Println(colorYellow + msgInfo + "Tidak ada barang dalam inventaris!" + colorReset)
		return
	}

	var jumlah int
	fmt.Print("Masukkan jumlah transaksi yang ingin digenerate: ")
	fmt.Scanln(&jumlah)

	var jenisTransaksi string
	fmt.Print("Generate transaksi (masuk/keluar): ")
	fmt.Scanln(&jenisTransaksi)

	jenisTransaksi = strings.ToLower(jenisTransaksi)
	if jenisTransaksi != "masuk" && jenisTransaksi != "keluar" {
		fmt.Println(colorRed + msgError + "Jenis transaksi tidak valid!" + colorReset)
		return
	}

	fmt.Printf("Anda akan men-generate %d transaksi %s. Lanjutkan? (y/n): ", jumlah, jenisTransaksi)
	var konfirmasi string
	fmt.Scanln(&konfirmasi)

	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println(colorYellow + msgInfo + "Generate transaksi dibatalkan!" + colorReset)
		return
	}

	successCount := 0
	for i := 0; i < jumlah; i++ {
		if totalTransaksi >= 100 {
			fmt.Println(colorRed + msgError + "Kapasitas transaksi penuh!" + colorReset)
			break
		}

		var newTransaksi Transaksi
		newTransaksi.IDTransaksi = fmt.Sprintf("TRX%d", time.Now().UnixNano())

		// Pilih barang secara random
		randomIdx := i % totalBarang
		newTransaksi.IDBarang = inventaris[randomIdx].ID
		newTransaksi.Jenis = jenisTransaksi
		newTransaksi.Jumlah = (i % 5) + 1 // Random jumlah 1-5
		newTransaksi.Tanggal = time.Now().Format("2006-01-02")

		// Update stok
		if jenisTransaksi == "keluar" {
			if inventaris[randomIdx].Stok < newTransaksi.Jumlah {
				continue // Skip jika stok tidak cukup
			}
			inventaris[randomIdx].Stok -= newTransaksi.Jumlah
		} else {
			inventaris[randomIdx].Stok += newTransaksi.Jumlah
		}

		transaksi[totalTransaksi] = newTransaksi
		totalTransaksi++
		successCount++
	}
	fmt.Printf(colorGreen+msgSuccess+"Berhasil generate %d dari %d transaksi %s!"+colorReset+"\n",
		successCount, jumlah, jenisTransaksi)
}
