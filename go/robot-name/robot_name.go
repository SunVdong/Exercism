package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var m map[int]int = make(map[int]int)

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	// form AA000 to ZZ999  sum  26*26*10*10*10
	total := 26 * 26 * 10 * 10 * 10
	if len(m) >= total {
		return "", errors.New("All names have been used")
	}

	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(total)
		if m[num] == 1 {
			continue
		}

		r.name = generateNameByNum(num)
		m[num] = 1
		return r.name, nil
	}
}

func (r *Robot) Reset() {
	// The following two lines: r.name can be assigned to a new robot.
	//num, _ := generateNumByName(r.name)
	//m[num] = 0

	r.name = ""
}

func generateNameByNum(num int) string {
	char1 := strconv.Itoa(num % 10)
	char2 := strconv.Itoa(num / 10 % 10)
	char3 := strconv.Itoa(num / 10 / 10 % 10)
	char4 := string(rune(num/10/10/10%26 + 'A'))
	char5 := string(rune(num/10/10/10/26%26 + 'A'))

	return char5 + char4 + char3 + char2 + char1
}

func generateNumByName(name string) (int, error) {
	if len(name) != 5 {
		return 0, errors.New("Wrong name")
	}

	num := 0
	for i, v := range name {
		switch i {
		case 0:
			num += int(v-'A') * 26 * 10 * 10 * 10
		case 1:
			num += int(v-'A') * 10 * 10 * 10
		case 2:
			num += int(v-'0') * 10 * 10
		case 3:
			num += int(v-'0') * 10
		case 4:
			num += int(v - '0')
		}
	}

	return num, nil
}
