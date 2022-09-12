package fabric

import "fmt"

// Create, Update, Read, Delete, CURD.

func CreateOrganization(configPath string, domain string) []string {
	var cmds []string
	// Check cryptogen.
	// Remove peer/ordererOrganization Folder if exists.
	cmds = append(cmds, fmt.Sprintf("cryptogen generate --config=%s --output=\"organizations\"", configPath))
	cmds = append(cmds, fmt.Sprintf("mkdir -p organizations/peerOrganizations/%s/", domain))
	cmds = append(cmds, fmt.Sprintf(""))
	return cmds
}
