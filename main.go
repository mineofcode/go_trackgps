package main

import (
	"fmt"

	"goyo.in/gpstracker/network"
	// import "goyo.in/gpstracker/network"
	// import "goyo.in/gpstracker/crc16"

	restservice "goyo.in/gpstracker/conf"
)

func main() {
	//Start TCP server
	en := network.TCPServer{Host: "0.0.0.0", Port: "6969", Timeout: 3000}
	err := en.Open()

	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
	}

	//Start Rest API & Socket.io server
	restservice.RestfulAPIServiceInit("HTTP")

	//compute crc
	// data := crc16.HexStringToCrcHexString("05130100")
	// //fmt.Println(Hex2Bin([]byte("154C")))
	// fmt.Println(data)
	// fmt.Println(hexToBin("154C","%016b"))
	// fmt.Println([]byte("154C"))

	// u := models.User{}

	// query := func(c *mgo.Collection) error {
	//     fn := c.Find(nil).All(&u)
	//     return fn
	// }

	// err := withCollection("trps",query)

	// if err != nil {
	//     // TODO: Do something about the error
	// } else {
	//     fmt.Println("Results All: ", u)
	// }

	//var results []mongomodel.User
	// session := mongo.GetSession();
	// c := session.DB("goyosch").C("trps")
	// err := c.Find(nil).All(&results)
	// if err != nil {
	// 	// TODO: Do something about the error
	// } else {
	// 	fmt.Println("Results All: ", results)
	// }

	// query := func(c *mgo.Collection) error {
	//     fn := c.Find(nil).All(&results)
	//     return fn
	// }
	// withCollection("trps", query)
	// fmt.Println("Results All: ", results)

}

// func withCollection(collection string, s func(*mgo.Collection) error) error {
// 	session := mongo.GetSession()
// 	defer session.Close()
// 	c := session.DB("goyosch").C(collection)
// 	return s(c)
// }
