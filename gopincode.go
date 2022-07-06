package gopincode

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenSix() (int, string) {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	c := rand.Intn(max-min+1) + min
	form := strconv.Itoa(c)
	s := fmt.Sprintf("%s %s", form[0:3], form[3:6])

	return c, s
}
