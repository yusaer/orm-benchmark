package benchs

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/naoina/genmai"
)

var genmaidb *genmai.DB

func init() {
	st := NewSuite("genmai")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GenmaiInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GenmaiInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GenmaiUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GenmaiRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GenmaiReadSlice)

		db, err := genmai.New(&genmai.PostgresDialect{}, ORM_SOURCE)
		checkErr(err)
		genmaidb = db
	}
}

func GenmaiInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = i
		if _, err := genmaidb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GenmaiInsertMulti(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func GenmaiUpdate(b *B) {
	panic(fmt.Errorf("in preparation"))
}

func GenmaiRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := genmaidb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		var results []Model
		if err := genmaidb.Select(&results); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GenmaiReadSlice(b *B) {
	panic(fmt.Errorf("in preparation"))
}
