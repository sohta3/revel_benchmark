# revel_benchmark

* Go: 1.5.2
* Revel: 0.12.0

```
$ revel run revel_benchmark prod

$ wrk -t4 -c100 -d30S --timeout 2000 "http://127.0.0.1:9000/welcome"
```
