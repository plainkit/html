package html

import twmerge "github.com/Oudwins/tailwind-merge-go"

func ClassMerge(classes ...string) string {
	return twmerge.Merge(classes...)
}
