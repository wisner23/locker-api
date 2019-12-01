package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserTable_20191130_171032 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20191130_171032{}
	m.Created = "20191130_171032"

	migration.Register("CreateUserTable_20191130_171032", m)
}

// Run the migrations
func (m *CreateUserTable_20191130_171032) Up() {
	m.SQL(`CREATE TABLE users (
				user_id              SERIAL PRIMARY KEY,
				email                VARCHAR(128) NOT NULL,
				username             VARCHAR(32) NULL,
				phone                VARCHAR(15) NULL
	  		);
	  `)

}

// Reverse the migrations
func (m *CreateUserTable_20191130_171032) Down() {
	m.SQL("DROP TABLE users")

}
