package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTrackCode(prefix string, length int) (string, error) {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(99999-10000) + 10000

	timestamp := time.Now().UnixNano()

	trackCode := fmt.Sprintf("%s%d%d", prefix, timestamp, randomNumber)

	if len(trackCode) > length {
		trackCode = trackCode[:length]
	}

	return trackCode, nil
}
