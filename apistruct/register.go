package apistruct

import "github.com/rs/xid"

type RegisterData struct{
	ID  	  string `json:"id,omitempty"`
	Name  	  string `json:"name,omitempty"`
	Token 	  xid.ID `json:"token,omitempty"`
	Tags 	  []string `json:"tags,omitempty"`
	Running		bool `json:"tags,omitempty"`
	EncryptKey xid.ID `json:"encryptKey,omitempty"`
}

type RegDataReturn struct {
	Token string `json:"token,omitempty"`
	EncryptKey xid.ID `json:"encryptKey,omitempty"`
}
