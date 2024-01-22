package textHandler

func generateProgressBar(percent float64) string {
	var progressBar string

	var percentCalc = int(percent * 25 / 100)

	progressBar += "["
	for i := 0; i < percentCalc; i++ {
		progressBar += "="
	}

	for i := 0; i < 25-percentCalc; i++ {
		progressBar += "."
	}

	progressBar += "]"

	return progressBar
}
