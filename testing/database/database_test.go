package database

import "testing"

type mockDb struct {
	Database
}

func (*mockDb) Insert(collection string, data interface{}) error {
	return nil
}

func TestInsert(t *testing.T) {
	mock := &mockDb{}

	err := Insert(mock, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}
