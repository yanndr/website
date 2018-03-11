package file

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/yanndr/website/pkg/model"
)

type fileType int

const (
	binary fileType = iota
	jsonFile
)

type profileFileRepository struct {
	f        io.ReadWriteSeeker
	mutex    *sync.RWMutex
	fileType fileType
}

func (r *profileFileRepository) Store(p *model.Profile) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	encoder := json.NewEncoder(r.f)
	return encoder.Encode(p)
}

func (r *profileFileRepository) Get() (*model.Profile, error) {
	p := &model.Profile{}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.f.Seek(0, 0)
	decoder := json.NewDecoder(r.f)
	err := decoder.Decode(p)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	return p, nil
}

//NewProfileRepository returns a new binary file profile repo.
func NewProfileRepository(rw io.ReadWriteSeeker) model.ProfileRepository {
	return &profileFileRepository{
		f:     rw,
		mutex: &sync.RWMutex{},
	}
}
