package fasttextgo

import (
	"strings"
	"testing"
	"time"
)

func TestDbpedia(t *testing.T) {
	t0 := time.Now()
	LoadModel("result/dbpedia.bin")
	t.Logf("Model loaded in %s\n", time.Since(t0))

	// Test the first row in test file
	s := "__label__11 , didelta , didelta is a genus of flowering plants in the daisy family . "
	t.Log(s)
	prob, label, err := Predict(s)
	if err != nil {
		t.Error(err)
	} else if !strings.HasPrefix(s, label) {
		t.Errorf("Label %q not matched %q", label, s)
	} else {
		t.Log(label, prob)
	}
}
