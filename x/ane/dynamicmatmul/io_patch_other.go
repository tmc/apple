//go:build !darwin

package dynamicmatmul

func writeFullTileInput(tile *tile) error {
	return tile.k.WriteInputF32(0, tile.inputPacked)
}

func writeTileRows(tile *tile, rows []int) error {
	return writeFullTileInput(tile)
}
