package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

//base struct of Oplog
type Oplog struct {
	Ts bson.MongoTimestamp `bson:"ts"`
	H  int                 `bson:"h"`
	V  int                 `bson:"v"`
	Op string              `bson:"op"`
	Ns string              `bson:"ns"`
}

//struct of Oplog insert object
type OplogInsert struct {
	Oplog `bson:",inline"`
	O     map[string]interface{} `bson:"o"`
}

//struct of Oplog update object
type OplogUpdate struct {
	Oplog `bson:",inline"`
	O2    map[string]bson.ObjectId `bson:"o2"`
	O     map[string]interface{}   `bson:"o`
}

type OplogDelete struct {
	Oplog `bson:",inline"`
	B     bool                     `bson:"b"`
	O     map[string]bson.ObjectId `bson:"o"`
}

//return all inserted oplog objects
func (c *Client) GetOplogsInsert() []OplogInsert {
	collection := c.Conn.DB("local").C("oplog.rs")

	var logs []OplogInsert
	collection.Find(bson.M{"op": "i", "ns": c.Conf.Database + "." + c.Conf.Collection, "ts": bson.M{"$type": 17}}).All(&logs)

	return logs
}

//return all updated oplog objects
func (c *Client) GetOplogsUpdate() []OplogUpdate {
	collection := c.Conn.DB("local").C("oplog.rs")

	var logs []OplogUpdate
	collection.Find(bson.M{"op": "u", "ns": c.Conf.Database + "." + c.Conf.Collection, "ts": bson.M{"$type": 17}}).All(&logs)

	return logs
}

//return all deleted oplog objects
func (c *Client) GetOplogsDelete() []OplogDelete {
	collection := c.Conn.DB("local").C("oplog.rs")

	var logs []OplogDelete
	collection.Find(bson.M{"op": "d", "ns": c.Conf.Database + "." + c.Conf.Collection, "ts": bson.M{"$type": 17}}).All(&logs)

	return logs
}
