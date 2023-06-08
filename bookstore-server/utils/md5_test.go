package utils

import (
	"fmt"
	"testing"
)

func TestMD5V(t *testing.T) {
	fmt.Println(MD5V([]byte("123456")))
}
