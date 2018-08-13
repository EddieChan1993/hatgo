package models

import (
	"log"
	"hatgo/pkg/link"
)

type Test struct {
	Id    int    `json:"id"`
	Type  int    `json:"type"`
	Name  string `json:"name"`
	CTime int    `json:"c_time"`
	UTime int    `json:"u_time"`
	DTime int    `json:"d_time"`
}

func AllTest() []Test {
	var test []Test
	err := link.Db.Find(&test)
	if err != nil {
		log.Println(err)
		return nil
	}

	return test
}