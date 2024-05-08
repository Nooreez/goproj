package utils

func generateTrackCode(prefix string, length int) string {
	// Generate a random number
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000) // Change 10000 to suit your requirements

	// Generate timestamp
	timestamp := time.Now().Unix()

	// Combine prefix, timestamp, and random number to form track code
	trackCode := fmt.Sprintf("%s%d%d", prefix, timestamp, randomNumber)

	// Truncate track code to the specified length
	if len(trackCode) > length {
		trackCode = trackCode[:length]
	}

	return trackCode
}