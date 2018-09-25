package models

type Test struct {
	Id    int      `json:"id" xorm:"pk autoinc"` //仅当自增Id不为int64及其名称部位Id时使用，已手动标识
	Type  int      `json:"type"`
	Name  string   `json:"name"`
	CTime int      `json:"c_time" xorm:"created"` //存入当前时间戳
	UTime int      `json:"u_time" xorm:"updated"`
	DTime xormDate `json:"d_time"` //存入日期,建议存时间戳，
}
