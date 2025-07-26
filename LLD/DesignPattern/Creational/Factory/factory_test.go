package factory

import (
	"testing"
)

func TestNewDatabase_MySQL(t *testing.T) {
	db := NewDatabase("mysql")
	if db == nil {
		t.Fatal("Expected non-nil Database, got nil")
	}
	if db.Connect() != "Connected to MySQL" {
		t.Errorf("Expected MySQL connection, got: %s", db.Connect())
	}
}

func TestNewDatabase_PostgreSQL(t *testing.T) {
	db := NewDatabase("postgresql")
	if db == nil {
		t.Fatal("Expected non-nil Database, got nil")
	}
	if db.Connect() != "Connected to PostgreSQL" {
		t.Errorf("Expected PostgreSQL connection, got: %s", db.Connect())
	}
}

func TestNewDatabase_Default(t *testing.T) {
	db := NewDatabase("unknown")
	if db == nil {
		t.Fatal("Expected non-nil Database for default case, got nil")
	}
	if db.Connect() != "Connected to MySQL" {
		t.Errorf("Expected default MySQL connection, got: %s", db.Connect())
	}
}
