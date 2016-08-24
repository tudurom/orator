# orator

## The simple static site generator

Orator is a static site generator written in Go. It is very fast, easy to use
and flexible. Orator takes a directory with content and renders it to HTML with
Go templates called "layouts".

Layouts and content files can include a yaml front matter that can be used in
the content. Additionally, Orator reads site-specific variables from a
`config.yaml` file. Variables defined there can be accessed in all layouts and
content files.

Orator runs on any platform where the Go tool chain can run like Plan 9, Linux, Windows,
Mac OS X and {DragonFly,Free,Open,Net}BSD.

### Installation

First, make sure that you have [Go](https://golang.org) installed.

Install Orator with the following command:

```bash
go get github.com/tudurom/orator
```

### Using Orator

The first thing you need to do is setting up the directory structure.
An Orator website needs a `config.yaml` file that stores site-wide configuration
and two directories: `layouts` for templates and `content` for your content.

#### Creating a layout

**Read more about Go templates [here](https://golang.org/pkg/text/template/)**.

A simple layout may look like this:

`default.html`

```html
{{ define "head" }}
<head>
	<meta charset="utf-8">
	{{ $title := index .Page.FrontMatter "title" }}
	<title>{{ .SiteConfig.Title }}{{ if $title }} - {{ $title }}{{end}}</title>
</head>
{{ end }}

{{ define "header" }}
<header>
	<h1><a href="/">{{ .SiteConfig.Title }}</a></h1>
</header>
{{ end }}

{{ define "default" }}
<!DOCTYPE html>
{{ template "head" . }}
<body>
	{{ template "header" . }}
	{{ .Page.Content }}
</body>
{{ end }}
```

As you can see each file can contain multiple templates. They are all loaded
anyway.

#### Creating content

Next up, we write some content for our site in the `content` directory. The
directory layout here is preserved in the generated site.

Content can be in any format. If the file's name ends in `.md`, it will be
rendered as markdown to html.

Content files can have a yaml front matter:

```markdown

---
layout: default
special_thing: false
---

Hey this a site!

{{ if index .Page.FrontMatter "special_thing" }}
	<h1>Hidden header</h1>
{{ end }}

* [users](/users)

```

In the example above, the header will not be shown because the
`special_thing` variable is set to false.

`layout` is a special variable that tells Orator what layout should this page
use.

#### Generating the site

`cd` into the site folder and run `./orator`. The final site will be generated
in the `gen` directory.
