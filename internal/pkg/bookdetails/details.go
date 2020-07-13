/* Package bookdetails defines data for any book */
package bookdetails

type BookDetails struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	Authors  []string `json:"authors"`
	Desc     string   `json:"desc"`
	Link     string   `json:"link"`
}
