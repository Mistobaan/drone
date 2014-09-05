package checkrun

import (
	"net/http"

	"github.com/Mistobaan/drone/pkg/build/docker"
	"github.com/Mistobaan/drone/pkg/handler"
)

type CheckRun struct {
	Status string
}

type CheckRunDashboard struct {
	Images []*docker.Images
}

func Dashboard(w http.ResponseWriter, req *http.Request) error {

	client := docker.New()

	images, err := client.Images.List()
	if err != nil {
		return err
	}

	return handler.RenderTemplate(w, "check_run.html", &CheckRunDashboard{images})
}
