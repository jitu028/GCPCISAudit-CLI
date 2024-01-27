// File: pkg/gemini/query.go

package gemini

import (
	"context"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func (g *GeminiClient) QueryCISRequirements(ctx context.Context, query string) (string, error) {
	// Replace 'modelName' with the appropriate Gemini AI model name
	modelName := "models/gemini-pro"
	model := g.Client.GenerativeModel(modelName)

	res, err := model.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		return "", err
	}

	// Assuming the response contains a slice of candidates with text content
	var responseBuilder strings.Builder
	for _, candidate := range res.Candidates {
		for _, part := range candidate.Content.Parts {
			if textPart, ok := part.(genai.Text); ok {
				responseBuilder.WriteString(string(textPart))
			}
		}
	}

	return responseBuilder.String(), nil
}
