package _929

import "testing"

func Test929(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	t.Log(numUniqueEmails(emails))
}
