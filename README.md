# GVM

#### ➡️ Go Version Manager, <s>only for windows</s>

## 📕 Purpose:

Automate the task of going to the [Go dl page](https://go.dev/dl), to download the Windows MSI executable to then open it and having to click 50 times on the "Next" button ＞﹏＜

Futhermore, this CLI will cache the MSI files to not having to redownload it later... so that you can switch between version very swiftly ⚡

## Made With:

1. **Elegance** ✅
2. `Go`, ironically
3. [`Cobra`](https://github.com/spf13/cobra), for CLI framework 💻

## Usage:

- To use this CLI you must be on windows (10/11 recommended).
- You can download the CLI executable in the Github Releases.

#### `switch` (Combines `dl` and `use`)

➡️ Let you switch of Go Main Version easily: uses the Go MSI executable (recommended for windows). It Downloads the msi executable from Go Official Website if not already downloaded, then it uninstalls the current Go version and finally it installs the newly/already downloaded Go version.

> TIP: use 'latest' arg to switch to the latest go version

```bash
gvm switch latest # This will switch your Go Version to the latest release
```

```bash
gvm switch 1.19 # This will switch your Go Version to go1.19
```

```bash
gvm switch 1.18.5 --no-cache # This will switch your Go Version to go1.18.5, without caching the file for later use (stores it in `/temp` and deletes it when installation finish)
```

...

#### `manager dl`

➡️ Downloads the specified Go MSI Version in the app 'AppData' Dir, for later use.

> TIP: use 'latest' arg to download the latest go version

```bash
gvm manager dl 1.18.5 # This will download the go1.18.5 MSI file on your disk (AppData/Roaming/gvm-windows)
```

...

#### `manager use`

➡️ Switch between multiples version of Go. If the specified Go Version is not downloaded the process exit.

> TIP: use 'latest' arg to use the latest go version

```bash
gvm manager use 1.18.5 # This will only switch your Go Version to go1.18.5 if go1.18.5 is already installed (in AppData/Roaming/gvm-windows)
```

...

#### `manager scan`

➡️ Scan and delete old downloaded Go MSI file based on the file creation date (default = 6 months ago).

```bash
gvm manager scan # Will scan to find old Go MSI file (old = >6 months) and then deletes it
```

```bash
gvm manager scan --date-limit=1638054000000 # will delete all the MSI file that were installed before the 11/28/2021 (date must be in unix timestamp millisecond format)
```
