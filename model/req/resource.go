package req

type SaveOrUpdateResource struct {
	ID            int    `json:"id"`
	Url           string `json:"url" mapstructure:"url"`
	RequestMethod string `json:"request_method" mapstructure:"request_method"`
	Name          string `json:"name" mapstructure:"name"`
	ParentId      int    `json:"parent_id" mapstructure:"parent_id,omitempty"`
}
