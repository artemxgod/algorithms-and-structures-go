package packages

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	TH32CS_SNAPPROCESS = 0x2
)

type Process struct {
	Size            uint32
	CntUsage        uint32
	ProcessID       uint32
	DefaultHeapID   uintptr
	ModuleID        uint32
	CntThreads      uint32
	ParentProcessID uint32
	PriClassBase    int32
	Flags           uint32
	ExeFile         [260]uint16
}

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	psapi               = syscall.NewLazyDLL("Psapi.dll")
	snapShotTool        = kernel32.NewProc("CreateToolhelp32Snapshot")
	firstProcess        = kernel32.NewProc("Process32FirstW")
	NextProcess         = kernel32.NewProc("Process32NextW")
	openProcess         = kernel32.NewProc("OpenProcess")
	getModuleFileNameEx = psapi.NewProc("GetModuleFileNameExW")
)

const (
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_VM_READ           = 0x0010
)

func ListProcesses() []uint32 {

	snapHandle, _, _ := snapShotTool.Call(TH32CS_SNAPPROCESS, 0)

	defer syscall.CloseHandle(syscall.Handle(snapHandle))

	var pe32 Process
	var pIDs []uint32
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	if success, _, err := firstProcess.Call(snapHandle, uintptr(unsafe.Pointer(&pe32))); success == 0 {
		fmt.Println("no processes found, error", err)
	}

	for {
		pIDs = append(pIDs, pe32.ProcessID)
		success, _, _ := NextProcess.Call(snapHandle, uintptr(unsafe.Pointer(&pe32)))
		if success == 0 {
			break
		}
	}

	return pIDs
}

func ProcsNameByID(procs []uint32) []string {
	result := make([]string, len(procs))
	var processName [1024]uint16

	for idx, val := range procs {
		processHandle, _, _ := openProcess.Call(PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, 0, uintptr(val))
		if processHandle == 0 {
			// fmt.Println("cant read process", val)
			continue
		}
		defer syscall.CloseHandle(syscall.Handle(processHandle))

		_, _, _ = getModuleFileNameEx.Call(processHandle, uintptr(unsafe.Pointer(&processName[0])), uintptr(len(processName)))

		result[idx] = syscall.UTF16ToString(processName[:])
	}

	return result
}
