package face

import (
	"testing"
)

func TestFaceMatcher(t *testing.T) {
	fm, _ := newFaceMatcher("./models")
	defer fm.Close()

	r, err := fm.CompareFaceImages("kofi1.jpeg", "gh_card.jpeg")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)

}
