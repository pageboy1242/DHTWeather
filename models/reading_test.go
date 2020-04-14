package models

import (
    "fmt"
    "testing"
)

func TestInvalidNodeId(t *testing.T) {
    _, err := NewReading(0, -1, "2020-01-01 12:00:00", 10, 20)
    
    fmt.Printf("TestInvalidNodeId")
    
    if err == nil {
        t.Errorf("Expected error")
    }
}
