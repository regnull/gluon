package rfc822

import (
	"fmt"
	"io"
)

type headerParser struct {
	header []byte
	offset int
}

func newHeaderParser(header []byte) headerParser {
	return headerParser{header: header}
}

// next will keep parsing until it collects a new entry. io.EOF is returned when there is nothing left to parse.
func (hp *headerParser) next() (parsedHeaderEntry, error) {
	headerLen := len(hp.header)

	if hp.offset >= headerLen {
		return parsedHeaderEntry{}, io.EOF
	}

	result := parsedHeaderEntry{
		keyStart:   hp.offset,
		keyEnd:     -1,
		valueStart: -1,
		valueEnd:   -1,
	}

	// Detect key, have to handle prelude case where there is no header information or last empty new line.
	{
		for hp.offset < headerLen {
			if hp.header[hp.offset] == ':' {
				prevOffset := hp.offset
				hp.offset++
				if hp.offset < headerLen && (hp.header[hp.offset] == ' ' || hp.header[hp.offset] == '\r' || hp.header[hp.offset] == '\n') {
					result.keyEnd = prevOffset

					// Validate the header key.
					for i := result.keyStart; i < result.keyEnd; i++ {
						if v := hp.header[i]; v < 33 || v > 126 {
							return parsedHeaderEntry{}, ErrNonASCIIHeaderKey
						}
					}

					break
				}
			} else if hp.header[hp.offset] == '\n' {
				hp.offset++

				result.keyEnd = result.keyStart
				result.valueStart = result.keyStart
				result.valueEnd = hp.offset

				return result, nil
			} else {
				hp.offset++
			}
		}

		if result.keyEnd == -1 {
			return parsedHeaderEntry{}, ErrKeyNotFound
		}
	}

	// collect value.
	searchOffset := result.keyEnd + 2
	result.valueStart = searchOffset

	for searchOffset < headerLen {
		b := hp.header[searchOffset]

		if b == '\r' {
			searchOffset++
			if searchOffset >= headerLen {
				return parsedHeaderEntry{}, io.ErrUnexpectedEOF
			}

			if hp.header[searchOffset] != '\n' {
				return parsedHeaderEntry{}, fmt.Errorf(`expected \n after \n`)
			}
			searchOffset++

			// If the next character after new line is a space, it's a fold
			if searchOffset < headerLen && isWSP(hp.header[searchOffset]) {
				continue
			}

			result.valueEnd = searchOffset

			break
		} else if b == '\n' {
			searchOffset++

			// If the next character after new line is a space, it's a fold
			if searchOffset < headerLen && isWSP(hp.header[searchOffset]) {
				continue
			}

			result.valueEnd = searchOffset

			break
		} else {
			searchOffset++
		}
	}

	hp.offset = searchOffset

	// handle case where we may have reached EOF without concluding any previous processing.
	if result.valueEnd == -1 && searchOffset >= headerLen {
		result.valueEnd = headerLen
	}

	return result, nil
}

func isWSP(b byte) bool {
	return b == ' ' || b == '\t'
}

type parsedHeaderEntry struct {
	keyStart   int
	keyEnd     int
	valueStart int
	valueEnd   int
}

func (p parsedHeaderEntry) hasKey() bool {
	return p.keyStart != p.keyEnd
}

func (p parsedHeaderEntry) getKey(header []byte) []byte {
	return header[p.keyStart:p.keyEnd]
}

func (p parsedHeaderEntry) getValue(header []byte) []byte {
	return header[p.valueStart:p.valueEnd]
}

func (p parsedHeaderEntry) getAll(header []byte) []byte {
	return header[p.keyStart:p.valueEnd]
}

func (p *parsedHeaderEntry) applyOffset(offset int) {
	p.keyStart += offset
	p.keyEnd += offset
	p.valueStart += offset
	p.valueEnd += offset
}
