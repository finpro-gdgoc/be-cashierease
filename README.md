# Cashierease - Sistem Kasir (Point of Sale) API

`Cashierease` adalah backend API lengkap untuk aplikasi sistem kasir (Point of Sale). Dibangun menggunakan Go (Golang) dengan framework Gin untuk performa tinggi dan GORM sebagai ORM untuk interaksi database yang efisien.

Proyek ini menyediakan endpoint RESTful untuk mengelola produk, pengguna (admin/kasir), pesanan, kupon diskon, dan statistik penjualan.

## ✨ Fitur Utama

* **Manajemen Pengguna & Autentikasi**:
    * Registrasi Pengguna (Admin, Kasir).
    * Login dengan sistem **JWT (Access Token & Refresh Token)**.
    * **Role-Based Access Control (RBAC)**: Middleware untuk membatasi akses endpoint hanya untuk `admin`.
    * CRUD lengkap untuk data pengguna (hanya admin).
* **Manajemen Produk**:
    * CRUD lengkap (Create, Read, Update, Delete) untuk produk.
    * Upload gambar untuk produk.
    * Pencarian produk berdasarkan nama (query `?nama=...`).
    * Pengambilan produk berdasarkan `slug` untuk URL yang *user-friendly*.
* **Manajemen Pesanan (Order)**:
    * Pembuatan pesanan baru dengan keranjang belanja (multiple item).
    * Kalkulasi otomatis `total harga`, `pajak (10%)`, dan `diskon` dari kupon.
    * Pengambilan riwayat semua pesanan.
* **Manajemen Kupon**:
    * CRUD lengkap untuk kupon diskon.
    * Validasi kupon berdasarkan kode, tanggal berlaku (awal & akhir), dan metode pembayaran yang digunakan.
    * Endpoint untuk melihat kupon yang sedang aktif.
* **Statistik & Laporan**:
    * Endpoint statistik umum (total pendapatan, total pesanan, dll.).
    * Data pendapatan mingguan untuk 1, 2, dan 3 bulan terakhir.
    * Data jumlah pelanggan harian untuk 7 hari, 31 hari, dan 1 tahun terakhir.
    * Endpoint untuk melihat menu terpopuler (Top 3) berdasarkan bulan.
* **Manajemen Toko**:
    * Endpoint untuk mengambil dan memperbarui informasi dasar toko (misal: nama toko).

## 🛠️ Teknologi yang Digunakan

* **Bahasa**: [Go (Golang)](https://go.dev/)
* **Framework**: [Gin (v1.10.1)](https://gin-gonic.com/)
* **Database**: [PostgreSQL](https://www.postgresql.org/)
* **ORM**: [GORM (v1.30.3)](https://gorm.io/)
* **Autentikasi**: [JWT (golang-jwt/v5)](https://github.com/golang-jwt/jwt)
* **Environment**: [godotenv](https://github.com/joho/godotenv)
* **Password Hashing**: [bcrypt](https://golang.org/x/crypto/bcrypt)
* **Slug Generator**: [slug](https://github.com/gosimple/slug)

## 🚀 Cara Menjalankan Proyek

### 1. Prasyarat

* [Go](https://go.dev/doc/install) (versi 1.25.0 atau lebih baru direkomendasikan)
* [PostgreSQL](https://www.postgresql.org/download/) Database yang sedang berjalan.

### 2. Instalasi

1.  **Clone repository ini:**
    ```bash
    git clone [https://github.com/username/cashierease.git](https://github.com/username/cashierease.git)
    cd cashierease
    ```

2.  **Install dependensi Go:**
    ```bash
    go mod tidy
    ```

3.  **Setup Environment Variables:**
    Buat file `.env` di root proyek dengan menyalin dari `.env.example`.
    ```bash
    cp .env.example .env
    ```
    Kemudian, sesuaikan isi file `.env` dengan konfigurasi lokalmu:
    ```env
    # Sesuaikan dengan koneksi PostgreSQL kamu
    DATABASE_URL="host=localhost user=postgres password=password_db dbname=cashierease port=5432 sslmode=disable"
    
    # Ganti dengan kunci rahasia yang kuat
    JWT_SECRET_KEY="rahasiabanget"
    ```

### 3. Menjalankan Database

Pastikan layanan PostgreSQL kamu berjalan dan kamu telah membuat database baru dengan nama yang sesuai (contoh: `cashierease`) seperti yang didefinisikan di `DATABASE_URL`.

Aplikasi akan **secara otomatis menjalankan migrasi** database saat pertama kali dijalankan. Ini akan membuat tabel-tabel berikut:
* `users`
* `produks`
* `coupons`
* `tokos`
* `orders`
* `order_items`

### 4. Menjalankan Aplikasi

Jalankan server utama:

```bash
go run cmd/server/main.go