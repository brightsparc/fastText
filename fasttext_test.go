package fasttextgo

import (
	"log"
	"strings"
	"testing"
	"time"
)

func init() {
	t0 := time.Now()
	LoadModel("result/dbpedia.bin")
	log.Printf("Model loaded in %s\n", time.Since(t0))
}

func TestPredict(t *testing.T) {
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

func TestPredictK(t *testing.T) {
	// Test the first row in test file
	s := "__label__11 , didelta , didelta is a genus of flowering plants in the daisy family . "
	t.Log(s)
	probs, labels, err := PredictK(s, 2)
	if err != nil {
		t.Error(err)
	} else if !strings.HasPrefix(s, labels[0]) {
		t.Errorf("Label %q not matched %q", labels[0], s)
	} else {
		t.Log(labels, probs)
	}
}
