package korwin

import (
	"math/rand"
	"strings"
)

func GenerateStatement() string {
	var quote []string

	for _, shard := range korwinQuotes {
		quote = append(quote, shard[rand.Intn(len(shard))])
	}

	return strings.Join(quote, " ")
}
