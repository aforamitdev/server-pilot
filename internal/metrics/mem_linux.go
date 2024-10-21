package metrics

import "os"

type Memory struct {
	MemTotal          uint64 "mem:MemTotal"
	MemFree           uint64 "mem:MemFree"
	MemAvailable      uint64 "mem:MemAvailable"
	Buffers           uint64 "mem:Buffers"
	Cached            uint64 "mem:Cached"
	SwapCached        uint64 "mem:SwapCached"
	Active            uint64 "mem:Active"
	Inactive          uint64 "mem:Inactive"
	Unevictable       uint64 "mem:Unevictable"
	Mlocked           uint64 "mem:Mlocked"
	SwapTotal         uint64 "mem:SwapTotal"
	SwapFree          uint64 "mem:SwapFree"
	Zswap             uint64 "mem:Zswap"
	Zswapped          uint64 "mem:Zswapped"
	Dirty             uint64 "mem:Dirty"
	Writeback         uint64 "mem:Writeback"
	AnonPages         uint64 "mem:AnonPages"
	Mapped            uint64 "mem:Mapped"
	Shmem             uint64 "mem:Shmem"
	KReclaimable      uint64 "mem:KReclaimable"
	Slab              uint64 "mem:Slab"
	SReclaimable      uint64 "mem:SReclaimable"
	SUnreclaim        uint64 "mem:SUnreclaim"
	KernelStack       uint64 "mem:KernelStack"
	PageTables        uint64 "mem:PageTables"
	SecPageTables     uint64 "mem:SecPageTables"
	NFS_Unstable      uint64 "mem:NFS_Unstable"
	Bounce            uint64 "mem:Bounce"
	WritebackTmp      uint64 "mem:WritebackTmp"
	CommitLimit       uint64 "mem:CommitLimit"
	Committed_AS      uint64 "mem:Committed_AS"
	VmallocTotal      uint64 "mem:VmallocTotal"
	VmallocUsed       uint64 "mem:VmallocUsed"
	VmallocChunk      uint64 "mem:VmallocChunk"
	Percpu            uint64 "mem:Percpu"
	HardwareCorrupted uint64 "mem:HardwareCorrupted"
	AnonHugePages     uint64 "mem:AnonHugePages"
	ShmemHugePages    uint64 "mem:ShmemHugePages"
	ShmemPmdMapped    uint64 "mem:ShmemPmdMapped"
	FileHugePages     uint64 "mem:FileHugePages"
	FilePmdMapped     uint64 "mem:FilePmdMapped"
	Unaccepted        uint64 "mem:Unaccepted"
	HugePagesTotal    uint64 "mem:HugePages_Total"
	HugePagesFree     uint64 "mem:HugePages_Free"
	HugePagesRsvd     uint64 "mem:HugePages_Rsvd"
	HugePagesSurp     uint64 "mem:HugePages_Surp"
	Hugepagesize      uint64 "mem:Hugepagesize"
	Hugetlb           uint64 "mem:Hugetlb"
	DirectMap4k       uint64 "mem:DirectMap4k"
	DirectMap2M       uint64 "mem:DirectMap2M"
	DirectMap1G       uint64 "mem:DirectMap1G"
}

func NewMemory() *Memory {
	buf, err := os.ReadFile("/proc/meminfo")

	if err != nil {
		return &Memory{}
	}
	mem := &Memory{}
	return mem

}
