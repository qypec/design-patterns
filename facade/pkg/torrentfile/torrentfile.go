package torrentfile

import (
	"errors"
	"fmt"
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
	const endgame = 1
	const separation = 2

	if f.name != filename {
		err = errors.New(fmt.Sprintf("%v: file not found\n", filename))
		return
	}

	f.mode = endgame
	if f.size < 1024 {
		f.mode = separation
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
