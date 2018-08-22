package search

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type searchController struct {
	db		*sqlx.DB
	log		*logrus.Logger
	searchConfig	*config.searchConfig
}

type Search struct {
	Search		string
	Use_regex	bool
}

func (u userController) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonSearch Search
	response, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(response, &jsonSearch)
	if err != nil {
		u.log.WithError(err).Warn("Error while unmarshalling json request.")
		http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
		return
	}
	// TODO Combine the resume token search with GraphQL search
}

func (u userController) registrationNameSearch() ([]RegistrationParticipant) {
	// TODO
}

func Load(r *mux.Router, db *sqlx.DB, log *logrus.Logger, searchConfig *config.SearchConfig, authConfig *config.AuthenticationConfig) {
	s := &searchController{db, log, searchConfig}
	createSearchSubR := r.PathPrefix("").Methods("POST").Subrouter()
	createSearchSubR.PathPrefix("").HandlerFunc(u.Search)
	createSearchSubR.Use(auth.RequireAuthentication(log, authConfig))
}
