### Disk IO test


## write 1Gb file and read back to buffer

Write took 401.385971ms
Written 1073741824, size 1073741824, <nil>
Read: read buffer size 1073741824, write buffer size 1073741824, took 707.545744ms 1073741824
Read: read buffer size 1073741824, write buffer size 1048576, took 479.352146ms 1073741824
Read: read buffer size 1073741824, write buffer size 1024, took 748.170493ms 1073741824
Read: read buffer size 4096, write buffer size 1073741824, took 277.277482ms 1073741824

## Write 1Mb file and read back
Write took 496.22µs
Written 1048576, size 1048576, <nil>
Read: read buffer size 1048576, write buffer size 1048576, took 285.939µs 1048576
Read: read buffer size 1048576, write buffer size 1024, took 322.396µs 1048576
Read: read buffer size 4096, write buffer size 1048576, took 275.646µs 1048576


## write 10MB file and read back
Write took 4.176671ms
Written 10485760, size 10485760, <nil>
Read: read buffer size 1048576, write buffer size 10485760, took 2.490423ms 10485760
Read: read buffer size 1048576, write buffer size 1024, took 5.114667ms 10485760
Read: read buffer size 4096, write buffer size 10485760, took 3.584089ms 10485760

## System spec
1. SSD disk
2. Page size 4096

