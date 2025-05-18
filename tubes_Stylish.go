package main

import "fmt"

const NMAX int = 100

type outfit struct {
	atasan    string
	bawah     string
	sepatu    string
	aksesoris string
}

type tabOutfit [NMAX]outfit

func tambahOutfit(A *tabOutfit, n *int) {
	var tambah string
	var lanjut bool = true

	for tambah && *n < NMAX {
		fmt.Println("Masukkan data outfit:")
		fmt.Print("Atasan: ")
		fmt.Scan(&A[*n].atasan)
		fmt.Print("\nBawahan: ")
		fmt.Scan(&A[*n].bawah)
		fmt.Print("\nSepatu: ")
		fmt.Scan(&A[*n].sepatu)
		fmt.Print("\nAksesoris: ")
		fmt.Scan(&A[*n].aksesoris)
		*n++

		if *n < NMAX {
			fmt.Print("Tambah outfit lagi? (y/n): ")
			fmt.Scan(&tambah)
			if tambah != "y" {
				lanjut = false
			}
		} else {
			fmt.Println("Lemari penuh!")
			lanjut = false
		}
	}
}

func cetakLemari(A tabOutfit, n int) {
	var i int
	if n == 0 {
		fmt.Println("Lemari kosong.")
	} else {
		fmt.Println("\nIsi Lemari:")
		for i = 0; i < n; i++ {
			fmt.Printf("OUTFIT%d: %s - %s - %s - %s\n", i+1, A[i].atasan, A[i].bawah, A[i].sepatu, A[i].aksesoris)
		}
	}
}
func urutkanPakaian(A *tabOutfit, n int){
	var i int
	
	for i = 0; i < n-1; i++{
		for j:=	0; j < n-i-1; j++{
			if j[i].atasan, > A[j+1].atasan{
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
	fmt.Println("Outfit sudah diurutkan sesuai nama atasan.")
}
func cariPakaian(A tabOutfit, n int, keyword string){
	var cari bool

	cari = false
	for i:= 0; i < n; i++{
		if A[i].atasan == keyword{
			fmt.Printf("ditemukan di indeks %d: \n", i)
			fmt.Printf("%s - %s - %s - %s\n", A[i].atasan, A[i].bawah, A[i].sepatu, A[i].aksesoris)
			cari = true
			break
		}
	}
	if !cari{
		fmt.Println("outfit dengan atasan tersebut tidak tersedia")
	}
}
func main(){
	var lemari tabOufit
	var n int
	var pilhan int
	var keyword string

	fmt.Println("\n=== MENU ===")
	fmt.Println("1. Tampilkan Semua Pakaian")
	fmt.Println("2. Tambah Pakaian)
	fmt.Println("3. Urutkan Pakaian")
	fmt.Println("4. Cari Pakaian")
	fmt.Println("5. Keluar")
	fmt.Println("Pilih Menu: ")

	fmt.Scan(&pilihan)
	
	switch pilihan {
		case 1 : 
			tambahOutfit(&lemari, &n)
		case 2 :
			cetakLemari(lemari, n)
		case 3 :
			urutkanPakaian(&lemari, n)
		case 4 : 
			fmt.Println("Masukkan Atasan Yang Dipilih:")
			fmt.Scan(&keyword)
			cariPakaian(lemari, n, keyword)
		case 5 : 
			fmt.Println("Terima Kasih!")
			return
		default : 
		fmt.Println("Pilihan Tak Tersedia")
	}
}

