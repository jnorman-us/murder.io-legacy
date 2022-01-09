package dummy

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"log"
)

type Listener struct {
}

func (l *Listener) GetChannel() string {
	return "PLAYER_INPUT"
}

func (l *Listener) HandleData(d *gob.Decoder) {
	inputs := &types.Input{}
	var err = d.Decode(inputs)

	if err != nil {
		log.Fatal("decoder error in lsitener:", err)
	}

	fmt.Println("Received: ", inputs)
}
