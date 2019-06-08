package apistruct


type RegisterData struct{
	ID  	  string `json:"id,omitempty"`
	Name  	  string `json:"name,omitempty"`
	Token 	  string `json:"token,omitempty"`
	Tags 	  []string `json:"tags,omitempty"`
	Running		bool `json:"tags,omitempty"`
	EncryptKey string `json:"encryptKey,omitempty"`
}

type RegDataReturn struct {
	Token string `json:"token,omitempty"`
	EncryptKey string `json:"encryptKey,omitempty"`
}
