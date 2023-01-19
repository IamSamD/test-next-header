package randomstring

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strconv"
)

func GenerateRandomString(size string) string {
	stringSizeInt, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}

	byteArr := make([]byte, stringSizeInt)

	rand.Read(byteArr)

	return base64.StdEncoding.EncodeToString(byteArr)
}
