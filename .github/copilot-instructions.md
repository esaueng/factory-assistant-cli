# Working agreements

Complete the requested work. Infer scope from context, make reasonable
implementation choices, and continue until finished or genuinely blocked.
Keep investigation-only requests read-only.

An explicit request or prior approval authorizes that action. Do not ask
again. Ask only when consequential ambiguity remains or an additional action
falls outside the authorized scope. Complete independent work while waiting.

Follow repository conventions. Keep changes focused. Preserve units,
tolerances, defaults, formats, and public contracts unless changes are
required. Label approximations and flag breaking changes or new dependencies.

Run required checks and tests proportional to the change. Add regression
coverage for bug fixes. Do not weaken assertions or bypass protections.
Attempt routine recovery from missing tools or Git refs before reporting a
blocker. Distinguish local validation from production verification.

Keep secrets, private instructions, and identifying information out of shared
artifacts. Inspect staged content and metadata before committing. Use the
approved repository Git identity.

## Delivery

- Use a branch and a ready-for-review pull request for repository changes.
- A merge request authorizes the identified pull request’s merge and its
  existing automatic deployment. Mention that consequence and proceed
  without another confirmation.
- Before merging, verify the current head, mergeability, reviews, and required
  checks. Wait for pending checks. Existing approval remains valid for the
  authorized change.
- Bypassing CI requires explicit authorization for the specific pull request.
  Disclose affected checks and risks; preserve review requirements.
- Standalone production deployments and manual migrations require
  authorization. An explicit request to perform them is sufficient.
- Resolve the target repository, environment, and account from configuration
  and context. Ask only if the target remains ambiguous.

Report the outcome briefly, followed by verification, actual blockers, and
relevant links. Never claim an unverified result passed.

---

# Home Assistant CLI

## Project Overview
This is the official Home Assistant CLI tool written in Go, providing command-line
interface to interact with the Home Assistant Supervisor. The CLI enables users to
manage apps, control the core system, handle audio/network settings, manage
backups, and perform various system operations.

The CLI is communicating with the Supervisor using the Supervisor's HTTP REST API.

## Repository Structure
- **`main.go`** - Entry point of the application
- **`cmd/`** - Contains all CLI command implementations using Cobra framework
- **`client/`** - HTTP client functionality for API communication
- **`spinner/`** - Progress spinner implementation
- **Root files** - Configuration and documentation

## Key Technologies
- **Language**: Go (use modern syntax)
- **CLI Framework**: Cobra (github.com/spf13/cobra)
- **HTTP Client**: Resty (github.com/go-resty/resty/v2)
- **Configuration**: Viper (github.com/spf13/viper)
- **Logging**: Go stdlib log/slog

## Available Commands
The CLI provides the following main command categories:
- `apps` - Install, update, remove and configure Home Assistant apps
- `audio` - Audio plug-in management
- `authentication` - Authentication for Home Assistant users
- `cli` - CLI plug-in management
- `core` - Home Assistant Core control
- `dns` - DNS plug-in management
- `docker` - Docker related configuration
- `hardware` - System hardware information
- `host` - Host OS control
- `info` - General Home Assistant information
- `multicast` - Multicast plug-in configuration
- `network` - Network configuration and management
- `observer` - Observer plug-in management
- `os` - Home Assistant OS specific operations
- `resolution` - Resolution center for issues and solutions
- `backups` - Backup creation, restoration, and management
- `supervisor` - Supervisor monitoring and control

## Development Environment
- **API Endpoint**: Configurable via `SUPERVISOR_ENDPOINT` environment variable
- **Authentication**: Uses API tokens via `SUPERVISOR_API_TOKEN` environment variable
- **Config File**: Optional config file support (default: `$HOME/.homeassistant.yaml`)

## Build Commands
- **Build**: `CGO_ENABLED=0 go build -ldflags="-s -w" -o "ha"`
- **Test**: `go test ./...`
- **Format**: `gofmt -s`

## File Organization Patterns
- Command files follow naming pattern: `<component>_<action>.go`
- Each command typically has its own file in the `cmd/` directory
- Helper functions are in `client/helper.go`
- Main client logic is in `client/client.go`

## Architecture Notes
- Uses Cobra for command structure and flag parsing
- Resty for HTTP API calls to Home Assistant Supervisor
- Viper for configuration management
- Go stdlib log/slog for structured logging
- Custom spinner implementation for progress indication

## Testing
- Unit tests available in `client/helper_test.go`
- Test command: `go test ./...`

## Contributing Guidelines
1. Create feature branch
2. Commit changes
3. Rebase against master
4. Run tests with `go test ./...`
5. Format code with `gofmt -s`
6. Create Pull Request

This CLI is designed to work with Home Assistant Supervisor API and is commonly used in
Home Assistant Operating System environments, SSH apps, and development setups.
