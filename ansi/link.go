package ansi

import (
	"github.com/muesli/termenv"
	"io"
	"net/url"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	Text    string
	BaseURL string
	URL     string
	Child   ElementRenderer
}

func (e *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	u, err := url.Parse(e.URL)
	if err == nil {
		text := termenv.Hyperlink(resolveRelativeURL(e.BaseURL, e.URL), e.Text)
		if "#"+u.Fragment == e.URL {
			text = e.Text // if the URL is just an anchor ignore it as we don't support anchors
		}
		el := &BaseElement{
			Token: text,
			Style: ctx.options.Styles.LinkText,
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
