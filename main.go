package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	flagverbose       bool
	flagdebug         bool
	flagstrength      uint
	flagpasswordnum   uint
	flagpassphrasenum uint
	flagkeynum        uint
	flagencrypt       bool
	flagdecrypt       bool
)

func init() {
	flag.BoolVar(&flagverbose, "v", false, "Enables verbosity to default logger")
	flag.BoolVar(&flagdebug, "vv", false, "Enables debug level verbosity to default logger")
	flag.UintVar(&flagstrength, "str", 160, "Use generation strength of S bits (default: 160)")
	flag.UintVar(&flagpasswordnum, "pw", 1, "Generate N random password of the specified strength")
	flag.UintVar(&flagpassphrasenum, "pp", 1, "Generate N random passphrases of the specified strength")
	flag.UintVar(&flagkeynum, "k", 1, "Generate N random keys of the specified strength")
	flag.BoolVar(&flagencrypt, "enc", false, "Encrypt generated random key using a random password that is at least a 16 characters (256-bit AES) and write to file")
	flag.BoolVar(&flagdecrypt, "decrypt", false, "Decrypt encrypted key file using a random password that is at least a 16 characters and save as text file")
	flag.Parse()

	if flagverbose {
		log.SetLevel(log.InfoLevel)
	} else if flagdebug {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	printver()
}

func main() {
	fmt.Println("Go-RandPassGenerator")

}
