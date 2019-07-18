package main

import (
	"sync"
	"testing"
)

func BenchmarkOnce(b *testing.B) {
	var count int

	onceBody := func() {
		count++
	}

	done := make(chan bool)

	var o sync.Once

	for i := 0; i <= b.N; i++ {
		once(onceBody, done, &o)
	}
}

func BenchmarkBranch(b *testing.B) {
	var count int

	onceBody := func() {
		count++
	}

	done := make(chan bool)

	for i := 0; i <= b.N; i++ {
		branch(onceBody, done)
	}
}
