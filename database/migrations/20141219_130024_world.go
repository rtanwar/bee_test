package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type World_20141219_130024 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &World_20141219_130024{}
	m.Created = "20141219_130024"
	migration.Register("World_20141219_130024", m)
}

// Run the migrations
func (m *World_20141219_130024) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE world(`i_d` int(11) DEFAULT NULL)")
}

// Reverse the migrations
func (m *World_20141219_130024) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `world`")
}
