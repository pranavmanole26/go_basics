package main

import "testing"

func BenchmarkSyncWithWaitGroups(t *testing.B) {
	for i := 0; i < t.N; i++ {
		syncWithWaitGroups()
	}
}

func BenchmarkSyncWithChannels(t *testing.B) {
	for i := 0; i < t.N; i++ {
		syncWithChannels()
	}
}
