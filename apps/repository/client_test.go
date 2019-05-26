package repository

import (
	"fmt"
	"testing"
)

func TestClientRepository_Index(t *testing.T) {
	clients, _ := ClientIndex()
	for _, element := range clients {
		fmt.Println(element.Name)
	}
}
