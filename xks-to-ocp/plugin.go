package xkstoocp

import "github.com/konveyor/crane-lib/transform"

const PluginVersion = "v0.0.0"

// Plugin is the Crane transform plugin for migrations from XKS to OpenShift.
type Plugin struct{}

func (p *Plugin) Metadata() transform.PluginMetadata {
	return transform.PluginMetadata{
		Name:            "xKStoOCP",
		Version:         PluginVersion,
		RequestVersion:  []transform.Version{transform.V1},
		ResponseVersion: []transform.Version{transform.V1},
	}
}

// Run currently passes every resource through without changes.
func (p *Plugin) Run(_ transform.PluginRequest) (transform.PluginResponse, error) {
	return transform.PluginResponse{}, nil
}
