package openai

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestResponseBodyStruct(t *testing.T) {

	buf := `{"id":"cmpl-6DoPudv6kgf3AaZPAg5SQW0gSZSLl","object":"text_completion","created":1668750086,"model":"text-davinci-002","choices":[{"text":"\n\nI'd like to","index":0,"logprobs":null,"finish_reason":"length"}],"usage":{"prompt_tokens":1,"completion_tokens":6,"total_tokens":7}}`

	var result ResponBody

	err := json.Unmarshal([]byte(buf), &result)
	if err != nil {
		t.Errorf("Wrong result %v", err)
	}

	if len(result.Choices) == 0 {
		t.Errorf("Wrong result 0")
	}

	text := strings.Trim(result.Choices[0].Text, "\n")
	if text != "I'd like to" {
		t.Errorf("Wrong result %v", text)
	}
}
