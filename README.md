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
        raw:    17.52s       438119 ns/op     568 B/op     25 allocs/op
         pg:    22.98s       574414 ns/op     745 B/op     17 allocs/op
       xorm:    37.30s       932480 ns/op    4050 B/op    116 allocs/op
  beego_orm:    38.85s       971251 ns/op    2713 B/op     78 allocs/op
       gorm:    53.69s      1342298 ns/op    8124 B/op    189 allocs/op
 
  10000 times - MultiInsert 100 row
        raw:    24.67s      2466865 ns/op  115807 B/op   1529 allocs/op
  beego_orm:    28.90s      2890410 ns/op  176681 B/op   3264 allocs/op
       xorm:    41.46s      4145856 ns/op 2373195 B/op   8093 allocs/op
         pg:     Not support multi insert
       gorm:     Not support multi insert
 
  40000 times - Update
        raw:     8.44s       211057 ns/op     568 B/op     28 allocs/op
         pg:    21.82s       545607 ns/op     608 B/op     11 allocs/op
  beego_orm:    38.99s       974828 ns/op    2281 B/op     71 allocs/op
       xorm:    42.24s      1055920 ns/op    4177 B/op    138 allocs/op
       gorm:    63.87s      1596629 ns/op    8252 B/op    209 allocs/op
 
  80000 times - Read
       gorm:     0.08s          955 ns/op     896 B/op      6 allocs/op
        raw:    20.15s       251837 ns/op     896 B/op     29 allocs/op
       xorm:    49.50s       618774 ns/op    7747 B/op    240 allocs/op
         pg:    50.19s       627314 ns/op     960 B/op     38 allocs/op
  beego_orm:    113.16s      1414474 ns/op    3081 B/op    108 allocs/op
 
  40000 times - MultiRead limit 100
        raw:    24.67s       616858 ns/op   36561 B/op   1509 allocs/op
         pg:    30.17s       754234 ns/op   25393 B/op    942 allocs/op
       xorm:    43.82s      1095545 ns/op    3169 B/op     80 allocs/op
  beego_orm:    49.98s      1249520 ns/op   91153 B/op   4601 allocs/op
       gorm:    84.87s      2121732 ns/op  314110 B/op   8509 allocs/op
```
 
#### Sample 2
 
```
 Reports: 
 
  40000 times - Insert
        raw:    16.58s       414555 ns/op     568 B/op     25 allocs/op
         pg:    22.18s       554521 ns/op     745 B/op     17 allocs/op
       xorm:    38.71s       967634 ns/op    4050 B/op    116 allocs/op
       gorm:    44.30s      1107443 ns/op    8124 B/op    189 allocs/op
  beego_orm:    46.71s      1167843 ns/op    2713 B/op     78 allocs/op
 
  10000 times - MultiInsert 100 row
        raw:    24.60s      2459817 ns/op  115808 B/op   1529 allocs/op
  beego_orm:    32.41s      3240596 ns/op  176672 B/op   3264 allocs/op
       xorm:    44.12s      4412202 ns/op 2373297 B/op   8095 allocs/op
         pg:     Not support multi insert
       gorm:     Not support multi insert
 
  40000 times - Update
        raw:     6.60s       164943 ns/op     568 B/op     28 allocs/op
         pg:    20.56s       514111 ns/op     608 B/op     11 allocs/op
       xorm:    37.94s       948610 ns/op    4178 B/op    138 allocs/op
  beego_orm:    41.45s      1036365 ns/op    2281 B/op     71 allocs/op
       gorm:    58.51s      1462819 ns/op    8252 B/op    209 allocs/op
 
  80000 times - Read
       gorm:     0.08s          947 ns/op     896 B/op      6 allocs/op
        raw:    16.79s       209922 ns/op     896 B/op     29 allocs/op
       xorm:    45.78s       572294 ns/op    7747 B/op    240 allocs/op
         pg:    47.14s       589196 ns/op     960 B/op     38 allocs/op
  beego_orm:    113.85s      1423104 ns/op    3081 B/op    108 allocs/op
 
  40000 times - MultiRead limit 100
        raw:    28.18s       704543 ns/op   36561 B/op   1509 allocs/op
         pg:    30.35s       758747 ns/op   25393 B/op    942 allocs/op
  beego_orm:    47.52s      1187878 ns/op   91146 B/op   4601 allocs/op
       xorm:    53.10s      1327411 ns/op    3169 B/op     80 allocs/op
       gorm:    83.21s      2080268 ns/op  314110 B/op   8509 allocs/op
```

#### Sample3

```
Reports: 

 40000 times - Insert
       raw:    16.56s       414094 ns/op     568 B/op     25 allocs/op
        pg:    21.34s       533419 ns/op     745 B/op     17 allocs/op
 beego_orm:    41.07s      1026854 ns/op    2713 B/op     78 allocs/op
      xorm:    43.96s      1099001 ns/op    4050 B/op    116 allocs/op
      gorm:    45.62s      1140381 ns/op    8124 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
       raw:    24.59s      2458837 ns/op  115807 B/op   1529 allocs/op
 beego_orm:    28.96s      2895750 ns/op  176681 B/op   3264 allocs/op
      xorm:    43.66s      4365794 ns/op 2373306 B/op   8095 allocs/op
        pg:     Not support multi insert
      gorm:     Not support multi insert

 40000 times - Update
       raw:     7.44s       185891 ns/op     568 B/op     28 allocs/op
        pg:    22.08s       552122 ns/op     608 B/op     11 allocs/op
 beego_orm:    38.61s       965344 ns/op    2281 B/op     71 allocs/op
      xorm:    40.71s      1017843 ns/op    4177 B/op    138 allocs/op
      gorm:    62.20s      1555081 ns/op    8252 B/op    209 allocs/op

 80000 times - Read
      gorm:     0.08s          973 ns/op     896 B/op      6 allocs/op
       raw:    17.66s       220694 ns/op     896 B/op     29 allocs/op
        pg:    42.99s       537392 ns/op     960 B/op     38 allocs/op
      xorm:    48.00s       600047 ns/op    7747 B/op    240 allocs/op
 beego_orm:    108.05s      1350636 ns/op    3081 B/op    108 allocs/op

 40000 times - MultiRead limit 100
       raw:    25.53s       638330 ns/op   36561 B/op   1509 allocs/op
        pg:    31.40s       785118 ns/op   25393 B/op    942 allocs/op
 beego_orm:    48.31s      1207847 ns/op   91153 B/op   4601 allocs/op
      xorm:    52.77s      1319151 ns/op    3169 B/op     80 allocs/op
      gorm:    84.07s      2101702 ns/op  314110 B/op   8509 allocs/op
```