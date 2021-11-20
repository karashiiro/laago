package laago_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/karashiiro/laago"
)

func buildDummy() error {
	cmd := exec.Command("make", "build_dummy")
	return cmd.Run()
}

func cleanup() {
	_ = os.Remove("dummy.exe")
}

func TestIsLargeAddressAware(t *testing.T) {
	err := buildDummy()
	if err != nil {
		t.Error(err)
	}

	if isLaa, err := laago.IsLargeAddressAware("dummy.exe"); isLaa {
		t.Error(err)
	}

	cleanup()
}

func TestSetLargeAddressAware(t *testing.T) {
	err := buildDummy()
	if err != nil {
		t.Error(err)
	}

	if err := laago.SetLargeAddressAware("dummy.exe"); err != nil {
		t.Error(err)
	}

	if isLaa, err := laago.IsLargeAddressAware("dummy.exe"); !isLaa {
		t.Error(err)
	}

	cleanup()
}

func TestUnsetLargeAddressAware(t *testing.T) {
	err := buildDummy()
	if err != nil {
		t.Error(err)
	}

	if err := laago.UnsetLargeAddressAware("dummy.exe"); err != nil {
		t.Error(err)
	}

	if isLaa, err := laago.IsLargeAddressAware("dummy.exe"); isLaa {
		t.Error(err)
	}

	cleanup()
}
