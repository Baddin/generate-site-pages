package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type site struct {
	header header
	body   body
	footer footer
}

type header struct {
	icon   string
	navbar []string
	title  string
	//backgroundImage string // make it struct with all it's proberty
	paragraph string
}

type body struct {
	content string //to be deleted
	divs    []string
}

type footer struct {
	links      []string
	copyrights string //add more later
}

func (s *site) GenerateHeader() string {

	//header data
	headerico := s.header.icon
	headernavbar := s.header.navbar
	headerTitle := s.header.title
	headerP := s.header.paragraph
	//headerBI := s.header.backgroundImage

	var navbarS string
	for _, element := range headernavbar {
		navbarS += element + " "
	}
	//header structure, when u fill the navbar with data put the <li> tag with the data
	header := `
		<div class="header">
			<div class = "ico">
				<img src="%s"></img>
			</div>
			<div class= "heading">
				<h1>%s</h1>
				<p><em>%s</em></p>
			</div>
			<div class="navbar">
				<ul>
					%s
				</ul>
			</div>
		</div>
	`
	//putting the data on the structure, respect the order for now TODO: find better way
	header = fmt.Sprintf(header, headerico, headerTitle, headerP, navbarS)
	return header

}

func (s *site) GenerateBody() string {
	//body data
	bodyDivs := s.body.divs       // the body sections tags(divs) has to be set manualy for now
	bodyContent := s.body.content // the other content beside the divs and I know it sounds silly

	var divsS string

	for _, element := range bodyDivs {
		divsS += element + " "
	}

	//body structure
	body := `
			%s
		<div class="content">
			%s
		</div>
	`
	//putting the data in the structure
	body = fmt.Sprintf(body, divsS, bodyContent)
	return body
}

func (s *site) GenerateFooter() string {
	//footer data
	footerLinks := s.footer.links
	footerCopyrights := s.footer.copyrights

	var linksS string

	for _, element := range footerLinks {
		linksS += element + " "
	}

	//footer structure very simple for now and add <li> to the links when filling it with data
	footer := `
		<div class="footer">
			<ul>
				%s
			</ul>
			%s
		</div>
	`

	//putting the data int the structure remmber to respect the order for now
	footer = fmt.Sprintf(footer, linksS, footerCopyrights)
	return footer
}

func (s *site) Generate(title string) string {

	header := s.GenerateHeader()
	body := s.GenerateBody()
	footer := s.GenerateFooter()

	//site structure
	generatedSite := `
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>%s</title>
	</head>
	<body>
		%s
		%s
		%s
	</body>
	`

	//putting the headers and the body and footer in the data structure
	generatedSite = fmt.Sprintf(generatedSite, title, header, body, footer)
	return generatedSite

}

func newSite(icon string, navbar []string, title string /*backgroundImage string,*/, paragraph string, content string, divs []string, links []string, copyrights string) *site {
	return &site{
		header{
			icon:   icon,
			navbar: navbar,
			title:  title,
			//backgroundImage: backgroundImage,
			paragraph: paragraph,
		},
		body{
			content: content,
			divs:    divs,
		},
		footer{
			links:      links,
			copyrights: copyrights,
		},
	}
}

func (s *site) WriteGeneratedSiteToFile(filename string) error {
	data := s.Generate("Test")

	//f, err := os.Create(filename)

	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//var aSite *site
	var siteName string
	navbar := []string{"<li>Index</li>", "<li>Brows</li>", "<li>About</li>"}
	links := []string{"<li><a href=\"#\">Test</a></li>", "<li><a href=\"#\">Test2</a></li>"}

	c := make(chan *site)
	go func(c chan *site) {
		for i := 0; i < 10; i++ {
			aSite := newSite("./ico.png", navbar, "Test site", "it is a generated site", "<p>test<p>", nil, links, "©copyrights smthng 2018")
			c <- aSite
		}
	}(c)
	go func(c chan *site) {
		for i := 0; i < 10; i++ {
			siteChan := <-c
			siteName = fmt.Sprintf("GeneratedSite%d.html", i)
			siteChan.WriteGeneratedSiteToFile(siteName)
			fmt.Println("site created: ", siteName)
		}
	}(c)

	var input string
	fmt.Scanln(&input)
}
