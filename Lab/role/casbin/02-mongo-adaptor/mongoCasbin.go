package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
)

func main() {
	// Initialize a MongoDB adapter and use it in a Casbin enforcer:
	// The adapter will use the database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := mongodbadapter.NewAdapter("127.0.0.1:27017/db") // Your MongoDB URL.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := mongodbadapter.NewAdapter("127.0.0.1:27017/abc", true)

	e := casbin.NewEnforcer("Resource/auth_model.conf", a)
	//p, alice, data1, read
	// Load the policy from DB.
	e.LoadPolicy()
	e.AddPolicy("p", "alice", "datal", "read")
	e.SavePolicy()
	e.LoadPolicy()
	e.LoadPolicy()
	e.AddPolicy("p", "admin", "book", "read")
	e.SavePolicy()
	e.LoadPolicy()
	// Check the permission.
	e.Enforce("alice", "datal", "read")
	fmt.Println("............................")
	e.Enforce("p", "admin", "book", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
