package common

import "testing"

import "os"

func TestItHasDefautPort(t *testing.T) {
	settings := Settings()

	expected := "8088"
	result := settings.Port

	if expected != result {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestItHasDefautFilePath(t *testing.T) {
	settings := Settings()

	expected := "../../../../../data/example"
	result := settings.FilePath

	if expected != result {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestUpdateByEnv(t *testing.T) {
	port := "5000"
	dir := "somedir"

	os.Setenv("PORT", port)
	os.Setenv("FILEPATH", dir)

	settings := Settings()

	if settings.FilePath != dir {
		t.Error("Expected directory %s got %s", dir, settings.FilePath)
	}

	if settings.Port != port {
		t.Error("Expected port %s got %s", dir, settings.Port)
	}
}
