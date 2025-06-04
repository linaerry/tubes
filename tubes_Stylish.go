package main

import "fmt"

const NMAX int = 100

type outfit struct {
	atasan    string
	bawah     string
	sepatu    string
	aksesoris string
	kategori  int    //1 = kasual, 2 = semi, 3 = formal
	cuaca     string //"panas", "dingin", "hujan"
	terakhir  int    //terakhir digunakan
}

type tabOutfit [NMAX]outfit

func main() {
	var pakaian tabOutfit
	var nPakaian, menu, cari, search, sort, idx int
	var out, cuaca string

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1. Tambah Outfit")
		fmt.Println("2. Tampilkan Lemari")
		fmt.Println("3. Edit Outfit")
		fmt.Println("4. Hapus Outfit")
		fmt.Println("5. Rekomendasi Outfit")
		fmt.Println("6. Searching")
		fmt.Println("7. Sorting")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			inputOutfit(&pakaian, &nPakaian)
		case 2:
			tampilLemari(pakaian, nPakaian)
		case 3:
			editOutfit(&pakaian, nPakaian)
		case 4:
			hapusOutfit(&pakaian, &nPakaian)
		case 5:
			rekomendasiOutfit(pakaian, nPakaian)
		case 6:
			fmt.Println("1. Cari Berdasarkan Kategori")
			fmt.Println("2. Cari Berdasarkan Cuaca")
			fmt.Print("Masukkan yang ingin dicari: ")
			fmt.Scan(&search)
			switch search {
			case 1:
				fmt.Print("Masukkan kategori yang ingin dicari (1=kasual, 2=semi, 3=formal): ")
				fmt.Scan(&cari)
				idx = cariKategoriBinary(pakaian, nPakaian, cari)
				if idx != -1 {
					fmt.Println("Outfit dengan kategori tersebut ditemukan")
					tampilLemari(pakaian, nPakaian)
				} else {
					fmt.Println("Outfit dengan kategori tersebut tidak ditemukan")
				}
			case 2:
				fmt.Print("Masukkan cuaca yang ingin dicari (panas/dingin/hujan): ")
				fmt.Scan(&cuaca)
				if cariCuacaSequential(pakaian, nPakaian, cuaca) {
					fmt.Println("Outfit dengan cuaca tersebut ditemukan")
					tampilLemari(pakaian, nPakaian)
				} else {
					fmt.Println("Outfit dengan cuaca tersebut tidak ditemukan")
				}
			default:
				fmt.Println("Kembali ke menu.")
			}
		case 7:
			fmt.Println("1. Urutkan Berdasarkan Kategori")
			fmt.Println("2. Urutkan Berdasarkan Terakhir digunakan")
			fmt.Print("Masukkan yang ingin diurutkan: ")
			fmt.Scan(&sort)
			switch sort {
			case 1:
				fmt.Print("Data Outfit diurutkan dari formal ke kasual")
				selectionSortkategori(&pakaian, nPakaian)
				tampilLemari(pakaian, nPakaian)
			case 2:
				fmt.Print("Data Outfit diurutkan dari tanggal terbaru ke lama")
				insertionSortTerakhirDipakai(&pakaian, nPakaian)
				tampilLemari(pakaian, nPakaian)
			}
		case 0:
			fmt.Print("Yakin mau keluar? (y/n): ")
			fmt.Scan(&out)
			if out == "y" || out == "Y" {
				fmt.Println("Terima kasih sudah pakai program ini!")
				return
			}
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func tanggal(tgl string) int {
	var hari, bulan, tahun int
	fmt.Sscanf(tgl, "%2d/%2d/%4d", &hari, &bulan, &tahun)
	return tahun*10000 + bulan*100 + hari
}

//menginput atau nambah baju
func inputOutfit(A *tabOutfit, n *int) {
	var tambah, tgl string
	tambah = "y"

	for (tambah == "y" || tambah == "Y") && *n < NMAX {
		fmt.Println("Masukkan data outfit( Mohon ganti spasi dengan '_'):")
		fmt.Print("Atasan: ")
		fmt.Scan(&A[*n].atasan)
		fmt.Print("Bawahan: ")
		fmt.Scan(&A[*n].bawah)
		fmt.Print("Sepatu: ")
		fmt.Scan(&A[*n].sepatu)
		fmt.Print("Aksesoris: ")
		fmt.Scan(&A[*n].aksesoris)
		fmt.Print("Tingkat kategori (1=kasual, 2=semi, 3=formal): ")
		fmt.Scan(&A[*n].kategori)
		fmt.Print("Cocok digunakan di cuaca apa?(panas, dingin, hujan): ")
		fmt.Scan(&A[*n].cuaca)
		fmt.Print("Tanggal terakhir dipakai (dd/mm/yyyy): ")
		fmt.Scan(&tgl)
		A[*n].terakhir = tanggal(tgl)

		*n = *n + 1

		if *n < NMAX {
			fmt.Print("Tambah outfit lagi? (y/n): ")
			fmt.Scan(&tambah)
		} else {
			fmt.Println("Lemari penuh!")
			tambah = "n"
		}
	}
}

//output
func tampilLemari(A tabOutfit, n int) {
	var i int
	if n != 0 {
		for i = 0; i < n; i++ {
			fmt.Printf("OUTFIT-%d: %s - %s - %s - %s\n", i+1, A[i].atasan, A[i].bawah, A[i].sepatu, A[i].aksesoris)
			fmt.Printf("Kategori (1=kasual, 2=semi, 3=formal): %d || Cuaca: %s || Terakhir dipakai: %d\n", A[i].kategori, A[i].cuaca, A[i].terakhir)
		}
	} else {
		fmt.Println("Lemari kosong")
	}
}

//kalo pengen ngedit isi lemari
func editOutfit(A *tabOutfit, n int) {
	var i, kode, found int
	var tgl string
	found = -1

	if n != 0 {
		fmt.Println("Daftar Outfit:")
		tampilLemari(*A, n)
		fmt.Println()
		fmt.Print("Masukkan nomor outfit yang ingin diedit: ")
		fmt.Scan(&kode)

		i = 0
		for i < n && found == -1 {
			if i+1 == kode {
				found = i
			}
			i++
		}

		if found != -1 {
			fmt.Println("\nMasukkan data baru untuk outfit:")
			fmt.Print("Atasan: ")
			fmt.Scan(&A[found].atasan)
			fmt.Print("Bawahan: ")
			fmt.Scan(&A[found].bawah)
			fmt.Print("Sepatu: ")
			fmt.Scan(&A[found].sepatu)
			fmt.Print("Aksesoris: ")
			fmt.Scan(&A[found].aksesoris)
			fmt.Print("Tingkat kategori (1=kasual, 2=semi, 3=formal): ")
			fmt.Scan(&A[found].kategori)
			fmt.Print("Cuaca: ")
			fmt.Scan(&A[found].cuaca)
			fmt.Print("Tanggal terakhir dipakai (dd/mm/yyyy): ")
			fmt.Scan(&tgl)
			A[found].terakhir = tanggal(tgl)

			fmt.Println("\nOutfit berhasil diupdate")
		} else {
			fmt.Println("Nomor outfit tidak valid")
		}

	} else {
		fmt.Println("Belum ada outfit yang bisa diedit")
	}
}

//kalo mau ngilangin outfit
func hapusOutfit(A *tabOutfit, n *int) {
	var i, idx int

	if *n == 0 {
		fmt.Println("Outfit tidak tersedia")
		return
	}

	tampilLemari(*A, *n)
	fmt.Print("Pilih outfit yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < *n {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("Outfit sudah dihapus")
	}
}

//nyari cuaca menggunakan sequential
func cariCuacaSequential(A tabOutfit, n int, cuaca string) bool {
	var k int
	var isKetemu bool

	if n == 0 {
		fmt.Println("Belum ada outfit yang tersimpan.")
		return false
	}

	isKetemu = false
	k = 0
	for !isKetemu && k < n {
		if A[k].cuaca == cuaca {
			isKetemu = true
		}
		k++
	}

	if isKetemu {
		fmt.Println("Cuaca tesebut tersedia")
	} else {
		fmt.Println("Cuaca tesebut tidak tersedia")
	}
	return isKetemu
}

//ngurutin outfit berdasarkan kategori baju formal ke kasual
func selectionSortkategori(A *tabOutfit, n int) {
	var pass, idx, i int
	var temp outfit

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if A[i].kategori < A[idx].kategori {
				idx = i
			}
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}

	fmt.Println("Outfit sudah diurutkan berdasarkan kategori")
}

//nyari kategori menggunakan binary
func cariKategoriBinary(A tabOutfit, n int, cari int) int {
	selectionSortkategori(&A, n)

	var left, right, mid, found int

	left = 0
	right = n - 1
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if A[mid].kategori == cari {
			found = mid
		} else if cari < A[mid].kategori {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return found
}

//ngurutin outfit yang terakhir dipakai ke yang paling baru dipake
func insertionSortTerakhirDipakai(A *tabOutfit, n int) {
	var i, pass int
	var temp outfit

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]

		//pengurutan dari tanggal baru ke lama (descending)
		for i > 0 && temp.terakhir > A[i-1].terakhir {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	fmt.Println("Outfit sudah diurutkan sesuai terakhir digunakan")
}

//rekomen or sequential search
func rekomendasiOutfit(A tabOutfit, n int) {
	var i, preferensi int
	var idx int = -1
	var cuaca string

	if n == 0 {
		fmt.Println("Tidak tersedia")
		return
	}

	fmt.Print("\nMasukkan tingkat kategori yang diinginkan (1=kasual, 2=semi, 3=formal): ")
	fmt.Scan(&preferensi)
	fmt.Print("Masukkan kondisi cuaca (panas/dingin/hujan): ")
	fmt.Scan(&cuaca)

	for i = 0; i <= n-1; i++ {
		if A[i].kategori == preferensi && A[i].cuaca == cuaca {
			if idx == -1 || A[i].terakhir < A[idx].terakhir {
				idx = i
			}
		}
	}

	if idx != -1 {
		fmt.Println("\nRekomendasi outfit terbaik untukmu:")
		fmt.Printf("OUTFIT-%d: %s - %s - %s - %s\n", idx+1, A[idx].atasan, A[idx].bawah, A[idx].sepatu, A[idx].aksesoris)
		fmt.Printf("Kategori: %d - Cocok untuk cuaca: %s - Terakhir dipakai: %d\n", A[idx].kategori, A[idx].cuaca, A[idx].terakhir)
	} else {
		fmt.Println("Tidak ada outfit yang cocok")
	}
	fmt.Println()
}
