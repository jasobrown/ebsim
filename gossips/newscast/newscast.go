package newscast

import (
	"github.com/jasobrown/ebsim/core"
)

type Newscast struct {
}

func NewNewscast() (newscast *Newscast, config *core.Config) {
	config = core.NewConfig()
	config.NodeCnt = uint16(5000)

	return &Newscast{}, config
}

func (newscast *Newscast) NextRound() {

}

func (newscast *Newscast) Execute() {

}
