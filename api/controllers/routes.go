package controllers

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", s.Home).Methods("GET")
	s.Router.HandleFunc("/tasks", s.CreateTask).Methods("POST")
	s.Router.HandleFunc("/tasks/{id}", s.GetTask).Methods("GET")
}
