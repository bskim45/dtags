package common

type Endpoint struct {
	// The base_client URL for requests
	URL string
}

var DockerHub = Endpoint{
	URL: "https://registry.hub.docker.com/",
}
