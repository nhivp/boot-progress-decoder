// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024 Nhi Pham

package edk2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func decodeStatusValue(value uint32) EFIStatusCodeValue {
	return EFIStatusCodeValue{
		Class:     uint8((value & EFI_STATUS_CODE_CLASS_MASK) >> 24),
		Subclass:  uint8((value & EFI_STATUS_CODE_SUBCLASS_MASK) >> 16),
		Operation: uint16(value & EFI_STATUS_CODE_OPERATION_MASK),
	}
}

func extractStatusCodeValue(codeString string) (EFIStatusCodeValue, error) {
	codeString = strings.TrimPrefix(codeString, "V")

	value, err := strconv.ParseUint(codeString, 16, 32)
	if err != nil {
		return EFIStatusCodeValue{}, fmt.Errorf("invalid status code format: %v", err)
	}

	return decodeStatusValue(uint32(value)), nil
}

func decodeStatusType(value uint32) EFIStatusCodeType {
	return EFIStatusCodeType{
		Type:     uint8(value & EFI_STATUS_CODE_TYPE_MASK),
		Severity: uint8((value & EFI_STATUS_CODE_SEVERITY_MASK) >> 24),
	}
}

func extractStatusCodeType(codeString string) (EFIStatusCodeType, error) {
	codeString = strings.TrimPrefix(codeString, "C")

	value, err := strconv.ParseUint(codeString, 16, 32)
	if err != nil {
		return EFIStatusCodeType{}, fmt.Errorf("invalid status code type format: %v", err)
	}

	return decodeStatusType(uint32(value)), nil
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func decodeOperation(statusValue EFIStatusCodeValue, isError bool) string {
	switch statusValue.Class {
	case 0x00: // Computing
		if statusValue.Operation < 0x1000 {
			if isError {
				return commonCUErrorCodeDesc[statusValue.Operation]
			} else {
				return commonCUProgressCodeDesc[statusValue.Operation]
			}
		}

		if statusValue.Operation >= 0x8000 {
			if isError {
				return "OEM Specific Error Code"
			} else {
				return "OEM Specific Progress Code"
			}
		}

		subclassOperation := statusValue.Operation &^ 0x1000
		if isError {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Computing Operation"
			case 0x01: // Host Processor
				return cUHPErrorCodeDesc[subclassOperation]
			case 0x02: // Firmware Processor
				return cUFPErrorCodeDesc[subclassOperation]
			case 0x03: // I/O Processor
				return "I/O Processor Progress Code"
			case 0x04: // Cache
				return cUCacheErrorCodeDesc[subclassOperation]
			case 0x05: // Memory
				return cUMemoryErrorCodeDesc[subclassOperation]
			case 0x06: // Chipset
				return cUChipsetErrorCodeDesc[subclassOperation]
			default:
				return "Unknown Computing Error Code"
			}
		} else {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Computing Operation"
			case 0x01: // Host Processor
				return cUHPProgressCodeDesc[subclassOperation]
			case 0x02: // Firmware Processor
				return "Firmware Processor Progress Code"
			case 0x03: // I/O Processor
				return "I/O Processor Progress Code"
			case 0x04: // Cache
				return cUCacheProgressCodeDesc[subclassOperation]
			case 0x05: // Memory
				return cUMemoryProgressCodeDesc[subclassOperation]
			case 0x06: // Chipset
				return cUChipsetProgressCodeDesc[subclassOperation]
			default:
				return "Unknown Computing Progress Code"
			}
		}
	case 0x01: // Peripheral
		if statusValue.Operation < 0x1000 {
			if isError {
				return commonPErrorCodeDesc[statusValue.Operation]
			} else {
				return commonPProgressCodeDesc[statusValue.Operation]
			}
		}

		if statusValue.Operation >= 0x8000 {
			if isError {
				return "OEM Specific Error Code"
			} else {
				return "OEM Specific Progress Code"
			}
		}

		subclassOperation := statusValue.Operation &^ 0x1000
		if isError {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Peripheral Operation"
			case 0x01: // Keyboard
				return pKeyBoardErrorCodeDesc[subclassOperation]
			case 0x02: // Mouse
				return pMouseErrorCodeDesc[subclassOperation]
			case 0x03: // Local Console
				return "Unknown Local Console Error Code"
			case 0x04: // Remote Console
				return "Unknown Remote Console Error Code"
			case 0x05: // Serial Port
				return "Unknown Serial Port Error Code"
			case 0x06: // Parallel Port
				return "Unknown Parallel Port Error Code"
			case 0x07: // Fixed Media
				return "Unknown Fixed Media Error Code"
			case 0x08: // Removable Media
				return "Unknown Removable Media Error Code"
			case 0x09: // Audio Input
				return "Unknown Audio Input Error Code"
			case 0x0A: // Audio Output
				return "Unknown Audio Output Error Code"
			case 0x0B: // LCD Device
				return "Unknown LCD Device Error Code"
			case 0x0C: // Network
				return "Unknown Network Error Code"
			case 0x0D: // Docking
				return "Unknown Docking Error Code"
			case 0x0E: // TPM
				return "Unknown TPM Error Code"
			}
		} else {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Peripheral Progress Code"
			case 0x01: // Keyboard
				return pKeyBoardProgressCodeDesc[subclassOperation]
			case 0x02: // Mouse
				return pMouseProgressCodeDesc[subclassOperation]
			case 0x03: // Local Console
				return "Unknown Local Console Progress Code"
			case 0x04: // Remote Console
				return "Unknown Remote Console Progress Code"
			case 0x05: // Serial Port
				return pSerialPortProgressCodeDesc[subclassOperation]
			case 0x06: // Parallel Port
				return "Unknown Parallel Port Progress Code"
			case 0x07: // Fixed Media
				return "Unknown Fixed Media Progress Code"
			case 0x08: // Removable Media
				return "Unknown Removable Media Progress Code"
			case 0x09: // Audio Input
				return "Unknown Audio Input Progress Code"
			case 0x0A: // Audio Output
				return "Unknown Audio Output Progress Code"
			case 0x0B: // LCD Device
				return "Unknown LCD Device Progress Code"
			case 0x0C: // Network
				return "Unknown Network Progress Code"
			case 0x0D: // Docking
				return "Unknown Docking Progress Code"
			case 0x0E: // TPM
				return "Unknown TPM Progress Code"
			}
		}
	case 0x02: // I/O Bus
		if statusValue.Operation < 0x1000 {
			if isError {
				return commonIOBErrorCodeDesc[statusValue.Operation]
			} else {
				return commonIOBProgressCodeDesc[statusValue.Operation]
			}
		}

		if statusValue.Operation >= 0x8000 {
			if isError {
				return "OEM Specific Error Code"
			} else {
				return "OEM Specific Progress Code"
			}
		}

		subclassOperation := statusValue.Operation &^ 0x1000
		if isError {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified I/O Bus Error Code"
			case 0x01: // PCI
				return iOBPciErrorCodeDesc[subclassOperation]
			case 0x02: // USB
				return "Unknown USB Error Code"
			case 0x06: // LPC
				return "Unknown LPC Error Code"
			case 0x07: // SCSI
				return "Unknown SCSI Error Code"
			case 0x08: // ATAPI
				return iOBAtaErrorCodeDesc[subclassOperation]
			case 0x0B: // SMBUS
				return "Unknown SMBUS Error Code"
			case 0x0C: // I2C
				return "Unknown I2C Error Code"
			}
		} else {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified I/O Bus Progress Code"
			case 0x01: // PCI
				return iOBPciProgressCodeDesc[subclassOperation]
			case 0x02: // USB
				return "Unknown USB Progress Code"
			case 0x06: // LPC
				return "Unknown LPC Progress Code"
			case 0x07: // SCSI
				return "Unknown SCSI Progress Code"
			case 0x08: // ATAPI
				return iOBAtaProgressCodeDesc[subclassOperation]
			case 0x0B: // SMBUS
				return "Unknown SMBUS Progress Code"
			case 0x0C: // I2C
				return "Unknown I2C Progress Code"
			}
		}
	case 0x03: // Software
		if statusValue.Operation < 0x1000 {
			if isError {
				return commonSWErrorCodeDesc[statusValue.Operation]
			} else {
				return commonSWProgressCodeDesc[statusValue.Operation]
			}
		}

		if statusValue.Operation >= 0x8000 {
			if isError {
				return "OEM Specific Error Code"
			} else {
				return "OEM Specific Progress Code"
			}
		}

		subclassOperation := statusValue.Operation &^ 0x1000
		if isError {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Software Error"
			case 0x01: // SEC
				return "Unknown SEC Error Code"
			case 0x02: // PEI Core
				return swPeiCoreErrorCodeDesc[subclassOperation]
			case 0x03: // PEI Driver
				return swPeiErrorCodeDesc[subclassOperation]
			case 0x04: // DXE Core
				return swDxeFoundationErrorCodeDesc[subclassOperation]
			case 0x05: // DXE Boot Driver
				return swDxeBsErrorCodeDesc[subclassOperation]
			case 0x06: // DXE Runtime Driver
				return "Unknown DXE Runtime Driver Error Code"
			case 0x07: // SMM Driver
				return "Unknown SMM Driver Error Code"
			case 0x08: // EFI Application
				return "Unknown EFI Application Error Code"
			case 0x09: // OS Loader
				return "Unknown OS Loader Error Code"
			case 0x0C: // EBC Exception
				return swEBCErrorCodeDesc[subclassOperation]
			case 0x0D: // IA32 Exception
				return swIA32ErrorCodeDesc[subclassOperation]
			case 0x0E: // IPF Exception
				return swIPFErrorCodeDesc[subclassOperation]
			case 0x0F: // PEI Service
				return swPeiServiceErrorCodeDesc[subclassOperation]
			case 0x10: // UEFI Boot Service
				return swDxeServiceErrorCodeDesc[subclassOperation]
			case 0x11: // UEFI Runtime Service
				return swDxeRtDriverProgressCodeDesc[subclassOperation]
			case 0x12: // DXE Service
				return "Unknown DXE Service Error Code"
			case 0x13: // X64 Exception
				return swX64ExceptionErrorCodeDesc[subclassOperation]
			case 0x14: // ARM Exception
				return swArmExceptionErrorCodeDesc[subclassOperation]
			default:
				return "Unknown Software Error Code"
			}
		} else {
			switch statusValue.Subclass {
			case 0x00: // Unspecified
				return "Unspecified Software Progress"
			case 0x01: // SEC
				return swSecProgressCodeDesc[subclassOperation]
			case 0x02: // PEI Core
				return swPeiCoreProgressCodeDesc[subclassOperation]
			case 0x03: // PEI Driver
				return swPeiProgressCodeDesc[subclassOperation]
			case 0x04: // DXE Core
				return swDxeCoreProgressCodeDesc[subclassOperation]
			case 0x05: // DXE Boot Driver
				return swDxeBsProgressCodeDesc[subclassOperation]
			case 0x06: // DXE Runtime Driver
				return swDxeRtProgressCodeDesc[subclassOperation]
			case 0x07: // SMM Driver
				return "Unknown SMM Driver Progress Code"
			case 0x08: // EFI Application
				return "Unknown EFI Application Progress Code"
			case 0x09: // OS Loader
				return "Unknown OS Loader Progress Code"
			case 0x0C: // EBC Exception
				return "Unknown EBC Exception Progress Code"
			case 0x0D: // IA32 Exception
				return "Unknown IA32 Exception Progress Code"
			case 0x0E: // IPF Exception
				return "Unknown IPF Exception Progress Code"
			case 0x0F: // PEI Service
				return swPeiServicesProgressCodeDesc[subclassOperation]
			case 0x10: // UEFI Boot Service
				return swBootServicesProgressCodeDesc[subclassOperation]
			case 0x11: // UEFI Runtime Service
				return swRuntimeServicesProgressCodeDesc[subclassOperation]
			case 0x12: // DXE Service
				return swDxeServicesProgressCodeDesc[subclassOperation]
			case 0x13: // X64 Exception
				return "Unknown X64 Exception Progress Code"
			case 0x14: // ARM Exception
				return "Unknown ARM Exception Progress Code"
			default:
				return "Unknown Software Progress Code"
			}
		}
	default:
		return "Unknown"
	}
	return "Unknown"
}

func DecodeStatusValue(statusCodeValue string, isError bool) (string, string, string, error) {
	statusCodeValue = strings.TrimSpace(statusCodeValue)

	value, err := extractStatusCodeValue(statusCodeValue)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to extract status code value: %v", err)
	}

	classDesc, ok := classCodeDesc[value.Class]
	if !ok {
		classDesc = "Unknown"
	}

	subclassDesc := ""
	switch value.Class {
	case 0x00:
		subclassDesc, ok = subclassComputingCodeDesc[value.Subclass]
	case 0x01:
		subclassDesc, ok = subclassPeripheralCodeDesc[value.Subclass]
	case 0x02:
		subclassDesc, ok = subclassIOBusCodeDesc[value.Subclass]
	case 0x03:
		subclassDesc, ok = subclassSoftwareCodeDesc[value.Subclass]
	}
	if !ok {
		subclassDesc = "Unknown"
	}

	operationDesc := decodeOperation(value, isError)
	if operationDesc == "" {
		operationDesc = "Unknown"
	}

	return classDesc, subclassDesc, operationDesc, nil
}

func DecodeStatusType(statusCodeType string) (string, string, error) {
	statusCodeType = strings.TrimSpace(statusCodeType)

	codeType, err := extractStatusCodeType(statusCodeType)
	if err != nil {
		return "", "", fmt.Errorf("failed to extract status code type: %v", err)
	}

	typeDesc := statusTypeDesc[codeType.Type]
	severityDesc := errorSeverityDesc[codeType.Severity]

	return typeDesc, severityDesc, nil
}
