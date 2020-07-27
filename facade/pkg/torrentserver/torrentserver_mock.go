package torrentserver

import (
	"github.com/stretchr/testify/mock"
)

// MockTorrentServer ...
type MockTorrentServer struct {
	mock.Mock
}

// Upload ...
func (m *MockTorrentServer) Upload(filename string) (err error) {
	args := m.Called(filename)
	return args.Error(0)
}
