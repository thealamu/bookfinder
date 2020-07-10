/* Package bookdetails defines data for any book */
package bookdetails

type BookDetails struct {
	ID       string
	Title    string
	Subtitle string
	Authors  []string
	Desc     string
	Link     string
}
