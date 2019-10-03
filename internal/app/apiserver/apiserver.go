package apiserver

// APIServer ...
type APIServer struct {
	config *Config
}

// New ...
func New() *APIServer {
	// init server
	return &APIServer{
		config: config,
	}
}

// Start Server
// DB init etc..
func (s *APIServer) Start() error {
	return nil
}
