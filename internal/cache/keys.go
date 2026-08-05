package cache

import "fmt"

func LinkKey(shortCode string) string {
	return fmt.Sprintf("link:%s", shortCode)
}
