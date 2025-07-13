package server

import "time"

// JWTInfo holds information specific to JWT credentials
type JWTInfo struct {
	Issue   string `json:"issue,omitempty"`
	Expires string `json:"expires,omitempty"`
	Tags    string `json:"tags,omitempty"`
}

// NATSInfo contains server connection information
type NATSInfo struct {
	Server     string `json:"server,omitempty"`
	ServerName string `json:"server_name,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Protocol   int    `json:"protocol,omitempty"`
}

// UserClaims holds information from the user JWT claims
type UserClaims struct {
	Sub         string                 `json:"sub,omitempty"`
	Subject     string                 `json:"sub,omitempty"` // Alternative name for Sub
	Issuer      string                 `json:"iss,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Audience    string                 `json:"aud,omitempty"`
	Expires     time.Time              `json:"exp,omitempty"`
	ExpiresAt   int64                  `json:"exp,omitempty"` // Alternative format for Expires
	IssuedAt    time.Time              `json:"iat,omitempty"`
	IssuedAtTS  int64                  `json:"iat,omitempty"` // Alternative format for IssuedAt
	Type        string                 `json:"type,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Permissions string                 `json:"permissions,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

// WSConnInfo contains information about WebSocket connections
type WSConnInfo struct {
	Protocol    string `json:"protocol,omitempty"`
	Compression bool   `json:"compression,omitempty"`
	URL         string `json:"url,omitempty"`
}

// MQTTConnMsg contains information about MQTT connections
type MQTTConnMsg struct {
	ClientID     string `json:"client_id,omitempty"`
	KeepAlive    int    `json:"keep_alive,omitempty"`
	CleanSession bool   `json:"clean_session,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
}

// ExtendedUserClaims holds information about user specific claims in a JWT with timestamp format.
type ExtendedUserClaims struct {
	Subject  string                 `json:"sub,omitempty"`
	Issuer   string                 `json:"iss,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Expires  int64                  `json:"exp,omitempty"`
	IssuedAt int64                  `json:"iat,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

// ConnectionInfo has detailed information on a per connection basis.
type ConnectionInfo struct {
	Cid            uint64              `json:"cid"`
	Kind           string              `json:"kind,omitempty"`
	Type           string              `json:"type,omitempty"`
	IP             string              `json:"ip"`
	Port           int                 `json:"port"`
	Start          time.Time           `json:"start"`
	LastActivity   time.Time           `json:"last_activity"`
	Stop           *time.Time          `json:"stop,omitempty"`
	Reason         string              `json:"reason,omitempty"`
	RTT            string              `json:"rtt,omitempty"`
	Uptime         string              `json:"uptime"`
	Idle           string              `json:"idle"`
	PendingBytes   int64               `json:"pending_bytes"`
	InMsgs         int64               `json:"in_msgs"`
	OutMsgs        int64               `json:"out_msgs"`
	InBytes        int64               `json:"in_bytes"`
	OutBytes       int64               `json:"out_bytes"`
	InBytesTotal   int64               `json:"in_bytes_total"`
	OutBytesTotal  int64               `json:"out_bytes_total"`
	NumSubs        uint32              `json:"subscriptions"`
	Name           string              `json:"name,omitempty"`
	Lang           string              `json:"lang,omitempty"`
	Version        string              `json:"version,omitempty"`
	TLSVersion     string              `json:"tls_version,omitempty"`
	TLSCipher      string              `json:"tls_cipher,omitempty"`
	AuthorizedUser string              `json:"authorized_user,omitempty"`
	Account        string              `json:"account,omitempty"`
	Subs           []string            `json:"subscriptions_list,omitempty"`
	SubsDetail     []SubDetail         `json:"subscriptions_list_detail,omitempty"`
	JetStream      *JSInfo             `json:"jetstream,omitempty"`
	JWT            *JWTInfo            `json:"jwt,omitempty"`
	NATS           *NATSInfo           `json:"nats,omitempty"`
	UserClaims     *UserClaims         `json:"user_claims,omitempty"`
	IssuerClaims   *ExtendedUserClaims `json:"issuer_claims,omitempty"`
	WSConnection   *WSConnInfo         `json:"ws_connection,omitempty"`
	MQTTConnection *MQTTConnMsg        `json:"mqtt_connection,omitempty"`
	LeafNode       *LeafInfo           `json:"leaf_node,omitempty"`
}
