# ORM Benchmark

### Environment

* go version go1.6 linux/amd64

### PostgreSQL

* PostgreSQL 9.5 for Linux on x86_64

### ORMs

All package run in no-cache mode.

* [dbr](https://github.com/gocraft/dbr) (in preparation)
* [genmai](https://github.com/naoina/genmai) (in preparation)
* [gorm](https://github.com/jinzhu/gorm)
* [gorp](https://github.com/go-gorp/gorp) (in preparation)
* [pg](https://github.com/go-pg/pg)
* [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
* [sqlx](https://github.com/jmoiron/sqlx) (in preparation)
* [xorm](https://github.com/xormplus/xorm)
	
### Run

```go
go get github.com/milkpod29/orm-benchmark
# build
go install
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=xorm
```

### Reports

#### Sample 1
 
```
 40000 times - Insert
        pg:    28.11s       702699 ns/op     680 B/op     17 allocs/op
       raw:    28.19s       704850 ns/op     568 B/op     25 allocs/op
      xorm:    52.28s      1307026 ns/op    4050 B/op    116 allocs/op
 beego_orm:    63.99s      1599813 ns/op    2713 B/op     78 allocs/op
      gorm:    78.97s      1974368 ns/op    8189 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    27.27s      2726560 ns/op   15164 B/op    317 allocs/op
 beego_orm:    33.12s      3312423 ns/op  176680 B/op   3264 allocs/op
       raw:    33.84s      3383613 ns/op  115808 B/op   1529 allocs/op
      xorm:    50.42s      5042414 ns/op 2373289 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    16.10s       402576 ns/op     568 B/op     28 allocs/op
        pg:    33.24s       830982 ns/op     608 B/op     11 allocs/op
      xorm:    58.68s      1466943 ns/op    4177 B/op    138 allocs/op
 beego_orm:    68.03s      1700812 ns/op    2281 B/op     71 allocs/op
      gorm:    98.61s      2465242 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    41.77s       522184 ns/op     896 B/op     29 allocs/op
      gorm:    59.69s       746157 ns/op    8195 B/op    194 allocs/op
      xorm:    67.30s       841274 ns/op    7747 B/op    240 allocs/op
        pg:    98.98s      1237223 ns/op     960 B/op     38 allocs/op
 beego_orm:    151.66s      1895795 ns/op    3081 B/op    108 allocs/op

 40000 times - MultiRead limit 100
       raw:    35.98s       899513 ns/op   36561 B/op   1509 allocs/op
        pg:    38.77s       969259 ns/op   25393 B/op    942 allocs/op
      xorm:    55.26s      1381464 ns/op    3169 B/op     80 allocs/op
 beego_orm:    75.09s      1877161 ns/op   91153 B/op   4601 allocs/op
      gorm:    197.95s      4948706 ns/op  628190 B/op  17018 allocs/op
```
 
#### Sample 2
 
```
 40000 times - Insert
       raw:    26.54s       663378 ns/op     568 B/op     25 allocs/op
        pg:    27.96s       699040 ns/op     745 B/op     17 allocs/op
      xorm:    52.40s      1310069 ns/op    4050 B/op    116 allocs/op
 beego_orm:    64.94s      1623420 ns/op    2713 B/op     78 allocs/op
      gorm:    96.02s      2400406 ns/op    8124 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    25.57s      2556601 ns/op   15164 B/op    317 allocs/op
       raw:    26.32s      2632441 ns/op  115808 B/op   1529 allocs/op
 beego_orm:    30.97s      3096774 ns/op  176681 B/op   3264 allocs/op
      xorm:    47.30s      4730103 ns/op 2373300 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    24.04s       601121 ns/op     568 B/op     28 allocs/op
        pg:    26.71s       667631 ns/op     608 B/op     11 allocs/op
      xorm:    45.58s      1139446 ns/op    4177 B/op    138 allocs/op
 beego_orm:    48.22s      1205619 ns/op    2281 B/op     71 allocs/op
      gorm:    96.84s      2420984 ns/op    8252 B/op    209 allocs/op

 80000 times - Read
       raw:    38.05s       475586 ns/op     896 B/op     29 allocs/op
      gorm:    61.08s       763527 ns/op    8195 B/op    194 allocs/op
      xorm:    62.42s       780222 ns/op    7748 B/op    240 allocs/op
        pg:    81.14s      1014239 ns/op     960 B/op     38 allocs/op
 beego_orm:    97.41s      1217656 ns/op    3081 B/op    108 allocs/op

 40000 times - MultiRead limit 100
       raw:    31.86s       796471 ns/op   36561 B/op   1509 allocs/op
        pg:    33.30s       832391 ns/op   25393 B/op    942 allocs/op
      xorm:    57.09s      1427295 ns/op    3169 B/op     80 allocs/op
 beego_orm:    63.73s      1593203 ns/op   91154 B/op   4601 allocs/op
      gorm:    209.57s      5239374 ns/op  628185 B/op  17018 allocs/op
```
