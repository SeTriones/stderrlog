package stderrlog

import (
	"testing"
)

func TestLog(t *testing.T) {
	Infof("digi=%d,string=%s", 1, "test")
	Errorf("digi=%d,string=%s", 1, "test")
}
