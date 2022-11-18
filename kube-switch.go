package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kube-switch",
		Short: "Interactive Kubernetes context switcher",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the current context
			currentContext, err := getCurrentContext()
			if err != nil {
				fmt.Printf("Error fetching current context: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Current context: %s\n", currentContext)

			// Get the list of contexts
			contexts, err := getContexts()
			if err != nil {
				fmt.Printf("Error fetching contexts: %v\n", err)
				os.Exit(1)
			}

			// Prompt the user to select a context
			selectedContext, err := promptForContext(contexts, currentContext)
			if err != nil {
				fmt.Printf("Error selecting context: %v\n", err)
				os.Exit(1)
			}

			// Switch to the selected context
			err = switchContext(selectedContext)
			if err != nil {
				fmt.Printf("Error switching context: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Switched to context: %s\n", selectedContext)
		},
	}

	rootCmd.Execute()
}

// Fetch the current Kubernetes context
func getCurrentContext() (string, error) {
	cmd := exec.Command("kubectl", "config", "current-context")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to fetch current context: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// Fetch all Kubernetes contexts
func getContexts() ([]string, error) {
	cmd := exec.Command("kubectl", "config", "get-contexts", "-o", "name")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch contexts: %v", err)
	}

	contexts := strings.Split(strings.TrimSpace(string(output)), "\n")
	return contexts, nil
}

// Prompt the user to select a Kubernetes context
func promptForContext(contexts []string, currentContext string) (string, error) {
	// Mark the current context in the list
	contextItems := make([]string, len(contexts))
	for i, context := range contexts {
		if context == currentContext {
			contextItems[i] = fmt.Sprintf("%s (current)", context)
		} else {
			contextItems[i] = context
		}
	}

	prompt := promptui.Select{
		Label: "Select Kubernetes Context",
		Items: contextItems,
		Searcher: func(input string, index int) bool {
			context := contextItems[index]
			return strings.Contains(context, input)
		},
		StartInSearchMode: true,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Strip "(current)" from the selected result, if present
	result = strings.Replace(result, " (current)", "", 1)

	return result, nil
}

// Switch to the specified Kubernetes context
func switchContext(context string) error {
	cmd := exec.Command("kubectl", "config", "use-context", context)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to switch context: %v\nOutput: %s", err, string(output))
	}

	return nil
}
