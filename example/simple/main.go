package main

import (
	"fmt"
	"github.com/Wall-js/nebula-core/project"
)

func main() {
	p := project.NewProject()
	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
