package yogaverbose

import (
	"github.com/GitHub-hyj/Yoga-Go/yogautil/yogatime"
)

func TimePrefix() string {
	return "[" + yogatime.BeijingTimeOption("Refer") + "]"
}
