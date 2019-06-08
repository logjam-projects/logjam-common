package handlers

import (
	"fmt"
	"log"
	"testing"
)

func TestEncryption(t *testing.T) {
	log.Println("Starting")

	token := "convertone"
	phrase := "this is a simple test, oh yes it is!!!"

	outputString := EncodeData(token, phrase)
	backOutputString := DecodeData(token, outputString)

	fmt.Println("output :" + backOutputString)

}
