# TTS, from Google Translate

Google Translate API wrapper in Go for converting Text to Speech.

## Example

### Save speech to a file

```
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nubunto/tts"
)

func main() {
	s, err := tts.Speak(tts.Config{
		Speak:    "olá!",
		Language: "pt-BR",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("output.mp3", s.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
```

**Note:** Highly experimental. Not sure if Google imposes any kind of rate-limiting, but I've decided not to treat for that for now. Documentation is scarse (non-existent, to be honest) and tests are, well, optmistic, to say the least. Use it at your own risk.
