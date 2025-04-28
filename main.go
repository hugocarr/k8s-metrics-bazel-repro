package main

import (
	"fmt"

	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func main() {
	// Just demonstrate that we can use the metrics types
	podMetrics := &metricsv1beta1.PodMetrics{}
	fmt.Printf("Created a PodMetrics object: %T\n", podMetrics)
}
