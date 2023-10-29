package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)/2; i++ {
		podium[i], podium[len(podium)-i-1] = podium[len(podium)-i-1], podium[i]
	}
	return podium
}
