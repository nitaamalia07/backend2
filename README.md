📋 
## Deskripsi Proyek
Medpoint API adalah RESTful API yang dikembangkan untuk mengelola data dokter dan jadwal praktik mereka. Proyek ini mencakup:

-CRUD (Create, Read, Update, Delete) untuk data dokter dan jadwal.

-Dokumentasi API menggunakan Swagger.

-Implementasi autentikasi dengan JWT.

-Penerapan Role-Based Access Control (RBAC) menggunakan Row Level Security (RLS) di Supabase.


🔐 
# Role-Based Access Control (RBAC)
Setiap endpoint (kecuali autentikasi) memerlukan JWT dan akses dibatasi berdasarkan peran pengguna:

-Super Admin: Mengelola semua data (dokter, jadwal, pembayaran, pengguna).

-Admin: Mengelola data dokter dan jadwal.

-Dokter: Mengakses data pribadi dan jadwal sendiri.

-User/Pasien: Melihat data dokter dan jadwal, serta membuat reservasi.


🛠️ 
# Teknologi yang Digunakan
-Backend: Go (Golang)

-Database: PostgreSQL melalui Supabase

-Autentikasi: Supabase Auth dengan JWT

-Dokumentasi API: Swagger/OpenAPI 3.0


📑 
# Dokumentasi API
Dokumentasi lengkap tersedia melalui Swagger. 


🚀 
# Menjalankan Proyek
Clone repository:
git clone https://github.com/nitaamalia07/backend2.git
cd backend2
Jalankan server:
raiden run
