package recommendations

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

type GearToCarry struct {
	Gear []string `json:"gear"`
}

func RecommendGearToCarry(weather string) string {
	projectID := os.Getenv("PROJECT_ID")
	locationEnv := os.Getenv("LOCATION")
	modelName := os.Getenv("MODEL_NAME")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, locationEnv)
	if err != nil {
		fmt.Println(err)
	}
	gemini := client.GenerativeModel(modelName)
	text := fmt.Sprintf("Based on the weather %s, recommend gear to carry. Provide a simple list of items.", weather)
	prompt := genai.Text(text)

	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return fmt.Sprintf("Error generating content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No content generated"
	}

	textResponse, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "Unexpected response format"
	}

	return string(textResponse)
}
