package tfile

import (
	"errors"
	"fmt"
)

// File represents everything about a torrent file
type TorrentFile interface {
	PrepareToDistribution(filename string) error
	GetName() string
}

type file struct {
	name string
	size int
	mode int
}

// GetName returns the file name
func (f *file)GetName() string {
	return f.name
}

// PrepareToDistribution validates the file name and defines the upload method
func (f *file)PrepareToDistribution(filename string) error {
	const endgame = 1
	const separation = 2

	if f.name != filename {
		return errors.New(fmt.Sprintf("%v: file not found\n", filename))
	}

	f.mode = endgame
	if f.size < 1024 {
		f.mode = separation
	}
	return nil
}

// NewTorrentFile creates an instance of the File
func NewTorrentFile(filename string, filesize int) TorrentFile {
	return &file{
		name: filename,
		size: filesize,
	}
}
