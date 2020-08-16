package main

import (
	"encoding/json"
	"fmt"

	"github.com/reddtsai/goreddprints/etcd/pkg/model"
	"github.com/reddtsai/goreddprints/etcd/pkg/scm"
)

func main() {
	addr := []string{"localhost:2379", "localhost:12379", "localhost:22379"}
	m, err := scm.NewSCM(addr)
	fmt.Println("scm ...")
	if err != nil {
		fmt.Println(err)
	}
	defer m.Close()
	sql := &model.Sql{
		Addr:   "0.0.0.0:3306",
		Db:     "test",
		User:   "admin",
		Passwd: "1234",
	}
	p, err := json.Marshal(sql)
	if err != nil {
		fmt.Println(err)
	}
	err = m.Put("sql", string(p))
	if err != nil {
		fmt.Println(err)
	}
	buf, err := m.Get("sql")
	fmt.Println(string(buf))
}
