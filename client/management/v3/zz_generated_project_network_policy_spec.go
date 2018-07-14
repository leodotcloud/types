package client

const (
	ProjectNetworkPolicySpecType             = "projectNetworkPolicySpec"
	ProjectNetworkPolicySpecFieldDescription = "description"
	ProjectNetworkPolicySpecFieldEnabled     = "enabled"
	ProjectNetworkPolicySpecFieldProjectId   = "projectId"
)

type ProjectNetworkPolicySpec struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ProjectId   string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
