//go:build !glhf_no_default
// +build !glhf_no_default

package glhf

// importing [glhf/driver/ebiten] already sets up the driver using init functions
import _ "github.com/crazyinfin8/glhf/driver/ebiten"
