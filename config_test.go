package gocfg

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string `json:"name"`
	Year    int    `json:"year"`
	Sex     string `json:"sex"`
	Address string `json:"address"`
}

func TestLoad(t *testing.T) {
	// defaultps := Person{Name: "wzh", Year: 22, Sex: "男", Address: "湖北"}
	ps := new(Person)
	// Config.SetDefaultConfig(&defaultps)
	Config.Load(ps)
	fmt.Println(ps)
}

func TestSave(t *testing.T) {
	ps := &Person{Name: "wzh", Year: 20, Sex: "男", Address: "湖北"}
	Config.Save(ps)
}
