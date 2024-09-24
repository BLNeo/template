package util

import (
	"fmt"
	"math/rand"
)

func GenerateValidateCode() string {
	randomNumber := rand.Intn(999999)
	return fmt.Sprintf("%06d", randomNumber)
}
