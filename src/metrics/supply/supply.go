package supply

import (
	"os"
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
	metricsInstallDir := filepath.Join(s.Stager.DepDir(), "metrics")

	tarDownloadDest := filepath.Join(metricsInstallDir, "tmp", "telegraf.tar.gz")

	metricsBinDir := filepath.Join(metricsInstallDir, "bin")

	s.Log.Info("Downloading Telegraf...")
	dependency, err := s.Manifest.DefaultVersion("telegraf")
	if err != nil {
		return err
	}

	if err := s.Manifest.FetchDependency(dependency, tarDownloadDest); err != nil {
		return err
	}

	s.Log.Info("Extracting Telegraf...")
	if err := libbuildpack.ExtractTarGz(tarDownloadDest, metricsBinDir); err != nil {
		return err
	}

	if err := os.Remove(tarDownloadDest); err != nil {
		s.Log.Info("Error Removing Tar file: %s", err.Error())
	}

	if err := os.Remove(filepath.Join(metricsInstallDir, "tmp")); err != nil {
		s.Log.Info("Error Removing tmp directory: %s", err.Error())
	}

	if err := s.Stager.AddBinDependencyLink(filepath.Join(metricsBinDir, "telegraf", "telegraf"), "telegraf"); err != nil {
		return err
	}

	return s.Stager.WriteEnvFile("PATH", metricsBinDir)
}
