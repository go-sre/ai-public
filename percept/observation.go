package percept

type Measure struct {
	Signal string // Saturation, Status Flag, Status Code
	//// Timeout, RateLimiter, Retry, Proxy, Saturation, Upstream Saturation
	Percent int
	Value   int // Based on the Signal, this can be a number, percentage, percentile
	Slope   int // Shows the direction and magnitude of the signal. 1 = increasing, 0 = no change, -1 = decreasing
}

type Observation struct {
	Source   string // Access Log Ingress, Egress, Ping, or Activity
	Origin   string
	Measures []Measure
}
