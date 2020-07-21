package facade

import (
	"fmt"
	"github.com/design-patterns/facade/pkg/tfile"
	"github.com/design-patterns/facade/pkg/tserver"
)


// TorrentDistributor controls the files distribution
type TorrentDistributor interface {
	TurnDistribution(filename string) error
}

type torrentClient struct {
	server		tserver.Uploader
	file		tfile.TorrentFile
}

// Turn Distributor starts the file distribution process.
// Before distribution you need to prepare the file and establish a connection to the torrent server.
func (t *torrentClient)TurnDistribution(filename string) (err error) {
	err = t.file.PrepareToDistribution(filename)
	if err != nil {
		return
	}

	err = t.server.Upload(t.file)
	if err != nil {
		return
	}
	fmt.Println("File distribution started successfully")
	return
}

// NewTorrentClient creates an instance of the TorrentDistributor
func NewTorrentClient(serverAddress string, fileToDist tfile.TorrentFile) TorrentDistributor {
	return &torrentClient{
		server: tserver.NewUploader(serverAddress),
		file: fileToDist,
	}
}
