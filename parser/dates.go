package parser

import (
	"bileti_go/utils"
	"time"
)

const DateLayout = "02.01.2006"

func GetLocation() *time.Location {

	loc := utils.Must(time.LoadLocation("Europe/Sofia"))

	return loc
}
