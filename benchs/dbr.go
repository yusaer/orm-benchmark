package benchs

import (
	"fmt"

	"github.com/gocraft/dbr"
)

var dbrsession *dbr.Session

func init() {
	st := NewSuite("dbr")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, DbrInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, DbrInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, DbrUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, DbrRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, DbrReadSlice)

		conn, _ := dbr.Open("postgres", ORM_SOURCE, nil)
		sess := conn.NewSession(nil)
		dbrsession = sess
	}
}

func DbrInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := dbrsession.InsertInto("model").Columns("name", "title", "fax", "web", "age", "right", "counter").Record(m).Exec(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
//
func DbrInsertMulti(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func DbrUpdate(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func DbrRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := dbrsession.InsertInto("model").Columns("name", "title", "fax", "web", "age", "right", "counter").Record(m).Exec(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		var m []Model
		if _, err := dbrsession.Select("*").From("model").Load(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func DbrReadSlice(b *B) {
	panic(fmt.Errorf("in preparation"))
}
