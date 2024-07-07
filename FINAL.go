package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struktur data untuk menyimpan informasi akun
type Akun struct {
	Username string
	Password string
	Status   string // "Pending", "Approved", "Rejected"
}

// Struktur data untuk pesan
type Pesan struct {
	Pengirim string
	Penerima string
	Isi      string
}

// Struktur data untuk grup chatting
type GrupChat struct {
	NamaGrup      string
	Anggota       [100]string
	JumlahAnggota int
	Pesan         [100]Pesan
	JumlahPesan   int
}

// Array statis untuk menyimpan data akun
var daftarAkun [100]Akun
var jumlahAkun int

// Array statis untuk menyimpan pesan pribadi
var pesanPribadi [100]Pesan
var jumlahPesanPribadi int

// Array statis untuk menyimpan grup chatting
var daftarGrupChat [10]GrupChat
var jumlahGrup int

// Fungsi untuk memeriksa apakah pengguna sudah terdaftar dan disetujui
func isUserApproved(username string) bool {
	for i := 0; i < jumlahAkun; i++ {
		if daftarAkun[i].Username == username && daftarAkun[i].Status == "Approved" {
			return true
		}
	}
	return false
}

// Fungsi untuk registrasi akun baru
func registrasiAkun() {
	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)
	if jumlahAkun < len(daftarAkun) { // Memeriksa apakah jumlah akun saat ini kurang dari kapasitas array
		daftarAkun[jumlahAkun] = Akun{Username: username, Password: password, Status: "Pending"}
		jumlahAkun++
		fmt.Println("Registrasi berhasil. Akun menunggu persetujuan.")
	} else {
		fmt.Println("Kapasitas penyimpanan akun penuh.") // Pesan ketika array sudah penuh
	}
}

// Fungsi untuk menyetujui atau menolak registrasi akun oleh admin
func tinjauAkun() {
	var username, aksi string
	var found bool // Menyimpan status apakah akun ditemukan
	fmt.Print("Masukkan username yang akan ditinjau: ")
	fmt.Scanln(&username)
	fmt.Print("Setujui atau tolak (setuju/tolak): ")
	fmt.Scanln(&aksi)
	for i := 0; i < jumlahAkun && !found; i++ { // Looping sampai akun ditemukan atau semua akun telah dicek
		if daftarAkun[i].Username == username {
			found = true
			if aksi == "setuju" {
				daftarAkun[i].Status = "Approved"
				fmt.Println("Akun", username, "disetujui.")
			} else if aksi == "tolak" {
				daftarAkun[i].Status = "Rejected"
				fmt.Println("Akun", username, "ditolak.")
			}
		}
	}
	if !found {
		fmt.Println("Akun tidak ditemukan.")
	}
}

// Fungsi untuk mencetak daftar akun
func cetakDaftarAkun() {
	for i := 0; i < jumlahAkun; i++ {
		fmt.Printf("Username: %s, Status: %s\n", daftarAkun[i].Username, daftarAkun[i].Status)
	}
}

// Fungsi untuk mengirim pesan pribadi
func kirimPesan(username string) {
	var penerima, isi string
	fmt.Print("Masukkan username penerima: ")
	fmt.Scanln(&penerima)
	fmt.Print("Masukkan isi pesan: ")
	reader := bufio.NewReader(os.Stdin)
	isi, _ = reader.ReadString('\n')
	isi = strings.TrimSpace(isi)

	if !isUserApproved(penerima) {
		fmt.Println("Penerima belum terdaftar atau belum disetujui.")
		return
	}

	if jumlahPesanPribadi < len(pesanPribadi) { // Memeriksa apakah jumlah pesan pribadi saat ini kurang dari kapasitas array
		pesanPribadi[jumlahPesanPribadi] = Pesan{Pengirim: username, Penerima: penerima, Isi: isi}
		jumlahPesanPribadi++
		fmt.Println("Pesan terkirim.")
	} else {
		fmt.Println("Kapasitas penyimpanan pesan penuh.") // Pesan ketika array sudah penuh
	}
}

// Fungsi untuk membuat grup chatting baru
func buatGrupChat(username string) {
	var namaGrup string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)

	if jumlahGrup < len(daftarGrupChat) { // Memeriksa apakah jumlah grup saat ini kurang dari kapasitas array
		daftarGrupChat[jumlahGrup] = GrupChat{NamaGrup: namaGrup, Anggota: [100]string{username}, JumlahAnggota: 1}
		jumlahGrup++
		fmt.Println("Grup", namaGrup, "berhasil dibuat.")
	} else {
		fmt.Println("Kapasitas penyimpanan grup penuh.") // Pesan ketika array sudah penuh
	}
}

// Fungsi untuk menambahkan anggota ke dalam grup chatting
func tambahAnggotaKeGrup() {
	var namaGrup, anggotaBaru string
	var found bool // Menyimpan status apakah grup ditemukan
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)
	fmt.Print("Masukkan username anggota baru: ")
	fmt.Scanln(&anggotaBaru)

	if !isUserApproved(anggotaBaru) {
		fmt.Println("Anggota baru belum terdaftar atau belum disetujui.")
		return
	}

	for i := 0; i < jumlahGrup && !found; i++ { // Looping sampai grup ditemukan atau semua grup telah dicek
		if daftarGrupChat[i].NamaGrup == namaGrup {
			found = true
			if daftarGrupChat[i].JumlahAnggota < len(daftarGrupChat[i].Anggota) { // Memeriksa apakah jumlah anggota kurang dari kapasitas array
				daftarGrupChat[i].Anggota[daftarGrupChat[i].JumlahAnggota] = anggotaBaru
				daftarGrupChat[i].JumlahAnggota++
				fmt.Println("Anggota", anggotaBaru, "berhasil ditambahkan ke grup", namaGrup)
			} else {
				fmt.Println("Kapasitas anggota grup penuh.") // Pesan ketika array sudah penuh
			}
		}
	}
	if !found {
		fmt.Println("Grup tidak ditemukan.")
	}
}

// Fungsi untuk mengirim pesan ke grup chatting
func kirimPesanGrup(username string) {
	var namaGrup, isi string
	var found bool // Menyimpan status apakah grup ditemukan
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)
	fmt.Print("Masukkan isi pesan: ")
	reader := bufio.NewReader(os.Stdin)
	isi, _ = reader.ReadString('\n')
	isi = strings.TrimSpace(isi)

	for i := 0; i < jumlahGrup && !found; i++ { // Looping sampai grup ditemukan atau semua grup telah dicek
		if daftarGrupChat[i].NamaGrup == namaGrup {
			found = true
			if daftarGrupChat[i].JumlahPesan < len(daftarGrupChat[i].Pesan) { // Memeriksa apakah jumlah pesan kurang dari kapasitas array
				daftarGrupChat[i].Pesan[daftarGrupChat[i].JumlahPesan] = Pesan{Pengirim: username, Penerima: namaGrup, Isi: isi}
				daftarGrupChat[i].JumlahPesan++
				fmt.Println("Pesan terkirim ke grup", namaGrup)
			} else {
				fmt.Println("Kapasitas penyimpanan pesan grup penuh.") // Pesan ketika array sudah penuh
			}
		}
	}
	if !found {
		fmt.Println("Grup tidak ditemukan.")
	}
}

// Fungsi untuk melihat anggota dalam grup
func lihatAnggotaGrup() {
	var namaGrup string
	var found bool // Menyimpan status apakah grup ditemukan
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)
	for i := 0; i < jumlahGrup && !found; i++ { // Looping sampai grup ditemukan atau semua grup telah dicek
		if daftarGrupChat[i].NamaGrup == namaGrup {
			found = true
			fmt.Println("Anggota dari", namaGrup, ":")
			for j := 0; j < daftarGrupChat[i].JumlahAnggota; j++ { // Looping untuk mencetak anggota grup
				fmt.Println(daftarGrupChat[i].Anggota[j])
			}
		}
	}
	if !found {
		fmt.Println("Grup tidak ditemukan.")
	}
}

// Fungsi untuk menampilkan daftar pesan berdasarkan akun tertentu
func tampilkanPesan(username string) {
	fmt.Println("Pesan untuk", username, ":")

	// Tampilkan pesan pribadi
	for i := 0; i < jumlahPesanPribadi; i++ {
		if pesanPribadi[i].Pengirim == username || pesanPribadi[i].Penerima == username {
			fmt.Printf("Dari: %s, Kepada: %s, Isi: %s\n", pesanPribadi[i].Pengirim, pesanPribadi[i].Penerima, pesanPribadi[i].Isi)
		}
	}

	// Tampilkan pesan grup
	for i := 0; i < jumlahGrup; i++ {
		for j := 0; j < daftarGrupChat[i].JumlahPesan; j++ {
			if daftarGrupChat[i].Pesan[j].Pengirim == username {
				fmt.Printf("[Grup: %s] Dari: %s, Isi: %s\n", daftarGrupChat[i].NamaGrup, daftarGrupChat[i].Pesan[j].Pengirim, daftarGrupChat[i].Pesan[j].Isi)
			}
		}
	}
}

// Fungsi untuk memeriksa autentikasi admin
func authAdminCheck(username, password string) bool {
	return (username == "hauzanrafi" && password == "hauzanrafi123") || (username == "taufiq" && password == "taufiq123")
}

// Fungsi untuk mode admin
func adminPOV() {
	// Meminta username dan password admin
	fmt.Println("\nAnda masuk sebagai admin.")
	fmt.Print("Masukkan username admin: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Masukkan password admin: ")
	var password string
	fmt.Scanln(&password)

	// Memeriksa autentikasi admin
	if authAdminCheck(username, password) {
		running := true
		for running {
			fmt.Println("\nPilih opsi:")
			fmt.Println("1. Tinjau akun")
			fmt.Println("2. Cetak daftar akun")
			fmt.Println("3. Cetak daftar akun terurut")
			fmt.Println("4. Cari akun")
			fmt.Println("5. Tukar Mode")
			fmt.Println("0. Keluar")

			var pilihan string
			fmt.Print("Masukkan pilihan: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case "1":
				tinjauAkun()
			case "2":
				cetakDaftarAkun()
			case "3":
				cetakDaftarAkunTerurut()
			case "4":
				cariAkun()
			case "5":
				fmt.Println("Tukar mode ke pengguna.")
				userPOV() // Memanggil fungsi userPOV secara rekursif
			case "0":
				fmt.Println("Keluar dari mode admin.")
				running = false
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else {
		fmt.Println("Autentikasi gagal. Anda bukan admin.")
	}
}

// Fungsi untuk mode pengguna
func userPOV() {
	running := true
	for running {
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Login")
		fmt.Println("2. Registrasi")
		fmt.Println("0. Kembali")

		var pilihan string
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			loginUser()
		case "2":
			registrasiAkun()
		case "0":
			running = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func loginUser() {
	fmt.Println("\nLogin Pengguna")
	fmt.Print("Masukkan username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	var password string
	fmt.Scanln(&password)

	if !isUserApproved(username) {
		fmt.Println("Autentikasi gagal. Anda belum terdaftar atau belum disetujui.")
		return
	}

	for i := 0; i < jumlahAkun; i++ {
		if daftarAkun[i].Username == username && daftarAkun[i].Password == password {
			userMenu(username)
			return
		}
	}
	fmt.Println("Autentikasi gagal. Username atau password salah.")
}

func userMenu(username string) {
	running := true
	for running {
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Kirim pesan pribadi")
		fmt.Println("2. Buat grup chat")
		fmt.Println("3. Tambah anggota ke grup")
		fmt.Println("4. Hapus anggota dari grup")
		fmt.Println("5. Edit anggota grup")
		fmt.Println("6. Kirim pesan ke grup")
		fmt.Println("7. Lihat anggota grup")
		fmt.Println("8. Tampilkan pesan")
		fmt.Println("9. Tukar Mode")
		fmt.Println("0. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			kirimPesan(username)
		case "2":
			buatGrupChat(username)
		case "3":
			tambahAnggotaKeGrup()
		case "4":
			hapusAnggotaDariGrup()
		case "5":
			editAnggotaGrup()
		case "6":
			kirimPesanGrup(username)
		case "7":
			lihatAnggotaGrup()
		case "8":
			tampilkanPesan(username)
		case "9":
			fmt.Println("Tukar mode ke admin.")
			adminPOV() // Memanggil fungsi adminPOV secara rekursif
			return     // Keluar dari fungsi userPOV setelah selesai tukar mode
		case "0":
			fmt.Println("Keluar dari mode pengguna.")
			running = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk meminta peran pengguna
func chooseRole() string {
	fmt.Println("\nSelamat datang!")
	fmt.Println("Apakah Anda masuk sebagai admin atau pengguna?")
	fmt.Println("1. Admin")
	fmt.Println("2. Pengguna")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ")

	var peran string
	fmt.Scanln(&peran)

	return peran
}

// Implementasi Binary Search
func binarySearchUser(username string) int {
	low, high := 0, jumlahAkun-1
	for low <= high {
		mid := (low + high) / 2
		if daftarAkun[mid].Username == username {
			return mid
		} else if daftarAkun[mid].Username < username {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Implementasi Selection Sort untuk mengurutkan akun berdasarkan username
func selectionSortAkun(ascending bool) {
	for i := 0; i < jumlahAkun-1; i++ {
		idx := i
		for j := i + 1; j < jumlahAkun; j++ {
			if (ascending && daftarAkun[j].Username < daftarAkun[idx].Username) ||
				(!ascending && daftarAkun[j].Username > daftarAkun[idx].Username) {
				idx = j
			}
		}
		daftarAkun[i], daftarAkun[idx] = daftarAkun[idx], daftarAkun[i]
	}
}

// Implementasi Insertion Sort untuk mengurutkan akun berdasarkan status
func insertionSortAkun(ascending bool) {
	for i := 1; i < jumlahAkun; i++ {
		key := daftarAkun[i]
		j := i - 1
		for j >= 0 && ((ascending && daftarAkun[j].Status > key.Status) ||
			(!ascending && daftarAkun[j].Status < key.Status)) {
			daftarAkun[j+1] = daftarAkun[j]
			j--
		}
		daftarAkun[j+1] = key
	}
}

// Fungsi untuk mencetak daftar akun yang sudah terurut
func cetakDaftarAkunTerurut() {
	fmt.Println("\nPilih metode pengurutan:")
	fmt.Println("1. Berdasarkan Username (Ascending)")
	fmt.Println("2. Berdasarkan Username (Descending)")
	fmt.Println("3. Berdasarkan Status (Ascending)")
	fmt.Println("4. Berdasarkan Status (Descending)")
	fmt.Print("Masukkan pilihan: ")

	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		selectionSortAkun(true)
	case "2":
		selectionSortAkun(false)
	case "3":
		insertionSortAkun(true)
	case "4":
		insertionSortAkun(false)
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}
	cetakDaftarAkun()
}

// Fungsi untuk mencari akun
func cariAkun() {

	fmt.Print("Masukkan username yang dicari: ")
	var username string
	fmt.Scanln(&username)

	var index int
	selectionSortAkun(true)
	index = binarySearchUser(username)

	if index != -1 {
		fmt.Printf("Akun ditemukan. Username: %s, Status: %s\n", daftarAkun[index].Username, daftarAkun[index].Status)
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

// Fungsi untuk menghapus anggota dari grup chatting
func hapusAnggotaDariGrup() bool {
	var namaGrup, anggotaDihapus string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)
	fmt.Print("Masukkan username anggota yang akan dihapus: ")
	fmt.Scanln(&anggotaDihapus)

	for i := 0; i < jumlahGrup; i++ {
		if daftarGrupChat[i].NamaGrup == namaGrup {
			for j := 0; j < daftarGrupChat[i].JumlahAnggota; j++ {
				if daftarGrupChat[i].Anggota[j] == anggotaDihapus {
					// Geser anggota ke kiri untuk menutup celah
					for k := j; k < daftarGrupChat[i].JumlahAnggota-1; k++ {
						daftarGrupChat[i].Anggota[k] = daftarGrupChat[i].Anggota[k+1]
					}
					// Reset anggota terakhir untuk menghapusnya
					daftarGrupChat[i].Anggota[daftarGrupChat[i].JumlahAnggota-1] = ""
					// Kurangi jumlah anggota
					daftarGrupChat[i].JumlahAnggota--
					fmt.Printf("Anggota %s berhasil dihapus dari grup %s\n", anggotaDihapus, namaGrup)
					return true
				}
			}
			fmt.Println("Anggota tidak ditemukan dalam grup.")
			return false
		}
	}
	fmt.Println("Grup tidak ditemukan.")
	return false
}

// Fungsi untuk mengedit anggota dalam grup chatting
func editAnggotaGrup() bool {
	var namaGrup, anggotaLama, anggotaBaru string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scanln(&namaGrup)
	fmt.Print("Masukkan username anggota yang ingin diedit: ")
	fmt.Scanln(&anggotaLama)
	fmt.Print("Masukkan username anggota baru: ")
	fmt.Scanln(&anggotaBaru)

	for i := 0; i < jumlahGrup; i++ {
		if daftarGrupChat[i].NamaGrup == namaGrup {
			for j := 0; j < daftarGrupChat[i].JumlahAnggota; j++ {
				if daftarGrupChat[i].Anggota[j] == anggotaLama {
					daftarGrupChat[i].Anggota[j] = anggotaBaru
					fmt.Printf("Anggota %s berhasil diubah menjadi %s dalam grup %s\n", anggotaLama, anggotaBaru, namaGrup)
					return true
				}
			}
			fmt.Println("Anggota tidak ditemukan dalam grup.")
			return false
		}
	}
	fmt.Println("Grup tidak ditemukan.")
	return false
}

// Fungsi utama
func main() {
	exit := false

	for !exit {
		peran := chooseRole()

		if peran == "1" {
			adminPOV()
		} else if peran == "2" {
			userPOV()
		} else if peran == "0" {
			fmt.Println("Keluar dari aplikasi.")
			exit = true // Mengubah nilai variabel exit agar loop berhenti
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
