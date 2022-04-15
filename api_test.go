package googletranslate

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	res, err := Translate("hello", En, Zh)
	if err != nil {
		t.Errorf("Translate() error = %v", err)
		return
	}
	if res != "你好" {
		t.Errorf("Translate() got = %v, want %v", res, "你好")
	}
}
