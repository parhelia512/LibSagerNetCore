//go:build !reproducible

package libcore

import (
	"golang.org/x/mobile/asset"
)

func assetOpen(name string) (assetFile, error) {
	return asset.Open(name)
}
