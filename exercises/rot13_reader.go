package exercises

import (
	"io"
	"os"
	"strings"
)

//Exercise: rot13Reader

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (n int, err error) {
	// reading data from nested io.Reader into rot13Reader
	n, err = rr.r.Read(p)

	if err != nil {
		return n, err
	}

	rot13 := func(b byte) byte {
		switch {
		case b >= 'A' && b <= 'Z':
			return 'A' + (b-'A'+13)%26
		case b >= 'a' && b <= 'z':
			return 'a' + (b-'a'+13)%26
		default:
			return b
		}
	}

	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}

	return n, nil
}

func Testrot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
