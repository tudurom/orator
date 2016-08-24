package gen

import (
	"bytes"
	"html/template"
	"io"

	yaml "gopkg.in/yaml.v2"

	"github.com/russross/blackfriday"
	"github.com/tudurom/orator/config"
	"github.com/tudurom/orator/util"
)

// Representation of a page
type Page struct {
	FrontMatter map[string]interface{}
	Content     template.HTML
}

// Data supplied to the template
type PageData struct {
	Page       *Page
	SiteConfig *config.SiteConfig
	Layout     Layout
}

// Reads the page's content and returns the generated output
func (p *Page) GeneratePage(input io.Reader, fileExt string, conf *config.SiteConfig,
	rootTpl *template.Template, fm *util.FrontMatter,
	layouts map[string]Layout) (generatedPage string, err error) {

	front, body, err := fm.Parse(input)

	if err != nil {
		return "", err
	}

	p.FrontMatter = make(map[string]interface{})
	yaml.Unmarshal([]byte(front), p.FrontMatter)

	// Handle different formats
	p.Content = template.HTML(renderBody(body, fileExt))

	generatedPage = ""
	// If the layout is missing the return it as is
	if p.FrontMatter["layout"] == "" {
		generatedPage = body
	} else {
		tplName := p.FrontMatter["layout"].(string)
		buf := new(bytes.Buffer)
		data := PageData{p, conf, layouts[tplName]}
		err := rootTpl.ExecuteTemplate(buf, tplName, data)
		if err != nil {
			return "", err
		}
		generatedPage = buf.String()
	}
	return generatedPage, nil
}

// Handle different file extensions
func renderBody(body, fileExt string) (output string) {
	if fileExt == "md" {
		output = string(blackfriday.MarkdownCommon([]byte(body)))
	} else {
		output = body
	}

	return output
}
