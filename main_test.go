package main

import (
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

func TestFoo(t *testing.T) {}
