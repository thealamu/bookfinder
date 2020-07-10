package finder

import "testing"

func TestNewGoogleBooksFinder(t *testing.T) {
	gbFinder := NewGoogleBooksFinder()
	if gbFinder == nil {
		t.Errorf("NewGoogleBooksFinder returns nil, expected a new GoogleBooks instance")
	}
}

func TestFind(t *testing.T) {
	gbFinder := NewGoogleBooksFinder()
	search := "test"

	details, err := gbFinder.Find(search)
	if err != nil {
		t.Errorf("GoogleBooks find returns err: %v", err)
	}

	if details == nil {
		t.Errorf("GoogleBooks find returns nil value, expected book details")
	}
}
