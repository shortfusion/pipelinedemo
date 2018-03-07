package main

import (
	"testing"
)

func TestSizeSite(t *testing.T) {
	testURL := "https://google.com/robots.txt"
	testSiteSize := sizeSite(testURL)
	testThreshold := 5000
	if testSiteSize < testThreshold {
		t.Errorf("Site size is too small, got: %d, want > %d.", testSiteSize, testThreshold)
	}
}
