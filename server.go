package speelcraft_mgr

import "net/http"

// Server represents an instance of a web server that allows clients to download the mods directly from the Minecraft
// server's 'mods' folder.
type Server struct {
	// server is the internal http.Server instance.
	server http.Server

	// Path is the absolute system path to the folder containing the server executable (`.jar`) and `mods` folder.
	Path string `json:"path"`

	// UseHTTPS is a boolean value that determines whether the web server will use TLS encryption.
	UseHTTPS bool `json:"useHTTPS"`

	// KeyFile is the absolute system path to the key file to be used for TLS encryption (HTTPS).
	KeyFile string `json:"keyfile"`

	// CertFile is the absolute system path to the certificate to be used for TLS encryption (HTTPS).
	CertFile string `json:"certfile"`
}

// Start starts the web server.
// The function will return an error if:
//   - Server.Path is not a valid path, is not a folder, or does not contain a `mods` subdirectory.
//   - Server.UseHTTPS is `true` but Server.KeyFile and/or Server.CertFile are not
//     properly configured.
func (s *Server) Start() error {
	return nil
}
