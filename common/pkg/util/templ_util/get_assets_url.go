package templ_util

import (
	"context"
	
	"github.com/a-h/templ"
	"github.com/spf13/cast"
	
	"common/pkg/config/assets_config"
	"common/pkg/facade"
)

func GetAssetsURL(c context.Context, extension string) templ.SafeURL {
	entryName := facade.Config(c).GetString(assets_config.EntryName)
	name := cast.ToString(c.Value(entryName + "." + extension))
	return templ.SafeURL("/static/dist/" + name)
}
