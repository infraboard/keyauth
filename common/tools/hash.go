package tools

import (
	"fmt"
	"hash/fnv"
)

func GenHashID(data string) string {
	h := fnv.New32a()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum32())
}
