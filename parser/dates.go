package parser

import (
	"log"
	"time"
)

const DateLayout = "02.01.2006"

func GetLocation() *time.Location {

	loc, err := time.LoadLocation("Europe/Sofia")

	if err != nil {
		log.Fatal(err)
	}

	return loc
}
