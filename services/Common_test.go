package services

import (
	"fmt"
	"testing"
)

func TestAddRootTopic(t *testing.T) {
	p := AddRootTopic("sadasdasd", "/")
	fmt.Println(p)
}
