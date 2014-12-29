package main

import (
	"html/template"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"subst": "injectedval"}
	htmlTemplate.Execute(w, data)
}

var templateFile string
var htmlTemplate *template.Template

func main() {
	htmlTemplate = template.Must(template.New("fred").Parse(htmlSource))
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":9876", nil)
}

//----------------------------------------------------------------------------

var htmlSource = `

<html lang="en"
   xmlns="http://www.w3.org/1999/xhtml">
   <head>
      <meta name="generator"
         content=
         "HTML Tidy for Linux/x86 (vers 25 March 2009), see www.w3.org" />
      <meta charset="utf-8" />
      <meta http-equiv="X-UA-Compatible"
         content="IE=edge" />
      <meta name="viewport"
         content="width=device-width, initial-scale=1" />
      <title>
         Bootstrap 101 Template
      </title>
      <!-- Bootstrap -->
      <link rel="stylesheet"
         href=
         "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css"
         type="text/css" />
      <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
      <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
      <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
      <![endif]-->
   </head>
   <body>
      <h1>
         Hello, world!
      </h1>
      <!-- Bootstrap core JavaScript
         ================================================== -->
      <!-- Placed at the end of the document so the pages load faster -->
      <script src=
         "https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"
         type="text/javascript"></script><script src=
         "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"
         type="text/javascript"></script>
   </body>
</html>

`
