package internal

import (
	"fmt"

	"github.com/speps/go-hashids"
)

func EncodeID(id int, userKey string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = userKey // Use the user-specific key as the salt
	hd.MinLength = 7  // Set the minimum length to 7 characters
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	e, err := h.Encode([]int{id})
	if err != nil {
		return "", err
	}
	return e, nil
}

func DecodeID(hash string, userKey string) (int, error) {
	hd := hashids.NewData()
	hd.Salt = userKey
	hd.MinLength = 7 // Ensure the same minimum length is used
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return 0, err
	}
	ids, err := h.DecodeWithError(hash)
	if err != nil {
		return 0, err
	}
	if len(ids) == 0 {
		return 0, fmt.Errorf("no ID decoded")
	}
	return ids[0], nil
}

func ReverseSring(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
