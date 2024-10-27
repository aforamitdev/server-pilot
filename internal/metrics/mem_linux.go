package metrics

import (
	"fmt"
	"os"
)

type Memory struct {
	MemTotal     uint64 "mem:MemTotal"
	MemFree      uint64 "mem:MemFree"
	MemAvailable uint64 "mem:MemAvailable"
	Buffers      uint64 "mem:Buffers"
	Cached       uint64 "mem:Cached"
	SwapCached   uint64 "mem:SwapCached"
	Active       uint64 "mem:Active"
	Inactive     uint64 "mem:Inactive"
	Unevictable  uint64 "mem:Unevictable"
	Mlocked      uint64 "mem:Mlocked"
	SwapTotal    uint64 "mem:SwapTotal"
	SwapFree     uint64 "mem:SwapFree"
	Zswap        uint64 "mem:Zswap"
	Zswapped     uint64 "mem:Zswapped"
	Dirty        uint64 "mem:Dirty"
}

func NewMemory() *Memory {
	buf, err := os.ReadFile("/proc/meminfo")
	fmt.Println(buf)
	if err != nil {
		return &Memory{}
	}
	mem := &Memory{}
	return mem

}
