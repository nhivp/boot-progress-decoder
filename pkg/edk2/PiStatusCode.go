// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024 Nhi Pham

package edk2

// Below are definitions of EFI_STATUS_CODE_TYPE
//
// 0         7                           24         31
// ┌─────────┬────────────────────────────┬──────────┐
// │         │                            │          │
// │  TYPE   │          RESERVED          │ SEVERITY │
// │         │                            │          │
// └────┬────┴────────────────────────────┴─────┬────┘
//      │                                       │
//      │                                       │   │For Error Code:
//      │                                       │   │0x40: Minor
//      │                                       └───┤0x80: Major
//      │                                           │0x90: Unrecovered
//      │                                           │0xA0: Uncontained
//      │
//      │                                           │1: Progress Code
//      └───────────────────────────────────────────┤2: Error Code
//                                                  │3: Debug Code
//
// Reference: https://github.com/tianocore/edk2/blob/master/MdePkg/Include/Pi/PiStatusCode.h
//

// represents the type and severity of a status code
type EFIStatusCodeType struct {
	Type     uint8
	Severity uint8
}

// Status Code Type Masks
const (
	EFI_STATUS_CODE_TYPE_MASK     uint32 = 0x000000FF
	EFI_STATUS_CODE_SEVERITY_MASK uint32 = 0xFF000000
	EFI_STATUS_CODE_RESERVED_MASK uint32 = 0x00FFFF00
)

// Status Type mappings
var statusTypeDesc = map[uint8]string{
	0x01: "Progress Code",
	0x02: "Error Code",
	0x03: "Debug Code",
}

// Error Severity mappings
var errorSeverityDesc = map[uint8]string{
	0x40: "Minor Error",
	0x80: "Major Error",
	0x90: "Uncovered Error",
	0xA0: "Uncontained Error",
}

//
// Below are definitions of EFI_STATUS_CODE_VALUE
//
//
// 0                         16           24          31
// ┌──────────────────────────┬────────────┬───────────┐
// │                          │            │           │
// │        OPERATION         │  SUBCLASS  │   CLASS   │
// │                          │            │           │
// └──────────────────────────┴─────┬──────┴─────┬─────┘
//                                  │            │
//                                  │            │     │0x00: Computing
//                                  │            └─────┤0x01: Peripheral
//                                  │                  │0x02: I/O Bus
//                                  │                  │0x03: Software
//                                  │
//                                  │           │For I/O Bus:
//        ┌─────────────────────────┴────────┐  │ 0x00: Unspecified
//        │                                  │  │ 0x01: PCI
//        │   │For Computing:                │  │ 0x02: USB
//        │   │ 0x00: Unspecified            │  │ 0x06: LPC
//        │   │ 0x01: Host Processor         │  │ 0x07: SCSI
//        │   │ 0x02: Firmware Processor     │  │ 0x08: ATAPI
//        │   │ 0x03: I/O Processor          │  │ 0x0B: SMBUS
//        │   │ 0x04: Cache                  │  │ 0x0C: I2C
//        │   │ 0x05: Memory                 │  │
//        │   │ 0x06: Chipset                │  │
//        │   │                              │  │For Software:
//        └───┤                              └──┤ 0x00: Unspecified
//            │For Peripheral:                  │ 0x01: SEC
//            │ 0x00: Unspecified               │ 0x02: PEI Core
//            │ 0x01: Keyboard                  │ 0x03: PEI Driver
//            │ 0x02: Mouse                     │ 0x04: DXE Core
//            │ 0x03: Local Console             │ 0x05: DXE Boot Driver
//            │ 0x04: Remote Console            │ 0x06: DXE Runtime Driver
//            │ 0x05: Serial Port               │ 0x07: SMM Driver
//            │ 0x06: Parallel Port             │ 0x08: EFI Application
//            │ 0x07: Fixed Media               │ 0x09: OS Loader
//            │ 0x08: Removable Media           │ 0x0C: EBC Exception
//            │ 0x09: Audio Input               │ 0x0D: X86 Exception
//            │ 0x0A: Audio Output              │ 0x0F: PEI Service
//            │ 0x0B: LCD Device                │ 0x10: UEFI Boot Service
//            │ 0x0C: Network                   │ 0x11: UEFI Runtime Service
//                                              │ 0x12: DXE Service
//
// Reference: https://github.com/tianocore/edk2/blob/master/MdePkg/Include/Pi/PiStatusCode.h
//

// represents the structure of EFI_STATUS_CODE_VALUE
type EFIStatusCodeValue struct {
	Class     uint8
	Subclass  uint8
	Operation uint16
}

// Status Code Value Masks
const (
	EFI_STATUS_CODE_CLASS_MASK     uint32 = 0xFF000000
	EFI_STATUS_CODE_SUBCLASS_MASK  uint32 = 0x00FF0000
	EFI_STATUS_CODE_OPERATION_MASK uint32 = 0x0000FFFF
)

// Class mappings
var classCodeDesc = map[uint8]string{
	0x00: "Computing",
	0x01: "Peripheral",
	0x02: "I/O Bus",
	0x03: "Software",
}

// Subclass mappings for Computing subclass
var subclassComputingCodeDesc = map[uint8]string{
	0x00: "Unspecified",
	0x01: "Host Processor",
	0x02: "Firmware Processor",
	0x03: "I/O Processor",
	0x04: "Cache",
	0x05: "Memory",
	0x06: "Chipset",
}

// Subclass mappings for Peripheral subclass
var subclassPeripheralCodeDesc = map[uint8]string{
	0x00: "Unspecified",
	0x01: "Keyboard",
	0x02: "Mouse",
	0x03: "Local Console",
	0x04: "Remote Console",
	0x05: "Serial Port",
	0x06: "Parallel Port",
	0x07: "Fixed Media",
	0x08: "Removable Media",
	0x09: "Audio Input",
	0x0A: "Audio Output",
	0x0B: "LCD Device",
	0x0C: "Network",
	0x0D: "Docking",
	0x0E: "TPM",
}

// Subclass mappings for I/O Bus subclass
var subclassIOBusCodeDesc = map[uint8]string{
	0x00: "Unspecified",
	0x01: "PCI",
	0x02: "USB",
	0x06: "LPC",
	0x07: "SCSI",
	0x08: "ATAPI",
	0x0B: "SMBUS",
	0x0C: "I2C",
}

// Subclass mappings for Software subclass
var subclassSoftwareCodeDesc = map[uint8]string{
	0x00: "Unspecified",
	0x01: "SEC",
	0x02: "PEI Core",
	0x03: "PEI Driver",
	0x04: "DXE Core",
	0x05: "DXE Boot Driver",
	0x06: "DXE Runtime Driver",
	0x07: "SMM Driver",
	0x08: "EFI Application",
	0x09: "OS Loader",
	0x0C: "EBC Exception",
	0x0D: "X86 Exception",
	0x0F: "PEI Service",
	0x10: "UEFI Boot Service",
	0x11: "UEFI Runtime Service",
	0x12: "DXE Service",
	0x13: "X64 Exception",
	0x14: "ARM Exception",
}

//
// Below are mappings for common or subclass specific operation
//

var commonCUProgressCodeDesc = map[uint16]string{
	0x0000: "Initialization Begin",
	0x0001: "Initialization End",
}

var cUHPProgressCodeDesc = map[uint16]string{
	0x0000: "Power On Init",
	0x0001: "Cache Init",
	0x0002: "RAM Init",
	0x0003: "Memory Controller Init",
	0x0004: "IO Init",
	0x0005: "BSP Select",
	0x0006: "BSP Reselect",
	0x0007: "AP Init",
	0x0008: "SMM Init",
}

var cUCacheProgressCodeDesc = map[uint16]string{
	0x0000: "Presence Detect",
	0x0001: "Configuration",
}

var cUMemoryProgressCodeDesc = map[uint16]string{
	0x0000: "SPD Read",
	0x0001: "Presence Detect",
	0x0002: "Timing",
	0x0003: "Configuring",
	0x0004: "Optimizing",
	0x0005: "Init",
	0x0006: "Test",
}

var cUChipsetProgressCodeDesc = map[uint16]string{
	0x0000: "PEI CAR South Bridge Initialization",
	0x0001: "PEI CAR North Bridge Initialization",
	0x0002: "PEI MEM South Bridge Initialization",
	0x0003: "PEI MEM North Bridge Initialization",
	0x0004: "DXE PCI Host Bridge Initialization",
	0x0005: "DXE North Bridge Initialization",
	0x0006: "DXE North Bridge SMM Initialization",
	0x0007: "DXE South Bridge Runtime Services Initialization",
	0x0008: "DXE South Bridge Initialization",
	0x0009: "DXE South Bridge SMM Initialization",
	0x000A: "DXE South Bridge Devices Initialization",
}

var commonCUErrorCodeDesc = map[uint16]string{
	0x0000: "Unspecified",
	0x0001: "Disabled",
	0x0002: "Not Supported",
	0x0003: "Not Detected",
	0x0004: "Not Configured",
}

var cUHPErrorCodeDesc = map[uint16]string{
	0x0000: "Invalid Type",
	0x0001: "Invalid Speed",
	0x0002: "Mismatch",
	0x0003: "Timer Expired",
	0x0004: "Self Test",
	0x0005: "Internal",
	0x0006: "Thermal",
	0x0007: "Low Voltage",
	0x0008: "High Voltage",
	0x0009: "Cache",
	0x000A: "Microcode Update",
	0x000B: "Correctable",
	0x000C: "Uncorrectable",
	0x000D: "No Microcode Update",
}

var cUFPErrorCodeDesc = map[uint16]string{
	0x0000: "Hard Fail",
	0x0001: "Soft Fail",
	0x0002: "Common Error",
}

var cUCacheErrorCodeDesc = map[uint16]string{
	0x0000: "Invalid Type",
	0x0001: "Invalid Speed",
	0x0002: "Invalid Size",
	0x0003: "Mismatch",
}

var cUMemoryErrorCodeDesc = map[uint16]string{
	0x0000: "Invalid Type",
	0x0001: "Invalid Speed",
	0x0002: "Correctable",
	0x0003: "Uncorrectable",
	0x0004: "SPD Fail",
	0x0005: "Invalid Size",
	0x0006: "Mismatch",
	0x0007: "S3 Resume Fail",
	0x0008: "Update Fail",
	0x0009: "None Detected",
	0x000A: "None Useful",
}

var cUChipsetErrorCodeDesc = map[uint16]string{
	0x0000: "Bad Battery",
	0x0001: "DXE North Bridge Error",
	0x0002: "DXE South Bridge Error",
	0x0003: "Intruder Detect",
}

var commonPProgressCodeDesc = map[uint16]string{
	0x0000: "Init",
	0x0001: "Reset",
	0x0002: "Disable",
	0x0003: "Presence Detect",
	0x0004: "Enable",
	0x0005: "Reconfig",
	0x0006: "Detected",
	0x0007: "Removed",
}

var pKeyBoardProgressCodeDesc = map[uint16]string{
	0x0000: "Clear Buffer",
	0x0001: "Self Test",
}

var pMouseProgressCodeDesc = map[uint16]string{
	0x0000: "Self Test",
}

var pSerialPortProgressCodeDesc = map[uint16]string{
	0x0000: "Clear Buffer",
}

var commonPErrorCodeDesc = map[uint16]string{
	0x0000: "Non Specific",
	0x0001: "Disabled",
	0x0002: "Not Supported",
	0x0003: "Not Detected",
	0x0004: "Not Configured",
	0x0005: "Interface Error",
	0x0006: "Controller Error",
	0x0007: "Input Error",
	0x0008: "Output Error",
	0x0009: "Resource Conflict",
}

var pKeyBoardErrorCodeDesc = map[uint16]string{
	0x0000: "Locked",
	0x0001: "Stuck Key",
	0x0002: "Buffer Full",
}

var pMouseErrorCodeDesc = map[uint16]string{
	0x0000: "Locked",
}

var commonIOBProgressCodeDesc = map[uint16]string{
	0x0000: "Init",
	0x0001: "Reset",
	0x0002: "Disable",
	0x0003: "Detect",
	0x0004: "Enable",
	0x0005: "Reconfig",
	0x0006: "Hotplug",
}

var iOBPciProgressCodeDesc = map[uint16]string{
	0x0000: "PCI Bus Enumeration",
	0x0001: "PCI Resource Allocation",
	0x0002: "PCI HPC Initialization",
}

var iOBAtaProgressCodeDesc = map[uint16]string{
	0x0000: "SMART Enable",
	0x0001: "SMART Disable",
	0x0002: "SMART Overthreshold",
	0x0003: "SMART Underthreshold",
}

var commonIOBErrorCodeDesc = map[uint16]string{
	0x0000: "Non Specific",
	0x0001: "Disabled",
	0x0002: "Not Supported",
	0x0003: "Not Detected",
	0x0004: "Not Configured",
	0x0005: "Interface Error",
	0x0006: "Controller Error",
	0x0007: "Read Error",
	0x0008: "Write Error",
	0x0009: "Resource Conflict",
}

var iOBPciErrorCodeDesc = map[uint16]string{
	0x0000: "PCI PERR",
	0x0001: "PCI SERR",
}

var iOBAtaErrorCodeDesc = map[uint16]string{
	0x0000: "ATA Bus SMART Not Supported",
	0x0001: "ATA Bus SMART Disabled",
}

var commonSWProgressCodeDesc = map[uint16]string{
	0x0000: "Init",
	0x0001: "Load",
	0x0002: "Init Begin",
	0x0003: "Init End",
	0x0004: "Authenticate Begin",
	0x0005: "Authenticate End",
	0x0006: "Input Wait",
	0x0007: "User Setup",
}

var swSecProgressCodeDesc = map[uint16]string{
	0x0000: "SEC Entry Point",
	0x0001: "SEC Handoff To Next",
}

var swPeiCoreProgressCodeDesc = map[uint16]string{
	0x0000: "PEI Core Entry Point",
	0x0001: "PEI Core Handoff To Next",
	0x0002: "PEI Core Return To Last",
}

var swPeiProgressCodeDesc = map[uint16]string{
	0x0000: "PEI Recovery Begin",
	0x0001: "PEI Capsule Load",
	0x0002: "PEI Capsule Start",
	0x0003: "PEI Recovery User",
	0x0004: "PEI Recovery Auto",
	0x0005: "PEI S3 Boot Script",
	0x0006: "PEI OS Wake",
	0x0007: "PEI S3 Started",
}

var swDxeCoreProgressCodeDesc = map[uint16]string{
	0x0000: "DXE Core Entry Point",
	0x0001: "DXE Core Handoff To Next",
	0x0002: "DXE Core Return To Last",
	0x0003: "DXE Core Start Driver",
	0x0004: "DXE Core Arch Ready",
}

var swDxeBsProgressCodeDesc = map[uint16]string{
	0x0000: "DXE BS Legacy OpROM Init",
	0x0001: "DXE BS Ready To Boot Event",
	0x0002: "DXE BS Legacy Boot Event",
	0x0003: "DXE BS Exit Boot Services Event",
	0x0004: "DXE BS Virtual Address Change Event",
	0x0005: "DXE BS Variable Services Init",
	0x0006: "DXE BS Variable Reclaim",
	0x0007: "DXE BS Attempt Boot Order Event",
	0x0008: "DXE BS Config Reset",
	0x0009: "DXE BS CSM Init",
}

var swDxeRtProgressCodeDesc = map[uint16]string{
	0x0000: "EFI RT Entry Point",
	0x0001: "EFI RT Handoff To Next",
	0x0002: "EFI RT Return To Last",
}

var swPeiServicesProgressCodeDesc = map[uint16]string{
	0x0000: "PEI Service Install PPI",
	0x0001: "PEI Service Reinstall PPI",
	0x0002: "PEI Service Locate PPI",
	0x0003: "PEI Service Notify PPI",
	0x0004: "PEI Service Get Boot Mode",
	0x0005: "PEI Service Set Boot Mode",
	0x0006: "PEI Service Get HOB List",
	0x0007: "PEI Service Create HOB",
	0x0008: "PEI Service FFS Find Next Volume",
	0x0009: "PEI Service FFS Find Next File",
	0x000A: "PEI Service FFS Find Section Data",
	0x000B: "PEI Service Install PEI Memory",
	0x000C: "PEI Service Allocate Pages",
	0x000D: "PEI Service Allocate Pool",
	0x000E: "PEI Service Copy Mem",
	0x000F: "PEI Service Set Mem",
	0x0010: "PEI Service Reset System",
	0x0013: "PEI Service FFS Find File By Name",
	0x0014: "PEI Service FFS Get File Info",
	0x0015: "PEI Service FFS Get Volume Info",
	0x0016: "PEI Service FFS Register For Shadow",
}

var swBootServicesProgressCodeDesc = map[uint16]string{
	0x0000: "EFI BS Raise TPL",
	0x0001: "EFI BS Restore TPL",
	0x0002: "EFI BS Allocate Pages",
	0x0003: "EFI BS Free Pages",
	0x0004: "EFI BS Get Memory Map",
	0x0005: "EFI BS Allocate Pool",
	0x0006: "EFI BS Free Pool",
	0x0007: "EFI BS Create Event",
	0x0008: "EFI BS Set Timer",
	0x0009: "EFI BS Wait For Event",
	0x000A: "EFI BS Signal Event",
	0x000B: "EFI BS Close Event",
	0x000C: "EFI BS Check Event",
	0x000D: "EFI BS Install Protocol Interface",
	0x000E: "EFI BS Reinstall Protocol Interface",
	0x000F: "EFI BS Uninstall Protocol Interface",
	0x0010: "EFI BS Handle Protocol",
	0x0011: "EFI BS PC Handle Protocol",
	0x0012: "EFI BS Register Protocol Notify",
	0x0013: "EFI BS Locate Handle",
	0x0014: "EFI BS Install Configuration Table",
	0x0015: "EFI BS Load Image",
	0x0016: "EFI BS Start Image",
	0x0017: "EFI BS Exit",
	0x0018: "EFI BS Unload Image",
	0x0019: "EFI BS Exit Boot Services",
	0x001A: "EFI BS Get Next Monotonic Count",
	0x001B: "EFI BS Stall",
	0x001C: "EFI BS Set Watchdog Timer",
	0x001D: "EFI BS Connect Controller",
	0x001E: "EFI BS Disconnect Controller",
	0x001F: "EFI BS Open Protocol",
	0x0020: "EFI BS Close Protocol",
	0x0021: "EFI BS Open Protocol Information",
	0x0022: "EFI BS Protocols Per Handle",
	0x0023: "EFI BS Locate Handle Buffer",
	0x0024: "EFI BS Locate Protocol",
	0x0025: "EFI BS Install Multiple Interfaces",
	0x0026: "EFI BS Uninstall Multiple Interfaces",
	0x0027: "EFI BS Calculate CRC32",
	0x0028: "EFI BS Copy Mem",
	0x0029: "EFI BS Set Mem",
	0x002A: "EFI BS Create Event Ex",
}

var swRuntimeServicesProgressCodeDesc = map[uint16]string{
	0x0000: "EFI RS Get Time",
	0x0001: "EFI RS Set Time",
	0x0002: "EFI RS Get Wakeup Time",
	0x0003: "EFI RS Set Wakeup Time",
	0x0004: "EFI RS Set Virtual Address Map",
	0x0005: "EFI RS Convert Pointer",
	0x0006: "EFI RS Get Variable",
	0x0007: "EFI RS Get Next Variable Name",
	0x0008: "EFI RS Set Variable",
	0x0009: "EFI RS Get Next High Monotonic Count",
	0x000A: "EFI RS Reset System",
	0x000B: "EFI RS Update Capsule",
	0x000C: "EFI RS Query Capsule Capabilities",
	0x000D: "EFI RS Query Variable Info",
}

var swDxeServicesProgressCodeDesc = map[uint16]string{
	0x0000: "EFI DS Add Memory Space",
	0x0001: "EFI DS Allocate Memory Space",
	0x0002: "EFI DS Free Memory Space",
	0x0003: "EFI DS Remove Memory Space",
	0x0004: "EFI DS Get Memory Space Descriptor",
	0x0005: "EFI DS Set Memory Space Attributes",
	0x0006: "EFI DS Get Memory Space Map",
	0x0007: "EFI DS Add IO Space",
	0x0008: "EFI DS Allocate IO Space",
	0x0009: "EFI DS Free IO Space",
	0x000A: "EFI DS Remove IO Space",
	0x000B: "EFI DS Get IO Space Descriptor",
	0x000C: "EFI DS Get IO Space Map",
	0x000D: "EFI DS Dispatch",
	0x000E: "EFI DS Schedule",
	0x000F: "EFI DS Trust",
	0x0010: "EFI DS Process Firmware Volume",
}

var commonSWErrorCodeDesc = map[uint16]string{
	0x0000: "Non-specific",
	0x0001: "Load Error",
	0x0002: "Invalid Parameter",
	0x0003: "Unsupported",
	0x0004: "Invalid Buffer",
	0x0005: "Out of Resources",
	0x0006: "Aborted",
	0x0007: "Illegal Software State",
	0x0008: "Illegal Hardware State",
	0x0009: "Start Error",
	0x000A: "Bad Date Time",
	0x000B: "CFG Invalid",
	0x000C: "CFG CLR Request",
	0x000D: "CFG Default",
	0x000E: "PWD Invalid",
	0x000F: "PWD CLR Request",
	0x0010: "PWD Cleared",
	0x0011: "Event Log Full",
	0x0012: "Write Protected",
	0x0013: "FV Corrupted",
	0x0014: "Inconsistent Memory Map",
}

var swPeiCoreErrorCodeDesc = map[uint16]string{
	0x0000: "DXE Core Corrupt",
	0x0001: "DXEIPL Not Found",
	0x0002: "Memory Not Installed",
}

var swPeiErrorCodeDesc = map[uint16]string{
	0x0000: "No Recovery Capsule",
	0x0001: "Invalid Capsule Descriptor",
	0x0002: "S3 Resume PPI Not Found",
	0x0003: "S3 Boot Script Error",
	0x0004: "S3 OS Wake Error",
	0x0005: "S3 Resume Failed",
	0x0006: "Recovery PPI Not Found",
	0x0007: "Recovery Failed",
	0x0008: "S3 Resume Error",
	0x0009: "Invalid Capsule",
}

var swDxeFoundationErrorCodeDesc = map[uint16]string{
	0x0000: "No Arch",
}

var swDxeBsErrorCodeDesc = map[uint16]string{
	0x0000: "Legacy OpROM No Space",
	0x0001: "Invalid Password",
	0x0002: "Boot Option Load Error",
	0x0003: "Boot Option Failed",
	0x0004: "Invalid IDE Password",
}

var swEBCErrorCodeDesc = map[uint16]string{
	0x0000: "Undefined",
	0x0001: "Divide Error",
	0x0002: "Debug",
	0x0003: "Breakpoint",
	0x0004: "Overflow",
	0x0005: "Invalid Opcode",
	0x0006: "Stack Fault",
	0x0007: "Alignment Check",
	0x0008: "Instruction Encoding",
	0x0009: "Bad Break",
	0x000A: "Step",
}

var swIA32ErrorCodeDesc = map[uint16]string{
	0x0000: "Divide Error",
	0x0001: "Debug",
	0x0002: "NMI",
	0x0003: "Breakpoint",
	0x0004: "Overflow",
	0x0005: "Bound",
	0x0006: "Invalid Opcode",
	0x0008: "Double Fault",
	0x000A: "Invalid TSS",
	0x000B: "Segment Not Present",
	0x000C: "Stack Fault",
	0x000D: "GP Fault",
	0x000E: "Page Fault",
	0x0010: "FP Error",
	0x0011: "Alignment Check",
	0x0012: "Machine Check",
	0x0013: "SIMD",
}

var swIPFErrorCodeDesc = map[uint16]string{
	0x0000: "ALT DTLB",
	0x0001: "DNESTED TLB",
	0x0002: "Breakpoint",
	0x0003: "External Interrupt",
	0x0004: "Gen Except",
	0x0005: "NAT Consumption",
	0x0006: "Debug Except",
	0x0007: "Unaligned Access",
	0x0008: "FP Fault",
	0x0009: "FP Trap",
	0x000A: "Taken Branch",
	0x000B: "Single Step",
}

var swPeiServiceErrorCodeDesc = map[uint16]string{
	0x0000: "Reset Not Available",
	0x0001: "Memory Installed Twice",
}

var swDxeServiceErrorCodeDesc = map[uint16]string{
	0x0005: "Begin Connecting Drivers",
	0x0006: "Verifying Password",
}

var swDxeRtDriverProgressCodeDesc = map[uint16]string{
	0x0000: "S0",
	0x0001: "S1",
	0x0002: "S2",
	0x0003: "S3",
	0x0004: "S4",
	0x0005: "S5",
}

var swX64ExceptionErrorCodeDesc = map[uint16]string{
	0x0000: "Divide Error",
	0x0001: "Debug",
	0x0002: "NMI",
	0x0003: "Breakpoint",
	0x0004: "Overflow",
	0x0005: "Bound",
	0x0006: "Invalid Opcode",
	0x0008: "Double Fault",
	0x000A: "Invalid TSS",
	0x000B: "Segment Not Present",
	0x000C: "Stack Fault",
	0x000D: "GP Fault",
	0x000E: "Page Fault",
	0x0010: "FP Error",
	0x0011: "Alignment Check",
	0x0012: "Machine Check",
	0x0013: "SIMD",
}

var swArmExceptionErrorCodeDesc = map[uint16]string{
	0x0000: "Reset",
	0x0001: "Undefined Instruction",
	0x0002: "Software Interrupt",
	0x0003: "Prefetch Abort",
	0x0004: "Data Abort",
	0x0005: "Reserved",
	0x0006: "IRQ",
	0x0007: "FIQ",
}
