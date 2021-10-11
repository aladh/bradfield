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
BenchmarkMetrics/Average_payment-4           198           6018441 ns/op (4.7x faster)
BenchmarkMetrics/Payment_stddev-4            135           8814534 ns/op (6.2x faster)
PASS
ok      example.com     6.710s
```

### Store ages in an array
```
BenchmarkMetrics/Average_age-4              2132            571574 ns/op (6.2x faster)
BenchmarkMetrics/Average_payment-4           194           6032774 ns/op
BenchmarkMetrics/Payment_stddev-4            134           8766550 ns/op
PASS
ok      example.com     6.322s
```

### Incrementally compute average age
```
BenchmarkMetrics/Average_age-4             10000            104049 ns/op (5.5x faster)
BenchmarkMetrics/Average_payment-4           199           6032538 ns/op
BenchmarkMetrics/Payment_stddev-4            135           8847441 ns/op
PASS
ok      example.com     6.139s
```

### 2x2 loop unrolling for average age
```
BenchmarkMetrics/Average_age-4             19341             61916 ns/op (1.7x faster)
BenchmarkMetrics/Average_payment-4           194           6290909 ns/op
BenchmarkMetrics/Payment_stddev-4            126           9053286 ns/op
PASS
ok      example.com     6.980s
```

### Use primitive value instead of payment struct
```
BenchmarkMetrics/Average_age-4             19519             62083 ns/op
BenchmarkMetrics/Average_payment-4           210           5758736 ns/op (1.1x faster)
BenchmarkMetrics/Payment_stddev-4            176           6860588 ns/op (1.3x faster)
PASS
ok      example.com     5.893s
```

### Incrementally compute average payment amount
```
BenchmarkMetrics/Average_age-4             19333             62614 ns/op
BenchmarkMetrics/Average_payment-4          1140           1048805 ns/op (5.5x faster)
BenchmarkMetrics/Payment_stddev-4            572           2108189 ns/op (3.3x faster)
PASS
ok      example.com     4.953s
```

### 2x2 loop unrolling for average payment
```
BenchmarkMetrics/Average_age-4             18614             61803 ns/op
BenchmarkMetrics/Average_payment-4          1875            623110 ns/op (1.7x faster)
BenchmarkMetrics/Payment_stddev-4            675           1672364 ns/op (1.3x faster)
PASS
ok      example.com     4.731s
```

### 2x2 loop unrolling for std dev payment amount
```
BenchmarkMetrics/Average_age-4             19417             62665 ns/op
BenchmarkMetrics/Average_payment-4          1876            625494 ns/op
BenchmarkMetrics/Payment_stddev-4            927           1283776 ns/op (1.3x faster)
PASS
ok      example.com     4.808s
```
