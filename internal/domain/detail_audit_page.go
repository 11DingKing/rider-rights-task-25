package domain

const DetailAuditPageSize = 100

func NextDetailAuditPage(offset, received, total int) (int, bool) {
	next := offset + received
	// Stop once we have read every record up to the reported total, or as soon
	// as a page comes back empty so pagination always terminates safely even
	// when the total shifts between fetches.
	if received == 0 || next >= total {
		return next, true
	}
	return next, false
}
