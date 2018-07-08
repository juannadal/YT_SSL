package main

import (
	"net/http"
	"html/template"
	"log"
	"path/filepath"
)



//This file will not work on localhost must run on a server such as digitalocean.com

func handler(w http.ResponseWriter, r *http.Request) {
	//find current directory to set Absolute path
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Fatal(err)
	}
	t, _ :=template.ParseFiles(dir+"/index.html")
	t.Execute(w,nil)
}
//redirect all HTTP traffic to HTTPS server on port :443
func redirectTLS(w http.ResponseWriter, r *http.Request)  {

	http.Redirect(w,r,"https://www.useYouDomain.com:443"+r.RequestURI,http.StatusMovedPermanently)
}


func main()  {

	http.HandleFunc("/", handler)
	//For Namecheap.com you will handler pointing to this path to validate your certificate
	//.well-known/pki-validation/

	//create goroutine to handle traffic going to port :443

	go http.ListenAndServeTLS(":443","/root/slhacker_club.crt","/root/slhackerclub.key",nil)

	//Port :80 is set of all default traffic
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)))

}
