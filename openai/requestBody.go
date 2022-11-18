package openai

/*
{
  "model": "text-davinci-002",
  "prompt": "Say this is a test",
  "max_tokens": 6,
  "temperature": 0,
  "top_p": 1,
  "n": 1,
  "stream": false,
  "logprobs": null,
  "stop": "\n"
}

*/

type RequestBody struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxToken    int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
	// TopP        int     `json:"top_p"`
	// N           int     `json:"n"`
	// Stream      bool    `json:"stream"`
	// Logprobs    string  `json:"logprobs"`
	// Stop        string  `json:"stop"`
}

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
