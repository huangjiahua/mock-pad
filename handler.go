package main

import "net/http"

func createSession(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(4096)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	ses := Session{}
	ses.owner = req.FormValue("owner")
	ses.problem = req.FormValue("problem")
	ses.id = RandStringRunes(16)
	for db.Has(ses.id) {
		ses.id = RandStringRunes(16)
	}

}
