package benchs

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var sqlxdb *sqlx.DB

func init() {
	st := NewSuite("sqlx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, SqlxInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, SqlxInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, SqlxUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, SqlxRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, SqlxReadSlice)

		db, err := sqlx.Connect("postgres", ORM_SOURCE);
		checkErr(err)
		sqlxdb = db
	}
}

func SqlxInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})
	for i := 0; i < b.N; i++ {
		sqlxdb.MustExec(`INSERT INTO model (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7)`, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	}
}

func SqlxInsertMulti(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func SqlxUpdate(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func SqlxRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		sqlxdb.MustExec(`INSERT INTO model (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7)`, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	})
	for i := 0; i < b.N; i++ {
		m := []Model{}
		if err := sqlxdb.Select(&m, "SELECT * FROM model"); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxReadSlice(b *B) {
	panic(fmt.Errorf("in preparation"))
}
