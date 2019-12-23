# gaeadb_bench - tencent cloud - ubuntu 18.04

    use redis-benchmark for benchmark

## HDD Single Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 1 -p 6378 -t set,get,lrange -n 100000 -q

### bolt

    SET: 403.57 requests per second
    GET: 23562.68 requests per second
    LPUSH (needed to benchmark LRANGE): 309.47 requests per second
    LRANGE_100 (first 100 elements): 951.12 requests per second
    LRANGE_300 (first 300 elements): 324.69 requests per second
    LRANGE_500 (first 450 elements): 217.72 requests per second
    LRANGE_600 (first 600 elements): 162.86 requests per second

### badger

    SET: 283.61 requests per second
    GET: 18928.64 requests per second
    LPUSH (needed to benchmark LRANGE): 278.46 requests per second
    LRANGE_100 (first 100 elements): 792.28 requests per second
    LRANGE_300 (first 300 elements): 279.24 requests per second
    LRANGE_500 (first 450 elements): 189.58 requests per second
    LRANGE_600 (first 600 elements): 144.15 requests per second

### gaeadb

    SET: 662.96 requests per second
    GET: 5804.50 requests per second
    LPUSH (needed to benchmark LRANGE): 10341.26 requests per second
    LRANGE_100 (first 100 elements): 885.50 requests per second
    LRANGE_300 (first 300 elements): 315.52 requests per second
    LRANGE_500 (first 450 elements): 215.28 requests per second
    LRANGE_600 (first 600 elements): 163.86 requests per second

## HDD Multiple Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 100 -p 6378 -t set,get,lrange -n 10000000 -q

### bolt

    SET: 424.44 requests per second
    GET: 83125.52 requests per second
    LPUSH (needed to benchmark LRANGE): 317.22 requests per second
    LRANGE_100 (first 100 elements): 3734.69 requests per second
    LRANGE_300 (first 300 elements): 1302.52 requests per second
    LRANGE_500 (first 450 elements): 891.91 requests per second
    LRANGE_600 (first 600 elements): 705.84 requests per second

### badger

    SET: 6691.09 requests per second
    GET: 81641.98 requests per second
    LPUSH (needed to benchmark LRANGE): 6679.08 requests per second
    LRANGE_100 (first 100 elements): 3689.25 requests per second
    LRANGE_300 (first 300 elements): 1270.22 requests per second
    LRANGE_500 (first 450 elements): 866.00 requests per second
    LRANGE_600 (first 600 elements): 672.38 requests per second

### gaeadb

    SET: 4153.55 requests per second
    GET: 21363.78 requests per second
    LPUSH (needed to benchmark LRANGE): 20776.15 requests per second
    LRANGE_100 (first 100 elements): 3444.59 requests per second
    LRANGE_300 (first 300 elements): 1296.65 requests per second
    LRANGE_500 (first 450 elements): 911.50 requests per second
    LRANGE_600 (first 600 elements): 720.85 requests per second

## SSD Single Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 1 -p 6378 -t set,get,lrange -r 100000 -n 100000 -q

### bolt

    SET: 403.57 requests per second
    GET: 23562.68 requests per second
    LPUSH (needed to benchmark LRANGE): 309.47 requests per second
    LRANGE_100 (first 100 elements): 951.12 requests per second
    LRANGE_300 (first 300 elements): 324.69 requests per second
    LRANGE_500 (first 450 elements): 217.72 requests per second
    LRANGE_600 (first 600 elements): 162.86 requests per second

### badger

    SET: 341.04 requests per second
    GET: 23196.47 requests per second
    LPUSH (needed to benchmark LRANGE): 341.33 requests per second
    LRANGE_100 (first 100 elements): 961.58 requests per second
    LRANGE_300 (first 300 elements): 331.11 requests per second
    LRANGE_500 (first 450 elements): 224.09 requests per second
    LRANGE_600 (first 600 elements): 168.26 requests per second

### gaeadb

    SET: 4 requests per second
    GET: 6538.94 requests per second
    LPUSH (needed to benchmark LRANGE): 11210.76 requests per second
    LRANGE_100 (first 100 elements): 899.44 requests per second
    LRANGE_300 (first 300 elements): 316.47 requests per second
    LRANGE_500 (first 450 elements): 216.01 requests per second
    LRANGE_600 (first 600 elements): 162.18 requests per second

## SSD Multiple Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 100 -p 6378 -t set,get,lrange -r 10000000 -n 10000000 -q

### bolt

    SET: 493.11 requests per second
    GET: 92250.92 requests per second
    LPUSH (needed to benchmark LRANGE): 358.42 requests per second
    LRANGE_100 (first 100 elements): 4261.30 requests per second
    LRANGE_300 (first 300 elements): 1454.65 requests per second
    LRANGE_500 (first 450 elements): 1006.22 requests per second
    LRANGE_600 (first 600 elements): 777.87 requests per second

### badger

    SET: 7818.44 requests per second
    GET: 88004.93 requests per second
    LPUSH (needed to benchmark LRANGE): 7854.47 requests per second
    LRANGE_100 (first 100 elements): 3984.57 requests per second
    LRANGE_300 (first 300 elements): 1346.18 requests per second
    LRANGE_500 (first 450 elements): 903.01 requests per second
    LRANGE_600 (first 600 elements): 722.38 requests per second

### gaeadb

    SET: 4313.69 requests per second
    GET: 22168.08 requests per second
    LPUSH (needed to benchmark LRANGE): 19106.91 requests per second
    LRANGE_100 (first 100 elements): 3526.71 requests per second
    LRANGE_300 (first 300 elements): 1300.00 requests per second
    LRANGE_500 (first 450 elements): 910.75 requests per second
    LRANGE_600 (first 600 elements): 725.24 requests per second
