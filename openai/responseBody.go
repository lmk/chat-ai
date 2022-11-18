package openai

/*
{
  "id": "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7",
  "object": "text_completion",
  "created": 1589478378,
  "model": "text-davinci-002",
  "choices": [
    {
      "text": "\n\nThis is a test",
      "index": 0,
      "logprobs": null,
      "finish_reason": "length"
    }
  ],
  "usage": {
    "prompt_tokens": 5,
    "completion_tokens": 6,
    "total_tokens": 11
  }
}
*/

type ResponBody struct {
	Id      string         `json:"id"`
	Object  string         `json:"object"`
	Created uint32         `json:"created"`
	Model   string         `json:"model"`
	Choices []ResponChoice `json:"choices"`
	Usage   ResponUsage    `json:"usage"`
}

type ResponChoice struct {
	Text         string `json:"text"`
	Index        uint32 `json:"index"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type ResponUsage struct {
	PromptTokens     uint32 `json:"prompt_tokens"`
	CompletionTokens uint32 `json:"completion_tokens"`
	TotalTokens      uint32 `json:"total_tokens"`
}
