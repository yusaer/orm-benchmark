package benchs

import (
	"fmt"
	"gopkg.in/pg.v4"
)

var pgdb *pg.DB

func init() {
	st := NewSuite("pg")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, PgInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, PgInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, PgUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, PgRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, PgReadSlice)

		pgdb = pg.Connect(&pg.Options{
			Addr:     "localhost",
			User:     "postgres",
			Password: "postgres",
			Database: "test",
		})
	}
}

func PgInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := pgdb.Create(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func PgUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		pgdb.Create(&m)
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Update(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		pgdb.Create(&m)
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Select(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := pgdb.Create(&m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model

		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(100).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
