package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func (crypto *Crypto) createWalletKeystore(password string) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	for _, acc := range ks.Accounts() {
		if acc.Address.Hex() == account.Address.Hex() {
			find, err := ks.Find(acc)
			if err != nil {
				fmt.Println(err.Error())
			}

			if b, err := ioutil.ReadFile(find.URL.Path); err != nil {
				fmt.Println(err.Error())
			} else {
				err = ioutil.WriteFile("./wallets/"+acc.Address.Hex()+".wallet.json", b, 0644)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			ks.Delete(acc, password)
		}
	}

	fmt.Println(account.Address.Hex())
}

func (crypto *Crypto) importKeystore(address string, password string) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile("./wallets/" + address + ".wallet.json")
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	ks.Delete(account, password)

}

func (crypto *Crypto) updateKs(address string, password string, newpassword string) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile("./wallets/" + address + ".wallet.json")
	if err != nil {
		log.Fatal(err)
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ks.Update(account, password, newpassword)
	if err != nil {
		fmt.Println(err.Error())
	}
	find, err := ks.Find(account)
	if err != nil {
		fmt.Println(err.Error())
	}

	if b, err := ioutil.ReadFile(find.URL.Path); err != nil {
		fmt.Println(err.Error())
	} else {
		err = ioutil.WriteFile("./wallets/"+account.Address.Hex()+".wallet.json", b, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	ks.Delete(account, newpassword)
}
