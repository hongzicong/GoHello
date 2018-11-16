# GoHello
A easy web program based on Golang

# curl test

```bash
$curl http://localhost:8080/test
{"message":"Hello! My name is GoHello!"}
C:\Users\DELL-PC>curl -v http://localhost:8080/test
*   Trying ::1...
* TCP_NODELAY set
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /test HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Fri, 16 Nov 2018 06:55:42 GMT
< Content-Length: 40
<
{"message":"Hello! My name is GoHello!"}* Connection #0 to host localhost left intact
```

# ab test

```bash
$ ./ab -n 1000 -c 100 http://localhost:8080/test
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /test
Document Length:        40 bytes

Concurrency Level:      100
Time taken for tests:   1.851 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      163000 bytes
HTML transferred:       40000 bytes
Requests per second:    540.38 [#/sec] (mean)
Time per request:       185.054 [ms] (mean)
Time per request:       1.851 [ms] (mean, across all concurrent requests)
Transfer rate:          86.02 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       3
Processing:     3   78  20.0     80     216
Waiting:        1   54  23.6     55     185
Total:          3   78  20.0     80     219

Percentage of the requests served within a certain time (ms)
  50%     80
  66%     87
  75%     89
  80%     90
  90%     95
  95%    100
  98%    107
  99%    143
 100%    219 (longest request)
```