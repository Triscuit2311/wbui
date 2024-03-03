package models

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Store struct {
	data map[int][]byte
	vId  int
}

func NewStore(alloc int) Store {
	return Store{
		data: make(map[int][]byte, alloc),
		vId:  0,
	}
}

func SetKV[T any](store *Store, v T) int {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(v)

	store.vId++
	store.data[store.vId] = buf.Bytes()
	return store.vId
}

func GetKV[T any](store *Store, key int) T {
	var buf bytes.Buffer

	b := store.data[key]
	_, err := buf.Write(b)
	if err != nil {
		log.Fatalf("Bad write to buffer from key: %d", key)
		panic("Invalid Write")
	}

	var ret T
	enc := gob.NewDecoder(&buf)
	enc.Decode(&ret)

	return ret
}
