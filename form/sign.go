package form

import "xyb/model"

func NewSignForm(info *model.SystemInfo) *Forms {
	forms := Forms{
		Form{"traineeId", info.TraineeId},
		Form{"adcode", info.Adcode},
		Form{"lat", info.Lat},
		Form{"lng", info.Lng},
		Form{"address", info.Address},
		Form{"deviceName", "microsoft"},
		Form{"punchInStatus", "1"},
		Form{"clockStatus", "1"},
	}
	return &forms
}
