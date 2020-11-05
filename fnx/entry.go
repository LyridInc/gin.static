package fnx

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Initialize() *gin.Engine {

	r := gin.Default()

	use_gzip, _ := strconv.ParseBool(os.Getenv("USE_GZIP")) // defaults to false if not exist
	if use_gzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	r.Static("/", "./public")

	return r
}
