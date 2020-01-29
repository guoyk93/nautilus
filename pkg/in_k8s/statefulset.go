package in_k8s

import (
	"os"
	"strconv"
	"strings"
)

// GetStatefulSetSequenceID get the sequence id from a stateful set hostname
// on failure, returns 0
// on success, returns sequence number started with 1
func GetStatefulSetSequenceID() uint64 {
	var err error
	var hostname string
	if hostname, err = os.Hostname(); err != nil {
		return 0
	}
	return extractStatefulSetSequenceID(hostname)
}

func extractStatefulSetSequenceID(hostname string) (id uint64) {
	var err error
	splits := strings.Split(hostname, "-")
	if len(splits) == 0 {
		return
	}
	if id, err = strconv.ParseUint(splits[len(splits)-1], 10, 64); err != nil {
		return
	}
	id = id + 1
	return
}
