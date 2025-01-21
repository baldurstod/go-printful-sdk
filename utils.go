package printfulsdk

import "fmt"

// Scan the X-RateLimit-Policy header and return the rate in tokens / s
func getRateFromPolicy(policy string) float64 {
	var quota, seconds uint

	n, err := fmt.Sscanf(policy, "%d;w=%d", &quota, &seconds)

	if err == nil && n == 2 {
		return float64(quota) / float64(seconds)
	}

	return 1
}
