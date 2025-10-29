<<<<<<< HEAD
// test komen git sekian terimakasih
=======
>>>>>>> 30862e50be49c442a317349aaf55a877e925a007
package main

import "fmt"

type Sampah struct {
	Jenis     string // Organik / Anorganik / B3
	Berat     float64
	Lokasi    string
	DaurUlang bool
}

const NMAX = 1000

type dataSampah [NMAX]Sampah

var totalSampah int

func main() {
	//var sampah Sampah
	//var data dataSampah
	for {
		login()
		/*fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH ===")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Hapus Data Sampah")
		fmt.Println("3. Lihat Data Sampah")
		fmt.Println("4. Tambah Data Daur Ulang")
		fmt.Println("5. Statistik Sampah")
		fmt.Println("6. Keluar")
		fmt.Println("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahSampah()
		case 2:
			hapusSampah()
		case 3:
			lihatSampah()
		case 4:
			tambahDaurUlang()
		case 5:
			statistikSampah()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}*/
	}
}

func login() {

	var pilih int

	fmt.Println("\n=== SELAMAT DATANG DI APLIKASI PENGELOLAAN SAMPAH ===")
	fmt.Println("1. Log In Sebagai Admin") // bisa semua fitur
	fmt.Println("2. Log In Sebagai Guest") // hanya bisa jual bahan daur ulang, dan melihat data sampah saja
	fmt.Println("3. Keluar")
	fmt.Println("=======================================================")
	fmt.Println("Pilih Menu: ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		loginadmin()
	case 2:
		loginguest()
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid!")

	}
}

func loginadmin() {
	var usn, pw string
	fmt.Println("Username: ")
	fmt.Scan(&usn)
	fmt.Print("Password: ")
	fmt.Scan(&pw)

	if usn == "Admin" && pw == "Admin123" {

		fmt.Println("Log In Berhasil Sebagai Admin!")
		menuAdmin()

	} else {
		fmt.Println("Username atau Password Salah!")
	}
}

func menuAdmin() {

	var pilihan int
	var sampah Sampah
	var data dataSampah

	fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH ===")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Hapus Data Sampah")
	fmt.Println("3. Lihat Data Sampah") // perincian jumlah dari masing masing jenis sampah
	fmt.Println("4. Statistik Sampah")  // mengurutkan jenis sampah tertinggi ke terendah dan total yg bisa didaur ulang

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahSampah(sampah)
	case 2:
		hapusSampah(data)
	case 3:
		lihatSampah()
	case 4:
		statistikSampah()
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func loginguest() {

	fmt.Println("Log In Sebagai Guest")
	menuGuest()

}

func menuGuest() {

	var pilihan int
	var sampah Sampah

	fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH ===")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Lihat Data Sampah")
	fmt.Println("3. Jual Daur Ulang Sampah")
	fmt.Println("4. Statistik Sampah")

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahSampah(sampah)
	case 2:
		lihatSampah()
	case 3:
		tambahDaurUlang()
	case 4:
		statistikSampah()
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func tambahSampah(x Sampah) {

	/*var jenis string
	var berat float64
	var lokasi string*/
	var pilih int

	totalSampah = 0

	fmt.Println("\n===Jenis Sampah===")
	fmt.Println("1. Organik\n 2. Anorganik\n 3. B3 (Bahan Berbahaya & Beracun)")
	fmt.Println("\nPilih jenis:")

	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		x.Jenis = "Organik"

	case 2:
		x.Jenis = "Anorganik"

	case 3:
		x.Jenis = "B3 (Bahan Berbahaya & Beracun)"

	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("\nMasukkan Berat:")
	fmt.Scan(&x.Berat)

	fmt.Println("\nMasukkan Lokasi:")
	fmt.Scan(&x.Lokasi)

	dataSampah[totalSampah] = Sampah{}
	totalSampah++

	fmt.Println("Data Berhasil Ditambah!")
}

func hapusSampah(x *dataSampah, n *int) {
	var i, j int
	fmt.Println("input data yang akan dihapus")
	fmt.Scan(&j)
	for i := 0; i < *n; i++ {
		x[j-1] = x[j]
		j++
	}
	*n = i - 1
}

func lihatSampah(x dataSampah, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d.%s %f %s\n", i+1, x[i].Jenis, x[i].Berat, x[i].Lokasi)
	}

}

func tambahDaurUlang() {

}

func statistikSampah() {

}
