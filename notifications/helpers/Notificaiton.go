package helpers

import (
	"gopkg.in/loremipsum.v1"
	"math/rand"
)

var loremIpsumGenerator = loremipsum.New()

func RandSentence() *string {
	rnd := loremIpsumGenerator.Sentence()
	return &(rnd)
}

func RandSentences(n int) *string {
	rnd := loremIpsumGenerator.Sentences(rand.Intn(n))
	return &(rnd)
}
