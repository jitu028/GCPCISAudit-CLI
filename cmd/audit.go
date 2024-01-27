// File: cmd/audit.go

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jitu028/GCPCisAudit-CLI/pkg/gemini"
	"github.com/spf13/cobra"
)

var cisQuery string

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit GCP resources against CIS benchmarks using Gemini AI",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// Retrieve the API key from the environment variable
		apiKey := os.Getenv("GOOGLE_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("GOOGLE_API_KEY environment variable is not set")
		}

		geminiClient, err := gemini.NewGeminiClient(ctx, apiKey)
		if err != nil {
			return fmt.Errorf("failed to create Gemini client: %w", err)
		}

		response, err := geminiClient.QueryCISRequirements(ctx, cisQuery)
		if err != nil {
			return fmt.Errorf("failed to query CIS requirements: %w", err)
		}

		fmt.Println("Gemini AI Response:", response)
		// Further process the response to generate summary reports and list of commands
		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)

	// Define flag for the CIS benchmark query
	auditCmd.Flags().StringVarP(&cisQuery, "query", "q", "", "CIS benchmark query")
	auditCmd.MarkFlagRequired("query")
}
