# kube-switch

`kube-switch` is a lightweight, customizable CLI tool for interactively switching Kubernetes contexts. It offers a fast, user-friendly interface to manage Kubernetes contexts directly from your terminal.

## Features
- Displays the current context for clarity.
- Searchable, interactive menu for selecting contexts.
- Cross-platform support as a standalone binary.

## How to Use

1. **Build the Binary**:
   ```bash
   go build -o kube-switch
   ```

2. **Run the Tool**:
   ```bash
   ./kube-switch
   ```

3. **Select Context**:
   - The current context will be displayed.
   - Use the interactive menu to select and switch to another context.

## Compile for Different Architectures

To compile for other platforms:

1. For Linux:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o kube-switch-linux
   ```

2. For macOS:
   ```bash
   GOOS=darwin GOARCH=amd64 go build -o kube-switch-macos
   ```

3. For Windows:
   ```bash
   GOOS=windows GOARCH=amd64 go build -o kube-switch.exe
   ```

## Why Not Use `kubectx`?

While `kubectx` is popular, `kube-switch` provides:
- **Customization**: Modify the tool to fit specific workflows or integrate additional features.
- **Portability**: A single, self-contained Go binary with no dependency on shell scripts.
- **Performance**: Faster context fetching by reading the Kubernetes configuration directly.

`kube-switch` is ideal for teams looking for a simpler, faster, and more customizable alternative to `kubectx`.