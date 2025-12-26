package migration

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/genai"
)

func generateSummary(description string) string {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  "",
		Location: "",
		Backend:  genai.BackendVertexAI,
		HTTPOptions: genai.HTTPOptions{
			APIVersion: "v1",
			Headers: http.Header{
				// Options:
				// - "dedicated": Use Provisioned Throughput
				// - "shared": Use pay-as-you-go
				// https://cloud.google.com/vertex-ai/generative-ai/docs/use-provisioned-throughput
				"X-Vertex-AI-LLM-Request-Type": []string{"shared"},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-pro-preview",
		genai.Text(fmt.Sprintf(`You are a news summarization service.
			Summarize the following article in 2–3 sentences.
			Rules:
			- Return ONLY plain text
			- Do NOT use headings, markdown, bullets, or formatting
			- Do NOT add introductory phrases
			- Do NOT add conclusions or opinions
			Content:
			%s`, description)),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return result.Text()
}
