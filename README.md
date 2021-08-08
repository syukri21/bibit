

## BIBIT TECHNICAL TEST

### 1.  SQL

```sql
SELECT
"USER"."ID",
"USER"."UserName",
( CASE "USER"."Parent" WHEN 0 THEN NULL ELSE ( SELECT "U"."UserName" FROM "USER" "U" WHERE "ID" = "USER"."ID" ) END ) AS "ParentUserName"
FROM
"USER"
```

### 2. Microcervices

- at folder `2`
- Read `README.md` Disarankan menggunakan `Docker` untuk menjalankan Aplikasi
- HTTP service at port `1323`
- gRPC service at port `1313`
- postgres untuk log ada di port `5432`


### 3. Refractor
```golang
package main

import "strings"

func findFirstStringInBrackets(s string) string {
	firstIndex := strings.IndexByte(s, '(')
	if firstIndex < 0 {
		return ""
	}
	firstIndex++
	lastIndex := strings.IndexByte(s[firstIndex:], ')')
		if lastIndex < 0 {
		return ""
	}
	return s[firstIndex : firstIndex+lastIndex]
}
```

### 4. Grouping Anagram
```golang
package anagrams

import "sorts"

func groupAnagrams(values []string) [][]string {
	m := make(map[string][]string)
	for _, v := range values {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ans [][]string
	for e := range m {
		ans = append(ans, m[e])
	}
	return ans
}
```

