package benchs

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

var dbmap *gorp.DbMap

func init() {
	st := NewSuite("gorp")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GorpInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GorpInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GorpUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GorpRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GorpReadSlice)

		db, err := sql.Open("postgres", ORM_SOURCE)
		checkErr(err)
		d := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
		dbmap = d
	}
}

func GorpInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})
	for i := 0; i < b.N; i++ {
		m.Id = 0
		d := dbmap.Insert(&m)
		if d.Error != nil {
			fmt.Println(d.Error())
			b.FailNow()
		}
	}
}

func GorpInsertMulti(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func GorpUpdate(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func GorpRead(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func GorpReadSlice(b *B) {
	panic(fmt.Errorf("in preparation"))
}
