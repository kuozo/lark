package pkg

import (
	"testing"
)

func TestGitHub(t *testing.T){

	rels, err := GitHub()
	if err != nil {
		t.Error(err)
	}else {
		t.Log(len(rels))
	}
}