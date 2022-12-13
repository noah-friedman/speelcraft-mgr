package speelcraft_mgr_test

import (
	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Define default options
var opts = godog.Options{
	Concurrency: runtime.NumCPU(),
	Format:      "progress",
	Paths:       []string{"features"},
	Randomize:   -1,
}

// Bind godog command-line flags
func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}
func TestMain(m *testing.M) {
	pflag.Parse()

	os.Exit(m.Run())
}

func TestFeatures(t *testing.T) {
	// Get paths for all feature files
	filePaths := make(map[string]bool)
	for _, p := range opts.Paths {
		if strings.TrimSpace(p) != "" {
			if f, e := os.Stat(p); e != nil {
				t.Error(e)
			} else if f.IsDir() {
				if g, e := filepath.Glob(p + "/*"); e != nil {
					t.Error(e)
				} else {
					for _, p := range g {
						filePaths[p] = true
					}
				}
			} else {
				filePaths[p] = true
			}
		}
	}

	// Get ready to run test suite(s)
	run := func(opts *godog.Options, t *testing.T) {
		(*opts).TestingT = t

		s := godog.TestSuite{
			ScenarioInitializer: stepdefs,
			Options:             opts,
		}

		if c := s.Run(); c != 0 {
			t.Errorf("godog returned non-zero exit code: %d", c)
		}
	}

	// If no feature files were found, just pass the provided array to godog runner
	if len(filePaths) == 0 {
		run(&opts, t)
	} else { // Otherwise, run a test suite for each feature file
		for f := range filePaths {
			t.Run(strings.TrimSuffix(path.Base(f), ".feature"), func(t *testing.T) {
				opts := opts
				opts.Paths = []string{f}

				run(&opts, t)
			})
		}
	}
}
