module example.com/main

go 1.22.1

replace example.com/dbpart => ../dbpart

replace example.com/model => ../model

require example.com/dbpart v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.9 // indirect
