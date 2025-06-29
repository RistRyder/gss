# gss - Go Subtitle Suite
This library aims to provide functionality for manipulating subtitles and is currently pre-release.

The audio transcription API is currently able to generate text from an input audio or video file using the following:
* [Vosk](https://alphacephei.com/vosk/)
* [Whisper on Windows](https://github.com/Purfview/whisper-standalone-win)

The following utilities are currently available:
* [Join](https://github.com/RistRyder/gss/blob/main/examples/util/join/main.go) - Join a collection of subtitles
* [Overlap](https://github.com/RistRyder/gss/blob/main/examples/fix/overlap/main.go) - Fix overlapping subtitle lines e.g. after a Join operation

## Examples
See [Here](https://github.com/RistRyder/gss/blob/main/examples/)

## License
`gss` is licensed under the GNU LESSER GENERAL PUBLIC LICENSE Version 3, 
so it free to use for commercial software, as long as you don't modify the library itself. 
LGPL 3.0 allows linking to the library in a way that doesn't require you to open source your own code. 
This means that if you use `gss` in your project, you can keep your own code private, 
as long as you don't modify `gss` itself.