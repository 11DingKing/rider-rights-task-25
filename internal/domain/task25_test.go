package domain

import "testing"

func TestDetailAuditContinuesAfterFirstPage(t *testing.T) {
	next, done := NextDetailAuditPage(0, DetailAuditPageSize, 250)
	if done {
		t.Fatal("detail audit stopped after first page")
	}
	if next != DetailAuditPageSize {
		t.Fatalf("next offset = %d", next)
	}
}
