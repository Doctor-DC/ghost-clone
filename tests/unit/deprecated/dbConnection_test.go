package deprecated

import (
	"github.com/cyantarek/ghost-clone-project/db"
	"testing"
)

func TestDBConnection(t *testing.T) {
	_, err := db.NewPsql()
	if err != nil {
		t.Fatal(err)
	}
}
