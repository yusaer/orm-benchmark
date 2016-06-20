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

左から

* ORM名
* 総実行時間（少ないほど良い）
* １回の実行にかかった時間（少ないほど良い）
* 実行ごとに割り当てられたメモリのサイズ（少ないほど良い）
* １回の実行でメモリアロケーション（メモリ割り当て/配分）が行われた回数（少ないほど良い）

#### First time

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
 
#### Second time
 
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

### Another environment.

* Mac OS X Yosemite 10.10.5
* 2.5 GHz Intel Core i5
* 16GB 1600MHz DDR3

#### First time

```
 40000 times - Insert
       raw:    31.12s       778024 ns/op     568 B/op     25 allocs/op
        pg:    35.09s       877176 ns/op     746 B/op     17 allocs/op
      xorm:    47.64s      1190920 ns/op    4050 B/op    116 allocs/op
 beego_orm:    50.80s      1270017 ns/op    2713 B/op     78 allocs/op
      gorm:    84.36s      2108979 ns/op    8125 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    41.00s      4099767 ns/op   15164 B/op    317 allocs/op
       raw:    43.40s      4339534 ns/op  115809 B/op   1529 allocs/op
 beego_orm:    50.50s      5049734 ns/op  176684 B/op   3264 allocs/op
      xorm:    65.46s      6545607 ns/op 2373346 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    11.56s       289081 ns/op     568 B/op     28 allocs/op
        pg:    48.30s      1207478 ns/op     608 B/op     11 allocs/op
      xorm:    49.56s      1238995 ns/op    4178 B/op    138 allocs/op
 beego_orm:    66.21s      1655274 ns/op    2281 B/op     71 allocs/op
      gorm:    84.11s      2102763 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    25.65s       320640 ns/op     896 B/op     29 allocs/op
        pg:    30.32s       378952 ns/op     960 B/op     38 allocs/op
 beego_orm:    65.18s       814704 ns/op    3081 B/op    108 allocs/op
      xorm:    75.87s       948316 ns/op    7813 B/op    241 allocs/op
      gorm:    75.89s       948676 ns/op    8196 B/op    194 allocs/op

 40000 times - MultiRead limit 100
      xorm:    32.75s       818868 ns/op    3265 B/op     82 allocs/op
       raw:    55.66s      1391485 ns/op   36561 B/op   1509 allocs/op
        pg:    64.12s      1603056 ns/op   25394 B/op    942 allocs/op
 beego_orm:    85.69s      2142218 ns/op   91159 B/op   4601 allocs/op
      gorm:    262.57s      6564279 ns/op  628239 B/op  17019 allocs/op
```

#### Second time

```
 40000 times - Insert
       raw:    28.57s       714132 ns/op     568 B/op     25 allocs/op
        pg:    34.95s       873690 ns/op     680 B/op     17 allocs/op
      xorm:    45.54s      1138420 ns/op    4050 B/op    116 allocs/op
 beego_orm:    60.91s      1522682 ns/op    2713 B/op     78 allocs/op
      gorm:    77.55s      1938706 ns/op    8190 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
       raw:    40.89s      4088782 ns/op  115809 B/op   1529 allocs/op
        pg:    42.50s      4249919 ns/op   15164 B/op    317 allocs/op
 beego_orm:    48.39s      4839005 ns/op  176683 B/op   3264 allocs/op
      xorm:    64.71s      6470947 ns/op 2373252 B/op   8094 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    10.63s       265667 ns/op     568 B/op     28 allocs/op
        pg:    51.84s      1295906 ns/op     608 B/op     11 allocs/op
      xorm:    64.63s      1615842 ns/op    4177 B/op    138 allocs/op
 beego_orm:    64.95s      1623816 ns/op    2281 B/op     71 allocs/op
      gorm:    86.03s      2150721 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    22.54s       281776 ns/op     896 B/op     29 allocs/op
        pg:    37.49s       468683 ns/op     960 B/op     38 allocs/op
      xorm:    84.79s      1059842 ns/op    7812 B/op    241 allocs/op
 beego_orm:    86.27s      1078420 ns/op    3081 B/op    108 allocs/op
      gorm:    87.23s      1090387 ns/op    8196 B/op    194 allocs/op

 40000 times - MultiRead limit 100
      xorm:    38.86s       971585 ns/op    3265 B/op     82 allocs/op
       raw:    47.20s      1179997 ns/op   36561 B/op   1509 allocs/op
        pg:    68.46s      1711417 ns/op   25394 B/op    942 allocs/op
 beego_orm:    76.58s      1914539 ns/op   91157 B/op   4601 allocs/op
      gorm:    251.53s      6288155 ns/op  628245 B/op  17019 allocs/op
```
