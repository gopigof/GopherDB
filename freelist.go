package main

type freelist struct {
	maxPage       pgnum
	releasedPages []pgnum
}

func newFreeList() *freelist {
	return &freelist{
		maxPage:       0,
		releasedPages: []pgnum{},
	}
}

func (fr *freelist) getNextPage() pgnum {
	if len(fr.releasedPages) != 0 {
		pageID := fr.releasedPages[len(fr.releasedPages)-1]
		fr.releasedPages = fr.releasedPages[:len(fr.releasedPages)-1]
		return pageID
	}
	fr.maxPage += 1
	return fr.maxPage
}

func (fr *freelist) releasePage(page pgnum) {
	fr.releasedPages = append(fr.releasedPages, page)
}
