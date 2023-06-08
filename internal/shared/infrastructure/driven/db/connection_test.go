package db

import (
	"errors"
	"testing"
)

func TestNewMySQLConnected(t *testing.T) {
	t.Run("Given a correct configuration, MySQL DB will connect", func(t *testing.T) {
		_, err := LoadSqlConnection(NewMySQLDBConfiguration())

		if !errors.Is(err, nil) {
			t.Errorf("%v error was expected, but %v error was obtained", nil, err)
		}
	})
}
