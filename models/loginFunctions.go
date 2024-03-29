package models

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// login
func loginRequest(object Login) Login {
	var obj Login
	//var err error
	c := accessDB(vhclsCollection)
	err := c.Find(bson.M{"username": object.Username, "password": object.Password}).One(&obj)
	fmt.Printf("\nerror: %s\n\n", err)
	js, __ := json.Marshal(obj)
	verifyErr(__)
	fmt.Printf("\nusuario logueado:\n%s\n\n", js)
	return obj
}
