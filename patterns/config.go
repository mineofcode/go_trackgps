package patterns

import (
	"time"

	"gopkg.in/mgo.v2"
)

const (
	dbip   = "localhost"
	dbport = "27017"
	//Dbname Mongodb database name
	Dbname = "goyosch"
	//ColVhcls 'vhcls' Collection
	ColVhcls = "vhcls"
	//ColVhtrps 'vhtrps' Collection
	ColVhtrps = "vhtrps"

	urldb        = dbip + ":" + dbport
	authUserName = ""
	authPassword = ""
)

var mongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{urldb},
	Timeout:  60 * time.Second,
	Database: Dbname,
	Username: authUserName,
	Password: authPassword,
}
