package report

import (
	"fmt"
	"testing"
)

func TestGetTRPGEngineReport(t *testing.T) {
	trpgReport := GetTRPGEngineReport()

	fmt.Printf(trpgReport)
}
