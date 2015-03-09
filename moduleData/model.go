package data

const FriendKind = "friend"

type Friend struct {
  Name string `json:"name" datastore:",noindex"`
  Phone string `json:"phone" datastore:",noindex"`
}

func (p *Friend) GetKey() string {
  return p.Name
}
func (p *Friend) GetKind() string {
  return FriendKind
}
