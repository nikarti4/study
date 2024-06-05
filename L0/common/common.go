package common

import (
	"log"
)

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}