package _929

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	// split witch @
	data := make(map[string]struct{})
	for _, email := range emails {
		res := strings.Split(email, "@")
		local, domain := res[0], res[1]
		localRes := strings.Split(local, "+")[0]
		localRes = strings.Replace(localRes, ".", "", -1)
		email = fmt.Sprintf("%s@%s", localRes, domain)
		data[email] = struct{}{}
	}

	fmt.Printf("%+v\n", data)

	// 使用 map 去重
	return len(data)
}
