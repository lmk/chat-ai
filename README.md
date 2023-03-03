# chat-ai
간단하게 AI와 채팅을 구현 하고자 합니다. 
AI는 `https://beta.openai.com/docs/guides/completion/conversation`의 내용을 사용합니다.
openai.com 에서는 영문만 지원하기 때문에, naver의 papago api를 사용해서 번역합니다.

## 제약사항
- openai.com은 가입하면 베타 버전이라 3달간만 무료로 사용할 수 있습니다. 서비스가 아닌 스터디 목적이니 3달만 사용하려 합니다. ~~(만기: 23년 1월 중순)~~ 만료됨
- naver papago api는 하루 10,000 글자만 번역이 무료 입니다. 번역 대상은 `내가 입력한 내용` + `AI가 대답한 내용` 입니다. 10,000글자를 초과하면 translater fail이 발생합니다. 

## app.yaml

설정을 위해 `app.yaml` 파일이 필요합니다. 

```yaml
listenPort: 8888
naverClientId: Naver Client ID
naverClientSecret: Naver Client Secret
openai:
  apiKey: OpenAI Client Secret
  model: text-davinci-003
  max_tokens: 150
  temperature: 0.9
  top_p: 1
  frequency_penalty: 0
  presence_penalty: 0.6
```

- `https://openapi.naver.com/v1/papago/n2mt` 연동을 위한 ID와 key가 필요합니다.
- `https://openai.com/api/` 연동을 위해 API key가 필요합니다.
- 구현을 시작할땐 `text-davinci-002`를 사용했는데, 구현 중간에 `text-davinci-003`이 나왔습니다. 답변이 더 향상 되었다고 합니다. 단점은 답변이 길이졌네요. 
- 기타 openai 관련 설정 값은, `https://beta.openai.com/playground/p/default-chat`을 참고했습니다.

## 구조

```text
main ㅡ openai
     ㄴ translater
     ㄴ public
```

main 패키지 아래 openai, translater 두개의 패키지로 분리했습니다. public 에는 채팅 UI를 구현한 html과 css 파일을 위치했습니다.

## 흐름

- 내 채팅은 `web-browser` -> `rest api` -> `papago api` -> `openai api` 순으로 요청하고, AI 채팅은 역순입니다.
- 소스 파일 기준으로는 `public/index.html` -> `App.go` -> `translater/translater.go` -> `openai/completion.go` 순입니다.

## code review

### `public/index.html`

- 처음에는 native javascript로 구현하려 했는데, 코드가 길어지고 가독성이 떨어져서 jquery를 사용했습니다.

- 종이 비행기 버튼 사용을 위해 fontawesome을 사용했습니다.

```html
<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
<button class="btn fa fa-paper-plane" type="submit"></button>
```

- AI와 연결된 대화를 위해 전체 대화 내용이 필요합니다. `chat-all`라는 hidden 엘리먼트에 저장합니다.

```html
<textarea class="hidden"  id="chat-all"></textarea>
```

- 채팅 인터페이스는 최대한 간단하게 li 태그를 사용했고, child로 div에 영문을 넣었습니다. chat-box가 처음에는 내용이 없지만, 채팅을 입력하면 동적으로 추가됩니다.

```html 
<ul id="chat-box">
    <li class="chat-box me">반가워!
        <div class="chat-appendix">Nice to meet you!</div>
    </li>
    <li class="chat-box ai">저도 만나서 반가워요!
        <div class="chat-appendix">Nice to meet you too!</div>
    </li>
```

- 시각적인 효과는 css를 활용했습니다.

```css
.chat-box {
  padding: 20px 30px;
  margin: 4px 4px;
  border-radius: 10px;
  clear:both;
}

.me {
  background-color: #a9edc1;
  float: right;
}

.ai {
  background-color: #e9e9e9;
  float: left;
}
```

- 대화 내용이 입력되면, rest api를 호출 합니다.

```javascript
    // 채팅 데이터 전송
    function requestMsg(msg) {
      var chatAll =  $('#chat-all').val()
      $.post( "/api/v1/chat", { text: chatAll, say: msg })
        .done(function( data ) {
```

### `App.go`

go에서 가장 많이 사용한다는 gin을 사용해서 rest api 서버를 구현했습니다.

```go
	r.Run(fmt.Sprintf(":%d", conf.ListenPort))
```

하위 디렉토리 라우팅은 아래와 같이 합니다.

```go
r.Use(static.Serve("/", static.LocalFile("public", true)))
```

POST 라우팅과 Request 데이터를 받는 방법은 아래와 같습니다.

```go
func StartWeb() {
	r := gin.Default()
	r.POST("/api/v1/chat", func(c *gin.Context) {

		text := c.PostForm("text")
    reqKrMsg := c.PostForm("say")
```

Response은 아래와 같이 합니다.

```go
		c.JSON(200, gin.H{
			"text":   text,
			"req-en": reqEnMsg,
			"req-kr": reqKrMsg,
			"res-en": resEnMsg,
			"res-kr": resKrMsg,
		})
```

연결된 대화를 위해 모든 대화 내용을 학습합니다.

```go
		reqKrMsg := c.PostForm("say")
		reqEnMsg, err := translater.Ko2En(reqKrMsg)
		text += openai.ME + reqEnMsg + "\n"
		resEnMsg, err := openai.Chat(text)
```

### `translater/translater.go`

rest client는 resty를 사용했습니다. 
파파고 api는 body에 소스 언어, 타켓 언어 그리고 번역할 텍스트를 전달합니다.

```go
	reqUrl := NewReqUrl("").
		SetParam("source", src).
		SetParam("target", dst).
		SetParam("text", msg)

	client := resty.New()
	respJson, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("X-Naver-Client-Id", NaverClientId).
		SetHeader("X-Naver-Client-Secret", NaverClientSecret).
		SetBody(reqUrl.Get()).
		Post(URI)
```

결과는 body에 json으로 담겨 옵니다.
json 파싱은 로직에서 하지 않고, 구조체를 만들어서 `json.Unmarshal()` 합니다.

```go
	var result NaverPapagoRespon

	err = json.Unmarshal(respJson.Body(), &result)
```

### `openai/completion.go`

요청 파라미터는 `https://beta.openai.com/examples/default-friend-chat`의 값을 참고했습니다.

```json
  "Model": "text-davinci-003",
  "MaxTokens": 60,
  "Temperature": 0.5,
  "TopP": 1,
  "FrequencyPenalty": 0.5,
  "PresencePenalty": 0,
  "stop": ["You:"]
```

연결된 대화를 위해, 매번 전체 메시지를 학습합니다.

```text
You:Nice to meet you!
Friend:Nice to meet you too!
You:Who are you?
Friend:I'm your new neighbor. I just moved in down the street.
You:Congratulations!
Friend:Thank you!
```

AI의 이름이 대화를 시작할때마다 다르게 와서, ":" 구분자로 메시지를 잘랐습니다.

```go
func StripPrefix(text string) string {
	i := strings.LastIndex(text, ":")

	if i >= 0 {
		text = strings.Trim(text[i+1:], " ")
	}

	return text
}
```

## 참조

- https://beta.openai.com/docs/guides/completion/conversation
- https://developers.naver.com/products/papago/nmt/nmt.md
- ~~https://chatai.newtype.dev/~~ 만료됨
- https://github.com/lmk/chat-ai
