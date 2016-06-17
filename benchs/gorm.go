package benchs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var gormdb *gorm.DB

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GormInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GormInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GormUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GormRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GormReadSlice)

		conn, err := gorm.Open("postgres", ORM_SOURCE)
		if err != nil {
			fmt.Println(err)
		}
		gormdb = conn
	}
}

func GormInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		gormdb.Create(&m)
	}
}

func GormInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func GormUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gormdb.Create(&m)
	})

	for i := 0; i < b.N; i++ {
		if err := gormdb.Update(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gormdb.Create(&m)
	})

	for i := 0; i < b.N; i++ {
		if err := gormdb.Select(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			gormdb.Create(&m)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model

		if err := gormdb.Where("id > ?", 0).Limit(100).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
