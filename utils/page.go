package utils

import "math"

type Page struct {
	// current page
	Page int `json:"page"`
	// page size
	Size int `json:"size"`
	// total pages
	Total int 
}

func (p *Page) GetPage() int {
	// the maximum number in one page
	max := int(math.Ceil(float64(p.Total) / float64(p.Size)))

	if p.Page > max {
		p.Page = max
	}
	return p.Page
}

// Get the start page number of one page
func (p * Page) GetStartPage() int {
	return (p.Page - 1) * p.Size
}