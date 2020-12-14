package openapi

var (
	UserDB = make([]LoginData, 0)
	Subscribers = make([]SubsData, 0)
	UEs = make([]UE, 0)
	RANs = make([]RanInfo, 0)
)

type UE struct {
	PlmnId int32 `json:"plmnId,omitempty"`

	UeId string `json:"ueId,omitempty"`

	Online bool `json:"online,omitempty"`
}