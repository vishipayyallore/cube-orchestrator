# Build System Documentation

This document describes the build system for the cube-orchestrator project.

## 📁 Directory Structure

```text
scripts/
├── build.sh           # Main build script with timestamps
└── cleanup-builds.sh  # Build cleanup utility

builds/                # Build output directory (in .gitignore)
├── cube-orchestrator_YYYYMMDD_HHMMSS  # Timestamped executables
└── cube-orchestrator_latest           # Symlink to latest build
```

## 🔨 Building the Application

### Option 1: Using the Build Script (Recommended)

```bash
# From project root
./scripts/build.sh
```

This will:

- Create a timestamped executable in `builds/`
- Update the `cube-orchestrator_latest` symlink
- Show build information and directory contents

### Option 2: Manual Build

```bash
# From orchestrator directory
cd src/orchestrator
go build -o ../../builds/cube-orchestrator_$(date +"%Y%m%d_%H%M%S") ./cmd
```

### Option 3: VS Code Task

1. Open Command Palette (`Ctrl+Shift+P`)
2. Type "Tasks: Run Task"
3. Select "Build Cube Orchestrator (Timestamped)"

## 🚀 Running the Application

### Run Latest Build

```bash
# From project root
./builds/cube-orchestrator_latest
```

### Run Specific Build

```bash
# From project root
./builds/cube-orchestrator_20250804_061039
```

## 🧹 Build Management

### Cleanup Old Builds

Keep only the latest 5 builds (default):

```bash
./scripts/cleanup-builds.sh
```

Keep only the latest 3 builds:

```bash
./scripts/cleanup-builds.sh 3
```

Keep only the latest 10 builds:

```bash
./scripts/cleanup-builds.sh 10
```

## 📝 Build Features

### Timestamped Builds

- Format: `cube-orchestrator_YYYYMMDD_HHMMSS`
- Example: `cube-orchestrator_20250804_061039`
- Allows keeping multiple versions for testing/rollback

### Latest Symlink

- `cube-orchestrator_latest` always points to the newest build
- Makes it easy to run the latest version without remembering timestamps
- Updated automatically by build script

### Automated Cleanup

- Prevents `builds/` directory from growing too large
- Configurable retention policy
- Preserves latest N builds (default: 5)

## 🎯 Integration with VS Code

The build system integrates with VS Code through tasks:

- **Build Task**: Runs the timestamped build script
- **Problem Matcher**: Integrates with Go compiler errors
- **Keyboard Shortcuts**: Can be assigned for quick building

## 🔧 Customization

### Build Script Modifications

Edit `scripts/build.sh` to:

- Change binary name
- Add build flags (e.g., `-ldflags`, `-tags`)
- Add cross-compilation targets
- Include version information

### Cleanup Script Modifications

Edit `scripts/cleanup-builds.sh` to:

- Change default retention count
- Add build size reporting
- Include build metrics

## 📋 Example Workflow

1. **Development**: Make code changes in `src/`
2. **Build**: Run `./scripts/build.sh` or use VS Code task
3. **Test**: Run `./builds/cube-orchestrator_latest`
4. **Iterate**: Repeat as needed
5. **Cleanup**: Periodically run `./scripts/cleanup-builds.sh`

## 🧪 Docs Quality CI

Documentation quality is enforced in CI:

- Workflow: `.github/workflows/docs-quality.yml`
- Lint config: `.markdownlint.json` at repo root
- Link checker config: `lychee.toml` at repo root

The workflow runs on PRs and pushes (main) that modify `docs/**` or `README.md`.

## 🏗️ CI: Build Orchestrator

A minimal CI workflow builds the Go orchestrator and publishes a binary artifact.

- Workflow: `.github/workflows/build-orchestrator.yml`
- Triggers:
  - Push to `main` affecting `src/orchestrator/**`
  - Pull requests touching `src/orchestrator/**`
- Permissions: least privilege (`contents: read`)
- Runner: `ubuntu-latest`

Steps (high level):

1. Checkout repository (`actions/checkout@v4`).

2. Setup Go (`actions/setup-go@v5`) with go-version `1.24.6` (from `src/orchestrator/go.mod` toolchain), cache enabled, and `cache-dependency-path` set to `src/orchestrator/go.sum`.

3. Download modules: `go mod download` (working dir: `src/orchestrator`).

4. Build binary: `go build -v -o ../builds/cube-orchestrator_ci main.go` (working dir: `src/orchestrator/cmd`).

5. Upload artifact (`actions/upload-artifact@v4`) as `cube-orchestrator_ci-linux-amd64` from `src/orchestrator/builds/cube-orchestrator_ci`.

Notes:

- The runner builds a Linux amd64 binary by default (suitable for downloading from the PR/Run page).
- The workflow uses least-privilege token permissions and Go module caching for faster builds.
- Future: add a separate frontend job (React) or a matrix if `src/frontend/` is introduced.

## 🛠️ Troubleshooting

### Build Fails

- Check Go installation: `go version`
- Verify dependencies: `go mod tidy`
- Check for syntax errors: `go build ./src`

### Permission Issues

- Make scripts executable: `chmod +x scripts/*.sh`
- Check file permissions in `builds/` directory

### Symlink Issues

- Remove broken symlink: `rm builds/cube-orchestrator_latest`
- Rebuild to recreate symlink: `./scripts/build.sh`

---

**Note**: The `builds/` directory is included in `.gitignore` to prevent committing binary files to the repository.
