package gmf

import "testing"

var (
// inputSampleFilename  string = "examples/tests-sample.mp4"
// outputSampleFilename string = "examples/tests-output.mp4"
// inputSampleWidth     int    = 320
// inputSampleHeight    int    = 200
)

func TestGetFilter(t *testing.T) {
	f := assert(GetFilter("buffer")).(*AVFilter)
	if f.avFilter == nil {
		t.Fatal("AVFilter is not initialized")
	}
}
