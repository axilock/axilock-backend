package util

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(mininum, maximum int64) int64 {
	return mininum + rand.Int63n(maximum-mininum+1) //nolint:gosec
}

func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for range n {
		c := alphabet[rand.Intn(k)] //nolint:gosec
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@test-%s.com", RandString(8), RandString(5))
}

func RandomUUID() string {
	return uuid.NewString()
}
