package types

import (
	"fmt"
	"github.com/parvez0/wabacli/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Media object for storing the information about file,
// which will be uploaded to whatsapp
type Media struct {
	Path     string `validate:"required"`
	Size     int64
	MimeType MediaType
	Data   	 []byte
}

// NewFileReader returns a new object of media type
// it will also convert the file path to absolute
func NewFileReader(path string) (*Media, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &Media{
		Path: abs,
	}, nil
}

// Read checks if the file exits by checking
// it's stats if file is not present it will
// return an error other wise it will process
// link to find the mime type and returns a
// []bytes for the files data
func (f *Media) Read() error {
	stats, err := os.Stat(f.Path)
	if err != nil {
		return err
	}
	reader, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer reader.Close()
	f.Size = stats.Size()
	err = f.findMimeType()
	if err != nil {
		return err
	}
	f.Data, err = ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return nil
}

// findMimeType will map the extension of the file
// to proper mime type, if extension not found in
// MediaTypeMapping it will return an unsupported
// file error
func (f *Media) findMimeType() error {
	if f.MimeType != ""{
		return nil
	}
	log.Debug("finding file extension for '", f.Path, "'")
	ext := filepath.Ext(f.Path)
	log.Debug(fmt.Sprintf("file has extension '%s'", ext))
	f.MimeType = MediaTypeMapping[ext]
	if f.MimeType == "" {
		return fmt.Errorf("upsupported media type, please refer '%s/api/media'", FacebookSupportUrl)
	}
	log.Debug(fmt.Sprintf("file has mime-type '%s'", f.MimeType))
	return nil
}

// GetMimeType will return the file mime type
// which is of type MediaType converted to string
func (f *Media) GetMimeType() string {
	return string(f.MimeType)
}