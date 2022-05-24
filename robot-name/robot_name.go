package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	name string
}

var names = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	var output strings.Builder

	randSeed := time.Now().UnixNano()
	randSource := rand.NewSource(randSeed)
	randGen := rand.New(randSource)

	for i := 0; i <= 1; i++ {
		// uppercase A-Z.
		char := 'A' + rune(randGen.Intn(26))
		output.WriteRune(char)
	}

	for i := 0; i <= 2; i++ {
		num := randGen.Intn(9)
		output.WriteString(strconv.Itoa(num))
	}

	name := output.String()

	_, exists := names[name]
	names[name] = true

	if exists {
		return r.Name()
	} else {
		r.name = name
		return name, nil
	}
}

func (r *Robot) Reset() {
	r.name = ""
	name, _ := r.Name()
	r.name = name
}
