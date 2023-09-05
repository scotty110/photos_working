package sets

import (
	"golang.a2z.com/consolidate/internal/hash"
)

func DiffCopyOrig(alpha map[string]string, beta map[string]string) map[string]hash.FileMeta {
	m := make(map[string]hash.FileMeta)
	for k, v := range alpha {
		if _, exists := beta[k]; !exists {
			nv := hash.FileMeta{FileName: v, FileHash: k}
			m[k] = nv
		}
	}
	return m
}

func DiffCopy(alpha map[string]hash.FileMeta, beta map[string]hash.FileMeta) map[string]hash.FileMeta {
	m := make(map[string]hash.FileMeta)
	for k, v := range alpha {
		if _, exists := beta[k]; !exists {
			m[k] = v
		}
	}
	return m
}
