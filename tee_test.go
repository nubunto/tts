package tts

import "testing"

func TestSpeak(t *testing.T) {
	speech, err := Speak(Config{Language: "en-US", Speak: "hello"})
	if err != nil {
		t.Error("Speak should not return a error")
	}

	audio := speech.Bytes()
	if len(audio) == 0 {
		t.Error("Should not return an empty audio")
	}
}
