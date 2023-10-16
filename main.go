package main

import (
	"flag"
	"fmt"
)

var (
	CryptPw     string
	flagFile    string
	decryptPtr  bool
	encryptPtr  bool
	verbose     bool
	debugPtr    bool
	absPath     string
	fileName    string
	workingPath string
)

func init() {
	flag.BoolVar(&encryptPtr, "e", false, "Encrypt the input data.")
	flag.BoolVar(&decryptPtr, "d", false, "Decrypt the input data.")
	flag.StringVar(&flagFile, "in", "", "The input filename, standard input by default.")
	flag.StringVar(&CryptPw, "k", "", "The password to derive the key from.")
	flag.BoolVar(&verbose, "v", false, "Enables verbosity to default logger")
	flag.BoolVar(&debugPtr, "debug", false, "Enables debug output to default logger")
	flag.Parse()

	printver()
}

func main() {
	fmt.Println("Go-RandPassGenerator")
}
