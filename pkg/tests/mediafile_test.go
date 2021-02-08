package tests

import (
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"gotest.tools/assert"
	"path/filepath"
	"testing"
)

func TestLocalMediaProcess(t *testing.T)  {
	files := []string{ "../../assets/test.jpeg", "test.jpeg" }
	path, _ := filepath.Abs("../assets/test.jpeg")
	for i, v := range files {
		t.Run(fmt.Sprintf("Running Test case %d", i), func(t *testing.T) {
			file, err := types.NewFileReader(v)
			if err != nil {
				t.Errorf("failed to create file object: %v", err)
				t.Failed()
			}
			err = file.Read()
			if err != nil && i == 0 {
				assert.NilError(t, err)
			} else {
				return
			}
			want := types.Media{
				Path:     path,
				MimeType: types.MediaTypeMapping[".jpeg"],
			}
			t.Logf("%+v ---- %+v", file, want)
			assert.Equal(t, want.MimeType, file.MimeType)
			assert.Assert(t, file.Data != nil)
		})
	}
}

func TestUploadMedia(t *testing.T) {
	conf, err := config.GetConfig()
	assert.NilError(t, err, "failed to initialize config")
	file, err := types.NewFileReader("../../assets/test.jpeg")
	assert.NilError(t, err, "reader not created")
	err = file.Read()
	assert.NilError(t, err, "failed to read file")
	byts, err := helpers.UploadMedia(&conf.CurrentCluster, file)
	assert.NilError(t, err, "failed to upload media")
	t.Logf("uploaded file - %s", string(byts))
}