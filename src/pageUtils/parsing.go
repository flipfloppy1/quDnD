package pageUtils

import (
	"encoding/xml"
)

type ParseHTMLPage struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   struct {
			Root string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}

type RestPageSearchResults struct {
	ID           int    `json:"id"`
	Key          string `json:"key"`
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	MatchedTitle any    `json:"matched_title"`
	Description  any    `json:"description"`
	Thumbnail    struct {
		Mimetype string `json:"mimetype"`
		Size     int    `json:"size"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		Duration any    `json:"duration"`
		URL      string `json:"url"`
	} `json:"thumbnail"`
}
type RestPageSearch struct {
	Pages []RestPageSearchResults `json:"pages"`
}

type PageXMLOutput struct {
	XMLName xml.Name `xml:"div"`
	Text    string   `xml:",chardata"`
	Class   string   `xml:"class,attr"`
	Div     []struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Style string `xml:"style,attr"`
		Div   []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			Style string `xml:"style,attr"`
			B     struct {
				Text string `xml:",chardata"`
				Span struct {
					Text  string `xml:",chardata"`
					Style string `xml:"style,attr"`
				} `xml:"span"`
			} `xml:"b"`
			Div []struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
				A     struct {
					Text  string `xml:",chardata"`
					Href  string `xml:"href,attr"`
					Class string `xml:"class,attr"`
					Img   struct {
						Text     string `xml:",chardata"`
						Alt      string `xml:"alt,attr"`
						Src      string `xml:"src,attr"`
						Decoding string `xml:"decoding,attr"`
						Width    string `xml:"width,attr"`
						Height   string `xml:"height,attr"`
					} `xml:"img"`
				} `xml:"a"`
				Div []struct {
					Text  string `xml:",chardata"`
					Class string `xml:"class,attr"`
					Span  struct {
						Text  string `xml:",chardata"`
						Style string `xml:"style,attr"`
						B     struct {
							Text string `xml:",chardata"`
							Span struct {
								Text  string `xml:",chardata"`
								Style string `xml:"style,attr"`
							} `xml:"span"`
						} `xml:"b"`
					} `xml:"span"`
					P struct {
						Text string `xml:",chardata"`
						Span struct {
							Text  string `xml:",chardata"`
							Style string `xml:"style,attr"`
						} `xml:"span"`
					} `xml:"p"`
				} `xml:"div"`
				Span struct {
					Text  string `xml:",chardata"`
					Style string `xml:"style,attr"`
				} `xml:"span"`
			} `xml:"div"`
			Hr    string `xml:"hr"`
			Table struct {
				Text  string `xml:",chardata"`
				Style string `xml:"style,attr"`
				Class string `xml:"class,attr"`
				Tbody struct {
					Text string `xml:",chardata"`
					Tr   struct {
						Text string `xml:",chardata"`
						Td   []struct {
							Text  string `xml:",chardata"`
							Class string `xml:"class,attr"`
							A     struct {
								Text  string `xml:",chardata"`
								Href  string `xml:"href,attr"`
								Class string `xml:"class,attr"`
								Title string `xml:"title,attr"`
								Img   struct {
									Text     string `xml:",chardata"`
									Alt      string `xml:"alt,attr"`
									Src      string `xml:"src,attr"`
									Decoding string `xml:"decoding,attr"`
									Width    string `xml:"width,attr"`
									Height   string `xml:"height,attr"`
									Srcset   string `xml:"srcset,attr"`
								} `xml:"img"`
							} `xml:"a"`
							I struct {
								Text string `xml:",chardata"`
								A    []struct {
									Text  string `xml:",chardata"`
									Href  string `xml:"href,attr"`
									Title string `xml:"title,attr"`
									Rel   string `xml:"rel,attr"`
									Class string `xml:"class,attr"`
								} `xml:"a"`
							} `xml:"i"`
						} `xml:"td"`
					} `xml:"tr"`
				} `xml:"tbody"`
			} `xml:"table"`
			I struct {
				Text string `xml:",chardata"`
				A    []struct {
					Text  string `xml:",chardata"`
					Href  string `xml:"href,attr"`
					Title string `xml:"title,attr"`
					Rel   string `xml:"rel,attr"`
					Class string `xml:"class,attr"`
				} `xml:"a"`
			} `xml:"i"`
		} `xml:"div"`
		Table struct {
			Text        string `xml:",chardata"`
			Class       string `xml:"class,attr"`
			Cellspacing string `xml:"cellspacing,attr"`
			Cellpadding string `xml:"cellpadding,attr"`
			Tbody       struct {
				Text string `xml:",chardata"`
				Tr   []struct {
					Text string `xml:",chardata"`
					Td   struct {
						Text    string `xml:",chardata"`
						Class   string `xml:"class,attr"`
						Colspan string `xml:"colspan,attr"`
						Div     []struct {
							Text  string `xml:",chardata"`
							Class string `xml:"class,attr"`
							Style string `xml:"style,attr"`
							Div   []struct {
								Text  string `xml:",chardata"`
								Class string `xml:"class,attr"`
								Div   []struct {
									Text  string `xml:",chardata"`
									Class string `xml:"class,attr"`
									Img   struct {
										Text     string `xml:",chardata"`
										Alt      string `xml:"alt,attr"`
										Src      string `xml:"src,attr"`
										Decoding string `xml:"decoding,attr"`
										Width    string `xml:"width,attr"`
										Height   string `xml:"height,attr"`
										Srcset   string `xml:"srcset,attr"`
									} `xml:"img"`
								} `xml:"div"`
								Span struct {
									Text  string `xml:",chardata"`
									Class string `xml:"class,attr"`
									A     struct {
										Text  string `xml:",chardata"`
										Href  string `xml:"href,attr"`
										Class string `xml:"class,attr"`
										Title string `xml:"title,attr"`
									} `xml:"a"`
								} `xml:"span"`
							} `xml:"div"`
							P struct {
								Text string `xml:",chardata"`
								Br   string `xml:"br"`
							} `xml:"p"`
							B  string `xml:"b"`
							Ul struct {
								Text string `xml:",chardata"`
								Li   []struct {
									Text string `xml:",chardata"`
									A    struct {
										Text  string `xml:",chardata"`
										Href  string `xml:"href,attr"`
										Class string `xml:"class,attr"`
										Title string `xml:"title,attr"`
									} `xml:"a"`
								} `xml:"li"`
							} `xml:"ul"`
							Span struct {
								Text  string `xml:",chardata"`
								Class string `xml:"class,attr"`
							} `xml:"span"`
						} `xml:"div"`
						P struct {
							Text string `xml:",chardata"`
							B    struct {
								Text string `xml:",chardata"`
								Span struct {
									Text  string `xml:",chardata"`
									Style string `xml:"style,attr"`
								} `xml:"span"`
							} `xml:"b"`
							A struct {
								Text  string `xml:",chardata"`
								Href  string `xml:"href,attr"`
								Title string `xml:"title,attr"`
							} `xml:"a"`
						} `xml:"p"`
						Hr string `xml:"hr"`
					} `xml:"td"`
					Th struct {
						Text string `xml:",chardata"`
						Div  struct {
							Text  string `xml:",chardata"`
							Style string `xml:"style,attr"`
							Class string `xml:"class,attr"`
							Span  struct {
								Text  string `xml:",chardata"`
								Class string `xml:"class,attr"`
								Sup   string `xml:"sup"`
								Span  struct {
									Text  string `xml:",chardata"`
									Class string `xml:"class,attr"`
									A     struct {
										Text  string `xml:",chardata"`
										Href  string `xml:"href,attr"`
										Title string `xml:"title,attr"`
									} `xml:"a"`
								} `xml:"span"`
							} `xml:"span"`
						} `xml:"div"`
					} `xml:"th"`
				} `xml:"tr"`
			} `xml:"tbody"`
		} `xml:"table"`
		I struct {
			Text string `xml:",chardata"`
			A    []struct {
				Text  string `xml:",chardata"`
				Href  string `xml:"href,attr"`
				Title string `xml:"title,attr"`
				Rel   string `xml:"rel,attr"`
				Class string `xml:"class,attr"`
			} `xml:"a"`
		} `xml:"i"`
		Br []string `xml:"br"`
	} `xml:"div"`
}
