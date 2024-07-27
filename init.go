package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func PemToDer(pemPath string) ([]byte, error) {
	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block from file: %s", pemPath)
	}

	return block.Bytes, nil
}

func WriteDer(derData []byte, derPath string) error {
	return ioutil.WriteFile(derPath, derData, 0644)
}

func initt() {
	certPemPath := "innoversecert.pem"
	keyPemPath := "innoversekey.pem"

	certDerPath := "innoversecert.der"
	keyDerPath := "innoversekey.der"

	certDer, err := PemToDer(certPemPath)
	if err != nil {
		fmt.Printf("Error converting innoversecert.pem to DER: %v\n", err)
		return
	}

	err = WriteDer(certDer, certDerPath)
	if err != nil {
		fmt.Printf("Error writing innoversecert.der: %v\n", err)
		return
	}

	keyDer, err := PemToDer(keyPemPath)
	if err != nil {
		fmt.Printf("Error converting innoversekey.pem to DER: %v\n", err)
		return
	}

	err = WriteDer(keyDer, keyDerPath)
	if err != nil {
		fmt.Printf("Error writing innoversekey.der: %v\n", err)
		return
	}

	fmt.Println("Conversion from PEM to DER completed successfully!")
}
