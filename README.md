# gaeadb_bench

    use redis-benchmark for benchmark

## HDD Single Thread

### CMD

redis-benchmark -h 127.0.0.1 -c 1 -p 6378 -t set,get -n 100000 -q

### bolt

    SET: 123.98 requests per second
    GET: 103.13 requests per second

### badger

    SET: 133.21 requests per second
    GET: 105.41 requests per second

### gaeadb

    SET: 76.56 requests per second
    GET: 77.15 requests per second

## HDD Multiple Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 100 -p 6378 -t set,get -n 10000000 -q

### badger

    SET: 20906.16 requests per second
    GET: 15417.82 requests per second

### gaeadb

    SET: 10279.43 requests per second
    GET: 10353.88 requests per second

## SSD Single Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 1 -p 6378 -t set,get -n 100000 -q

### bolt

    SET: 599.68 requests per second
    GET: 94339.62 requests per second

### badger

    SET: 153.15 requests per second
    GET: 82169.27 requests per second

### gaeadb

    SET: 32362.46 requests per second
    GET: 3672.69 requests per second

## SSD Multiple Thread

### CMD

    redis-benchmark -h 127.0.0.1 -c 100 -p 6378 -t set,get -n 10000000 -q

### badger

    SET: 3665.78 requests per second
    GET: 82493.27 requests per second

### gaeadb

    SET: 85941.66 requests per second
    GET: 85460.59 requests per second
