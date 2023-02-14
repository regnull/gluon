package command

import (
	"fmt"
	"github.com/ProtonMail/gluon/rfcparser"
)

type CloseCommand struct{}

func (l CloseCommand) String() string {
	return fmt.Sprintf("CLOSE")
}

func (l CloseCommand) SanitizedString() string {
	return l.String()
}

type CloseCommandParser struct{}

func (CloseCommandParser) FromParser(p *rfcparser.Parser) (Payload, error) {
	return &CloseCommand{}, nil
}
