package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

//Number of characters in the hash
const numOfCharacters int = 4

//Hash the key with MD5
func hashKey(key string) string {
	return MD5(key)
	}

//Hash the key using the MD5 crypto package
func MD5(key string) string {
	hasher := md5.New()
    hasher.Write([]byte(key))
    hashString := hex.EncodeToString(hasher.Sum(nil))
	return (hashString[:numOfCharacters]+"/"+key)
	}