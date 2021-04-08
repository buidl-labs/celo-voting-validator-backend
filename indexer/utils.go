package indexer

const blocksPerEpoch = 17280

func calculateCurrentEpoch(latestBlock int) int {
	return latestBlock / blocksPerEpoch
}
