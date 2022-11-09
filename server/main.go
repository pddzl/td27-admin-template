package main

import (
	"server/core"
	"server/global"
)

func main() {
	global.TD27_VP = core.Viper()
	global.TD27_LOG = core.Zap()
}
