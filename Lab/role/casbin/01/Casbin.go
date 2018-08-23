package main

import (
	"log"

	"github.com/casbin/casbin"
)

func main() {
	//http://casbin.org/editor/
	//https://doc.xuwenliang.com/docs/casbin/188
	//https://zupzup.org/casbin-http-role-auth/
	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	if authEnforcer.Enforce(sub, obj, act) == true {
		// permit alice to read data1
	} else {
		// deny the request, show an error
	}

}
