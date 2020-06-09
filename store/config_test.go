package store

import "testing"

func TestDatabaseConfig_String(t *testing.T) {
	want := "postgresql://postgres:root@localhost:5432/amakhosi_pumps?sslmode=disable"
	if got := dataSourceName(); got != want {
		t.Fatalf("Expected %s ,\n got %s", want, got)
	}
}
