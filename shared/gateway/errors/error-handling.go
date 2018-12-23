package errors

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func BadRequest(resp *restful.Response) {
	err := resp.WriteHeaderAndEntity(http.StatusBadRequest, &Error{
		Message: "Invalid request",
	})

	if err != nil {
		log.Printf("Error while writing to response, err: %s\n", err)
	}
}

func Internal(resp *restful.Response, err error) {
	log.Printf("Internal err: %s\n", err)

	err = resp.WriteHeaderAndEntity(http.StatusInternalServerError, &Error{
		Message: "Internal error",
	})

	if err != nil {
		log.Printf("Error while writing to response, err: %s", err)
	}
}
