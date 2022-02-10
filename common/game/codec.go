package game

import "encoding/gob"

func (g *Game) GetChannel() byte {
	return 0x03
}

func (g *Game) Flush() {

}

func (g *Game) GetData(encoder *gob.Encoder) error {
	return encoder.Encode(g)
}

func (g *Game) HandleData(decoder *gob.Decoder) error {
	var newGame = &Game{}

	err := decoder.Decode(newGame)
	if err != nil {
		return err
	}

	// copy the contents of newgame into the game

	return nil
}
