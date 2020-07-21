package torrentfile

import (
	"errors"

	"github.com/design-patterns/facade/pkg/api/v1"
)

// Accessor provides functionality for torrentFile structure
type Accessor interface {
	PrepareToDistribution(filename string) (err error)
	GetName() (name string)
}

type torrentFile struct {
	name string
	size int
	mode int
}

// GetName returns the file name
func (f *torrentFile)GetName() (name string) {
	name = f.name
	return
}

// PrepareToDistribution validates the file name and defines the upload method
func (f *torrentFile)PrepareToDistribution(filename string) (err error) {
	if f.name != filename {
		err = errors.New(filename + v1.FileNotFoundError)
		return
	}

	f.mode = v1.Endgame
	if f.size <= v1.MaxFileSizeForEndgame {
		f.mode = v1.Separation
	}
	return
}

// NewTorrentFile creates an instance of the Accessor
func NewTorrentFile(filename string, filesize int) Accessor {
	return &torrentFile{
		name: filename,
		size: filesize,
	}
}
