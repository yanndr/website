package file

import (
	"encoding/gob"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/yanndr/website/model"
)

type fileType int

const (
	binary fileType = iota
	jsonFile
)

type profileFileRepository struct {
	filePath string
	mutex    *sync.RWMutex
	fileType fileType
}

func (r *profileFileRepository) Store(p *model.Profile) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()
	file, err := os.OpenFile(r.filePath, os.O_WRONLY, 0755)
	defer file.Close()
	if err != nil {
		return err
	}

	encoder := r.getEncoder(file)
	return encoder.Encode(p)
}

func (r *profileFileRepository) Get() (*model.Profile, error) {
	p := &model.Profile{}
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	file, err := os.Open(r.filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	decoder := r.getDecoder(file)
	err = decoder.Decode(p)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	return p, nil
}

type decoder interface {
	Decode(interface{}) error
}
type encoder interface {
	Encode(interface{}) error
}

func (r *profileFileRepository) getDecoder(reader io.Reader) decoder {
	if r.fileType == binary {
		return gob.NewDecoder(reader)
	}

	return json.NewDecoder(reader)
}

func (r *profileFileRepository) getEncoder(w io.Writer) encoder {
	if r.fileType == binary {
		return gob.NewEncoder(w)
	}

	return json.NewEncoder(w)
}

//NewProfileRepository returns a new binary file profile repo.
func NewProfileRepository(filePath string) (model.ProfileRepository, error) {
	r, err := newfileRepository(filePath)
	if err != nil {
		return nil, err
	}
	r.fileType = binary

	return r, nil
}

func newfileRepository(filePath string) (*profileFileRepository, error) {
	r := &profileFileRepository{
		filePath: filePath,
		mutex:    &sync.RWMutex{},
	}

	err := createFileIfNotExist(filePath)
	if err != nil {
		return nil, err
	}

	return r, nil
}

//NewJSONProfileRepository returns a new Json file profile repo.
func NewJSONProfileRepository(filePath string) (model.ProfileRepository, error) {
	r, err := newfileRepository(filePath)
	if err != nil {
		return nil, err
	}
	r.fileType = jsonFile

	return r, nil
}

func createFileIfNotExist(filePath string) error {
	var file *os.File
	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	file, err = os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}

	return nil
}
