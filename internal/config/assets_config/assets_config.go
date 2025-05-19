package assets_config

import "github.com/spf13/viper"

const (
	EntryName    = "assets-entry-name"
	ManifestPath = "assets-manifest-path"
)

func Read(cfg *viper.Viper) {
	cfg.Set(EntryName, "bundle")
	cfg.Set(ManifestPath, "./public/static/dist/manifest.json")
}
