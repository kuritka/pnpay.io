package main

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"pnpay.io/internal/cli"
	"sort"
)

// The JS code in the browser detects changes when the SHA changes, and if it's different, the browser
// automatically refreshes. The SHA code is calculated from individual files in the folder where we deploy the output.
var initialSHA string

func startWebServer(config *cli.Config, err error) {
	logger.Info().Msg("Starting local server")
	// following code generates sha for teh index.html file, so web browser automatically reloads the page when the file changes
	initialSHA = computeSHA(config.LocalServerPath)
	logger.Info().Msgf("Initial SHA: %s", initialSHA)
	http.HandleFunc("/sha", shaHandler) // Endpoint serving the SHA

	server := http.FileServer(http.Dir(config.LocalServerPath))
	http.Handle("/", server)

	logger.Info().Msgf("Server runs on http://localhost:%s/\n", config.LocaLServerPort)
	err = http.ListenAndServe(":"+config.LocaLServerPort, nil)
	if err != nil {
		logger.Fatal().Msg("Server failed to start")
	}
}

// Compute SHA256 hash of the HTML file
func computeSHA(dirPath string) string {
	var hashes []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Ignore directories
		if d.IsDir() {
			return nil
		}
		// Read file content
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// Compute hash for file
		hash := sha256.Sum256(data)
		hashes = append(hashes, fmt.Sprintf("%x", hash))
		return nil
	})

	if err != nil {
		logger.Info().Msgf("Error computing SHA: %s", err)
		return ""
	}

	// Sort hashes to ensure consistent order
	sort.Strings(hashes)

	// Compute final hash from concatenated file hashes
	finalHash := sha256.Sum256([]byte(fmt.Sprintf("%v", hashes)))
	return fmt.Sprintf("%x", finalHash)
}

// Serve SHA hash (remains static during server lifetime)
func shaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, initialSHA)
}
