package controller


import (
	"context"
	"fmt"

	"github.com/pinguo/pgo2"
	"github.com/pinguo/pgo2/adapter"
	"github.com/pinguo/pgo2/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongodb struct {
pgo2.Controller
}

// curl -v http://127.0.0.1:8000/mongodb/insert
func (m *Mongodb) ActionInsert() {
// 获取mongo的上下文适配对象
mongo := m.GetObj(adapter.NewMongodb("test", "test")).(*adapter.Mongodb)

// 通过map插入
doc1 := map[string]interface{}{"f1": "val1", "f2": true, "f3": 99}
ret,err := mongo.InsertOne(doc1)
fmt.Println("insert one doc1 id:",ret.InsertedID,"err", err)
// 通过bson.M插入

doc2 := bson.M{"f1": "val2", "f2": false, "f3": 10}
ret,err = mongo.InsertOne(doc2)
fmt.Println("insert one doc2 id:",ret.InsertedID,"err", err)

// 通过struct插入
doc3 := struct {
F1 string `bson:"f1"`
F2 bool   `bson:"f2"`
F3 int    `bson:"f3" validation:"gte:3"` // 自动验证是否大于等于3
}{"val3", false, 6}
ret,err = mongo.InsertOne(doc3)
fmt.Println("insert one  doc3 id:",ret.InsertedID,"err", err)

// 批量插入
docs := []interface{}{
bson.M{"f1": "val4", "f2": true, "f3": 7},
bson.M{"f1": "val5", "f2": false, "f3": 8},
bson.M{"f1": "val6", "f2": true, "f3": 9},
}
rets,err := mongo.InsertMany(docs)
fmt.Println("insert all docs ids:",rets.InsertedIDs,"err", err)
}

// curl -v http://127.0.0.1:8000/mongodb/update
func (m *Mongodb) ActionUpdate() {
// 对象池获取mongo的上下文适配对象
mongo := m.GetObjBox(adapter.MongodbClass, "test", "test").(*adapter.Mongodb)

// 更新单个文档
query := bson.M{"f1": "val1"}
update := bson.M{"$inc": bson.M{"f3": 2}}
err := mongo.UpdateOne(query, update)
fmt.Println("update one f1==val1:", err)

// 更新多个文档
query = bson.M{"f3": bson.M{"$gte": 7}}
update = bson.M{"$set": bson.M{"f2": false}}
result,err := mongo.UpdateAll(query, update)
fmt.Println("update all f3>=7: update count :",result.ModifiedCount,"err", err)

// 更新或插入
query = bson.M{"f1": "val10"}
update = bson.M{"f3": 2}
	result,err = mongo.Upsert(query, update)
fmt.Println("update or insert f1==val10:",util.ToString(result),"err", err)
}

// curl -v http://127.0.0.1:8000/mongodb/query
func (m *Mongodb) ActionQuery() {
// 获取mongo的上下文适配对象
mongo := m.GetObj(adapter.NewMongodb("test", "test")).(*adapter.Mongodb)

// 查询单个文档(未指定结果类型，结果为bson.M)
var v1 interface{}
err := mongo.Find(bson.M{"f1": "val1"}).Select(bson.M{"f1": 1}).One( &v1)
fmt.Println("query one f1==val1:", v1, err)

// 查询单个文档(结果类型为map)
var v2 map[string]interface{}
err = mongo.Find(bson.M{"f1": "val2"}).One(&v2)
fmt.Println("query one f1==val2:", v2, err)

// 查询单个文档(结果类型为struct)
var v3 struct {
Id primitive.ObjectID `bson:"_id"`
F1 string        `bson:"f1"`
F2 bool          `bson:"f2"`
F3 int           `bson:"f3"`
}
err = mongo.Find(bson.M{"f1": "val3"}).One(&v3)
fmt.Println("query one f1==val3:", v3, err)

// 查询多个文档(指定结果为map)
var docs []map[string]interface{}
err = mongo.Find(bson.M{"f3": bson.M{"$gte": 6}}).Sort("-_id").All( &docs)
fmt.Println("query all f3>=6:", docs, err)
}

// curl -v http://127.0.0.1:8000/mongodb/transaction
// - version of mongoDB server >= v4.0
// - Topology of mongoDB server is not Single
func (m *Mongodb) ActionTransaction() {
	// 获取mongo的上下文适配对象
	mongo := m.GetObj(adapter.NewMongodb("test", "test")).(*adapter.Mongodb)
	ret,err:=mongo.DoTransaction(context.Background(), func(sessCtx context.Context) (i interface{}, e error) {
		// 通过map插入
		doc1 := map[string]interface{}{"f1": "val1", "f2": true, "f3": 99}
		ret,err := mongo.InsertOneCtx(sessCtx,doc1)
		if err !=nil {
			return ret,err
		}
		fmt.Println("insert one doc1 id:",ret.InsertedID,"err", err)
		// 通过bson.M插入

		doc2 := bson.M{"f1": "val2", "f2": false, "f3": 10}
		ret,err = mongo.InsertOneCtx(sessCtx,doc2)
		fmt.Println("insert one doc2 id:",ret.InsertedID,"err", err)
		return nil,nil
	})

	fmt.Println("itransaction:",ret,"err", err)
}