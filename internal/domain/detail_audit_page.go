package domain

const DetailAuditPageSize = 100

func NextDetailAuditPage(offset, received, total int) (int, bool) {
	if offset < 0 {
		offset = 0
	}
	if received < 0 {
		received = 0
	}
	next := offset + received
	if received == 0 || next >= total {
		return next, true
	}
	return next, false
}
