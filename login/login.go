package login

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
)

func relayLoginToAgent() {
	queue := cellnet.NewEventQueue()
	peer.NewGenericPeer("tcp.Connector")
}