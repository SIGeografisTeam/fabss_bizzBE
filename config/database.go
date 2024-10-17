package config

import (
	"plastiqu_co/helper/atdb"
	// "os"
)

var MongoString string = "mongodb+srv://fathyafathazzra:Mongodbatlas12@cluster0.8xvps.mongodb.net/"

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "plastiqu",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)