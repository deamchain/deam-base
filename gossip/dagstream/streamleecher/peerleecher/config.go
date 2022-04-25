package peerleecher

import (
	"time"

	"github.com/deamchain/deam-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
