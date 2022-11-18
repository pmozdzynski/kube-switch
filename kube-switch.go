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
			contexts, err := getContexts()
			if err != nil {
				fmt.Printf("Error fetching contexts: %v\n", err)
				os.Exit(1)
			}

			selectedContext, err := promptForContext(contexts)
			if err != nil {
				fmt.Printf("Error selecting context: %v\n", err)
				os.Exit(1)
			}

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

func getContexts() ([]string, error) {
	cmd := exec.Command("kubectl", "config", "get-contexts", "-o", "name")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch contexts: %v", err)
	}

	contexts := strings.Split(strings.TrimSpace(string(output)), "\n")
	return contexts, nil
}

func promptForContext(contexts []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select Kubernetes Context",
		Items: contexts,
		Searcher: func(input string, index int) bool {
			context := contexts[index]
			return strings.Contains(context, input)
		},
		StartInSearchMode: true,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func switchContext(context string) error {
	cmd := exec.Command("kubectl", "config", "use-context", context)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to switch context: %v\nOutput: %s", err, string(output))
	}

	return nil
}


