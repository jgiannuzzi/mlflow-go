//go:build mage

//nolint:wrapcheck
package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Start the mlflow-go dev server connecting to postgres.
func Dev() error {
	mg.Deps(Generate)

	return sh.RunWithV(
		map[string]string{
			"PYTHON_LOGGING_LEVEL":        "DEBUG",
			"MLFLOW_SQLALCHEMYSTORE_ECHO": "True",
		},
		"mlflow-go",
		"server",
		"--backend-store-uri",
		"postgresql://postgres:postgres@localhost:5432/postgres",
		"--go-opts",
		"log_level=debug,shutdown_timeout=5s",
	)
}
