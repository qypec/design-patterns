package facade

import (
	"fmt"
)

type accessor interface {
	PrepareToDistribution(filename string) (err error)
	GetName() (name string)
}

type uploader interface {
	Upload(filename string) (err error)
}

// TorrentDistributor controls the files distribution
type TorrentDistributor interface {
	TurnDistribution(filename string) (err error)
}

type torrentClient struct {
	server		uploader
	file		accessor
}

// Turn Distributor starts the file distribution process.
// Before distribution you need to prepare the file and establish a connection to the torrent server.
func (t *torrentClient)TurnDistribution(filename string) (err error) {
	err = t.file.PrepareToDistribution(filename)
	if err != nil {
		return
	}

	err = t.server.Upload(t.file.GetName())
	if err != nil {
		return
	}
	fmt.Println("File distribution started successfully")
	return
}

// NewTorrentClient creates an instance of the TorrentDistributor
func NewTorrentClient(server uploader, fileToDist accessor) TorrentDistributor {
	return &torrentClient{
		server: server,
		file: fileToDist,
	}
}
