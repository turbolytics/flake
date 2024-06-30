package flake

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ID structure
type ID struct {
	Timestamp  uint64 // 48 bits
	RegionID   uint16 // 16 bits
	MachineID  uint16 // 16 bits
	Sequence   uint32 // 32 bits
	Randomness uint16 // 16 bits
}

func GeneratorWithRegionID(regionID uint16) GeneratorOption {
	return func(fg *Generator) error {
		fg.regionID = &regionID
		return nil
	}
}

func GeneratorWithMachineID(machineID uint16) GeneratorOption {
	return func(fg *Generator) error {
		fg.machineID = &machineID
		return nil
	}
}

type GeneratorOption func(*Generator) error

// Generator holds state for generating IDs
type Generator struct {
	regionID  *uint16
	machineID *uint16
	sequence  uint32
	lastTime  uint64
	mutex     sync.Mutex
}

// NewGenerator initializes a Generator
func NewGenerator(opts ...GeneratorOption) (*Generator, error) {
	fg := &Generator{}

	// Apply options
	for _, opt := range opts {
		if err := opt(fg); err != nil {
			return nil, err
		}
	}

	if fg.regionID == nil || fg.machineID == nil {
		return nil, errors.New("regionID and machineID must be provided")
	}

	return fg, nil
}

// GenerateFlakeID creates a new ID
func (fg *Generator) GenerateFlakeID() (ID, error) {
	fg.mutex.Lock()
	defer fg.mutex.Unlock()

	now := uint64(time.Now().UnixMilli()) & 0xFFFFFFFFFFFF // 48 bits timestamp

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

	randomness := uint16(0)
	binary.Read(rand.Reader, binary.BigEndian, &randomness)

	return ID{
		Timestamp:  now,
		RegionID:   *fg.regionID,
		MachineID:  *fg.machineID,
		Sequence:   fg.sequence,
		Randomness: randomness,
	}, nil
}

// String representation for the ID
func (f ID) String() string {
	return fmt.Sprintf("%012X-%04X-%04X-%08X-%04X", f.Timestamp, f.RegionID, f.MachineID, f.Sequence, f.Randomness)
}

// NewIDFromString parses a Flake ID string into a FlakeID struct
func NewIDFromString(id string) (ID, error) {
	parts := strings.Split(id, "-")
	if len(parts) != 5 {
		return ID{}, fmt.Errorf("invalid ID format")
	}

	timestamp, err := strconv.ParseUint(parts[0], 16, 64)
	if err != nil {
		return ID{}, fmt.Errorf("invalid timestamp: %w", err)
	}

	regionID, err := strconv.ParseUint(parts[1], 16, 16)
	if err != nil {
		return ID{}, fmt.Errorf("invalid region ID: %w", err)
	}

	machineID, err := strconv.ParseUint(parts[2], 16, 16)
	if err != nil {
		return ID{}, fmt.Errorf("invalid machine ID: %w", err)
	}

	sequence, err := strconv.ParseUint(parts[3], 16, 32)
	if err != nil {
		return ID{}, fmt.Errorf("invalid sequence: %w", err)
	}

	randomness, err := strconv.ParseUint(parts[4], 16, 16)
	if err != nil {
		return ID{}, fmt.Errorf("invalid randomness: %w", err)
	}

	return ID{
		Timestamp:  timestamp,
		RegionID:   uint16(regionID),
		MachineID:  uint16(machineID),
		Sequence:   uint32(sequence),
		Randomness: uint16(randomness),
	}, nil
}
