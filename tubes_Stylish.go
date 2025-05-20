package main

import "fmt"

const NMAX int = 100

type outfit struct {
	atasan string
	bawah string
	sepatu string
	aksesoris string
	kategori int  //1 = kasual, 2 = semi, 3 = formal
	cuaca string  //"panas", "dingin", "hujan"
	terakhir int  //terakhir digunakan
}

type tabOutfit [NMAX]outfit

//menginput atau nambah baju
func inputOutfit(A *tabOutfit, n *int) {
	var tambah string
	var lanjut bool = true

	for tambah == "y" && *n < NMAX {
		fmt.Println("Masukkan data outfit:")
		fmt.Print("Atasan: ")
		fmt.Scan(&A[*n].atasan)
		fmt.Print("\nBawahan: ")
		fmt.Scan(&A[*n].bawah)
		fmt.Print("\nSepatu: ")
		fmt.Scan(&A[*n].sepatu)
		fmt.Print("\nAksesoris: ")
		fmt.Scan(&A[*n].aksesoris)
		fmt.Print("Tingkat kategori (1=kasual, 2=semi, 3=formal): ")
		fmt.Scan(&A[*n].kategori)
		fmt.Print("Tanggal terakhir dipakai (yyyymmdd): ")
		fmt.Scan(&A[*n].terakhir)
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

//output
func tampilLemari(A tabOutfit, n int) {
	var i int
	if n == 0 {
		fmt.Println("Lemari kosong")
	} else {
		for i = 0; i <= n-1; i++ {
			fmt.Printf("OUTFIT%d: %s - %s - %s - %s\n", i+1, A[i].atasan, A[i].bawah, A[i].sepatu, A[i].aksesoris)
		}
	}
}

//kalo pengen ngedit isi lemari
func editOutfit(A tabOutfit, n int) {
	var i, idx int
	if n != 0 {
		fmt.Println("\nDaftar Outfit:")
		for i = 0; i <= n-1; i++ {
			fmt.Printf("OUTFIT%d: %s - %s - %s - %s\n", i+1, A[i].atasan, A[i].bawah, A[i].sepatu, A[i].aksesoris)
		}

		fmt.Print("\nMasukkan nomor outfit yang ingin diedit: ")
		fmt.Scan(&idx)
		idx--

		if idx >= 0 && idx < n {
			fmt.Println("\nMasukkan data baru untuk outfit:")
			fmt.Print("Atasan: ")
			fmt.Scan(&A[idx].atasan)
			fmt.Print("Bawahan: ")
			fmt.Scan(&A[idx].bawah)
			fmt.Print("Sepatu: ")
			fmt.Scan(&A[idx].sepatu)
			fmt.Print("Aksesoris: ")
			fmt.Scan(&A[idx].aksesoris)
			fmt.Print("Tingkat kategori (1=kasual, 2=semi, 3=formal): ")
			fmt.Scan(&A[idx].kategori)
			fmt.Print("Tanggal terakhir dipakai (yyyymmdd): ")
			fmt.Scan(&A[idx].terakhir)

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

	if *n == 0{
		fmt.Println("Outfit tidak tersedia")
		return
	}
	
	tampilLemari(*A, *n)
	fmt.Println("Pilih outfit yang ingin dihapus:")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < *n{
		for i = idx; i < *n-1; i++{
			A[i] = A[i+1]
		} 
		*n--
		fmt.Println("Outfit sudah dihapus")
	}
}

//nyari warna menggunakan sequential
func cariWarnaSequential() {

}

//nyari kategori menggunakan binary
func cariKategoriBinary() {

}

//ngurutin outfit berdasarkan kategori baju
func selectionSortkategori() {

}

//ngurutin outfit yang terakhir dipakai ke yang paling baru dipake
func insertionSortTerakhirDipakai(A *tabOutfit, n int) { 
	var i, j int
	var temp tabOutfit

	for i = 0; i <= n-1; i++{
		temp = A[i]
		i--

		//pengurutan dari tanggal lama ke baru
		for j >= && A[i] > temp{
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
	var idx int 
	var cuaca string

	if n == 0{
		fmt.Println("Tidak tersedia")
		return
	}

	fmt.Print("\nMasukkan tingkat kategori yang diinginkan (1=kasual, 2=semi, 3=formal): ")
	fmt.Scan(&preferensi)
	fmt.Print("Masukkan kondisi cuaca (panas/dingin/hujan): ")
	fmt.Scan(&cuaca)

	for i = 0; i <= n-1; i++{
		if A[i].kategori == preferensi && A[i].cuaca == cuaca {
			if idx == -1 || A[i].terakhir < A[idx].terakhir {
				idx = i
			}
		}
	}

	if idx != -1 {
		fmt.Println("\nRekomendasi outfit terbaik untukmu:")
		fmt.Printf("Atasan: %s\n", A[idx].atasan)
		fmt.Printf("Bawahan: %s\n", A[idx].bawah)
		fmt.Printf("Sepatu: %s\n", A[idx].sepatu)
		fmt.Printf("Aksesoris: %s\n", A[idx].aksesoris)
		fmt.Printf("Cocok untuk cuaca: %s\n", A[idx].cuaca)
	} else {
		fmt.Println("Tidak ada outfit yang cocok")
	}
}

func main() {
	var pakaian tabOutfit
	var nPakaian int
	var menu int

	
}
