package random

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func Number(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteString(strconv.Itoa(r.Intn(10)))
	}
	return buffer.String()
}

func Hybrid(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	strs := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteByte(strs[r.Intn(len(strs))])
	}
	return buffer.String()
}
