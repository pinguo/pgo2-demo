package controller

import (
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
)

type MysqlController struct {
    pgo2.Controller
}

// curl -v http://127.0.0.1:8000/mysql/exec
// 使用db.Exec/db.ExecContext在主库上执行非查询操作，
// pgo默认使用10s超时，可通过context来自定义超时。
func (m *MysqlController) ActionExec() {
    // 获取db的上下文适配对象
    db := m.GetObj(adapter.NewDb()).(*adapter.Db)

    // 执行插入操作
    age := (time.Now().Nanosecond() / 1000) % 100
    name := fmt.Sprintf("name%d", age)
    query := "INSERT INTO test1 (name, age) VALUES (?, ?)"

    res := db.Exec(query, name, age)
    if res == nil {
        panic("exec insert failed, " + query)
    }

    lastId, _ := res.LastInsertId()
    numRow, _ := res.RowsAffected()
    fmt.Println("insert: ", lastId, numRow)

    // 执行修改操作
    query = "UPDATE test1 SET age=age+1 WHERE id=?"
    res = db.Exec(query, lastId)
    if res == nil {
        panic("exec update failed, " + query)
    }

    lastId1, _ := res.LastInsertId()
    numRow1, _ := res.RowsAffected()
    fmt.Println("update: ", lastId1, numRow1)

    // 执行删除操作
    query = "DELETE FROM test1 WHERE id=?"
    res = db.Exec(query, lastId)
    if res == nil {
        panic("exec delete failed, " + query)
    }

    lastId2, _ := res.LastInsertId()
    numRow2, _ := res.RowsAffected()
    fmt.Println("delete: ", lastId2, numRow2)
}

// curl -v http://127.0.0.1:8000/mysql/query
// 使用db.Query/QueryOne/QueryContext/QueryOneContext来查询数据，
// 若当前db对象未执行过任何操作，则使用从库进行查询，否则复用上一次获取的db连接。
func (m *MysqlController) ActionQuery() {

    // 对象池获取db的上下文适配对象
    db := m.GetObjBox(adapter.DbClass).(*adapter.Db)

    // 执行插入操作
    age := (time.Now().Nanosecond() / 1000) % 100
    name := fmt.Sprintf("name%d", age)
    query := "INSERT INTO test1 (name, age) VALUES (?, ?)"

    db.Exec(query, name, age)

    // 查询单条数据
    id, name, age := 0, "", 0
    query = "SELECT id, name, age FROM test1 WHERE id=?"
    err := db.QueryOne(query, 3).Scan(&id, &name, &age)
    if err != nil {
        fmt.Println("query one failed, " + err.Error())
    } else {
        fmt.Println("query one for id=3, ", id, name, age)
    }

    // 查询多条数据
    query = "SELECT id, name, age FROM test1 WHERE age>?"
    rows := db.Query(query, 10)
    if rows == nil {
        panic("query failed, query: " + query)
    }

    defer rows.Close()
    for rows.Next() {
        id, name, age := 0, "", 0
        if err := rows.Scan(&id, &name, &age); err != nil {
            panic("query scan faild, err: " + err.Error())
        }
        fmt.Println("user: ", id, name, age)
    }
}

// curl -v http://127.0.0.1:8000/mysql/prepare
// 使用db.PrepareSql/db.PrepareContext来执行批量操作，默认查询操作在
// 从库上进行，其余操作在主库上进行，若当前db对象有过其它操作，则查询
// 操作会复用之前的连接。
func (m *MysqlController) ActionPrepare() {
    // 获取db的上下文适配对象
    db := m.GetObj(adapter.NewDb()).(*adapter.Db)

    query := "INSERT INTO test1 (name, age) VALUES (?, ?)"
    stmt := db.PrepareSql(query)
    defer stmt.Close()

    for i := 0; i < 10; i++ {
        name := fmt.Sprintf("name%d", i)
        res := stmt.Exec(name, i)
        if res == nil {
            panic("stmt exec failed, " + query)
        }
    }
}

// curl -v http://127.0.0.1:8000/mysql/transaction
// 使用db.Begin/BeginContext/Commit/Rollback来进行事务操作
func (m *MysqlController) ActionTransaction() {
    // 获取db的上下文适配对象
    db := m.GetObj(adapter.NewDb()).(*adapter.Db)

    tx := db.Begin()
    defer func() {
        if err:=recover();err!=nil{
            tx.Rollback()
        }
    }()
    tx.Exec("INSERT INTO test1 (name, age) VALUES (?, ?)", "name1", 1)
    tx.Exec("UPDATE test1 SET age=age+1 WHERE id=?", 1)
    if tx.Commit() == false{
        tx.Rollback()
    }else {
        fmt.Println("Transaction ok")
    }

}
