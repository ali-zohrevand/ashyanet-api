package main

import (
	"fmt"
	"github.com/casbin/casbin"
	scas "github.com/qiangmzsx/string-adapter"
	"time"
)

func main() {
	KeyMatchRbac()
	//StringRbac()
	//UserRbac()
	//Read()
}

func KeyMatchRbac() {
	conf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub)  && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
	line := `
p, alice, /alice_data/*, (GET)|(POST)
p, alice, /alice_data/resource1, POST
p, data_group_admin, /admin/*, POST
p, ali, /*, GET
g, alice, data_group_admin
g, as, data_group_admin
g, asd, admin
`
	sa := scas.NewAdapter(line)
	e := casbin.NewEnforcer(casbin.NewModel(conf), sa)
	sub := "ali"
	obj := "/alice_data/logidf/dfn"
	act := "GET"
	if e.Enforce(sub, obj, act) == true {
		fmt.Println("**YES**")
	} else {
		fmt.Println("--NO--")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("..............")
	m := e.GetPolicy()
	fmt.Println("m ", m)
	fmt.Println("..............")
	fmt.Println(e.GetRolesForUser("alice"))
	s := e.GetGroupingPolicy()
	fmt.Println(s)
}

func Read() {
	conf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act`
	line := `
p, alice, data1, read
p, bob, data2, write
`
	sa := scas.NewAdapter(line)
	e := casbin.NewEnforcer(casbin.NewModel(conf), sa)
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	if e.Enforce(sub, obj, act) == true {
		fmt.Println("**YES**")
	} else {
		fmt.Println("--NO--")
	}
}

func StringRbac() {
	conf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	line := `
p, alice, data1, read
p, data_group_admin, data3, read
p, data_group_admin, data3, write
g, alice, data_group_admin
`
	sa := scas.NewAdapter(line)
	e := casbin.NewEnforcer(casbin.NewModel(conf), sa)
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.
	if e.Enforce(sub, obj, act) == true {
		fmt.Println("**YES**")
	} else {
		fmt.Println("--NO--")
	}
}
