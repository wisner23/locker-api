package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsernameConstraint_20191201_143739 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsernameConstraint_20191201_143739{}
	m.Created = "20191201_143739"

	migration.Register("CreateUsernameConstraint_20191201_143739", m)
}

// Run the migrations
func (m *CreateUsernameConstraint_20191201_143739) Up() {
	m.SQL("ALTER TABLE users ADD CONSTRAINT username_unique UNIQUE (username);")

}

// Reverse the migrations
func (m *CreateUsernameConstraint_20191201_143739) Down() {
	m.SQL("ALTER TABLE users DROP CONSTRAINT username_unique;")

}
