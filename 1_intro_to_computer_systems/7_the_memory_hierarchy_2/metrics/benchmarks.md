# Benchmarks

### Original
```
BenchmarkMetrics/Average_age-4               334           3537751 ns/op
BenchmarkMetrics/Average_payment-4            46          28203221 ns/op
BenchmarkMetrics/Payment_stddev-4             21          54813337 ns/op
PASS
ok      example.com     6.281s
```

### Store all payments in a single array

```
BenchmarkMetrics/Average_age-4               338           3564686 ns/op
BenchmarkMetrics/Average_payment-4           198           6018441 ns/op (79% faster)
BenchmarkMetrics/Payment_stddev-4            135           8814534 ns/op (84% faster)
PASS
ok      example.com     6.710s
```
