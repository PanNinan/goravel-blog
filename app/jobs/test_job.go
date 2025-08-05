package jobs

import (
	"fmt"
)

type TestJob struct {
}

// Signature The name and signature of the job.
func (receiver *TestJob) Signature() string {
	return "test_job"
}

// Handle Execute the job.
func (receiver *TestJob) Handle(args ...any) error {
	for i, i2 := range args {
		fmt.Printf("%d, 值=%v，类型=%T\n", i, i2, i2)
	}
	return nil
}
