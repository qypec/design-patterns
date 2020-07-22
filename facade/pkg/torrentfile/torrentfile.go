package torrentfile

import (
	"errors"

	"github.com/design-patterns/facade/pkg/api/v1"
)

// Accessor provides functionality for torrentFile structure
type Accessor interface {
	PrepareToDistribution(filename string) (err error)
	GetName() (name string)
	GetMode() (mode int)
	GetSize() (size int)
}

type torrentFile struct {
	name string
	size int
	mode int
}

// GetName returns the file name
func (f *torrentFile) GetName() (name string) {
	name = f.name
	return
}

// GetMode returns the file mode
func (f *torrentFile) GetMode() (mode int) {
	mode = f.mode
	return
}

// GetSize returns the file size
func (f *torrentFile) GetSize() (size int) {
	size = f.size
	return
}

// PrepareToDistribution validates the file name and defines the upload method
func (f *torrentFile) PrepareToDistribution(filename string) (err error) {
	if f.name != filename {
		err = errors.New(filename + v1.FileNotFoundError)
		return
	}

	f.mode = v1.Endgame
	if f.size > v1.MaxFileSizeForEndGame {
		f.mode = v1.Separation
	}
	return
}

// NewAccessor creates an instance of the Accessor
func NewAccessor(filename string, filesize int) Accessor {
	return &torrentFile{
		name: filename,
		size: filesize,
	}
}
