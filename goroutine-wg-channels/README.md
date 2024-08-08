Benchmark results are as below

pmanole@G1VWG396DJ goroutine-wg-channels % go test -bench=. -benchtime=5s
goos: darwin
goarch: arm64
pkg: goroutine-wg-channels
BenchmarkSyncWithWaitGroups-12           8532236               702.2 ns/op
BenchmarkSyncWithChannels-12            10002130               598.5 ns/op
PASS
ok      goroutine-wg-channels   16.359s

Remarks:
With the use of channels we get better results compare to waitgroups