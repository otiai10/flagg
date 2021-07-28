// +build go1.16

package largo

import (
	_ "embed"
)

//go:embed usage.go.tpl
var defaultUsageTemplate []byte
