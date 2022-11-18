package translater

/*
{
  "message": {
    "result": {
      "srcLangType": "ko",
      "tarLangType": "en",
      "translatedText": "How are you doing?",
      "engineType": "PRETRANS",
      "pivot": null,
      "dict": null,
      "tarDict": null
    },
    "@type": "response",
    "@service": "naverservice.nmt.proxy",
    "@version": "1.0.0"
  }
}
*/

type NaverPapagoResult struct {
	SrcLangType    string `json:"srcLangType"`
	TarLangType    string `json:"tarLangType"`
	TranslatedText string `json:"translatedText"`
	EngineType     string `json:"engineType"`
	Pivot          string `json:"pivot"`
	Dict           string `json:"dict"`
	TarDict        string `json:"tarDict"`
}

type NaverPapagoMessage struct {
	Result  NaverPapagoResult `json:"result"`
	MsgType string            `json:"@type"`
	Service string            `json:"@service"`
	Version string            `json:"@version"`
}

type NaverPapagoRespon struct {
	Message NaverPapagoMessage `json:"message"`
}
