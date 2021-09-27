package main

import (
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var expectedCounts = map[string]int{
	"blank":   0,
	"checker": 30,
	"corners": 4,
	"empty":   0,
	"huge":    6581317,
	"jagged":  1,
	"one":     1,
	"tiny":    1,
	"two":     2,
}

var namedGrids = loadData()

func loadData() map[string]Grid {
	namedGrids := map[string]Grid{}
	root := "data"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return nil
		}
		name := info.Name()
		var grid Grid
		if strings.HasSuffix(name, ".gz") {
			name = name[:len(name)-3]
			r, gzerr := gzip.NewReader(f)
			if gzerr != nil {
				return gzerr
			}
			grid, err = FromReader(r)
			r.Close()
		} else {
			grid, err = FromReader(f)
		}
		f.Close()
		if err != nil {
			return err
		}
		namedGrids[name] = grid
		return nil
	})
	if err != nil {
		panic(err)
	}
	return namedGrids
}

func TestGrids(t *testing.T) {
	for name, grid := range namedGrids {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, expectedCounts[name], countIslands(grid), name)
		})
	}
}

func BenchmarkGrids(b *testing.B) {
	for name, v := range namedGrids {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countIslands(v)
			}
		})
	}
}
