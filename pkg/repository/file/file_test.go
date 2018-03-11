package file_test

import (
	"log"
	"sync"
	"testing"

	"github.com/mattetti/filebuffer"
	"github.com/yanndr/website/pkg/model"
	"github.com/yanndr/website/pkg/repository/file"
)

func setup(t *testing.T) model.ProfileRepository {
	t.Parallel()
	return file.NewProfileRepository(filebuffer.New(nil))
}

func TestStore(t *testing.T) {
	const name = "yann"
	r := setup(t)

	err := r.Store(&model.Profile{Firstname: name})
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

func TestMultipleGet(t *testing.T) {
	const name = "yann"
	r := setup(t)

	err := r.Store(&model.Profile{Firstname: name})
	if err != nil {
		t.Error(err)
		return
	}

	const loop = 100
	wg := sync.WaitGroup{}
	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func(i int) {
			p, err := r.Get()
			log.Println("profile -", i, " :", p)
			if err != nil {
				t.Error(err)
				return
			}

			if p.Firstname != name {
				t.Errorf("expected %s, got %s", name, p.Firstname)
			}
			wg.Done()
		}(i + 1)
	}

	wg.Wait()
	log.Println("done")
}
