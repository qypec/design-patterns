package torrentfile

import (
	"github.com/stretchr/testify/mock"
)

// MockTorrentFile ...
type MockTorrentFile struct {
	mock.Mock
}

// GetName ...
func (m *MockTorrentFile) GetName() (name string) {
	args := m.Called()
	return args.String(0)
}

// GetMode ...
func (m *MockTorrentFile) GetMode() (mode int) {
	args := m.Called()
	return args.Int(0)
}

// GetSize ...
func (m *MockTorrentFile) GetSize() (size int) {
	args := m.Called()
	return args.Int(0)
}

// PrepareToDistribution ...
func (m *MockTorrentFile) PrepareToDistribution(filename string) (err error) {
	args := m.Called(filename)
	return args.Error(0)
}
