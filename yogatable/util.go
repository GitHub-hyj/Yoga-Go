package yogatable

import "github.com/iikira/BaiduPCS-Go/pcsutil/pcstime"

func TimePrefix() string {
	return "[" + pcstime.BeijingTimeOption("Refer") + "]"
}
