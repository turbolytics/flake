package flake

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// NewGenerator initializes a Generator
func NewGenerator(opts ...GeneratorOption) (*Generator, error) {
	fg := &Generator{
		nowFn: func() time.Time {
			return time.Now().UTC()
		},
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(fg); err != nil {
			return nil, err
		}
	}

	if fg.workerID == nil {
		return nil, errors.New("workerID must be provided")
	}

	return fg, nil
}

// max48Bits is the maximum value for a 48-bit unsigned integer
const max48Bits = (1 << 48) - 1

// IsWithin48BitsRange checks if the given uint64 value fits within 48 bits
func IsWithin48BitsRange(value uint64) bool {
	return value <= max48Bits
}

// ID structure
type ID struct {
	Timestamp uint64 `json:"timestamp"` // 64 bits
	WorkerID  uint64 `json:"worker_id"` // 48 bits used
	Sequence  uint16 `json:"sequence"`  // 16 bits
}

func GeneratorWithWorkerID(workerID uint64) GeneratorOption {
	return func(fg *Generator) error {
		if !IsWithin48BitsRange(workerID) {
			return errors.New("workerID must fit within 48 bits")
		}
		fg.workerID = &workerID
		return nil
	}
}

// GeneratorWithNowFn sets a custom now function in Generator
func GeneratorWithNowFn(nowFn func() time.Time) GeneratorOption {
	return func(fg *Generator) error {
		fg.nowFn = nowFn
		return nil
	}
}

type GeneratorOption func(*Generator) error

// Generator holds state for generating IDs
type Generator struct {
	workerID *uint64
	sequence uint16
	lastTime uint64
	nowFn    func() time.Time
	mutex    sync.Mutex
}

// GenerateFlakeID creates a new ID
func (fg *Generator) GenerateFlakeID() (ID, error) {
	fg.mutex.Lock()
	defer fg.mutex.Unlock()

	now := uint64(fg.nowFn().UnixMilli())

	// Handle clock going backwards
	if now < fg.lastTime {
		now = fg.lastTime
		fg.sequence++
	} else if now == fg.lastTime {
		fg.sequence++
	} else {
		fg.sequence = 0
		fg.lastTime = now
	}

	return ID{
		Timestamp: now,
		WorkerID:  *fg.workerID,
		Sequence:  fg.sequence,
	}, nil
}

// String representation for the ID
func (f ID) String() string {
	return fmt.Sprintf("%016X-%012X-%04X", f.Timestamp, f.WorkerID, f.Sequence)
}

func NewIDFromStr(idStr string) (ID, error) {
	var id ID

	// Split the string by the dashes
	parts := strings.Split(idStr, "-")
	if len(parts) != 3 {
		return id, fmt.Errorf("invalid ID format: %s", idStr)
	}

	// Parse each part
	timestamp, err := strconv.ParseUint(parts[0], 16, 64)
	if err != nil {
		return id, fmt.Errorf("invalid Timestamp: %s", err)
	}

	workerID, err := strconv.ParseUint(parts[1], 16, 64)
	if err != nil {
		return id, fmt.Errorf("invalid WorkerID: %s", err)
	}

	sequence, err := strconv.ParseUint(parts[2], 16, 16)
	if err != nil {
		return id, fmt.Errorf("invalid Sequence: %s", err)
	}

	// Construct the ID
	id = ID{
		Timestamp: timestamp,
		WorkerID:  workerID,
		Sequence:  uint16(sequence),
	}

	return id, nil
}
