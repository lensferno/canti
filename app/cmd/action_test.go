package cmd

import (
	"log"
	"testing"
)

func TestCreateService(t *testing.T) {
	err := createService(nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//fmt.Println("ok")
}
