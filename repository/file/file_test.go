package file_test

import (
	"os"
	"testing"

	"github.com/yanndr/website/model"
	"github.com/yanndr/website/repository/file"
)

func TestStore(t *testing.T) {
	const name = "yann"
	defer os.Remove("test.dat")
	r, err := file.NewProfileRepository("test.dat")
	if err != nil {
		t.Error(err)
		return
	}

	err = r.Store(&model.Profile{Firstname: name})
	if err != nil {
		t.Error(err)
		return
	}

	p, err := r.Get()
	if err != nil {
		t.Error(err)
		return
	}

	if p.Firstname != name {
		t.Errorf("expected %s, got %s", name, p.Firstname)
	}

}
