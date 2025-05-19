package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"config/assets_config"
)

func Assets(cfg *viper.Viper) gin.HandlerFunc {
	manifestPath := cfg.GetString(assets_config.ManifestPath)
	entryName := cfg.GetString(assets_config.EntryName)
	return func(c *gin.Context) {
		manifestBytes, readManifestErr := os.ReadFile(manifestPath)
		if readManifestErr != nil {
			c.AbortWithError(http.StatusInternalServerError, readManifestErr)
			return
		}
		manifest := make(map[string]string)
		if unmarshalManifestErr := json.Unmarshal(manifestBytes, &manifest); unmarshalManifestErr != nil {
			c.AbortWithError(http.StatusInternalServerError, unmarshalManifestErr)
			return
		}
		for _, value := range manifest {
			if strings.HasSuffix(value, ".css") {
				c.Set(entryName+".css", value)
			}
			if strings.HasSuffix(value, ".js") {
				c.Set(entryName+".js", value)
			}
		}
		c.Next()
	}
}
