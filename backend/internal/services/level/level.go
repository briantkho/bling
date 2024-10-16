package level

import "github.com/google/uuid"

type Level struct {
	LevelID        uuid.UUID `json:"levelID"`
	LevelNumber    int       `json:"levelNumber"`
	PointsRequired int       `json:"pointsRequired"`
	LevelName      string    `json:"levelName"`
	LevelIcon      string    `json:"levelIcon"`
}
