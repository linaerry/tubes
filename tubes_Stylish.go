package main

import "fmt"

const NMAX int = 100

type outfit struct {
	atasan string
	bawah string
	sepatu string
	aksesoris string
	kategori int    //1 = kasual, 2 = semi, 3 = formal
	cuaca string //"panas", "dingin", "hujan"
	terakhir int    //terakhir digunakan
}

type tabOutfit [NMAX]outfit

//menginput atau nambah baju
func inputOutfit(A *tabOutfit, n *int) {
	var tambah string
	tambah = "y"

	for (tambah == "y" || tambah == "Y") && *n < NMAX {
		fmt.Println("Masukkan data outfit:")
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
		}
	} else {
		fmt.Println("Lemari kosong")
	}
}

//kalo pengen ngedit isi lemari
func editOutfit(A *tabOutfit, n int) {
	var i, kode, found int
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
			fmt.Print("Tanggal terakhir dipakai (yyyymmdd): ")
			fmt.Scan(&A[found].terakhir)
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
	fmt.Println("Pilih outfit yang ingin dihapus:")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < *n {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Outfit sudah dihapus")
	}
}

//nyari warna menggunakan sequential
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
		if A[k].atasan == cuaca || A[k].bawah == cuaca || A[k].sepatu == cuaca || A[k].aksesoris == cuaca {
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

//nyari kategori menggunakan binary
func cariKategoriBinary(A tabOutfit, n int, cari int) bool {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if A[mid].kategori == cari {
			return true
		} else if cari < A[mid].kategori {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

//ngurutin outfit berdasarkan kategori baju
func selectionSortkategori(A *tabOutfit, n int) {
	var pass, idx, i int
	var temp outfit

	pass = 0
	for pass < n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].kategori < A[i].kategori {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
	fmt.Println("Outfit sudah diurutkan berdasarkan kategori")
}

//ngurutin outfit yang terakhir dipakai ke yang paling baru dipake
func insertionSortTerakhirDipakai(A *tabOutfit, n int) {
	var i, j int
	var temp outfit

	for i = 0; i <= n-1; i++ {
		temp = A[i]
		j = i - 1

		//pengurutan dari tanggal lama ke baru
		for j >= 0 && A[i].terakhir > temp.terakhir {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
	fmt.Println("Outfit sudah diurutkan sesuai terakhir digunakan")
}

//rekomen
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

	if A[i].kategori == preferensi && A[i].cuaca == cuaca {
		for i = 0; i <= n-1; i++ {
			if idx == -1 || A[i].terakhir < A[idx].terakhir {
				idx = i
			}
		}
	} else {
		fmt.Println("Tidak ada outfit yang cocok")
	}

	if idx != -1 {
		fmt.Println("\nRekomendasi outfit terbaik untukmu:")
		fmt.Printf("OUTFIT-%d: %s - %s - %s - %s\n", idx+1, A[idx].atasan, A[idx].bawah, A[idx].sepatu, A[idx].aksesoris)
		fmt.Printf("Cocok untuk cuaca: %s\n", A[idx].cuaca)
	} else {
		fmt.Println("Tidak ada outfit yang cocok")
	}
	fmt.Println()
}

func main() {
	var pakaian tabOutfit
	var nPakaian, menu, cari, search int
	var out, cuaca string

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1. Tambah Outfit")
		fmt.Println("2. Tampilkan Lemari")
		fmt.Println("3. Edit Outfit")
		fmt.Println("4. Hapus Outfit")
		fmt.Println("5. Rekomendasi Outfit")
		fmt.Println("6. Searching")
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
				if cariKategoriBinary(pakaian, nPakaian, cari) {
					fmt.Println("Outfit dengan kategori tersebut ditemukan")
					tampilLemari(pakaian, nPakaian)
				} else {
					fmt.Println("Outfit dengan kategori tersebut tidak ditemukan")
				}
			case 2:
				fmt.Print("Masukkan kategori yang ingin dicari (1=kasual, 2=semi, 3=formal): ")
				fmt.Scan(&cuaca)
				if cariCuacaSequential(pakaian, nPakaian, cuaca) {
					fmt.Println("Outfit dengan kategori tersebut ditemukan")
					tampilLemari(pakaian, nPakaian)
				} else {
					fmt.Println("Outfit dengan kategori tersebut tidak ditemukan")
				}
			default:
				fmt.Println("Kembali ke menu")
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
