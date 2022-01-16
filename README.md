# telegrambot-notify
Notification lib to be used as companion to telegrambot

## Sample use case

```go
func TestNotify(t *testing.T) {

	Init("https://url/pub", "yourkey")

	for i := 0; i < 10; i++ {
		Notify("a", "Some from test")
	}

}
```