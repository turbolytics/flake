package flake

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

// FlakeID structure
type FlakeID struct {
	Timestamp  uint64 // 48 bits
	RegionID   uint16 // 16 bits
	MachineID  uint16 // 16 bits
	Sequence   uint32 // 32 bits
	Randomness uint16 // 16 bits
}

// FlakeGenerator holds state for generating IDs
type FlakeGenerator struct {
	regionID  uint16
	machineID uint16
	sequence  uint32
	lastTime  uint64
	mutex     sync.Mutex
}

// NewFlakeGenerator initializes a FlakeGenerator
func NewFlakeGenerator(regionID, machineID uint16) *FlakeGenerator {
	return &FlakeGenerator{
		regionID:  regionID,
		machineID: machineID,
	}
}

// GenerateFlakeID creates a new FlakeID
func (fg *FlakeGenerator) GenerateFlakeID() FlakeID {
	fg.mutex.Lock()
	defer fg.mutex.Unlock()

	now := uint64(time.Now().UnixMilli()) & 0xFFFFFFFFFFFF // 48 bits timestamp
	if now == fg.lastTime {
		fg.sequence++
	} else {
		fg.sequence = 0
		fg.lastTime = now
	}

	randomness := uint16(0)
	binary.Read(rand.Reader, binary.BigEndian, &randomness)

	return FlakeID{
		Timestamp:  now,
		RegionID:   fg.regionID,
		MachineID:  fg.machineID,
		Sequence:   fg.sequence,
		Randomness: randomness,
	}
}

// String representation for the FlakeID
func (f FlakeID) String() string {
	return fmt.Sprintf("%012X-%04X-%04X-%08X-%04X", f.Timestamp, f.RegionID, f.MachineID, f.Sequence, f.Randomness)
}
