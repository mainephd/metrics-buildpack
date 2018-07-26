package supply

import (
	"metrics/data"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

type Supplier struct {
	Stager   *libbuildpack.Stager
	Manifest *libbuildpack.Manifest
	Log      *libbuildpack.Logger
}

func New(stager *libbuildpack.Stager, manifest *libbuildpack.Manifest, logger *libbuildpack.Logger) *Supplier {
	return &Supplier{
		Stager:   stager,
		Manifest: manifest,
		Log:      logger,
	}
}

func (s *Supplier) Run() error {
	return installTelegraf(s)
}

func installTelegraf(s *Supplier) error {

	metricsBinDir := filepath.Join(s.Stager.DepDir(), "metrics", "bin")

	dependency, err := s.Manifest.DefaultVersion("telegraf")
	if err != nil {
		return err
	}

	if err := s.Manifest.InstallDependency(dependency, metricsBinDir); err != nil {
		return err
	}

	if err := s.Stager.AddBinDependencyLink(filepath.Join(metricsBinDir, "telegraf", "telegraf"), "telegraf"); err != nil {
		return err
	}

	if err := s.Stager.WriteProfileD("telegraf.sh", data.TelegrafBackgroundScript()); err != nil {
		return err
	}

	return s.Stager.WriteEnvFile("PATH", metricsBinDir)
}
