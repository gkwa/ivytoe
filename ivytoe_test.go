package ivytoe_test

import (
	"os"
	"testing"

	"github.com/taylormonacelli/ivytoe"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"ivytoe": ivytoe.Main,
	}))
}

func TestIvytoe(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
