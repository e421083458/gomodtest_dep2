package gomodtest_dep2

import (
	"fmt"
	"github.com/e421083458/gomodtest_base"
)

func NewIntCollection() {
	fmt.Println("gomodtest_dep")
	gomodtest_base.NewIntCollection()
}
