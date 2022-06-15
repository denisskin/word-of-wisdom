package wow_server

import (
	"bytes"
	_ "embed"
	"math/rand"
	"time"
)

//go:embed wisdom-book.txt
var wisdomBookSource []byte

type DB [][]byte

func newDB() DB {
	return bytes.Split(wisdomBookSource, []byte("\n\n"))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (db DB) randomWisdom() []byte {
	return db[rand.Intn(len(db))]
}
