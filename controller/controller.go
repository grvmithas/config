package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personal/config/domain"

	"github.com/julienschmidt/httprouter"
)

// Controller type
type Controller struct {
	Config *domain.Config
}

//ReadRequest for readRequest route
func (c *Controller) ReadRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	serviceName := params.ByName("serviceName")

	config, err := c.Config.Get(serviceName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occured")
	}

	res, err := json.Marshal(&config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occured")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(res))

}
