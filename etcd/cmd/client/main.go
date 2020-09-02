package main

import (
	"fmt"

	"github.com/reddtsai/goreddprints/etcd/pkg/scm"
)

func main() {
	m, err := scm.NewSCM("localhost:2379", "localhost:12379", "localhost:22379")
	fmt.Println("scm ...")
	if err != nil {
		fmt.Println(err)
	}
	defer m.Close()
	err = m.Put("/foo", "hello")
	if err != nil {
		fmt.Println(err)
	}
	err = m.Put("/foo/bar", "world")
	if err != nil {
		fmt.Println(err)
	}
	buf, err := m.Get("/foo")
	for k, v := range buf {
		fmt.Println(k, "-", string(v))
	}
}
