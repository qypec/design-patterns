package main

import (
	"github.com/design-patterns/facade/pkg/facade"
	"github.com/design-patterns/facade/pkg/tfile"
)

const (
	filename = "confidential_data_wildberries.dump.sql"
	filesize = 64
	torrentServerIP = "127.0.0.1"
)

func main() {
	file := tfile.NewTorrentFile(filename, filesize)
	client := facade.NewTorrentClient(torrentServerIP, file)
	client.TurnDistribution(filename)
}
