package templates

import (
	"fmt"
	"io"
)

type HelloData struct {
	Name    string
	Email   string
	Lang    string
	Entries []*Entry
}

type Entry struct {
	Type  string
	Title string
}

func Hello(w io.Writer, d *HelloData) (err error) {
	_, err = io.WriteString(w, "<h1>")
	if err != nil {
		return
	}

	_, err = io.WriteString(w, fmt.Sprintf("Hi, %s", d.Name))
	if err != nil {
		return
	}

	_, err = io.WriteString(w, " ")
	if err != nil {
		return
	}

	err = d.MyEmail(w)
	if err != nil {
		return
	}

	_, err = io.WriteString(w, "</h1>\n")
	if err != nil {
		return
	}

	_, err = io.WriteString(w, "<ul>\n")
	if err != nil {
		return
	}

	for _, e := range d.Entries {
		_, err = io.WriteString(w, "<li>")
		if err != nil {
			return
		}
		err = EntryItem(w, e, d.Lang)
		if err != nil {
			return
		}
		_, err = io.WriteString(w, "</li>\n")
		if err != nil {
			return
		}
	}

	_, err = io.WriteString(w, "</ul>\n")
	if err != nil {
		return
	}

	return
}

func EntryItem(w io.Writer, e *Entry, lang string) (err error) {
	switch e.Type {
	case "post":
		_, err = io.WriteString(w, "<div class=\"post\">")
		if err != nil {
			return
		}
		_, err = io.WriteString(w, e.Title)
		if err != nil {
			return
		}
		_, err = io.WriteString(w, "</div>\n")
		if err != nil {
			return
		}
	case "comment":
		_, err = io.WriteString(w, "<div class=\"comment\">")
		if err != nil {
			return
		}
		_, err = io.WriteString(w, e.Title)
		if err != nil {
			return
		}
		_, err = io.WriteString(w, "</div>\n")
		if err != nil {
			return
		}
	default:
		_, err = io.WriteString(w, "<div class=\"entry\">")
		if err != nil {
			return
		}
		_, err = io.WriteString(w, e.Title)
		if err != nil {
			return
		}
		_, err = io.WriteString(w, "</div>\n")
		if err != nil {
			return
		}
	}

	if lang == "english" {
		_, err = io.WriteString(w, "Written in English\n")
		if err != nil {
			return
		}
	} else {
		_, err = io.WriteString(w, "Unknown Language\n")
		if err != nil {
			return
		}
	}
	return
}

func (h *HelloData) MyEmail(w io.Writer) (err error) {
	_, err = io.WriteString(w, "<em>")
	if err != nil {
		return
	}

	_, err = io.WriteString(w, h.Email)
	if err != nil {
		return
	}

	_, err = io.WriteString(w, "</em>")
	if err != nil {
		return
	}
	return
}
