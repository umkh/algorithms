package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrDuplicateKey    = errors.New("duplicate key")
	ErrKeyDoesNotExist = errors.New("key does not exist")
)

var array = make(map[int]interface{}, 0)

func main() {
	for i := 100; i < 150; i++ {
		if err := insert(i, fmt.Sprintf("%d", i)); err != nil {
			log.Fatal(err)
		}
	}

	look, err := lookup(100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(look)
}

// Add a (key, value) pair to the associative array
func insert(key int, value interface{}) error {
	if _, ok := array[key]; ok {
		return ErrDuplicateKey
	}

	array[key] = value

	return nil
}

// Remove the key's pair in the associative array
func remove(key int) error {
	if _, ok := array[key]; ok {

		delete(array, key)

		return nil
	}

	return ErrKeyDoesNotExist
}

// Assigns/updates the value to/of the existing key
func update(key int, value interface{}) error {
	if _, ok := array[key]; ok {
		delete(array, key)

		array[key] = value

		return nil
	}

	return ErrKeyDoesNotExist
}

// Returns the value assigned to this key
func lookup(key int) (interface{}, error) {
	if _, ok := array[key]; ok {
		return array[key], nil
	}

	return nil, ErrKeyDoesNotExist
}
