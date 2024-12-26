# Aplikasi Inventori Barang (CLI) - TUBES

Selamat datang di repositori **Aplikasi Inventori Barang**! Proyek ini merupakan tugas besar (TUBES) untuk membantu pengelolaan inventori barang menggunakan antarmuka Command Line Interface (CLI).

## Anggota Kelompok

1. **Muhammad Rafiful Hana** (2311102227)

---

### Fitur Utama

- **Tampilkan Barang**: Melihat daftar barang yang tersedia dalam inventaris, lengkap dengan status stok.
- **Tambah Barang**: Menambahkan data barang baru ke inventaris dengan detail seperti nama, kategori, harga, dan stok.
- **Edit Barang**: Memperbarui informasi barang yang sudah ada, termasuk kategori, harga, dan stok minimal.
- **Hapus Barang**: Menghapus barang dari inventaris.
- **Manajemen Kategori**: Menambahkan, mengedit, atau menghapus kategori barang.
- **Catat Transaksi**: Mencatat transaksi masuk dan keluar barang.
- **Generate Data Sample**: Membuat data contoh untuk kategori, barang, dan transaksi.

### Penjelasan Source Code

Source code aplikasi ini dibangun menggunakan bahasa pemrograman Go (Golang) dengan struktur sebagai berikut:

- **Struktur Data**:
  - `Barang`: Struktur untuk menyimpan informasi barang seperti ID, nama, kategori, harga, stok, dan batas minimum stok.
  - `Transaksi`: Struktur untuk mencatat transaksi barang (masuk atau keluar), termasuk ID transaksi, ID barang, jenis transaksi, jumlah, dan tanggal.

- **Fungsi Utama**:
  1. **displayItems()**: Menampilkan daftar barang yang tersedia dalam inventaris dengan status stok.
  2. **addItem()**: Menambahkan barang baru ke inventaris.
  3. **updateItem()**: Memperbarui data barang, termasuk nama, kategori, harga, dan batas minimum stok.
  4. **deleteItem()**: Menghapus barang berdasarkan ID.
  5. **recordTransaction()**: Mencatat transaksi masuk atau keluar barang dan memperbarui stok barang.
  6. **manajemenBarang()**: Menu utama untuk mengelola barang.
  7. **manajemenKategori()**: Menu untuk mengelola kategori barang, termasuk menambah, mengedit, dan menghapus kategori.
  8. **generateSampleData()**: Membuat data contoh untuk kategori, barang, dan transaksi.

- **Visualisasi dan Antarmuka**:
  - Menu ditampilkan menggunakan elemen visual sederhana seperti garis dan warna untuk memberikan pengalaman pengguna yang lebih menarik.
  - Pesan kesalahan dan sukses diwarnai dengan kode ANSI (misalnya, merah untuk kesalahan dan hijau untuk sukses).

- **Logika Tambahan**:
  - Pencarian barang berdasarkan ID atau nama.
  - Sorting barang berdasarkan harga atau nama menggunakan algoritma selection sort dan insertion sort.
  - Validasi data untuk memastikan input yang diterima sesuai.

Aplikasi ini menggunakan array sebagai tempat penyimpanan data, sehingga cocok untuk skala kecil tanpa memerlukan database eksternal.

---

