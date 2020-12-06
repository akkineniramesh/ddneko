package types

import "net/http"

type MemberProfile struct {
	Secret             string
	Name               string

	IsAdmin            bool
	CanLogin           bool
	CanConnect         bool
	CanWatch           bool
	CanHost            bool
	CanAccessClipboard bool
}

type MembersDatabase interface {
	Connect() error
	Disconnect() error

	Insert(id string, profile MemberProfile) error	
	Update(id string, profile MemberProfile) error	
	Delete(id string) error	
	Select() (map[string]MemberProfile, error)
}

type Session interface {
	ID() string

	VerifySecret(secret string) bool
	Name() string
	IsAdmin() bool
	CanLogin() bool
	CanConnect() bool
	CanWatch() bool
	CanHost() bool
	CanAccessClipboard() bool

	IsHost() bool
	IsConnected() bool
	IsReceiving() bool
	Disconnect(reason string) error

	SetWebSocketPeer(websocket_peer WebSocketPeer)
	SetWebSocketConnected(connected bool)
	Send(v interface{}) error

	SetWebRTCPeer(webrtc_peer WebRTCPeer)
	SetWebRTCConnected(connected bool)
	SignalAnswer(sdp string) error
}

type SessionManager interface {
	Connect() error
	Disconnect() error

	Create(id string, profile MemberProfile) (Session, error)
	Update(id string, profile MemberProfile) error
	Get(id string) (Session, bool)
	Delete(id string) error

	SetHost(host Session)
	GetHost() Session
	ClearHost()

	HasConnectedMembers() bool
	Members() []Session
	Broadcast(v interface{}, exclude interface{})
	AdminBroadcast(v interface{}, exclude interface{})

	OnHost(listener func(session Session))
	OnHostCleared(listener func(session Session))
	OnCreated(listener func(session Session))
	OnDeleted(listener func(session Session))
	OnConnected(listener func(session Session))
	OnDisconnected(listener func(session Session))
	OnProfileChanged(listener func(session Session))
	OnStateChanged(listener func(session Session))

	ImplicitHosting() bool

	Authenticate(r *http.Request) (Session, error)
}
