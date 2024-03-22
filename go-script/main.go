package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	markdownTable := `
| Consideration Point                   | OPA/Gatekeeper | Kyverno | Kubewarden | JsPolicy | Weight |
|---------------------------------------|----------------|---------|------------|----------|--------|
| Security and Isolation                | 3              | 2       | 3          | 3        | 3      |
| Flexibility in Policy Development     | 3              | 2       | 3          | 3        | 3      |
| Performance Efficiency                | 2              | 2       | 3          | 2        | 3      |
| Scalability                           | 3              | 2       | 3          | 2        | 3      |
| Community and Ecosystem Growth       | 3              | 2       | 2          | 1        | 1      |
| Alignment with Organizational Practices| 2             | 3       | 3          | 3        | 3      |
| Ease of Policy Management             | 1              | 3       | 3          | 3        | 2      |
| Learning Curve                        | 1              | 3       | 3          | 3        | 2      |
| Resource Utilization                  | 2              | 2       | 2          | 2        | 3      |
| Policy Language Compatibility         | 1              | 3       | 2          | 3        | 2      |
| High Availability and Fault Tolerance | 3              | 2       | 2          | 2        | 3      |
| Integration with Existing Tools       | 3              | 3       | 3          | 2        | 3      |
| Upgrade and Maintenance Path          | 3              | 1       | 3          | 2        | 2      |
| Support and Documentation             | 3              | 2       | 2          | 2        | 2      |
| Cost Implications                     | 3              | 3       | 3          | 3        | 1      |
| Compliance and Audit Capabilities     | 3              | 3       | 3          | 2        | 3      |
| Policy Execution Transparency         | 2              | 3       | 3          | 3        | 2      |
| Testing Capapabilities	            | 2              | 1       | 3          | 3        | 2      |
`

	calculateScores(markdownTable)
}

func calculateScores(markdownTable string) {
	lines := strings.Split(markdownTable, "\n")
	var weights []int
	scores := map[string][]int{
		"OPA/Gatekeeper": {},
		"Kyverno":        {},
		"Kubewarden":     {},
		"JsPolicy":       {},
	}

	for _, line := range lines[3:] {
		line = strings.Trim(line, "|")
		line = strings.TrimSpace(line)

		cells := strings.Split(line, "|")

		if len(cells) != 6 {
			continue
		}

		for i, cell := range cells {
			value, err := strconv.Atoi(strings.TrimSpace(cell))
			if err != nil {
				fmt.Printf("Error converting cell to integer: %v\n", err)
				continue
			}

			switch i {
			case 1:
				scores["OPA/Gatekeeper"] = append(scores["OPA/Gatekeeper"], value)
			case 2:
				scores["Kyverno"] = append(scores["Kyverno"], value)
			case 3:
				scores["Kubewarden"] = append(scores["Kubewarden"], value)
			case 4:
				scores["JsPolicy"] = append(scores["JsPolicy"], value)
			case 5:
				weights = append(weights, value)
			}
		}
	}

	for engine, engineScores := range scores {
		totalScore := 0
		for i, score := range engineScores {
			totalScore += score * weights[i]
		}
		fmt.Printf("%s: %d\n", engine, totalScore)
	}
}
