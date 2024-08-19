package val

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/f1w/golio/api"
	"github.com/f1w/golio/internal"
	"github.com/f1w/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), log.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
