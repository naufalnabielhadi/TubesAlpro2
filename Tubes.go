package main

import ("fmt";"math/rand")

const NMAX int = 1000
// NMAX untuk membatasi kapasitas array
type wallet struct{
	nomor int
	masterSeed string
	address string
	privateKey string
	publicKey string
	saldo float64
}
type receipt struct{
	address1 string
	address2 string
	nominal float64
	TXID string
}
var kombinasiHexa = [15]string{"1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}
var kombinasiHash = [61]string{"1","2","3","4","5","6","7","8","9","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
type tabWallet [NMAX]wallet
type tabBlockchain [NMAX]receipt
// Untuk setiap ruang ke-n (n adalah anggota himpunan bilangan real), array menyimpan 4 jenis elemen bertipe wallet.

func main(){
	
	var wlt tabWallet
	var tabBc tabBlockchain
	var i,opsi int = 0,0
	for i = 0; i < NMAX; i++{
		if i < NMAX {
			wlt[i].nomor = -1
		}
	}
	
	for i = 0; i < NMAX; i++{
		if i < NMAX {
			tabBc[i].TXID = "Null"
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
		case 4:
			transactionLobby(wlt,&opsi)
		case 5:
			transactionTab(&wlt,&tabBc,&opsi)
		case 6:
			blockchainOutput(tabBc,&opsi)
		case 7:
			cheatTab(&wlt,&opsi)
		}
		
		walletSort(&wlt)
	}
	
}

func exit(){
}

func seedSearch(wlt tabWallet, seed string){
// I.S prosedur menerima array bertipe tabWallet dan master seed dalam string sebagai input.
// F.S prosedur menampilkan semua address, private key, public key, dan saldo yang berasosiasi dengan master seed yang diberikan.
	var i,founded int = 0,0
	var saldoTotal float64 = 0
	
	for i < NMAX && i != -1{
		if wlt[i].nomor == -1{
			// i == -2 karena akan melalui i++
			i = -2
		} else if wlt[i].masterSeed == seed{
			fmt.Printf("Address     : %s\n", wlt[i].address)
			fmt.Printf("Private Key : %s\n", wlt[i].privateKey)
			fmt.Printf("Public Key  : %s\n", wlt[i].publicKey)
			fmt.Printf("Saldo (BTC) : %.8f\n", wlt[i].saldo)
			saldoTotal += wlt[i].saldo
			fmt.Println("-------------------")
			fmt.Printf("Saldo Total : %.8f\n", saldoTotal)
			founded++
			fmt.Println("")
		}
		i++
	}
	
	if founded == 0{
		fmt.Println("")
		fmt.Println("No address matched.")
		fmt.Println("")
	}
}

func masterScanner(masterSeed *string){
// I.S prosedur menampung variabel bertipe string pada parameter formal masterSeed.
// F.S prosedur menggenerasi master seed.
	var i1,i2,i3,notAlphabet int
	var tempSeed string
		
	for len(*masterSeed) == 0{
		notAlphabet = 0
		i1 = 0
		i2 = 0
		fmt.Println("Buat master seed dengan 4 kata yang dipisah satu spasi (gunakan '.' dipisah dengan spasi sebagai sentinel): ")
		for i1 == 0 && notAlphabet == 0{
			i3 = 0
			fmt.Scan(&tempSeed)
				
			for i3 < len(tempSeed) && tempSeed != "." {
				if !(('A' <= tempSeed[i3] && tempSeed[i3] <= 'Z') || ('a' <= tempSeed[i3] && tempSeed[i3] <= 'z')){
					notAlphabet++
				}
				i3++
			}
				
			if notAlphabet == 0 && tempSeed != "."{
				i2++
				*masterSeed = *masterSeed + tempSeed
			} else {
				i1++
			}
		}
			
		if notAlphabet > 0 {
			fmt.Println("")
			fmt.Println("Seed is invalid, contains a non-Alphabetic character")
			fmt.Println("")
			*masterSeed = ""
		} else if i2 < 4 {
			fmt.Println("")
			fmt.Println("Seed is invalid, words are less")
			fmt.Println("")
			*masterSeed = ""
		} else if i2 > 4 {
			fmt.Println("")
			fmt.Println("Seed is invalid, words are too many")
			fmt.Println("")
			*masterSeed = ""
		}
	}
}

func privateKeySearcher(wlt tabWallet, privateKey1 *string, keyLocation *int){
// I.S menerima dan membaca array bertipe tabWallet.
// F.S mengembalikan private key dan lokasi dari wallet yang dicari.

	var i1,i2 int
	var foundedKey, isLatest, finished bool = false,true,false
	var masterSeed1 string
	
	for !finished {
		i1 = 0
		i2 = 0
		foundedKey = false
		isLatest = true
		finished = false
		fmt.Scanln(&*privateKey1)
		
		for i1 < NMAX && !foundedKey && !finished{
			if wlt[i1].nomor == -1 {
				finished = true
			} else if wlt[i1].privateKey == *privateKey1{
				foundedKey = true
			} else {
				i1++
			}
		}
	
		masterSeed1 = wlt[i1].masterSeed
		finished = false
		
		for i2 < NMAX && isLatest && !foundedKey && !finished{
			if wlt[i2].nomor == -1{
				finished = true
			} else if masterSeed1 == wlt[i2].masterSeed && i2 > i1{
				isLatest = false
			} else {
				i2++
			}
		}
		
		if foundedKey && !isLatest{
			*privateKey1 = "Outdated"
			finished = true
		} else if !foundedKey{
			fmt.Println("")
			fmt.Println("Your private key is not found. Enter a key again")
			fmt.Println("")
			finished = false
		} else {
			*privateKey1 = wlt[i2].privateKey
			*keyLocation = i2
		}
	}
}

func publicKeySearcher(wlt tabWallet, publicKey *string, masterLocation int, keyLocation *int){
// I.S menerima input variabel bertipe tabWallet dan integer.
// F.S mengembalikan public key dan lokasi dari wallet yang dicari.
	var cek,i int
	var founded,self bool = false, false
	
	for !founded{
		cek = 0
		i = 0
		self = false
		fmt.Scanf("%s\n", &*publicKey)
				
		for i < NMAX && !founded && cek == 0{
			if wlt[i].publicKey == *publicKey && !self{
				founded = true
				*keyLocation = i
			} else if wlt[i].nomor == -1{
				cek++
			} else if wlt[i].nomor == wlt[masterLocation].nomor{
				self = true
			} else{
				i++
			}
		}
				
		if self {
			fmt.Println("")
			fmt.Println("Cannot send crypto to your own wallet. Enter other's wallet")
			fmt.Println("")
		} else if cek > 0 {
			fmt.Println("")
			fmt.Println("Cannot find a wallet with associated public key. Please enter a valid public key")
			fmt.Println("")
		}
	}
}

func balance(wlt tabWallet, keyLoc int, saldo *float64){
// I.S procedure menerima input berupa sebuah variabel bertipe tabWallet dan variabel bertipe IntegerType.
// F.S procedure mengembalikan saldo sesuai dengan input keyLoc yang diberikan.
	var i, cek1 int = 0,0
	
	for i < NMAX && cek1 == 0{
		
		if wlt[i].nomor == -1 {
			cek1++
		} else if wlt[i].masterSeed == wlt[keyLoc].masterSeed{
			*saldo += wlt[i].saldo
		} else {
			i++
		}
	}
}

func generatorHexa8(wlt tabWallet, kode *string){
// I.S prosedur menerima input berupa string (kode).
// F.S prosedur mengembalikan string yang sudah terisi oleh 8 karakter heksadesimal yang unik.
	var addKode string
	var i,j int = 0,0
	
	for len(*kode) != 8{
		i = 0
		j = 0
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

func generatorHash8(tabBc tabBlockchain, kode *string){
// I.S prosedur menerima input berupa string (kode).
// F.S prosedur mengembalikan string yang sudah terisi oleh 8 karakter hash yang unik.
	var addKode string
	var i,j int = 0,0
	
	for len(*kode) != 8{
		for j = 0; j < 8; j++{
			addKode = kombinasiHexa[rand.Intn(61)]
			*kode = *kode + addKode
		}
		
		for i < NMAX && *kode != ""{
			if tabBc[i].TXID == *kode{
				*kode = ""
			}
			i++
		}
	}
}

func blockchain(wlt *tabWallet, tabBc *tabBlockchain, keyLoc1, keyLoc2 int, nominal float64, status *bool){
// I.S menerima input berupa dua variabel integer dan satu variabel string.
// F.S mengisi array bertipe blockchain dengan catatan transaksi dan memperbarui data pada array bertipe tabWallet.
	var i1,i2,i3,i4,i5,isAvailable,addressNeeded int = 0,0,0,0,0,0,0
	var private1,private2,public1,public2,address1,address2,kode1 string
	var saldo1,saldo2,change float64 = 0,0,0
	var notExist bool = false
	
	if wlt[keyLoc1].saldo > nominal{
		addressNeeded = 3
		for i1 < NMAX && isAvailable <= addressNeeded{
			if tabBc[i1].TXID == "Null"{
				isAvailable++
			}
			i1++
		}
		isAvailable -= 1
		i1 -= addressNeeded+1
	} else if wlt[keyLoc1].saldo == nominal{
		addressNeeded = 2
		for i1 < NMAX && isAvailable <= addressNeeded && !notExist{
			if tabBc[i1].TXID == "Null"{
				isAvailable++
			}
			i1++
		}
		isAvailable -= 1
		i1 -= addressNeeded+1
	} else {
		addressNeeded = 2
		for i1 < NMAX && nominal > saldo1 && !notExist{
			if wlt[i1].nomor != -1 {
				if wlt[i1].masterSeed == wlt[keyLoc1].masterSeed && i1 != keyLoc1{
					saldo1 += wlt[i1].saldo
					addressNeeded++
				}
				i1++
			} else if nominal > saldo1 {
				saldo1 += wlt[i1].saldo
				addressNeeded++
				notExist = true
			} else {
				notExist = true
			}
		}
		
		i1 = 0
		
		if saldo1 == nominal {
			addressNeeded = 1
		} else {
			addressNeeded = 2
		}
		
		for i1 < NMAX && isAvailable <= addressNeeded {
			if tabBc[i1].TXID == "Null"{
				isAvailable++
			}
			i1++
		}
		
		isAvailable -= 1
		i1 -= addressNeeded+1
	}
	
	if isAvailable < addressNeeded {
		*status = false
	} else {
		if wlt[keyLoc1].saldo > nominal {
			change = wlt[keyLoc1].saldo - nominal
			generatorHash8(*tabBc,&kode1)
			tabBc[i1].TXID = kode1 
			tabBc[i1].nominal = wlt[keyLoc1].saldo
			tabBc[i1].address1 = wlt[keyLoc1].address
			tabBc[i1+1].TXID = kode1 
			tabBc[i1+1].nominal = nominal
			tabBc[i1+1].address2 = wlt[keyLoc2].address
			tabBc[i1+2].TXID = kode1 
			tabBc[i1+2].nominal = change
			tabBc[i1+2].address1 = wlt[keyLoc1].address
			wlt[keyLoc1].saldo = change
			wlt[keyLoc2].saldo += nominal
			for i4 < NMAX && i5 == 0{
				if wlt[i4].nomor != -1 {
					i5++
				} else {
					i4++
				}
			}
			generatorHexa8(*wlt,&private1)
			generatorHexa8(*wlt,&public1)
			generatorHexa8(*wlt,&address1)
			generatorHexa8(*wlt,&private2)
			generatorHexa8(*wlt,&public2)
			generatorHexa8(*wlt,&address2)
			wlt[i4].nomor = i4+1
			wlt[i4].masterSeed = wlt[keyLoc1].masterSeed
			wlt[i4].privateKey = private1
			wlt[i4].publicKey = public1
			wlt[i4].address = address1
			wlt[i4].saldo = 0
			wlt[i4+1].nomor = i4+2
			wlt[i4+1].masterSeed = wlt[keyLoc2].masterSeed
			wlt[i4+1].privateKey = private2
			wlt[i4+1].publicKey = public2
			wlt[i4+1].address = address2
			wlt[i4+1].saldo = 0
			*status = true
		} else if wlt[keyLoc1].saldo == nominal{
			generatorHash8(*tabBc,&kode1)
			tabBc[i1].TXID = kode1 
			tabBc[i1].nominal = wlt[keyLoc1].saldo
			tabBc[i1].address1 = wlt[keyLoc1].address
			tabBc[i1+1].TXID = kode1 
			tabBc[i1+1].nominal = nominal
			tabBc[i1+1].address2 = wlt[keyLoc2].address
			for i4 < NMAX && i5 == 0{
				if (*wlt)[i4].nomor != -1 {
					i5++
				} else {
					i4++
				}
			}
			wlt[keyLoc1].saldo = change
			wlt[keyLoc2].saldo += nominal
			generatorHexa8(*wlt,&private1)
			generatorHexa8(*wlt,&public1)
			generatorHexa8(*wlt,&address1)
			generatorHexa8(*wlt,&private2)
			generatorHexa8(*wlt,&public2)
			generatorHexa8(*wlt,&address2)
			wlt[i4].nomor = i4+1
			wlt[i4].masterSeed = wlt[keyLoc1].masterSeed
			wlt[i4].privateKey = private1
			wlt[i4].publicKey = public1
			wlt[i4].address = address1
			wlt[i4].saldo = 0
			wlt[i4+1].nomor = i4+2
			wlt[i4+1].masterSeed = wlt[keyLoc2].masterSeed
			wlt[i4+1].privateKey = private2
			wlt[i4+1].publicKey = public2
			wlt[i4+1].address = address2
			wlt[i4+1].saldo = 0
			*status = true
		} else {
			generatorHash8(*tabBc,&kode1)
			for i2 < NMAX && i3 < addressNeeded {
				if wlt[i2].nomor != -1{
					if i2 != keyLoc1 && wlt[i2].masterSeed == wlt[keyLoc1].masterSeed {
						tabBc[i1].TXID = kode1 
						tabBc[i1].nominal = wlt[i2].saldo
						tabBc[i1].address1 = wlt[i2].address
						saldo2 += wlt[keyLoc1].saldo
						i1++
						i3++
					}
				} else if saldo2 < nominal && i3 < addressNeeded{
					tabBc[i1].TXID = kode1 
					tabBc[i1].nominal = wlt[keyLoc1].saldo
					tabBc[i1].address1 = wlt[keyLoc1].address
					saldo2 += wlt[keyLoc1].saldo
					i1++
					i3++
				}
				i2++
			}
			i2--
			if saldo2 > nominal{
				change = saldo2 - nominal
				tabBc[i1].TXID = kode1 
				tabBc[i1].nominal = nominal
				tabBc[i1].address2 = wlt[keyLoc2].address
				tabBc[i1+1].TXID = kode1 
				tabBc[i1+1].nominal = change
				tabBc[i1+1].address2 = wlt[keyLoc1].address
				for i4 < NMAX && i5 == 0{
					if wlt[i4].nomor != -1 {
						i5++
					} else {
						i4++
					}
				}
				wlt[keyLoc1].saldo = change
				wlt[keyLoc2].saldo += nominal
				generatorHexa8(*wlt,&private1)
				generatorHexa8(*wlt,&public1)
				generatorHexa8(*wlt,&address1)
				generatorHexa8(*wlt,&private2)
				generatorHexa8(*wlt,&public2)
				generatorHexa8(*wlt,&address2)
				wlt[i4].nomor = i4+1
				wlt[i4].masterSeed = wlt[keyLoc1].masterSeed
				wlt[i4].privateKey = private1
				wlt[i4].publicKey = public1
				wlt[i4].address = address1
				wlt[i4].saldo = 0
				wlt[i4+1].nomor = i4+2
				wlt[i4+1].masterSeed = wlt[keyLoc2].masterSeed
				wlt[i4+1].privateKey = private2
				wlt[i4+1].publicKey = public2
				wlt[i4+1].address = address2
				wlt[i4+1].saldo = 0
				*status = true
			} else {
				change = saldo2 - nominal
				tabBc[i1].TXID = kode1 
				tabBc[i1].nominal = nominal
				tabBc[i1].address2 = wlt[keyLoc2].address
				for i4 < NMAX && i5 == 0{
					if wlt[i4].nomor != -1 {
						i5++
					} else {
						i4++
					}
				}
				wlt[keyLoc1].saldo = change
				wlt[keyLoc2].saldo += nominal
				generatorHexa8(*wlt,&private1)
				generatorHexa8(*wlt,&public1)
				generatorHexa8(*wlt,&address1)
				generatorHexa8(*wlt,&private2)
				generatorHexa8(*wlt,&public2)
				generatorHexa8(*wlt,&address2)
				wlt[i4].nomor = i4+1
				wlt[i4].masterSeed = wlt[keyLoc1].masterSeed
				wlt[i4].privateKey = private1
				wlt[i4].publicKey = public1
				wlt[i4].address = address1
				wlt[i4].saldo = 0
				wlt[i4+1].nomor = i4+2
				wlt[i4+1].masterSeed = wlt[keyLoc2].masterSeed
				wlt[i4+1].privateKey = private2
				wlt[i4+1].publicKey = public2
				wlt[i4+1].address = address2
				wlt[i4+1].saldo = 0
				*status = true
			}
		}
	}
}

func walletSort(wlt *tabWallet){
// I.S membaca array bertipe tabWallet.
// F.S mengurutkan elemen array tabWallet secara descending berdasarkan nomor wallet.
	var pass,i,idx int
	var temp1 int
	var temp2,temp3,temp4,temp5 string
	var temp6 float64
	var isDeadEnd bool = false
	
	pass = 1
	for pass < NMAX && !isDeadEnd{
		i = pass
		idx = pass
		temp1 = wlt[pass].nomor
		temp2 = wlt[pass].masterSeed
		temp3 = wlt[pass].address
		temp4 = wlt[pass].privateKey
		temp5 = wlt[pass].publicKey
		temp6 = wlt[pass].saldo
		for i > 0 {
			if temp1 > wlt[i-1].nomor{
				wlt[i].nomor = wlt[i-1].nomor
				wlt[i].masterSeed = wlt[i-1].masterSeed
				wlt[i].address = wlt[i-1].address
				wlt[i].privateKey = wlt[i-1].privateKey
				wlt[i].publicKey = wlt[i-1].publicKey
				wlt[i].saldo = wlt[i-1].saldo
				idx = i-1
			}
			i--
		}
		wlt[idx].nomor = temp1
		wlt[idx].masterSeed = temp2
		wlt[idx].address = temp3
		wlt[idx].privateKey = temp4
		wlt[idx].publicKey = temp5
		wlt[idx].saldo = temp6
		if wlt[pass].nomor == -1{
			isDeadEnd = true
		} else {
			pass++
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
		
	case 2:
		*opsi = 6
		
	case 3:
		*opsi = 4
		
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
// I.S prosedur memfasilitasi pengguna untuk membuat sebuah wallet dengan meminta input master seed yang ingin dibuat.
// F.S prosedur membuat sebuah wallet dengan menginput data bertipe tabWallet.

	var i1 int = 0
	var address,key1,key2,masterSeed1 string
	var isAvailable bool = false
	
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
	
	if wlt[NMAX-1].nomor == -1{
		masterScanner(&masterSeed1)
		wlt[i1].masterSeed = masterSeed1
		
		generatorHexa8(*wlt,&address)
		generatorHexa8(*wlt,&key1)
		generatorHexa8(*wlt,&key2)
		
		wlt[i1].address = address
		wlt[i1].privateKey = key1
		wlt[i1].publicKey = key2
		wlt[i1].saldo = 0
		
		fmt.Println("")
		fmt.Println("Wallet has been set up.")
		*opsi = 1
	} else {
		fmt.Println("Space is full.")
		*opsi = 0
	}
	fmt.Println("")
}

func walletLogIn(wlt tabWallet, opsi *int){
// I.S prosedur menerima input bertipe tabWallet dan meminta input master seed.
// F.S prosedur menampilkan informasi personal wallet berdasarkan input master seed.
	
	var tempSeed, masterSeed1 string
	var i1,i2,i3,notAlphabet int
		
	fmt.Println("-------------------------------")
	fmt.Println("     ENTER YOUR MASTER KEY")
	fmt.Println("-------------------------------")
	
	for len(masterSeed1) == 0{
		tempSeed = ""
		notAlphabet = 0
		i1 = 0
		i2 = 0
		for tempSeed != "." && notAlphabet == 0{
			i3 = 0
			fmt.Scan(&tempSeed)
			for i3 < len(tempSeed) && tempSeed != "."{
				if !(('A' <= tempSeed[i3] && tempSeed[i3] <= 'Z') || ('a' <= tempSeed[i3] && tempSeed[i3] <= 'z')){
					notAlphabet++
				}
				i3++
			}
			
			if notAlphabet == 0 && tempSeed != "."{
				i2++
				masterSeed1 += tempSeed
			} else {
				i1++
			}
		}
		
		if notAlphabet > 0{
			fmt.Println("")
			fmt.Println("Master seed is invalid. Contain a non-Alphabetic character.")
			fmt.Println("")
			masterSeed1 = ""
		} else if i2 < 4{
			fmt.Println("")
			fmt.Println("Master seed is invalid. Words are less.")
			fmt.Println("")
			masterSeed1 = ""
		} else if i2 > 4{
			fmt.Println("")
			fmt.Println("Master seed is invalid. Words are too many.")
			fmt.Println("")
			masterSeed1 = ""
		}
	}
	
	fmt.Println("------------------------------")
	fmt.Println("         YOUR ADDRESS")
	fmt.Println("------------------------------")
	
	seedSearch(wlt,masterSeed1)
	*opsi = 1
}

func transactionLobby (wlt tabWallet, opsi *int){
	*opsi = -1
	
	fmt.Println("-------------------------------------------")
	fmt.Println("     WELCOME TO THE TRANSACTION CENTER")
	fmt.Println("-------------------------------------------")
	fmt.Println("1.Transfer to other user")
	fmt.Println("2.Cheat")
	fmt.Println("3.Back to menu")
	fmt.Println("-------------------------------------------")
	for *opsi != 1 && *opsi != 2 && *opsi != 3{
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&*opsi)
		fmt.Println("")
		if *opsi != 1 && *opsi != 2 && *opsi != 3{
			fmt.Println("Option is invalid")
			fmt.Println("")
		}
	}
	
	switch *opsi {
		
	case 1:
		*opsi = 5
		
	case 2:
		*opsi = 7
		
	case 3:
		*opsi = 0
	}
	
}

func transactionTab (wlt *tabWallet, tabBc *tabBlockchain,  opsi *int){
	
	var i1,cek1,cek2 int = 1,0,0
	var key1,key2 string
	var keyLoc1, keyLoc2 int = -1,-1
	var isAccepted,isFull bool = false,false
	var status bool
	var saldo,nominal float64 = 0,0
	
	if wlt[NMAX-2].nomor != -1{
		isFull = true
	}
	
	for i1 < NMAX && wlt[0].nomor != -1 && cek1 == 0 && cek2 == 0 && !isFull{
		if wlt[0].masterSeed != wlt[i1].masterSeed {
			cek1++
		}
		if wlt[i1].nomor == -1{
			cek2++
		}
		i1++
	}
	
	if cek1 > 0 && wlt[0].nomor != -1 && wlt[1].nomor != -1 && key1 != "Outdated" && !isFull{
		fmt.Println("------------------------------")
		fmt.Println("    ENTER YOUR PRIVATE KEY")
		fmt.Println("------------------------------")
		fmt.Println("Petunjuk: gunakan '.' yang dispasi sebagai sentinel")
		
		privateKeySearcher(*wlt,&key1,&keyLoc1)
		balance(*wlt,keyLoc1,&saldo)
		
		if len(key1) == 8 && key1 != "Outdated" && keyLoc1 != -1 && saldo != 0 {
			fmt.Println("------------------------------")
			fmt.Println("       ENTER PUBLIC KEY")
			fmt.Println("------------------------------")
			
			publicKeySearcher(*wlt,&key2,keyLoc1,&keyLoc2)
			
			if keyLoc2 != -1{
				fmt.Println("------------------------------")
				fmt.Println("      ENTER YOUR NOMINAL")
				fmt.Println("------------------------------")
				fmt.Printf("Your balance: %.8f\n",saldo)
				fmt.Println("")
				for !isAccepted{
					fmt.Scanf("%f\n",&nominal)
					
					if nominal > saldo{
						fmt.Println("")
						fmt.Println("Your request exceeds your current balance")
						fmt.Println("")
					} else {
						blockchain(&*wlt,&*tabBc,keyLoc1,keyLoc2,nominal,&status)
						if status {
							fmt.Println("")
							fmt.Println("Your transaction has been concluded.")
							*opsi = 4
						} else {
							fmt.Println("")
							fmt.Println("Blockchain database is out of range.")
							*opsi = 4
						}
					}
				}
			}
		}
	} else if isFull{
		fmt.Println("")
		fmt.Println("The wallet bank is full.")
		*opsi = 4
	} else if cek1 == 0{
		fmt.Println("")
		fmt.Println("No valid user target.")
		*opsi = 4
	} else if key1 == "Outdated"{
		fmt.Println("")
		fmt.Println("Your private key is outdated.")
		*opsi = 4
	} else if saldo == 0{
		fmt.Println("")
		fmt.Println("Your balance is zero. Cannot do any further transaction in this category.")
		*opsi = 4
	}
	fmt.Println("")
}

func cheatTab(wlt *tabWallet, opsi *int){
	var i1,i2,nAccount int = 0,0,0
	var isDeadEnd,isValid bool = false,false
	var nominal float64
	var public1 string
	
	fmt.Println("-------------------------------------------")
	fmt.Println("           CRYPTO JALUR BELAKANG")
	fmt.Println("-------------------------------------------")
	
	for i1 < NMAX && !isDeadEnd{
		if wlt[i1].nomor != -1{
			nAccount++
		} else {
			isDeadEnd = true
		}
		i1++
	}
	i1--
	isDeadEnd = false
	
	if nAccount > 0{
		fmt.Println("------------------------------------------")
		fmt.Println("             ENTER PUBLIC KEY")
		fmt.Println("------------------------------------------")
		for !isValid{
			i2 = 0
			isDeadEnd = false
			fmt.Scan(&public1)
			for i2 < NMAX && !isDeadEnd && !isValid{
				if wlt[i2].publicKey == public1{
					isValid = true
				} else if wlt[i2].nomor == -1{
					isDeadEnd = true
				}
				i2++
			}
			
			if !isValid {
				fmt.Println("")
				fmt.Println("Public key is invalid. Please enter again.")
				fmt.Println("")
			}
		}
		isValid = false
		fmt.Println("------------------------------------------")
		fmt.Println("            ENTER YOUR NOMINAL")
		fmt.Println("------------------------------------------")
		for !isValid{
			fmt.Scan(&nominal)
			if nominal < 0{
				fmt.Println("")
				fmt.Println("Nominal is invalid.")
				fmt.Println("")
			} else {
				fmt.Println("")
				fmt.Println("Crypto has been send.")
				fmt.Println(i2)
				wlt[i2-1].saldo = nominal
				isValid = true
			}
		}
	} else {
		fmt.Println("")
		fmt.Println("No valid target. No account has been created.")
	}
	fmt.Println("")
	*opsi = 4
}

func blockchainOutput(tabBc tabBlockchain, opsi *int){
	var i int
	var isDeadEnd bool = false
	
	fmt.Println("------------------------------")
	fmt.Println(          "Blockchain")
	fmt.Println("------------------------------")
	
	if tabBc[0].TXID == "Null"{
		fmt.Println("")
		fmt.Println("No transaction recorded")
		fmt.Println("")
	} else {
		for i < NMAX && !isDeadEnd{
			if tabBc[i].TXID != "Null"{
				if len(tabBc[i].address1) > 0 && len(tabBc[i].address2) == 0{
					if i == 0{
						fmt.Printf("[%s] %.8f from %s\n",tabBc[i].TXID,tabBc[i].nominal,tabBc[i].address1)
					} else {
						fmt.Printf("         %.8f from %s\n",tabBc[i].nominal,tabBc[i].address1)
					}
				} else if len(tabBc[i].address1) == 0 && len(tabBc[i].address2) > 0{
					if i == 0{
						fmt.Printf("[%s] %.8f to %s\n",tabBc[i].TXID,tabBc[i].nominal,tabBc[i].address2)
					} else {
						fmt.Printf("         %.8f to %s\n",tabBc[i].nominal,tabBc[i].address2)
					}
				}
			} else {
				isDeadEnd = true
			}
		}
	}
	*opsi = 4
}