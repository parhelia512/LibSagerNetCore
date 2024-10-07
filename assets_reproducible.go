//go:build reproducible

package libcore

import (
	"github.com/sagernet/gomobile/asset"
)

func assetOpen(name string) (assetFile, error) {
	return asset.Open(name)
}
