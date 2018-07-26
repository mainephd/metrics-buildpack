package finalize

import (
	"github.com/cloudfoundry/libbuildpack"
)

type Finalizer struct {
	BuildDir string
	DepDir   string
	Log      *libbuildpack.Logger
}

func New(stager *libbuildpack.Stager, logger *libbuildpack.Logger) *Finalizer {
	return &Finalizer{
		BuildDir: stager.BuildDir(),
		DepDir:   stager.DepDir(),
		Log:      logger,
	}
}

func (f *Finalizer) Run(sf *Finalizer) error {
	return nil
}
