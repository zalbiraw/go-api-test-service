package helpers

import (
	"gopkg.in/loremipsum.v1"
	"math/rand"
)

var loremIpsumGenerator = loremipsum.New()

func RandSentence() *string {
	rnd := loremIpsumGenerator.Words(rand.Intn(5) + 1)
	return &rnd
}

func RandSentences(n int) *string {
	rnd := loremIpsumGenerator.Sentences(rand.Intn(n) + 1)
	return &rnd
}
