package worldout

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Datafiable interface {
	GetID() types.ID
	GetClass() types.Channel
	GetSchema() packets.Schema
	PopulateData(data *packets.Data)
}

func (m *Manager) AddDatafiable(id types.ID, d *Datafiable) {
	var channel = (*d).GetClass()
	var schema = (*d).GetSchema()
	var stream, ok = m.streams[channel]
	if !ok {
		stream = m.packets.CreateStream(channel, true)
		m.streams[channel] = stream
	}
	var data = stream.CreateData(id, schema)
	m.data[id] = data
	m.datafiables[id] = d
}

func (m *Manager) RemoveDatafiable(id types.ID) {
	var datafiable = m.datafiables[id]
	var channel = (*datafiable).GetClass()
	m.streams[channel].DeleteData(id)
	delete(m.data, id)
	delete(m.datafiables, id)
}
