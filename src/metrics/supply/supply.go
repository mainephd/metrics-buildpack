package supply

import (
	"fmt"
	"io"
	"net/http"
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

	if err := downloadFile("https://dl.influxdata.com/telegraf/releases/telegraf-1.7.2-static_linux_amd64.tar.gz", tarDownloadDest); err != nil {
		return err
	}

	if err := libbuildpack.ExtractTarGz(tarDownloadDest, metricsBinDir); err != nil {
		return err
	}

	if err := os.Remove(tarDownloadDest); err != nil {
		s.Log.Info("Error Removing Tar file: %s", err.Error())
	}

	if err := s.Stager.AddBinDependencyLink(filepath.Join(metricsBinDir, "telegraf", "telegraf"), "telegraf"); err != nil {
		return err
	}

	return s.Stager.WriteEnvFile("PATH", metricsBinDir)
}

func downloadFile(url, destFile string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("could not download: %d", resp.StatusCode)
	}

	return writeToFile(resp.Body, destFile, 0555)
}

func writeToFile(source io.Reader, destFile string, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(destFile), 0755)
	if err != nil {
		return err
	}

	fh, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fh, source)
	if err != nil {
		return err
	}

	return nil
}
