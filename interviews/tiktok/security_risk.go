package tiktok

func findSecurityRisk(videoHash, encryptionKey string) int {
	n := len(videoHash)

	for i := 0; i < n; i++ {
		if videoHash[i] < encryptionKey[i] {
			return 0
		} else if videoHash[i] > encryptionKey[i] {
			return i + 1
		} else {
			continue
		}
	}

	return -1
}
