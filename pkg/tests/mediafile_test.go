package tests

import (
	"fmt"
	"github.com/parvez0/wabacli/pkg/cmd/send"
	"gotest.tools/assert"
	"path/filepath"
	"testing"
)

func TestLocalMediaProcess(t *testing.T)  {
	files := []string{ "../../assets/test.jpeg", "test.jpeg" }
	path, _ := filepath.Abs("../assets/test.jpeg")
	for i, v := range files {
		t.Run(fmt.Sprintf("Running Test case %d", i), func(t *testing.T) {
			file, err := send.NewFileReader(v)
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
			want := send.Media{
				Path:     path,
				MimeType: send.MediaTypeMapping[".jpeg"],
			}
			t.Logf("%+v ---- %+v", file, want)
			assert.Equal(t, want.MimeType, file.MimeType)
			assert.Assert(t, file.Reader != nil)
		})
	}
}
