package util_test

import (
	"testing"
	"time"

	"github.com/mazafard/ferferi/pkg/util"
)

func TestExpiredMap_SetAndGet(t *testing.T) {
	em := util.NewExpiredMap()
	em.Set("key1", "value1", 2)
	em.Set("key2", "value2", 2)

	found1, val1 := em.Get("key1")
	if !found1 || val1 != "value1" {
		t.Errorf("Expected value1 for key1, but got %v", val1)
	}

	found2, val2 := em.Get("key2")
	if !found2 || val2 != "value2" {
		t.Errorf("Expected value2 for key2, but got %v", val2)
	}

	time.Sleep(3 * time.Second)

	found3, val3 := em.Get("key1")
	if found3 || val3 != nil {
		t.Errorf("Expected nil value for key1, but got %v", val3)
	}

	found4, val4 := em.Get("key2")
	if found4 || val4 != nil {
		t.Errorf("Expected nil value for key2, but got %v", val4)
	}
}

func TestExpiredMap_Length(t *testing.T) {
	em := util.NewExpiredMap()
	em.Set("key1", "value1", 2)
	em.Set("key2", "value2", 2)
	em.Set("key3", "value3", 3)

	length := em.Length()
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	time.Sleep(2 * time.Second)

	length = em.Length()
	if length != 1 {
		t.Errorf("Expected length 1, but got %d", length)
	}
}

func TestExpiredMap_Clear(t *testing.T) {
	em := util.NewExpiredMap()
	em.Set("key1", "value1", 2)
	em.Set("key2", "value2", 2)
	em.Set("key3", "value3", 3)

	length := em.Length()
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	em.Clear()

	length = em.Length()
	if length != 0 {
		t.Errorf("Expected length 0, but got %d", length)
	}
}
