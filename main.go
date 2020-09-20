package main

import (
	"html/template"
	"log"
	"os"
)

var (
	urlPrefix      = ""
	templatePrefix = ""
	localDBPrefix  = ""
	fbAPIKey       = ""
	fbAuthDomain   = ""
	fbDBUrl        = ""

	tpl *template.Template = nil
)

func init() {
	// get url prefix
	urlPrefix = os.Getenv("URL_PREFIX")
	if urlPrefix == "" {
		urlPrefix = "" // default to root URL
	}

	// get template files
	templatePrefix = os.Getenv("TEMPLATE_PREFIX")
	if templatePrefix == "" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		templatePrefix = path + string(os.PathSeparator) + "template" + string(os.PathSeparator)
		tpl = template.Must(template.ParseGlob(templatePrefix + "*")) // get all templates
	}

	// get local db path
	// this app use local files to store data
	localDBPrefix = os.Getenv("LOCAL_DB_PREFIX")
	if localDBPrefix == "" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		localDBPrefix = path
	}

	// get FireBase configuration
	fbAPIKey = os.Getenv("FIREBASE_API_KEY")
	fbAuthDomain = os.Getenv("FB_AUTH_DOMAIN")
	fbDBUrl = os.Getenv("FB_DB_URL")

	if fbAuthDomain == "" || fbDBUrl == "" || fbAPIKey == "" {
		log.Fatal("please include FireBase configuration in the environment variables")
	}
}

func url(s string) string {
	return urlPrefix + s
}

func main() {

}
