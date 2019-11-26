package yogaverbose

import (
	"Yoga-Go/yogautil/yogatime"
)

func TimePrefix() string {
	return "[" + yogatime.BeijingTimeOption("Refer") + "]"
}
