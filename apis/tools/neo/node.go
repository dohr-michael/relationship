package neo

import "github.com/dohr-michael/relationship/apis/tools/structure"

type Node struct {
	NodeIdentity int64                  `json:"NodeIdentity"`
	Labels       []string               `json:"Labels"`
	Properties   map[string]interface{} `json:"Properties"`
}

func HasNode(props map[string]interface{}) (Node, error) {
	res := Node{}
	err := structure.Decode(&res, props)
	return res, err
}
