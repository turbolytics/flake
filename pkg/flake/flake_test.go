package flake

import (
	"testing"
	"time"
)

// TestIs48Bits tests the Is48Bits function
func TestIs48Bits(t *testing.T) {
	cases := []struct {
		value    uint64
		expected bool
	}{
		{0xFFFFFFFFFFFF, true}, // 48-bit max value
		{1 << 49, false},       // 49-bit value
		{0x000000000001, true}, // Small value
	}

	for _, c := range cases {
		if got := IsWithin48BitsRange(c.value); got != c.expected {
			t.Errorf("Is48Bits(%X) = %v; want %v", c.value, got, c.expected)
		}
	}
}

// TestNewIDFromStr tests the parsing of a string into an ID
func TestNewIDFromStr(t *testing.T) {
	originalID := ID{
		Timestamp: 1234567890,
		WorkerID:  0x123456789ABC,
		Sequence:  0x0010,
	}
	idStr := originalID.String()

	parsedID, err := NewIDFromStr(idStr)
	if err != nil {
		t.Fatalf("NewIDFromStr failed: %v", err)
	}

	if originalID != parsedID {
		t.Errorf("Parsed ID = %+v; want %+v", parsedID, originalID)
	}
}

// TestNewGenerator tests the NewGenerator function
func TestNewGenerator(t *testing.T) {
	validWorkerID := uint64(0x123456)
	invalidWorkerID := uint64(1 << 49) // 49-bit value

	_, err := NewGenerator(GeneratorWithWorkerID(validWorkerID))
	if err != nil {
		t.Errorf("NewGenerator failed with valid workerID: %v", err)
	}

	_, err = NewGenerator(GeneratorWithWorkerID(invalidWorkerID))
	if err == nil {
		t.Errorf("NewGenerator should have failed with invalid workerID")
	}

	_, err = NewGenerator()
	if err == nil {
		t.Errorf("NewGenerator should have failed without workerID")
	}
}

func TestGenerateFlakeID(t *testing.T) {
	// Fixed timestamp for deterministic test
	fixedTime := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)

	// Custom now function returning fixed time
	customNow := func() time.Time {
		return fixedTime
	}

	workerID := uint64(0x123456)
	sequenceStart := uint16(0)

	// Initialize Generator with fixed now function
	gen, err := NewGenerator(
		GeneratorWithWorkerID(workerID),
		GeneratorWithNowFn(customNow),
	)
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	// Generate first ID
	id1, err := gen.GenerateFlakeID()
	if err != nil {
		t.Fatalf("Failed to generate first ID: %v", err)
	}

	expectedID1 := ID{
		Timestamp: uint64(fixedTime.UnixMilli()),
		WorkerID:  workerID,
		Sequence:  sequenceStart,
	}

	if id1 != expectedID1 {
		t.Errorf("ID1 mismatch: expected %v, got %v", expectedID1, id1)
	}

	// Generate second ID, expecting sequence increment
	id2, err := gen.GenerateFlakeID()
	if err != nil {
		t.Fatalf("Failed to generate second ID: %v", err)
	}

	expectedID2 := ID{
		Timestamp: uint64(fixedTime.UnixMilli()),
		WorkerID:  workerID,
		Sequence:  sequenceStart + 1,
	}

	if id2 != expectedID2 {
		t.Errorf("ID2 mismatch: expected %v, got %v", expectedID2, id2)
	}

	// Test that the string representation matches expected format
	expectedStr1 := "000001906B975C00-000000123456-0000"
	if id1.String() != expectedStr1 {
		t.Errorf("ID1 string mismatch: expected %s, got %s", expectedStr1, id1.String())
	}

	expectedStr2 := "000001906B975C00-000000123456-0001"
	if id2.String() != expectedStr2 {
		t.Errorf("ID2 string mismatch: expected %s, got %s", expectedStr2, id2.String())
	}
}

func BenchmarkGenerateFlakeID(b *testing.B) {
	// Initialize the FlakeGenerator
	fg, _ := NewGenerator(
		GeneratorWithWorkerID(1),
	)

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fg.GenerateFlakeID()
	}
}
