package speelcraft_mgr

// Config represents an imported JSON file with configuration information for the program.
type Config struct {
	// Path is the absolute system path to the folder containing the server executable (.jar) and 'mods' folder.
	Path string `json:"path"`

	// UseHTTPS is a boolean value that determines whether the web server will use TLS encryption.
	UseHTTPS bool `json:"useHTTPS"`

	// KeyFile is the absolute system path to the key file to be used for TLS encryption (HTTPS).
	KeyFile string `json:"keyfile"`

	// CertFile is the absolute system path to the certificate to be used for TLS encryption (HTTPS).
	CertFile string `json:"certfile"`
}
