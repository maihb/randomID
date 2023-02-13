package rid

import (
	"fmt"
	"testing"
	"time"
)

var base []byte = []byte{'a', 'S', 'n', 'u', 'U', 'P', 'x', 's', 'w', '2', 'g', 'v', 'c', 'C', 'Q', 'p', '3', 'l', '5', 'W', 'K', 'V', 'i', '1', 'I', '4', 'M', 'A', 'O', 'R', '8', 'k', 'y', 'E', 'D', 'j', 'G', 'm', 'Z', 'X', 'd', 'f', '6', 'H', 't', 'r', 'B', 'z', 'Y', 'J', 'o', 'L', 'h', 'F', '0', 'e', 'N', '7', 'b', 'T', '9', 'q'}

func TestNewBase(t *testing.T) {
	m := make(map[byte]struct{})
	for _, v := range base {
		m[v] = struct{}{}
	}
	str := "{"
	for k, _ := range m {
		if str == "{" {
			str += "'" + string(k) + "'"
		} else {
			str += ",'" + string(k) + "'"
		}
	}
	fmt.Println(str + "}")
}

func TestGenid(t *testing.T) {
	fmt.Println(generateAccount(base))
}
func TestMutiGenid(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(i, generateAccount(base))
		time.Sleep(1 * time.Millisecond)
	}
}
