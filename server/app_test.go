package server

import "testing"

func TestNewTableApp(t *testing.T) {
	a := NewTableApp()
	a.Start()
}
