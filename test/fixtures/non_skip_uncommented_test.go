package basecase

import (
	"os"
	"testing"

	"github.com/terraform-modules-krish/terratest/modules/logger"
	test_structure "github.com/terraform-modules-krish/terratest/modules/test-structure"
)

func TestWithStages(t *testing.T) {
	t.Parallel()

	// Uncomment the items below to skip certain parts of the test
	os.Setenv("THE_ANSWER", "42")
	//os.Setenv("SKIP_setup", "true")
	//os.Setenv("SKIP_deploy", "true")
	//os.Setenv("SKIP_validate", "true")
	//os.Setenv("SKIP_cleanup", "true")

	test_structure.RunTestStage(t, "setup", func() {
		logger.Logf(t, "setup")
	})

	defer test_structure.RunTestStage(t, "cleanup", func() {
		logger.Logf(t, "cleanup")
	})

	test_structure.RunTestStage(t, "deploy", func() {
		logger.Logf(t, "deploy")
	})

	test_structure.RunTestStage(t, "validate", func() {
		logger.Logf(t, "validate")
	})
}