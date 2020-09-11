package helpers

import (
	"math/rand"
	"time"
)

func GenrateRandomNumber() (randNumber int16) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randNumber = int16(r1.Intn(10000))
	//it make 4 digit OTP
	if randNumber <= 1000 || randNumber == 10000 {
		return GenrateRandomNumber()
	}
	return randNumber
}
