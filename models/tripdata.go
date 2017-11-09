package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
	patterns "goyo.in/gpstracker/patterns"
	"goyo.in/gpstracker/socketios"
)

type (
	// Vehicles represents the structure of our resource
	Vehicles struct {
		ID       bson.ObjectId `bson:"_id,omitempty" json:"-"`
		Sertm    time.Time     `bson:"sertm" json:"sertm"`
		Loctm    time.Time     `bson:"loctm" json:"loctm"`
		Loc      []float64     `bson:"loc" json:"loc"`
		Speed    int           `bson:"speed" json:"speed"`
		Bearing  int           `bson:"bearing" json:"bearing"`
		Appvr    string        `bson:"appvr" json:"appvr"`
		UID      string        `bson:"uid" json:"uid"`
		Btr      int           `bson:"btr" json:"btr"`
		Alwspeed int           `bson:"alwspeed" json:"alwspeed"`
		Flag     string        `bson:"flag" json:"flag"`
		Accr     int           `bson:"accr" json:"accr"`
		Alt      int           `bson:"alt" json:"alt"`
		Gpstm    string        `bson:"gpstm" json:"gpstm"`
		Actvt    string        `bson:"actvt" json:"actvt"`
		Acttm    string        `bson:"acttm" json:"acttm"`
		Vhid     string        `bson:"vhid" json:"vhid"`
		Sat      int           `bson:"sat" json:"sat"`
		Oe       int           `bson:"oe" json:"oe"`
		Gp       int           `bson:"gp" json:"gp"`
		Alm      string        `bson:"alm" json:"alm"`
		Chrg     int           `bson:"chrg" json:"chrg"`
		Acc      int           `bson:"acc" json:"acc"`
		Gsmsig   int           `bson:"gsmsig" json:"gsmsig"`
	}
)

// type GeoJson struct {
// 	Type        string    `json:"-"`
// 	Coordinates []float64 `json:"coordinates"`
// }

type (
	// ParamsTripdata represents the structure of our resource
	ParamsTripdata struct {
		Vhids string
	}
)

func init() {}

//GetLastStatus of vehicles
func GetLastStatus(trpparams ParamsTripdata) (ret []Vehicles, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var result []Vehicles
	c := col(_sn, patterns.ColVhcls)

	_vh := strings.Split(trpparams.Vhids, ",")
	err = c.Find(bson.M{"vhid": bson.M{"$in": _vh}}).All(&result)
	// if err != nil {
	// 	panic(err)
	// }
	return result, err
}

//UpdateData
func UpdateData(d bson.M, vhid string, f string) (err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, patterns.ColVhcls)
	_, err = c.UpsertId(bson.M{"vhid": vhid}, bson.M{"$set": d})

	// send in history
	if (f == "reg") || (f == "loc") {
		// insert in history table
		ch := col(_sn, patterns.ColVhtrps)
		err = ch.Insert(d)

	}

	// send to socket server
	go func() {

		if f == "reg" {
			delete(d, "appvr")
			delete(d, "imei")
			delete(d, "acttm")
		} else if f == "loc" {
			delete(d, "appvr")
			delete(d, "imei")
		} else if f == "hrt" {
			delete(d, "appvr")
			delete(d, "imei")
			delete(d, "lng")
		}
		socket := socketios.GetSocketIO()
		socket.BroadcastTo(vhid, "msgd", d)

	}()

	return err
}

//GetLastStatus of vehicles
// func GetLastStatus(trpparams ParamsTripdata) (ret []Vehicles, err error) {
// 	_sn := getDBSession().Copy()
// 	defer _sn.Close()

// 	var result []Vehicles
// 	c := col(_sn, patterns.ColVhcls)

// 	_vh := strings.Split(trpparams.Vhids, ",")
// 	err = c.Find(bson.M{"vhid": bson.M{"$in": _vh}}).All(&result)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	return result, err
// }
