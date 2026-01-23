package commands

import (
	"encoding/json"
	"os/exec"
	"fmt"
)

func FetchChaosRoles() []string {
	path, err := exec.LookPath("chaos")
	if err != nil || path == "" {path = "/usr/bin/chaos"}

	cmd := exec.Command("sh", "/usr/bin/chaos", "check", "roles", "-j")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []string{fmt.Sprintf("Error: %v", err), fmt.Sprintf("Output: %s", string(output))}
	}

	var roles []string
	if err := json.Unmarshal(output, &roles); err != nil {
		return []string{"Error parsing roles"}
	}

	return roles
}
