package telegrambotnotify

import "testing"

func TestNotify(t *testing.T) {

	Init("https://url/pub", "apikey")

	for i := 0; i < 10; i++ {
		Notify("a", "Some from test")
	}

}
