package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func pemToDer(pemPath string, derPath string) error {
	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return fmt.Errorf("failed to decode PEM block")
	}

	derData := block.Bytes

	err = ioutil.WriteFile(derPath, derData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func initt() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: pemtoder <input.pem> <output.der>")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	err := pemToDer(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Error converting PEM to DER: %v\n", err)
	}
}
