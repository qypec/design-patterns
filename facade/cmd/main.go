package main

import (
	"github.com/design-patterns/facade/pkg/facade"
	"github.com/design-patterns/facade/pkg/torrentfile"
	"github.com/design-patterns/facade/pkg/torrentserver"
)

const (
	filename = "confidential_data_wildberries.dump.sql"
	filesize = 64
	torrentServerIP = "127.0.0.1"
)

func main() {
	file := torrentfile.NewTorrentFile(filename, filesize)
	server := torrentserver.NewUploader([]byte(torrentServerIP))
	client := facade.NewTorrentClient(server, file)
	client.TurnDistribution(filename)
}
