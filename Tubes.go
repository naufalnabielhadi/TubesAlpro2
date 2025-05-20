package main

import ("fmt";"math/rand")

const NMAX int = 1000
// NMAX untuk membatasi kapasitas array
type wallet struct{
	nomor int
	masterSeed string
	address string
	publicKey string
	privateKey string
	saldo float64
}
var kombinasiHexa = [15]string{"1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}
type tabWallet [NMAX]wallet
// Untuk setiap ruang ke-n (n adalah anggota himpunan bilangan real), array menyimpan 4 jenis elemen bertipe wallet.

func main(){
	
	var wlt tabWallet
	var i,opsi int = 0,0
	for i = 0; i < NMAX; i++{
		if i < NMAX {
			wlt[i].nomor = -1
		}
	}
	
	for opsi != -1{
		
		switch opsi{
		
		case 0:
			menu(wlt,&opsi)
		case 1:
			walletLobby(wlt,&opsi)
		case 2:
			walletSignUp1(&wlt,&opsi)
		case 3:
			walletLogIn(wlt,&opsi)
		
		}
	}
	
}

func exit(){
}

func seedSearch(wlt tabWallet, seed string){
	var i,founded int = 0,0
	
	for i < NMAX && i != -1{
		if wlt[i].nomor == -1{
			i = -2
		} else if wlt[i].masterSeed == seed{
			fmt.Printf("Address     : %s\n", wlt[i].address)
			fmt.Printf("Private Key : %s\n", wlt[i].privateKey)
			fmt.Printf("Public Key  : %s\n", wlt[i].privateKey)
			founded++
		}
		i++
	}
	
	if founded == 0{
		fmt.Println("")
		fmt.Println("No address matched.")
	}
	fmt.Println("")
}

func GeneratorHexa8(wlt tabWallet, kode *string){
// I.S prosedur menerima input berupa string (kode).
// F.S prosedur mengembalikan string yang sudah terisi oleh 8 karakter heksadesimal yang unik.
	var addKode string
	var i,j int = 0,0
	
	for len(*kode) != 8{
		for j = 0; j < 8; j++{
			addKode = kombinasiHexa[rand.Intn(15)]
			*kode = *kode + addKode
		}
		
		for i < NMAX && *kode != ""{
			if wlt[i].address == *kode || wlt[i].privateKey == *kode || wlt[i].publicKey == *kode{
				*kode = ""
			}
			i++
		}
	}
}

func menu(wlt tabWallet, opsi *int){
// I.S prosedur "menu" menampilkan opsi-opsi utama yang dapat dipilih oleh user.
// F.S prosedur menuntun user ke prosedur lain sesuai pilihan user.
	*opsi = -1
	
	fmt.Println("------------------------------")
	fmt.Println("             MENU")
	fmt.Println("------------------------------")
	fmt.Println("1.Wallet")
	fmt.Println("2.Blockchain")
	fmt.Println("3.Transaction")
	fmt.Println("4.Keluar")
	fmt.Println("------------------------------")
	for *opsi != 1 && *opsi != 2 && *opsi != 3 && *opsi != 4{
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&*opsi)
		fmt.Println("")
		if *opsi != 1 && *opsi != 2 && *opsi != 3 && *opsi != 4{
			fmt.Println("Option is invalid")
			fmt.Println("")
		}
	}
	
	switch *opsi {
		
	case 1:
		*opsi = 1
		
	case 4:
		*opsi = -1
	}
}

func walletLobby(wlt tabWallet, opsi *int){
// I.S prosedur menampilkan menu wallet.
// F.S prosedur menuntun user ke prosedur selanjutnya sesuai dengan pilihan user.
	*opsi = -1
	
	fmt.Println("-------------------------------")
	fmt.Println("     WELCOME TO THE WALLET")
	fmt.Println("-------------------------------")
	fmt.Println("1.Enter master seed")
	fmt.Println("2.Create wallet")
	fmt.Println("3.Back to menu")	
	fmt.Println("------------------------------")
	for *opsi != 1 && *opsi != 2 && *opsi != 3{
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&*opsi)
		fmt.Println("")
		if *opsi != 1 && *opsi != 2 && *opsi != 3{
			fmt.Println("Option is invalid")
			fmt.Println("")
		}
	}
	
	switch *opsi{
		
	case 1:
		*opsi = 3
	
	case 2:
		*opsi = 2
	
	case 3:
		*opsi = 0
	}
}

func walletSignUp1(wlt *tabWallet, opsi *int){
	var i1 int = 0
	var i2,i3,i4 int
	var address,masterSeed1,tempSeed,key1,key2 string
	var isAvailable,isAlphabetic bool = false,true
	var back rune
	
	fmt.Println("-------------------------------")
	fmt.Println("         WALLET SET UP")
	fmt.Println("-------------------------------")
	for !isAvailable && i1 <= NMAX-1{
		if wlt[i1].nomor == -1{
			isAvailable = true
		} else {
			i1++
		}
	}
	wlt[i1].nomor=i1+1
	
	if wlt[999].nomor == -1{
		for len(masterSeed1) == 0{
			i2 = 0
			i3 = 0
			i4 = 0
			fmt.Println("Buat master seed dengan 4 kata yang dipisah satu spasi (gunakan '.' dipisah dengan spasi sebagai sentinel): ")
			for i2 == 0 && isAlphabetic{
				fmt.Scan(&tempSeed)
				
				for i4 < len(tempSeed) && isAlphabetic {
					if !(('A' <= tempSeed[i4] && tempSeed[i4] <= 'Z') || ('a' <= tempSeed[i4] && tempSeed[i4] <= 'z')){
						isAlphabetic = false
						fmt.Printf("%s",tempSeed[i4])
					}
					i4++
				}
				
				if tempSeed != "."{
					i3++
					masterSeed1 += tempSeed
				} else {
					i2++
				}
			}
			if i3 < 4 && isAlphabetic{
				fmt.Println("")
				fmt.Println("Seed is invalid, words are less")
				fmt.Println("")
				masterSeed1 = ""
			} else if i3 > 4 && isAlphabetic{
				fmt.Println("")
				fmt.Println("Seed is invalid, words are too many")
				fmt.Println("")
				masterSeed1 = ""
			}
		}
		wlt[i1].masterSeed = masterSeed1
		
		GeneratorHexa8(*wlt,&address)
		GeneratorHexa8(*wlt,&key1)
		GeneratorHexa8(*wlt,&key2)
		
		wlt[i1].address = address
		wlt[i1].privateKey = key1
		wlt[i1].publicKey = key2
		wlt[i1].saldo = 0
		
		fmt.Println("")
		fmt.Println("Wallet has been set up. Enter any key to back to wallet menu")
		fmt.Println("")
		fmt.Scan(&back)
		*opsi = 1
	} else {
		fmt.Println("Space is full. Enter anything to back to menu")
		fmt.Scan(&back)
		*opsi = 0
	}
}

func walletLogIn(wlt tabWallet, opsi *int){
	
	var masterSeed1 string
	var back rune
		
	fmt.Println("-------------------------------")
	fmt.Println("     ENTER YOUR MASTER KEY")
	fmt.Println("-------------------------------")
	
	fmt.Scan(&masterSeed1)
	
	fmt.Println("------------------------------")
	fmt.Println("         YOUR ADDRESS")
	fmt.Println("------------------------------")
	
	seedSearch(wlt,masterSeed1)
	fmt.Println("Enter any key to go back to menu")
	fmt.Scan(&back)
	fmt.Println("")
	*opsi = 1
}