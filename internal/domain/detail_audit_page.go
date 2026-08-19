package domain

const DetailAuditPageSize = 100

func NextDetailAuditPage(offset, received, total int) (int, bool) {
	next := offset + received
	return next, true
}
