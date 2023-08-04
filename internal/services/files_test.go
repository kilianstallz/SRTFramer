package services

import "testing"

func TestCreateOutputPath(t *testing.T) {
	outPath := "test"
	expected := "test/frames"

	actual := CreateOutputPath(outPath)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	outPath = "."
	expected = "frames"

	actual = CreateOutputPath(outPath)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
