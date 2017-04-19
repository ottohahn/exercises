package main

import (
	"testing"
)

var (
	args = []string{"this is first second third"}
)

func TestEcho1(t *testing.T) {
    if len(args) > 0 {
        Echo1(args)
    }
}

func TestEcho2(t *testing.T) {
    if len(args) > 0 {
        Echo2(args)
    }
}

func TestEcho3(t *testing.T) {
    if len(args) > 0 {
	 Echo3(args)
    }
}

func BenchmarkEcho1(b *testing.B) {
    if len(args) > 0 {
	for n := 0; n < b.N; n++ {
	    Echo1(args)
	}
    }
}

func BenchmarkEcho2(b *testing.B) {
    if len(args) > 0 {
	for n := 0; n < b.N; n++ {
	    Echo2(args)
	}
    }
}

func BenchmarkEcho3(b *testing.B) {
    if len(args) > 0 {
        for n := 0; n < b.N; n++ {
	    Echo3(args)
	}
    }
}
