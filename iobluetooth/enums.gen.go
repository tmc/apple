// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPCommandRejectReason
type BluetoothAMPCommandRejectReason uint32

const (
	KBluetoothAMPManagerCommandRejectReasonCommandNotRecognized BluetoothAMPCommandRejectReason = 0
)

func (e BluetoothAMPCommandRejectReason) String() string {
	switch e {
	case KBluetoothAMPManagerCommandRejectReasonCommandNotRecognized:
		return "KBluetoothAMPManagerCommandRejectReasonCommandNotRecognized"
	default:
		return fmt.Sprintf("BluetoothAMPCommandRejectReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPCreatePhysicalLinkResponseStatus
type BluetoothAMPCreatePhysicalLinkResponseStatus uint32

const (
	KBluetoothAMPManagerCreatePhysicalLinkResponseAMPDisconnectedPhysicalLinkRequestReceived BluetoothAMPCreatePhysicalLinkResponseStatus = 0x4
	KBluetoothAMPManagerCreatePhysicalLinkResponseCollisionOccurred                          BluetoothAMPCreatePhysicalLinkResponseStatus = 0x3
	KBluetoothAMPManagerCreatePhysicalLinkResponseInvalidControllerID                        BluetoothAMPCreatePhysicalLinkResponseStatus = 0x1
	KBluetoothAMPManagerCreatePhysicalLinkResponsePhysicalLinkAlreadyExists                  BluetoothAMPCreatePhysicalLinkResponseStatus = 0x5
	KBluetoothAMPManagerCreatePhysicalLinkResponseSecurityViolation                          BluetoothAMPCreatePhysicalLinkResponseStatus = 0x6
	KBluetoothAMPManagerCreatePhysicalLinkResponseSuccess                                    BluetoothAMPCreatePhysicalLinkResponseStatus = 0
	KBluetoothAMPManagerCreatePhysicalLinkResponseUnableToStartLinkCreation                  BluetoothAMPCreatePhysicalLinkResponseStatus = 0x2
)

func (e BluetoothAMPCreatePhysicalLinkResponseStatus) String() string {
	switch e {
	case KBluetoothAMPManagerCreatePhysicalLinkResponseAMPDisconnectedPhysicalLinkRequestReceived:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseAMPDisconnectedPhysicalLinkRequestReceived"
	case KBluetoothAMPManagerCreatePhysicalLinkResponseCollisionOccurred:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseCollisionOccurred"
	case KBluetoothAMPManagerCreatePhysicalLinkResponseInvalidControllerID:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseInvalidControllerID"
	case KBluetoothAMPManagerCreatePhysicalLinkResponsePhysicalLinkAlreadyExists:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponsePhysicalLinkAlreadyExists"
	case KBluetoothAMPManagerCreatePhysicalLinkResponseSecurityViolation:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseSecurityViolation"
	case KBluetoothAMPManagerCreatePhysicalLinkResponseSuccess:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseSuccess"
	case KBluetoothAMPManagerCreatePhysicalLinkResponseUnableToStartLinkCreation:
		return "KBluetoothAMPManagerCreatePhysicalLinkResponseUnableToStartLinkCreation"
	default:
		return fmt.Sprintf("BluetoothAMPCreatePhysicalLinkResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPDisconnectPhysicalLinkResponseStatus
type BluetoothAMPDisconnectPhysicalLinkResponseStatus uint32

const (
	KBluetoothAMPManagerDisconnectPhysicalLinkResponseInvalidControllerID BluetoothAMPDisconnectPhysicalLinkResponseStatus = 0x1
	KBluetoothAMPManagerDisconnectPhysicalLinkResponseNoPhysicalLink      BluetoothAMPDisconnectPhysicalLinkResponseStatus = 0x2
	KBluetoothAMPManagerDisconnectPhysicalLinkResponseSuccess             BluetoothAMPDisconnectPhysicalLinkResponseStatus = 0
)

func (e BluetoothAMPDisconnectPhysicalLinkResponseStatus) String() string {
	switch e {
	case KBluetoothAMPManagerDisconnectPhysicalLinkResponseInvalidControllerID:
		return "KBluetoothAMPManagerDisconnectPhysicalLinkResponseInvalidControllerID"
	case KBluetoothAMPManagerDisconnectPhysicalLinkResponseNoPhysicalLink:
		return "KBluetoothAMPManagerDisconnectPhysicalLinkResponseNoPhysicalLink"
	case KBluetoothAMPManagerDisconnectPhysicalLinkResponseSuccess:
		return "KBluetoothAMPManagerDisconnectPhysicalLinkResponseSuccess"
	default:
		return fmt.Sprintf("BluetoothAMPDisconnectPhysicalLinkResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPDiscoverResponseControllerStatus
type BluetoothAMPDiscoverResponseControllerStatus uint32

const (
	KBluetoothAMPManagerDiscoverResponseControllerStatusBluetoothOnly  BluetoothAMPDiscoverResponseControllerStatus = 0x1
	KBluetoothAMPManagerDiscoverResponseControllerStatusFullCapacity   BluetoothAMPDiscoverResponseControllerStatus = 0x6
	KBluetoothAMPManagerDiscoverResponseControllerStatusHighCapacity   BluetoothAMPDiscoverResponseControllerStatus = 0x5
	KBluetoothAMPManagerDiscoverResponseControllerStatusLowCapacity    BluetoothAMPDiscoverResponseControllerStatus = 0x3
	KBluetoothAMPManagerDiscoverResponseControllerStatusMediumCapacity BluetoothAMPDiscoverResponseControllerStatus = 0x4
	KBluetoothAMPManagerDiscoverResponseControllerStatusNoCapacity     BluetoothAMPDiscoverResponseControllerStatus = 0x2
	KBluetoothAMPManagerDiscoverResponseControllerStatusPoweredDown    BluetoothAMPDiscoverResponseControllerStatus = 0
)

func (e BluetoothAMPDiscoverResponseControllerStatus) String() string {
	switch e {
	case KBluetoothAMPManagerDiscoverResponseControllerStatusBluetoothOnly:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusBluetoothOnly"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusFullCapacity:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusFullCapacity"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusHighCapacity:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusHighCapacity"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusLowCapacity:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusLowCapacity"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusMediumCapacity:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusMediumCapacity"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusNoCapacity:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusNoCapacity"
	case KBluetoothAMPManagerDiscoverResponseControllerStatusPoweredDown:
		return "KBluetoothAMPManagerDiscoverResponseControllerStatusPoweredDown"
	default:
		return fmt.Sprintf("BluetoothAMPDiscoverResponseControllerStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPGetAssocResponseStatus
type BluetoothAMPGetAssocResponseStatus uint32

const (
	KBluetoothAMPManagerGetAssocResponseInvalidControllerID BluetoothAMPGetAssocResponseStatus = 0x1
	KBluetoothAMPManagerGetAssocResponseSuccess             BluetoothAMPGetAssocResponseStatus = 0
)

func (e BluetoothAMPGetAssocResponseStatus) String() string {
	switch e {
	case KBluetoothAMPManagerGetAssocResponseInvalidControllerID:
		return "KBluetoothAMPManagerGetAssocResponseInvalidControllerID"
	case KBluetoothAMPManagerGetAssocResponseSuccess:
		return "KBluetoothAMPManagerGetAssocResponseSuccess"
	default:
		return fmt.Sprintf("BluetoothAMPGetAssocResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPGetInfoResponseStatus
type BluetoothAMPGetInfoResponseStatus uint32

const (
	KBluetoothAMPManagerGetInfoResponseInvalidControllerID BluetoothAMPGetInfoResponseStatus = 0x1
	KBluetoothAMPManagerGetInfoResponseSuccess             BluetoothAMPGetInfoResponseStatus = 0
)

func (e BluetoothAMPGetInfoResponseStatus) String() string {
	switch e {
	case KBluetoothAMPManagerGetInfoResponseInvalidControllerID:
		return "KBluetoothAMPManagerGetInfoResponseInvalidControllerID"
	case KBluetoothAMPManagerGetInfoResponseSuccess:
		return "KBluetoothAMPManagerGetInfoResponseSuccess"
	default:
		return fmt.Sprintf("BluetoothAMPGetInfoResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAMPManagerCode
type BluetoothAMPManagerCode uint32

const (
	KBluetoothAMPManagerCodeAMPChangeNotify                   BluetoothAMPManagerCode = 0x4
	KBluetoothAMPManagerCodeAMPChangeResponse                 BluetoothAMPManagerCode = 0x5
	KBluetoothAMPManagerCodeAMPCommandReject                  BluetoothAMPManagerCode = 0x1
	KBluetoothAMPManagerCodeAMPCreatePhysicalLinkRequest      BluetoothAMPManagerCode = 0xa
	KBluetoothAMPManagerCodeAMPCreatePhysicalLinkResponse     BluetoothAMPManagerCode = 0xb
	KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkRequest  BluetoothAMPManagerCode = 0xc
	KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkResponse BluetoothAMPManagerCode = 0xd
	KBluetoothAMPManagerCodeAMPDiscoverRequest                BluetoothAMPManagerCode = 0x2
	KBluetoothAMPManagerCodeAMPDiscoverResponse               BluetoothAMPManagerCode = 0x3
	KBluetoothAMPManagerCodeAMPGetAssocRequest                BluetoothAMPManagerCode = 0x8
	KBluetoothAMPManagerCodeAMPGetAssocResponse               BluetoothAMPManagerCode = 0x9
	KBluetoothAMPManagerCodeAMPGetInfoRequest                 BluetoothAMPManagerCode = 0x6
	KBluetoothAMPManagerCodeAMPGetInfoResponse                BluetoothAMPManagerCode = 0x7
	KBluetoothAMPManagerCodeReserved                          BluetoothAMPManagerCode = 0
)

func (e BluetoothAMPManagerCode) String() string {
	switch e {
	case KBluetoothAMPManagerCodeAMPChangeNotify:
		return "KBluetoothAMPManagerCodeAMPChangeNotify"
	case KBluetoothAMPManagerCodeAMPChangeResponse:
		return "KBluetoothAMPManagerCodeAMPChangeResponse"
	case KBluetoothAMPManagerCodeAMPCommandReject:
		return "KBluetoothAMPManagerCodeAMPCommandReject"
	case KBluetoothAMPManagerCodeAMPCreatePhysicalLinkRequest:
		return "KBluetoothAMPManagerCodeAMPCreatePhysicalLinkRequest"
	case KBluetoothAMPManagerCodeAMPCreatePhysicalLinkResponse:
		return "KBluetoothAMPManagerCodeAMPCreatePhysicalLinkResponse"
	case KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkRequest:
		return "KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkRequest"
	case KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkResponse:
		return "KBluetoothAMPManagerCodeAMPDisconnectPhysicalLinkResponse"
	case KBluetoothAMPManagerCodeAMPDiscoverRequest:
		return "KBluetoothAMPManagerCodeAMPDiscoverRequest"
	case KBluetoothAMPManagerCodeAMPDiscoverResponse:
		return "KBluetoothAMPManagerCodeAMPDiscoverResponse"
	case KBluetoothAMPManagerCodeAMPGetAssocRequest:
		return "KBluetoothAMPManagerCodeAMPGetAssocRequest"
	case KBluetoothAMPManagerCodeAMPGetAssocResponse:
		return "KBluetoothAMPManagerCodeAMPGetAssocResponse"
	case KBluetoothAMPManagerCodeAMPGetInfoRequest:
		return "KBluetoothAMPManagerCodeAMPGetInfoRequest"
	case KBluetoothAMPManagerCodeAMPGetInfoResponse:
		return "KBluetoothAMPManagerCodeAMPGetInfoResponse"
	case KBluetoothAMPManagerCodeReserved:
		return "KBluetoothAMPManagerCodeReserved"
	default:
		return fmt.Sprintf("BluetoothAMPManagerCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothAuthenticationRequirementsValues
type BluetoothAuthenticationRequirementsValues uint32

const (
	KBluetoothAuthenticationRequirementsMITMProtectionNotRequired                 BluetoothAuthenticationRequirementsValues = 0
	KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredDedicatedBonding BluetoothAuthenticationRequirementsValues = 0x2
	KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredGeneralBonding   BluetoothAuthenticationRequirementsValues = 0x4
	KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredNoBonding        BluetoothAuthenticationRequirementsValues = 0
	KBluetoothAuthenticationRequirementsMITMProtectionRequired                    BluetoothAuthenticationRequirementsValues = 0x1
	KBluetoothAuthenticationRequirementsMITMProtectionRequiredDedicatedBonding    BluetoothAuthenticationRequirementsValues = 0x3
	KBluetoothAuthenticationRequirementsMITMProtectionRequiredGeneralBonding      BluetoothAuthenticationRequirementsValues = 0x5
	KBluetoothAuthenticationRequirementsMITMProtectionRequiredNoBonding           BluetoothAuthenticationRequirementsValues = 0x1
)

func (e BluetoothAuthenticationRequirementsValues) String() string {
	switch e {
	case KBluetoothAuthenticationRequirementsMITMProtectionNotRequired:
		return "KBluetoothAuthenticationRequirementsMITMProtectionNotRequired"
	case KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredDedicatedBonding:
		return "KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredDedicatedBonding"
	case KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredGeneralBonding:
		return "KBluetoothAuthenticationRequirementsMITMProtectionNotRequiredGeneralBonding"
	case KBluetoothAuthenticationRequirementsMITMProtectionRequired:
		return "KBluetoothAuthenticationRequirementsMITMProtectionRequired"
	case KBluetoothAuthenticationRequirementsMITMProtectionRequiredDedicatedBonding:
		return "KBluetoothAuthenticationRequirementsMITMProtectionRequiredDedicatedBonding"
	case KBluetoothAuthenticationRequirementsMITMProtectionRequiredGeneralBonding:
		return "KBluetoothAuthenticationRequirementsMITMProtectionRequiredGeneralBonding"
	default:
		return fmt.Sprintf("BluetoothAuthenticationRequirementsValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothCompanyIdentifers
type BluetoothCompanyIdentifers uint32

const (
	KBluetoothCompanyIdentifer3Com                                 BluetoothCompanyIdentifers = 5
	KBluetoothCompanyIdentifer3DSP                                 BluetoothCompanyIdentifers = 73
	KBluetoothCompanyIdentifer3DiJoy                               BluetoothCompanyIdentifers = 84
	KBluetoothCompanyIdentifer9SolutionsOy                         BluetoothCompanyIdentifers = 102
	KBluetoothCompanyIdentiferAAMPofAmerica                        BluetoothCompanyIdentifers = 190
	KBluetoothCompanyIdentiferAAndDEngineering                     BluetoothCompanyIdentifers = 105
	KBluetoothCompanyIdentiferAAndRCambridge                       BluetoothCompanyIdentifers = 124
	KBluetoothCompanyIdentiferACTSTechnologies                     BluetoothCompanyIdentifers = 232
	KBluetoothCompanyIdentiferAMICCOMElectronics                   BluetoothCompanyIdentifers = 192
	KBluetoothCompanyIdentiferAPT                                  BluetoothCompanyIdentifers = 79
	KBluetoothCompanyIdentiferARCHOS                               BluetoothCompanyIdentifers = 207
	KBluetoothCompanyIdentiferARPDevicesUnlimited                  BluetoothCompanyIdentifers = 168
	KBluetoothCompanyIdentiferAVMBerlin                            BluetoothCompanyIdentifers = 31
	KBluetoothCompanyIdentiferAboveAverageOutcomes                 BluetoothCompanyIdentifers = 238
	KBluetoothCompanyIdentiferAccelSemiconductor                   BluetoothCompanyIdentifers = 74
	KBluetoothCompanyIdentiferAceSensor                            BluetoothCompanyIdentifers = 188
	KBluetoothCompanyIdentiferAceUni                               BluetoothCompanyIdentifers = 248
	KBluetoothCompanyIdentiferAdidas                               BluetoothCompanyIdentifers = 195
	KBluetoothCompanyIdentiferAdvancedPANMOBILSystems              BluetoothCompanyIdentifers = 145
	KBluetoothCompanyIdentiferAirohaTechnology                     BluetoothCompanyIdentifers = 148
	KBluetoothCompanyIdentiferAlcatel                              BluetoothCompanyIdentifers = 36
	KBluetoothCompanyIdentiferAlpwise                              BluetoothCompanyIdentifers = 154
	KBluetoothCompanyIdentiferAplix                                BluetoothCompanyIdentifers = 189
	KBluetoothCompanyIdentiferApple                                BluetoothCompanyIdentifers = 76
	KBluetoothCompanyIdentiferAtherosCommunications                BluetoothCompanyIdentifers = 69
	KBluetoothCompanyIdentiferAtmel                                BluetoothCompanyIdentifers = 19
	KBluetoothCompanyIdentiferAustcoCommunicationsSystems          BluetoothCompanyIdentifers = 213
	KBluetoothCompanyIdentiferAutonetMobile                        BluetoothCompanyIdentifers = 127
	KBluetoothCompanyIdentiferAvagoTechnologies                    BluetoothCompanyIdentifers = 78
	KBluetoothCompanyIdentiferBDETechnology                        BluetoothCompanyIdentifers = 180
	KBluetoothCompanyIdentiferBandXIInternational                  BluetoothCompanyIdentifers = 100
	KBluetoothCompanyIdentiferBandspeed                            BluetoothCompanyIdentifers = 32
	KBluetoothCompanyIdentiferBangAndOlufson                       BluetoothCompanyIdentifers = 259
	KBluetoothCompanyIdentiferBeatsElectronics                     BluetoothCompanyIdentifers = 204
	KBluetoothCompanyIdentiferBeautifulEnterprise                  BluetoothCompanyIdentifers = 108
	KBluetoothCompanyIdentiferBekey                                BluetoothCompanyIdentifers = 178
	KBluetoothCompanyIdentiferBelkinInternational                  BluetoothCompanyIdentifers = 92
	KBluetoothCompanyIdentiferBinauricSE                           BluetoothCompanyIdentifers = 203
	KBluetoothCompanyIdentiferBioResearchAssociates                BluetoothCompanyIdentifers = 236
	KBluetoothCompanyIdentiferBiosentronics                        BluetoothCompanyIdentifers = 219
	KBluetoothCompanyIdentiferBitsplitters                         BluetoothCompanyIdentifers = 239
	KBluetoothCompanyIdentiferBlueRadios                           BluetoothCompanyIdentifers = 133
	KBluetoothCompanyIdentiferBluegiga                             BluetoothCompanyIdentifers = 71
	KBluetoothCompanyIdentiferBluetoothSIG                         BluetoothCompanyIdentifers = 63
	KBluetoothCompanyIdentiferBose                                 BluetoothCompanyIdentifers = 158
	KBluetoothCompanyIdentiferBriarTek                             BluetoothCompanyIdentifers = 109
	KBluetoothCompanyIdentiferBroadcom                             BluetoothCompanyIdentifers = 15
	KBluetoothCompanyIdentiferCATC                                 BluetoothCompanyIdentifers = 52
	KBluetoothCompanyIdentiferCONWISETechnology                    BluetoothCompanyIdentifers = 66
	KBluetoothCompanyIdentiferCTechnologies                        BluetoothCompanyIdentifers = 38
	KBluetoothCompanyIdentiferCaenRFID                             BluetoothCompanyIdentifers = 170
	KBluetoothCompanyIdentiferCambridgeSiliconRadio                BluetoothCompanyIdentifers = 10
	KBluetoothCompanyIdentiferCinetix                              BluetoothCompanyIdentifers = 175
	KBluetoothCompanyIdentiferClarinoxTechnologies                 BluetoothCompanyIdentifers = 179
	KBluetoothCompanyIdentiferColorfy                              BluetoothCompanyIdentifers = 156
	KBluetoothCompanyIdentiferCommil                               BluetoothCompanyIdentifers = 51
	KBluetoothCompanyIdentiferConexantSystems                      BluetoothCompanyIdentifers = 28
	KBluetoothCompanyIdentiferConnectBlueAB                        BluetoothCompanyIdentifers = 113
	KBluetoothCompanyIdentiferConnecteDevice                       BluetoothCompanyIdentifers = 151
	KBluetoothCompanyIdentiferContinentialAutomotiveSystems        BluetoothCompanyIdentifers = 75
	KBluetoothCompanyIdentiferCreativeTechnology                   BluetoothCompanyIdentifers = 118
	KBluetoothCompanyIdentiferCrystalCode                          BluetoothCompanyIdentifers = 250
	KBluetoothCompanyIdentiferDanlers                              BluetoothCompanyIdentifers = 225
	KBluetoothCompanyIdentiferDeLormePublishingCompany             BluetoothCompanyIdentifers = 128
	KBluetoothCompanyIdentiferDelphi                               BluetoothCompanyIdentifers = 252
	KBluetoothCompanyIdentiferDexcom                               BluetoothCompanyIdentifers = 208
	KBluetoothCompanyIdentiferDialogSemiconductor                  BluetoothCompanyIdentifers = 210
	KBluetoothCompanyIdentiferDigianswerAS                         BluetoothCompanyIdentifers = 12
	KBluetoothCompanyIdentiferEMMicroElectronicMarin               BluetoothCompanyIdentifers = 90
	KBluetoothCompanyIdentiferEclipse                              BluetoothCompanyIdentifers = 53
	KBluetoothCompanyIdentiferEcotest                              BluetoothCompanyIdentifers = 136
	KBluetoothCompanyIdentiferEdenSoftwareConsultants              BluetoothCompanyIdentifers = 229
	KBluetoothCompanyIdentiferElcometer                            BluetoothCompanyIdentifers = 246
	KBluetoothCompanyIdentiferElgatoSystems                        BluetoothCompanyIdentifers = 206
	KBluetoothCompanyIdentiferEquinux                              BluetoothCompanyIdentifers = 134
	KBluetoothCompanyIdentiferEricssonTechnologyLicensing          BluetoothCompanyIdentifers = 0
	KBluetoothCompanyIdentiferEvluma                               BluetoothCompanyIdentifers = 201
	KBluetoothCompanyIdentiferFree2Move                            BluetoothCompanyIdentifers = 83
	KBluetoothCompanyIdentiferFreshtemp                            BluetoothCompanyIdentifers = 230
	KBluetoothCompanyIdentiferFuGoo                                BluetoothCompanyIdentifers = 257
	KBluetoothCompanyIdentiferFunaiElectric                        BluetoothCompanyIdentifers = 144
	KBluetoothCompanyIdentiferGCTSemiconductor                     BluetoothCompanyIdentifers = 45
	KBluetoothCompanyIdentiferGNNetcom                             BluetoothCompanyIdentifers = 103
	KBluetoothCompanyIdentiferGNResound                            BluetoothCompanyIdentifers = 137
	KBluetoothCompanyIdentiferGarminInternational                  BluetoothCompanyIdentifers = 135
	KBluetoothCompanyIdentiferGeLo                                 BluetoothCompanyIdentifers = 200
	KBluetoothCompanyIdentiferGeneq                                BluetoothCompanyIdentifers = 194
	KBluetoothCompanyIdentiferGeneralMotors                        BluetoothCompanyIdentifers = 104
	KBluetoothCompanyIdentiferGennum                               BluetoothCompanyIdentifers = 59
	KBluetoothCompanyIdentiferGeoforce                             BluetoothCompanyIdentifers = 157
	KBluetoothCompanyIdentiferGibsonGuitars                        BluetoothCompanyIdentifers = 98
	KBluetoothCompanyIdentiferGimbal                               BluetoothCompanyIdentifers = 140
	KBluetoothCompanyIdentiferGoogle                               BluetoothCompanyIdentifers = 224
	KBluetoothCompanyIdentiferGreenThrottleGames                   BluetoothCompanyIdentifers = 172
	KBluetoothCompanyIdentiferGroupSense                           BluetoothCompanyIdentifers = 115
	KBluetoothCompanyIdentiferHanlynnTechnologies                  BluetoothCompanyIdentifers = 123
	KBluetoothCompanyIdentiferHarmonInternational                  BluetoothCompanyIdentifers = 87
	KBluetoothCompanyIdentiferHewlettPackard                       BluetoothCompanyIdentifers = 101
	KBluetoothCompanyIdentiferHitachi                              BluetoothCompanyIdentifers = 41
	KBluetoothCompanyIdentiferHosiden                              BluetoothCompanyIdentifers = 221
	KBluetoothCompanyIdentiferIBM                                  BluetoothCompanyIdentifers = 3
	KBluetoothCompanyIdentiferIPextreme                            BluetoothCompanyIdentifers = 61
	KBluetoothCompanyIdentiferITechDynamicGlobalDistribution       BluetoothCompanyIdentifers = 153
	KBluetoothCompanyIdentiferInMusicBrands                        BluetoothCompanyIdentifers = 227
	KBluetoothCompanyIdentiferInfineonTechnologiesAG               BluetoothCompanyIdentifers = 9
	KBluetoothCompanyIdentiferIngenieurSystemgruppeZahn            BluetoothCompanyIdentifers = 171
	KBluetoothCompanyIdentiferInnovativeYachtterSolutions          BluetoothCompanyIdentifers = 262
	KBluetoothCompanyIdentiferIntegratedSiliconSolution            BluetoothCompanyIdentifers = 65
	KBluetoothCompanyIdentiferIntegratedSystemSolution             BluetoothCompanyIdentifers = 57
	KBluetoothCompanyIdentiferIntel                                BluetoothCompanyIdentifers = 2
	KBluetoothCompanyIdentiferInteropIdentifier                    BluetoothCompanyIdentifers = 65535
	KBluetoothCompanyIdentiferInventel                             BluetoothCompanyIdentifers = 30
	KBluetoothCompanyIdentiferJandM                                BluetoothCompanyIdentifers = 82
	KBluetoothCompanyIdentiferJawbone                              BluetoothCompanyIdentifers = 138
	KBluetoothCompanyIdentiferJiangsuToppowerAutomotiveElectronics BluetoothCompanyIdentifers = 155
	KBluetoothCompanyIdentiferJohnsonControls                      BluetoothCompanyIdentifers = 185
	KBluetoothCompanyIdentiferJollyLogic                           BluetoothCompanyIdentifers = 237
	KBluetoothCompanyIdentiferKCTechnology                         BluetoothCompanyIdentifers = 22
	KBluetoothCompanyIdentiferKOUKAMM                              BluetoothCompanyIdentifers = 251
	KBluetoothCompanyIdentiferKSTechnologies                       BluetoothCompanyIdentifers = 231
	KBluetoothCompanyIdentiferKawantech                            BluetoothCompanyIdentifers = 212
	KBluetoothCompanyIdentiferKeiser                               BluetoothCompanyIdentifers = 258
	KBluetoothCompanyIdentiferKensingtonComputerProductsGroup      BluetoothCompanyIdentifers = 160
	KBluetoothCompanyIdentiferKentDisplays                         BluetoothCompanyIdentifers = 243
	KBluetoothCompanyIdentiferLGElectronics                        BluetoothCompanyIdentifers = 196
	KBluetoothCompanyIdentiferLSResearch                           BluetoothCompanyIdentifers = 228
	KBluetoothCompanyIdentiferLairdTechnologies                    BluetoothCompanyIdentifers = 119
	KBluetoothCompanyIdentiferLessWire                             BluetoothCompanyIdentifers = 121
	KBluetoothCompanyIdentiferLinak                                BluetoothCompanyIdentifers = 164
	KBluetoothCompanyIdentiferLucent                               BluetoothCompanyIdentifers = 7
	KBluetoothCompanyIdentiferLudusHelsinki                        BluetoothCompanyIdentifers = 132
	KBluetoothCompanyIdentiferMC10                                 BluetoothCompanyIdentifers = 202
	KBluetoothCompanyIdentiferMStarTechnologies                    BluetoothCompanyIdentifers = 122
	KBluetoothCompanyIdentiferMacronixInternational                BluetoothCompanyIdentifers = 44
	KBluetoothCompanyIdentiferMagnetiMarelli                       BluetoothCompanyIdentifers = 169
	KBluetoothCompanyIdentiferMansella                             BluetoothCompanyIdentifers = 33
	KBluetoothCompanyIdentiferMarvellTechnologyGroup               BluetoothCompanyIdentifers = 72
	KBluetoothCompanyIdentiferMatsushitaElectricIndustrial         BluetoothCompanyIdentifers = 58
	KBluetoothCompanyIdentiferMediaTek                             BluetoothCompanyIdentifers = 70
	KBluetoothCompanyIdentiferMesoInternational                    BluetoothCompanyIdentifers = 182
	KBluetoothCompanyIdentiferMetaWatch                            BluetoothCompanyIdentifers = 163
	KBluetoothCompanyIdentiferMewTelTechnology                     BluetoothCompanyIdentifers = 47
	KBluetoothCompanyIdentiferMiCommand                            BluetoothCompanyIdentifers = 99
	KBluetoothCompanyIdentiferMicrochipTechnology                  BluetoothCompanyIdentifers = 205
	KBluetoothCompanyIdentiferMicrosoft                            BluetoothCompanyIdentifers = 6
	KBluetoothCompanyIdentiferMindTree                             BluetoothCompanyIdentifers = 106
	KBluetoothCompanyIdentiferMisfitWearables                      BluetoothCompanyIdentifers = 223
	KBluetoothCompanyIdentiferMistubishiElectric                   BluetoothCompanyIdentifers = 20
	KBluetoothCompanyIdentiferMitelSemiconductor                   BluetoothCompanyIdentifers = 16
	KBluetoothCompanyIdentiferMobilian                             BluetoothCompanyIdentifers = 55
	KBluetoothCompanyIdentiferMonster                              BluetoothCompanyIdentifers = 112
	KBluetoothCompanyIdentiferMorseProject                         BluetoothCompanyIdentifers = 242
	KBluetoothCompanyIdentiferMotorola                             BluetoothCompanyIdentifers = 8
	KBluetoothCompanyIdentiferMusik                                BluetoothCompanyIdentifers = 222
	KBluetoothCompanyIdentiferNEC                                  BluetoothCompanyIdentifers = 34
	KBluetoothCompanyIdentiferNECLightning                         BluetoothCompanyIdentifers = 149
	KBluetoothCompanyIdentiferNautilus                             BluetoothCompanyIdentifers = 244
	KBluetoothCompanyIdentiferNewlogic                             BluetoothCompanyIdentifers = 23
	KBluetoothCompanyIdentiferNielsenKellerman                     BluetoothCompanyIdentifers = 234
	KBluetoothCompanyIdentiferNike                                 BluetoothCompanyIdentifers = 120
	KBluetoothCompanyIdentiferNokiaMobilePhones                    BluetoothCompanyIdentifers = 1
	KBluetoothCompanyIdentiferNordicSemiconductor                  BluetoothCompanyIdentifers = 89
	KBluetoothCompanyIdentiferNorwoodSystems                       BluetoothCompanyIdentifers = 46
	KBluetoothCompanyIdentiferODMTechnology                        BluetoothCompanyIdentifers = 150
	KBluetoothCompanyIdentiferOTLDynamics                          BluetoothCompanyIdentifers = 165
	KBluetoothCompanyIdentiferOmegawave                            BluetoothCompanyIdentifers = 174
	KBluetoothCompanyIdentiferOnsetComputer                        BluetoothCompanyIdentifers = 197
	KBluetoothCompanyIdentiferOpenInterface                        BluetoothCompanyIdentifers = 39
	KBluetoothCompanyIdentiferPLUSLocationSystems                  BluetoothCompanyIdentifers = 260
	KBluetoothCompanyIdentiferPandaOcean                           BluetoothCompanyIdentifers = 166
	KBluetoothCompanyIdentiferParrotSA                             BluetoothCompanyIdentifers = 67
	KBluetoothCompanyIdentiferParthusTechnologies                  BluetoothCompanyIdentifers = 14
	KBluetoothCompanyIdentiferPassifSemiconductor                  BluetoothCompanyIdentifers = 176
	KBluetoothCompanyIdentiferPayPal                               BluetoothCompanyIdentifers = 240
	KBluetoothCompanyIdentiferPeterSystemtechnik                   BluetoothCompanyIdentifers = 173
	KBluetoothCompanyIdentiferPhilipsSemiconductor                 BluetoothCompanyIdentifers = 37
	KBluetoothCompanyIdentiferPlantronics                          BluetoothCompanyIdentifers = 85
	KBluetoothCompanyIdentiferPolarElectroEurope                   BluetoothCompanyIdentifers = 209
	KBluetoothCompanyIdentiferPolarElectroOY                       BluetoothCompanyIdentifers = 107
	KBluetoothCompanyIdentiferProctorAndGamble                     BluetoothCompanyIdentifers = 220
	KBluetoothCompanyIdentiferQualcomm                             BluetoothCompanyIdentifers = 29
	KBluetoothCompanyIdentiferQualcommConnectedExperiences         BluetoothCompanyIdentifers = 216
	KBluetoothCompanyIdentiferQualcommInnovationCenter             BluetoothCompanyIdentifers = 184
	KBluetoothCompanyIdentiferQualcommTechnologies                 BluetoothCompanyIdentifers = 215
	KBluetoothCompanyIdentiferQuintic                              BluetoothCompanyIdentifers = 142
	KBluetoothCompanyIdentiferQuupa                                BluetoothCompanyIdentifers = 199
	KBluetoothCompanyIdentiferRDAMicroelectronics                  BluetoothCompanyIdentifers = 97
	KBluetoothCompanyIdentiferRFCMicroDevices                      BluetoothCompanyIdentifers = 40
	KBluetoothCompanyIdentiferRTXTelecom                           BluetoothCompanyIdentifers = 21
	KBluetoothCompanyIdentiferRalinkTechnology                     BluetoothCompanyIdentifers = 91
	KBluetoothCompanyIdentiferRealtekSemiconductor                 BluetoothCompanyIdentifers = 93
	KBluetoothCompanyIdentiferRedMCommunications                   BluetoothCompanyIdentifers = 50
	KBluetoothCompanyIdentiferRenesasTechnology                    BluetoothCompanyIdentifers = 54
	KBluetoothCompanyIdentiferResearchInMotion                     BluetoothCompanyIdentifers = 60
	KBluetoothCompanyIdentiferRivieraWaves                         BluetoothCompanyIdentifers = 96
	KBluetoothCompanyIdentiferRohdeandSchwarz                      BluetoothCompanyIdentifers = 25
	KBluetoothCompanyIdentiferSPowerElectronics                    BluetoothCompanyIdentifers = 187
	KBluetoothCompanyIdentiferSRMedizinelektronik                  BluetoothCompanyIdentifers = 161
	KBluetoothCompanyIdentiferSTMicroelectronics                   BluetoothCompanyIdentifers = 48
	KBluetoothCompanyIdentiferSamsungElectronics                   BluetoothCompanyIdentifers = 117
	KBluetoothCompanyIdentiferSarisCyclingGroup                    BluetoothCompanyIdentifers = 177
	KBluetoothCompanyIdentiferSeersTechnology                      BluetoothCompanyIdentifers = 125
	KBluetoothCompanyIdentiferSeikoEpson                           BluetoothCompanyIdentifers = 64
	KBluetoothCompanyIdentiferSelflyBV                             BluetoothCompanyIdentifers = 198
	KBluetoothCompanyIdentiferSemilink                             BluetoothCompanyIdentifers = 226
	KBluetoothCompanyIdentiferSennheiserCommunications             BluetoothCompanyIdentifers = 130
	KBluetoothCompanyIdentiferServerTechnology                     BluetoothCompanyIdentifers = 235
	KBluetoothCompanyIdentiferShangHaiSuperSmartElectronics        BluetoothCompanyIdentifers = 114
	KBluetoothCompanyIdentiferShenzhenExcelsecuDataTechnology      BluetoothCompanyIdentifers = 193
	KBluetoothCompanyIdentiferSiRFTechnology                       BluetoothCompanyIdentifers = 80
	KBluetoothCompanyIdentiferSigniaTechnologies                   BluetoothCompanyIdentifers = 27
	KBluetoothCompanyIdentiferSiliconWave                          BluetoothCompanyIdentifers = 11
	KBluetoothCompanyIdentiferSmartifier                           BluetoothCompanyIdentifers = 245
	KBluetoothCompanyIdentiferSocketCommunications                 BluetoothCompanyIdentifers = 68
	KBluetoothCompanyIdentiferSonyEricssonMobileCommunications     BluetoothCompanyIdentifers = 86
	KBluetoothCompanyIdentiferSoundID                              BluetoothCompanyIdentifers = 111
	KBluetoothCompanyIdentiferSportsTrackingTechnologies           BluetoothCompanyIdentifers = 126
	KBluetoothCompanyIdentiferStaccatoCommunications               BluetoothCompanyIdentifers = 77
	KBluetoothCompanyIdentiferStalmartTechnology                   BluetoothCompanyIdentifers = 191
	KBluetoothCompanyIdentiferStanleyBlackAndDecker                BluetoothCompanyIdentifers = 254
	KBluetoothCompanyIdentiferStarkeyLaboratories                  BluetoothCompanyIdentifers = 186
	KBluetoothCompanyIdentiferStickNFind                           BluetoothCompanyIdentifers = 249
	KBluetoothCompanyIdentiferStonestreetOne                       BluetoothCompanyIdentifers = 94
	KBluetoothCompanyIdentiferSummitDataCommunications             BluetoothCompanyIdentifers = 110
	KBluetoothCompanyIdentiferSuuntoOy                             BluetoothCompanyIdentifers = 159
	KBluetoothCompanyIdentiferSwirlNetworks                        BluetoothCompanyIdentifers = 181
	KBluetoothCompanyIdentiferSymbolTechnologies                   BluetoothCompanyIdentifers = 42
	KBluetoothCompanyIdentiferSynopsys                             BluetoothCompanyIdentifers = 49
	KBluetoothCompanyIdentiferSystemsAndChips                      BluetoothCompanyIdentifers = 62
	KBluetoothCompanyIdentiferTTPCom                               BluetoothCompanyIdentifers = 26
	KBluetoothCompanyIdentiferTZeroTechnologies                    BluetoothCompanyIdentifers = 81
	KBluetoothCompanyIdentiferTaixingbangTechnology                BluetoothCompanyIdentifers = 211
	KBluetoothCompanyIdentiferTelitWirelessSolutions               BluetoothCompanyIdentifers = 143
	KBluetoothCompanyIdentiferTenovis                              BluetoothCompanyIdentifers = 43
	KBluetoothCompanyIdentiferTerax                                BluetoothCompanyIdentifers = 56
	KBluetoothCompanyIdentiferTexasInstruments                     BluetoothCompanyIdentifers = 13
	KBluetoothCompanyIdentiferThinkOptics                          BluetoothCompanyIdentifers = 146
	KBluetoothCompanyIdentiferTimeKeepingSystems                   BluetoothCompanyIdentifers = 131
	KBluetoothCompanyIdentiferTimexGroup                           BluetoothCompanyIdentifers = 214
	KBluetoothCompanyIdentiferTomTomInternational                  BluetoothCompanyIdentifers = 256
	KBluetoothCompanyIdentiferTopconPositioningSystems             BluetoothCompanyIdentifers = 139
	KBluetoothCompanyIdentiferToshiba                              BluetoothCompanyIdentifers = 4
	KBluetoothCompanyIdentiferTransilica                           BluetoothCompanyIdentifers = 24
	KBluetoothCompanyIdentiferTreLab                               BluetoothCompanyIdentifers = 183
	KBluetoothCompanyIdentiferTypeProducts                         BluetoothCompanyIdentifers = 255
	KBluetoothCompanyIdentiferUbiquitousComputingTechnology        BluetoothCompanyIdentifers = 261
	KBluetoothCompanyIdentiferUniversalElectriconics               BluetoothCompanyIdentifers = 147
	KBluetoothCompanyIdentiferVSNTechnologies                      BluetoothCompanyIdentifers = 247
	KBluetoothCompanyIdentiferValenceTech                          BluetoothCompanyIdentifers = 253
	KBluetoothCompanyIdentiferVertu                                BluetoothCompanyIdentifers = 162
	KBluetoothCompanyIdentiferVisio                                BluetoothCompanyIdentifers = 88
	KBluetoothCompanyIdentiferVisteon                              BluetoothCompanyIdentifers = 167
	KBluetoothCompanyIdentiferVoyetraTurtleBeach                   BluetoothCompanyIdentifers = 217
	KBluetoothCompanyIdentiferVtrackSystems                        BluetoothCompanyIdentifers = 233
	KBluetoothCompanyIdentiferWavePlusTechnology                   BluetoothCompanyIdentifers = 35
	KBluetoothCompanyIdentiferWicentric                            BluetoothCompanyIdentifers = 95
	KBluetoothCompanyIdentiferWidcomm                              BluetoothCompanyIdentifers = 17
	KBluetoothCompanyIdentiferWilliamDemantHolding                 BluetoothCompanyIdentifers = 263
	KBluetoothCompanyIdentiferWitronTechnology                     BluetoothCompanyIdentifers = 241
	KBluetoothCompanyIdentiferWuXiVimicro                          BluetoothCompanyIdentifers = 129
	KBluetoothCompanyIdentiferZeevo                                BluetoothCompanyIdentifers = 18
	KBluetoothCompanyIdentiferZero1TV                              BluetoothCompanyIdentifers = 152
	KBluetoothCompanyIdentiferZomm                                 BluetoothCompanyIdentifers = 116
	KBluetoothCompanyIdentiferZscanSoftware                        BluetoothCompanyIdentifers = 141
	KBluetoothCompanyIdentifertxtrGMBH                             BluetoothCompanyIdentifers = 218
)

func (e BluetoothCompanyIdentifers) String() string {
	switch e {
	case KBluetoothCompanyIdentifer3Com:
		return "KBluetoothCompanyIdentifer3Com"
	case KBluetoothCompanyIdentifer3DSP:
		return "KBluetoothCompanyIdentifer3DSP"
	case KBluetoothCompanyIdentifer3DiJoy:
		return "KBluetoothCompanyIdentifer3DiJoy"
	case KBluetoothCompanyIdentifer9SolutionsOy:
		return "KBluetoothCompanyIdentifer9SolutionsOy"
	case KBluetoothCompanyIdentiferAAMPofAmerica:
		return "KBluetoothCompanyIdentiferAAMPofAmerica"
	case KBluetoothCompanyIdentiferAAndDEngineering:
		return "KBluetoothCompanyIdentiferAAndDEngineering"
	case KBluetoothCompanyIdentiferAAndRCambridge:
		return "KBluetoothCompanyIdentiferAAndRCambridge"
	case KBluetoothCompanyIdentiferACTSTechnologies:
		return "KBluetoothCompanyIdentiferACTSTechnologies"
	case KBluetoothCompanyIdentiferAMICCOMElectronics:
		return "KBluetoothCompanyIdentiferAMICCOMElectronics"
	case KBluetoothCompanyIdentiferAPT:
		return "KBluetoothCompanyIdentiferAPT"
	case KBluetoothCompanyIdentiferARCHOS:
		return "KBluetoothCompanyIdentiferARCHOS"
	case KBluetoothCompanyIdentiferARPDevicesUnlimited:
		return "KBluetoothCompanyIdentiferARPDevicesUnlimited"
	case KBluetoothCompanyIdentiferAVMBerlin:
		return "KBluetoothCompanyIdentiferAVMBerlin"
	case KBluetoothCompanyIdentiferAboveAverageOutcomes:
		return "KBluetoothCompanyIdentiferAboveAverageOutcomes"
	case KBluetoothCompanyIdentiferAccelSemiconductor:
		return "KBluetoothCompanyIdentiferAccelSemiconductor"
	case KBluetoothCompanyIdentiferAceSensor:
		return "KBluetoothCompanyIdentiferAceSensor"
	case KBluetoothCompanyIdentiferAceUni:
		return "KBluetoothCompanyIdentiferAceUni"
	case KBluetoothCompanyIdentiferAdidas:
		return "KBluetoothCompanyIdentiferAdidas"
	case KBluetoothCompanyIdentiferAdvancedPANMOBILSystems:
		return "KBluetoothCompanyIdentiferAdvancedPANMOBILSystems"
	case KBluetoothCompanyIdentiferAirohaTechnology:
		return "KBluetoothCompanyIdentiferAirohaTechnology"
	case KBluetoothCompanyIdentiferAlcatel:
		return "KBluetoothCompanyIdentiferAlcatel"
	case KBluetoothCompanyIdentiferAlpwise:
		return "KBluetoothCompanyIdentiferAlpwise"
	case KBluetoothCompanyIdentiferAplix:
		return "KBluetoothCompanyIdentiferAplix"
	case KBluetoothCompanyIdentiferApple:
		return "KBluetoothCompanyIdentiferApple"
	case KBluetoothCompanyIdentiferAtherosCommunications:
		return "KBluetoothCompanyIdentiferAtherosCommunications"
	case KBluetoothCompanyIdentiferAtmel:
		return "KBluetoothCompanyIdentiferAtmel"
	case KBluetoothCompanyIdentiferAustcoCommunicationsSystems:
		return "KBluetoothCompanyIdentiferAustcoCommunicationsSystems"
	case KBluetoothCompanyIdentiferAutonetMobile:
		return "KBluetoothCompanyIdentiferAutonetMobile"
	case KBluetoothCompanyIdentiferAvagoTechnologies:
		return "KBluetoothCompanyIdentiferAvagoTechnologies"
	case KBluetoothCompanyIdentiferBDETechnology:
		return "KBluetoothCompanyIdentiferBDETechnology"
	case KBluetoothCompanyIdentiferBandXIInternational:
		return "KBluetoothCompanyIdentiferBandXIInternational"
	case KBluetoothCompanyIdentiferBandspeed:
		return "KBluetoothCompanyIdentiferBandspeed"
	case KBluetoothCompanyIdentiferBangAndOlufson:
		return "KBluetoothCompanyIdentiferBangAndOlufson"
	case KBluetoothCompanyIdentiferBeatsElectronics:
		return "KBluetoothCompanyIdentiferBeatsElectronics"
	case KBluetoothCompanyIdentiferBeautifulEnterprise:
		return "KBluetoothCompanyIdentiferBeautifulEnterprise"
	case KBluetoothCompanyIdentiferBekey:
		return "KBluetoothCompanyIdentiferBekey"
	case KBluetoothCompanyIdentiferBelkinInternational:
		return "KBluetoothCompanyIdentiferBelkinInternational"
	case KBluetoothCompanyIdentiferBinauricSE:
		return "KBluetoothCompanyIdentiferBinauricSE"
	case KBluetoothCompanyIdentiferBioResearchAssociates:
		return "KBluetoothCompanyIdentiferBioResearchAssociates"
	case KBluetoothCompanyIdentiferBiosentronics:
		return "KBluetoothCompanyIdentiferBiosentronics"
	case KBluetoothCompanyIdentiferBitsplitters:
		return "KBluetoothCompanyIdentiferBitsplitters"
	case KBluetoothCompanyIdentiferBlueRadios:
		return "KBluetoothCompanyIdentiferBlueRadios"
	case KBluetoothCompanyIdentiferBluegiga:
		return "KBluetoothCompanyIdentiferBluegiga"
	case KBluetoothCompanyIdentiferBluetoothSIG:
		return "KBluetoothCompanyIdentiferBluetoothSIG"
	case KBluetoothCompanyIdentiferBose:
		return "KBluetoothCompanyIdentiferBose"
	case KBluetoothCompanyIdentiferBriarTek:
		return "KBluetoothCompanyIdentiferBriarTek"
	case KBluetoothCompanyIdentiferBroadcom:
		return "KBluetoothCompanyIdentiferBroadcom"
	case KBluetoothCompanyIdentiferCATC:
		return "KBluetoothCompanyIdentiferCATC"
	case KBluetoothCompanyIdentiferCONWISETechnology:
		return "KBluetoothCompanyIdentiferCONWISETechnology"
	case KBluetoothCompanyIdentiferCTechnologies:
		return "KBluetoothCompanyIdentiferCTechnologies"
	case KBluetoothCompanyIdentiferCaenRFID:
		return "KBluetoothCompanyIdentiferCaenRFID"
	case KBluetoothCompanyIdentiferCambridgeSiliconRadio:
		return "KBluetoothCompanyIdentiferCambridgeSiliconRadio"
	case KBluetoothCompanyIdentiferCinetix:
		return "KBluetoothCompanyIdentiferCinetix"
	case KBluetoothCompanyIdentiferClarinoxTechnologies:
		return "KBluetoothCompanyIdentiferClarinoxTechnologies"
	case KBluetoothCompanyIdentiferColorfy:
		return "KBluetoothCompanyIdentiferColorfy"
	case KBluetoothCompanyIdentiferCommil:
		return "KBluetoothCompanyIdentiferCommil"
	case KBluetoothCompanyIdentiferConexantSystems:
		return "KBluetoothCompanyIdentiferConexantSystems"
	case KBluetoothCompanyIdentiferConnectBlueAB:
		return "KBluetoothCompanyIdentiferConnectBlueAB"
	case KBluetoothCompanyIdentiferConnecteDevice:
		return "KBluetoothCompanyIdentiferConnecteDevice"
	case KBluetoothCompanyIdentiferContinentialAutomotiveSystems:
		return "KBluetoothCompanyIdentiferContinentialAutomotiveSystems"
	case KBluetoothCompanyIdentiferCreativeTechnology:
		return "KBluetoothCompanyIdentiferCreativeTechnology"
	case KBluetoothCompanyIdentiferCrystalCode:
		return "KBluetoothCompanyIdentiferCrystalCode"
	case KBluetoothCompanyIdentiferDanlers:
		return "KBluetoothCompanyIdentiferDanlers"
	case KBluetoothCompanyIdentiferDeLormePublishingCompany:
		return "KBluetoothCompanyIdentiferDeLormePublishingCompany"
	case KBluetoothCompanyIdentiferDelphi:
		return "KBluetoothCompanyIdentiferDelphi"
	case KBluetoothCompanyIdentiferDexcom:
		return "KBluetoothCompanyIdentiferDexcom"
	case KBluetoothCompanyIdentiferDialogSemiconductor:
		return "KBluetoothCompanyIdentiferDialogSemiconductor"
	case KBluetoothCompanyIdentiferDigianswerAS:
		return "KBluetoothCompanyIdentiferDigianswerAS"
	case KBluetoothCompanyIdentiferEMMicroElectronicMarin:
		return "KBluetoothCompanyIdentiferEMMicroElectronicMarin"
	case KBluetoothCompanyIdentiferEclipse:
		return "KBluetoothCompanyIdentiferEclipse"
	case KBluetoothCompanyIdentiferEcotest:
		return "KBluetoothCompanyIdentiferEcotest"
	case KBluetoothCompanyIdentiferEdenSoftwareConsultants:
		return "KBluetoothCompanyIdentiferEdenSoftwareConsultants"
	case KBluetoothCompanyIdentiferElcometer:
		return "KBluetoothCompanyIdentiferElcometer"
	case KBluetoothCompanyIdentiferElgatoSystems:
		return "KBluetoothCompanyIdentiferElgatoSystems"
	case KBluetoothCompanyIdentiferEquinux:
		return "KBluetoothCompanyIdentiferEquinux"
	case KBluetoothCompanyIdentiferEricssonTechnologyLicensing:
		return "KBluetoothCompanyIdentiferEricssonTechnologyLicensing"
	case KBluetoothCompanyIdentiferEvluma:
		return "KBluetoothCompanyIdentiferEvluma"
	case KBluetoothCompanyIdentiferFree2Move:
		return "KBluetoothCompanyIdentiferFree2Move"
	case KBluetoothCompanyIdentiferFreshtemp:
		return "KBluetoothCompanyIdentiferFreshtemp"
	case KBluetoothCompanyIdentiferFuGoo:
		return "KBluetoothCompanyIdentiferFuGoo"
	case KBluetoothCompanyIdentiferFunaiElectric:
		return "KBluetoothCompanyIdentiferFunaiElectric"
	case KBluetoothCompanyIdentiferGCTSemiconductor:
		return "KBluetoothCompanyIdentiferGCTSemiconductor"
	case KBluetoothCompanyIdentiferGNNetcom:
		return "KBluetoothCompanyIdentiferGNNetcom"
	case KBluetoothCompanyIdentiferGNResound:
		return "KBluetoothCompanyIdentiferGNResound"
	case KBluetoothCompanyIdentiferGarminInternational:
		return "KBluetoothCompanyIdentiferGarminInternational"
	case KBluetoothCompanyIdentiferGeLo:
		return "KBluetoothCompanyIdentiferGeLo"
	case KBluetoothCompanyIdentiferGeneq:
		return "KBluetoothCompanyIdentiferGeneq"
	case KBluetoothCompanyIdentiferGeneralMotors:
		return "KBluetoothCompanyIdentiferGeneralMotors"
	case KBluetoothCompanyIdentiferGennum:
		return "KBluetoothCompanyIdentiferGennum"
	case KBluetoothCompanyIdentiferGeoforce:
		return "KBluetoothCompanyIdentiferGeoforce"
	case KBluetoothCompanyIdentiferGibsonGuitars:
		return "KBluetoothCompanyIdentiferGibsonGuitars"
	case KBluetoothCompanyIdentiferGimbal:
		return "KBluetoothCompanyIdentiferGimbal"
	case KBluetoothCompanyIdentiferGoogle:
		return "KBluetoothCompanyIdentiferGoogle"
	case KBluetoothCompanyIdentiferGreenThrottleGames:
		return "KBluetoothCompanyIdentiferGreenThrottleGames"
	case KBluetoothCompanyIdentiferGroupSense:
		return "KBluetoothCompanyIdentiferGroupSense"
	case KBluetoothCompanyIdentiferHanlynnTechnologies:
		return "KBluetoothCompanyIdentiferHanlynnTechnologies"
	case KBluetoothCompanyIdentiferHarmonInternational:
		return "KBluetoothCompanyIdentiferHarmonInternational"
	case KBluetoothCompanyIdentiferHewlettPackard:
		return "KBluetoothCompanyIdentiferHewlettPackard"
	case KBluetoothCompanyIdentiferHitachi:
		return "KBluetoothCompanyIdentiferHitachi"
	case KBluetoothCompanyIdentiferHosiden:
		return "KBluetoothCompanyIdentiferHosiden"
	case KBluetoothCompanyIdentiferIBM:
		return "KBluetoothCompanyIdentiferIBM"
	case KBluetoothCompanyIdentiferIPextreme:
		return "KBluetoothCompanyIdentiferIPextreme"
	case KBluetoothCompanyIdentiferITechDynamicGlobalDistribution:
		return "KBluetoothCompanyIdentiferITechDynamicGlobalDistribution"
	case KBluetoothCompanyIdentiferInMusicBrands:
		return "KBluetoothCompanyIdentiferInMusicBrands"
	case KBluetoothCompanyIdentiferInfineonTechnologiesAG:
		return "KBluetoothCompanyIdentiferInfineonTechnologiesAG"
	case KBluetoothCompanyIdentiferIngenieurSystemgruppeZahn:
		return "KBluetoothCompanyIdentiferIngenieurSystemgruppeZahn"
	case KBluetoothCompanyIdentiferInnovativeYachtterSolutions:
		return "KBluetoothCompanyIdentiferInnovativeYachtterSolutions"
	case KBluetoothCompanyIdentiferIntegratedSiliconSolution:
		return "KBluetoothCompanyIdentiferIntegratedSiliconSolution"
	case KBluetoothCompanyIdentiferIntegratedSystemSolution:
		return "KBluetoothCompanyIdentiferIntegratedSystemSolution"
	case KBluetoothCompanyIdentiferIntel:
		return "KBluetoothCompanyIdentiferIntel"
	case KBluetoothCompanyIdentiferInteropIdentifier:
		return "KBluetoothCompanyIdentiferInteropIdentifier"
	case KBluetoothCompanyIdentiferInventel:
		return "KBluetoothCompanyIdentiferInventel"
	case KBluetoothCompanyIdentiferJandM:
		return "KBluetoothCompanyIdentiferJandM"
	case KBluetoothCompanyIdentiferJawbone:
		return "KBluetoothCompanyIdentiferJawbone"
	case KBluetoothCompanyIdentiferJiangsuToppowerAutomotiveElectronics:
		return "KBluetoothCompanyIdentiferJiangsuToppowerAutomotiveElectronics"
	case KBluetoothCompanyIdentiferJohnsonControls:
		return "KBluetoothCompanyIdentiferJohnsonControls"
	case KBluetoothCompanyIdentiferJollyLogic:
		return "KBluetoothCompanyIdentiferJollyLogic"
	case KBluetoothCompanyIdentiferKCTechnology:
		return "KBluetoothCompanyIdentiferKCTechnology"
	case KBluetoothCompanyIdentiferKOUKAMM:
		return "KBluetoothCompanyIdentiferKOUKAMM"
	case KBluetoothCompanyIdentiferKSTechnologies:
		return "KBluetoothCompanyIdentiferKSTechnologies"
	case KBluetoothCompanyIdentiferKawantech:
		return "KBluetoothCompanyIdentiferKawantech"
	case KBluetoothCompanyIdentiferKeiser:
		return "KBluetoothCompanyIdentiferKeiser"
	case KBluetoothCompanyIdentiferKensingtonComputerProductsGroup:
		return "KBluetoothCompanyIdentiferKensingtonComputerProductsGroup"
	case KBluetoothCompanyIdentiferKentDisplays:
		return "KBluetoothCompanyIdentiferKentDisplays"
	case KBluetoothCompanyIdentiferLGElectronics:
		return "KBluetoothCompanyIdentiferLGElectronics"
	case KBluetoothCompanyIdentiferLSResearch:
		return "KBluetoothCompanyIdentiferLSResearch"
	case KBluetoothCompanyIdentiferLairdTechnologies:
		return "KBluetoothCompanyIdentiferLairdTechnologies"
	case KBluetoothCompanyIdentiferLessWire:
		return "KBluetoothCompanyIdentiferLessWire"
	case KBluetoothCompanyIdentiferLinak:
		return "KBluetoothCompanyIdentiferLinak"
	case KBluetoothCompanyIdentiferLucent:
		return "KBluetoothCompanyIdentiferLucent"
	case KBluetoothCompanyIdentiferLudusHelsinki:
		return "KBluetoothCompanyIdentiferLudusHelsinki"
	case KBluetoothCompanyIdentiferMC10:
		return "KBluetoothCompanyIdentiferMC10"
	case KBluetoothCompanyIdentiferMStarTechnologies:
		return "KBluetoothCompanyIdentiferMStarTechnologies"
	case KBluetoothCompanyIdentiferMacronixInternational:
		return "KBluetoothCompanyIdentiferMacronixInternational"
	case KBluetoothCompanyIdentiferMagnetiMarelli:
		return "KBluetoothCompanyIdentiferMagnetiMarelli"
	case KBluetoothCompanyIdentiferMansella:
		return "KBluetoothCompanyIdentiferMansella"
	case KBluetoothCompanyIdentiferMarvellTechnologyGroup:
		return "KBluetoothCompanyIdentiferMarvellTechnologyGroup"
	case KBluetoothCompanyIdentiferMatsushitaElectricIndustrial:
		return "KBluetoothCompanyIdentiferMatsushitaElectricIndustrial"
	case KBluetoothCompanyIdentiferMediaTek:
		return "KBluetoothCompanyIdentiferMediaTek"
	case KBluetoothCompanyIdentiferMesoInternational:
		return "KBluetoothCompanyIdentiferMesoInternational"
	case KBluetoothCompanyIdentiferMetaWatch:
		return "KBluetoothCompanyIdentiferMetaWatch"
	case KBluetoothCompanyIdentiferMewTelTechnology:
		return "KBluetoothCompanyIdentiferMewTelTechnology"
	case KBluetoothCompanyIdentiferMiCommand:
		return "KBluetoothCompanyIdentiferMiCommand"
	case KBluetoothCompanyIdentiferMicrochipTechnology:
		return "KBluetoothCompanyIdentiferMicrochipTechnology"
	case KBluetoothCompanyIdentiferMicrosoft:
		return "KBluetoothCompanyIdentiferMicrosoft"
	case KBluetoothCompanyIdentiferMindTree:
		return "KBluetoothCompanyIdentiferMindTree"
	case KBluetoothCompanyIdentiferMisfitWearables:
		return "KBluetoothCompanyIdentiferMisfitWearables"
	case KBluetoothCompanyIdentiferMistubishiElectric:
		return "KBluetoothCompanyIdentiferMistubishiElectric"
	case KBluetoothCompanyIdentiferMitelSemiconductor:
		return "KBluetoothCompanyIdentiferMitelSemiconductor"
	case KBluetoothCompanyIdentiferMobilian:
		return "KBluetoothCompanyIdentiferMobilian"
	case KBluetoothCompanyIdentiferMonster:
		return "KBluetoothCompanyIdentiferMonster"
	case KBluetoothCompanyIdentiferMorseProject:
		return "KBluetoothCompanyIdentiferMorseProject"
	case KBluetoothCompanyIdentiferMotorola:
		return "KBluetoothCompanyIdentiferMotorola"
	case KBluetoothCompanyIdentiferMusik:
		return "KBluetoothCompanyIdentiferMusik"
	case KBluetoothCompanyIdentiferNEC:
		return "KBluetoothCompanyIdentiferNEC"
	case KBluetoothCompanyIdentiferNECLightning:
		return "KBluetoothCompanyIdentiferNECLightning"
	case KBluetoothCompanyIdentiferNautilus:
		return "KBluetoothCompanyIdentiferNautilus"
	case KBluetoothCompanyIdentiferNewlogic:
		return "KBluetoothCompanyIdentiferNewlogic"
	case KBluetoothCompanyIdentiferNielsenKellerman:
		return "KBluetoothCompanyIdentiferNielsenKellerman"
	case KBluetoothCompanyIdentiferNike:
		return "KBluetoothCompanyIdentiferNike"
	case KBluetoothCompanyIdentiferNokiaMobilePhones:
		return "KBluetoothCompanyIdentiferNokiaMobilePhones"
	case KBluetoothCompanyIdentiferNordicSemiconductor:
		return "KBluetoothCompanyIdentiferNordicSemiconductor"
	case KBluetoothCompanyIdentiferNorwoodSystems:
		return "KBluetoothCompanyIdentiferNorwoodSystems"
	case KBluetoothCompanyIdentiferODMTechnology:
		return "KBluetoothCompanyIdentiferODMTechnology"
	case KBluetoothCompanyIdentiferOTLDynamics:
		return "KBluetoothCompanyIdentiferOTLDynamics"
	case KBluetoothCompanyIdentiferOmegawave:
		return "KBluetoothCompanyIdentiferOmegawave"
	case KBluetoothCompanyIdentiferOnsetComputer:
		return "KBluetoothCompanyIdentiferOnsetComputer"
	case KBluetoothCompanyIdentiferOpenInterface:
		return "KBluetoothCompanyIdentiferOpenInterface"
	case KBluetoothCompanyIdentiferPLUSLocationSystems:
		return "KBluetoothCompanyIdentiferPLUSLocationSystems"
	case KBluetoothCompanyIdentiferPandaOcean:
		return "KBluetoothCompanyIdentiferPandaOcean"
	case KBluetoothCompanyIdentiferParrotSA:
		return "KBluetoothCompanyIdentiferParrotSA"
	case KBluetoothCompanyIdentiferParthusTechnologies:
		return "KBluetoothCompanyIdentiferParthusTechnologies"
	case KBluetoothCompanyIdentiferPassifSemiconductor:
		return "KBluetoothCompanyIdentiferPassifSemiconductor"
	case KBluetoothCompanyIdentiferPayPal:
		return "KBluetoothCompanyIdentiferPayPal"
	case KBluetoothCompanyIdentiferPeterSystemtechnik:
		return "KBluetoothCompanyIdentiferPeterSystemtechnik"
	case KBluetoothCompanyIdentiferPhilipsSemiconductor:
		return "KBluetoothCompanyIdentiferPhilipsSemiconductor"
	case KBluetoothCompanyIdentiferPlantronics:
		return "KBluetoothCompanyIdentiferPlantronics"
	case KBluetoothCompanyIdentiferPolarElectroEurope:
		return "KBluetoothCompanyIdentiferPolarElectroEurope"
	case KBluetoothCompanyIdentiferPolarElectroOY:
		return "KBluetoothCompanyIdentiferPolarElectroOY"
	case KBluetoothCompanyIdentiferProctorAndGamble:
		return "KBluetoothCompanyIdentiferProctorAndGamble"
	case KBluetoothCompanyIdentiferQualcomm:
		return "KBluetoothCompanyIdentiferQualcomm"
	case KBluetoothCompanyIdentiferQualcommConnectedExperiences:
		return "KBluetoothCompanyIdentiferQualcommConnectedExperiences"
	case KBluetoothCompanyIdentiferQualcommInnovationCenter:
		return "KBluetoothCompanyIdentiferQualcommInnovationCenter"
	case KBluetoothCompanyIdentiferQualcommTechnologies:
		return "KBluetoothCompanyIdentiferQualcommTechnologies"
	case KBluetoothCompanyIdentiferQuintic:
		return "KBluetoothCompanyIdentiferQuintic"
	case KBluetoothCompanyIdentiferQuupa:
		return "KBluetoothCompanyIdentiferQuupa"
	case KBluetoothCompanyIdentiferRDAMicroelectronics:
		return "KBluetoothCompanyIdentiferRDAMicroelectronics"
	case KBluetoothCompanyIdentiferRFCMicroDevices:
		return "KBluetoothCompanyIdentiferRFCMicroDevices"
	case KBluetoothCompanyIdentiferRTXTelecom:
		return "KBluetoothCompanyIdentiferRTXTelecom"
	case KBluetoothCompanyIdentiferRalinkTechnology:
		return "KBluetoothCompanyIdentiferRalinkTechnology"
	case KBluetoothCompanyIdentiferRealtekSemiconductor:
		return "KBluetoothCompanyIdentiferRealtekSemiconductor"
	case KBluetoothCompanyIdentiferRedMCommunications:
		return "KBluetoothCompanyIdentiferRedMCommunications"
	case KBluetoothCompanyIdentiferRenesasTechnology:
		return "KBluetoothCompanyIdentiferRenesasTechnology"
	case KBluetoothCompanyIdentiferResearchInMotion:
		return "KBluetoothCompanyIdentiferResearchInMotion"
	case KBluetoothCompanyIdentiferRivieraWaves:
		return "KBluetoothCompanyIdentiferRivieraWaves"
	case KBluetoothCompanyIdentiferRohdeandSchwarz:
		return "KBluetoothCompanyIdentiferRohdeandSchwarz"
	case KBluetoothCompanyIdentiferSPowerElectronics:
		return "KBluetoothCompanyIdentiferSPowerElectronics"
	case KBluetoothCompanyIdentiferSRMedizinelektronik:
		return "KBluetoothCompanyIdentiferSRMedizinelektronik"
	case KBluetoothCompanyIdentiferSTMicroelectronics:
		return "KBluetoothCompanyIdentiferSTMicroelectronics"
	case KBluetoothCompanyIdentiferSamsungElectronics:
		return "KBluetoothCompanyIdentiferSamsungElectronics"
	case KBluetoothCompanyIdentiferSarisCyclingGroup:
		return "KBluetoothCompanyIdentiferSarisCyclingGroup"
	case KBluetoothCompanyIdentiferSeersTechnology:
		return "KBluetoothCompanyIdentiferSeersTechnology"
	case KBluetoothCompanyIdentiferSeikoEpson:
		return "KBluetoothCompanyIdentiferSeikoEpson"
	case KBluetoothCompanyIdentiferSelflyBV:
		return "KBluetoothCompanyIdentiferSelflyBV"
	case KBluetoothCompanyIdentiferSemilink:
		return "KBluetoothCompanyIdentiferSemilink"
	case KBluetoothCompanyIdentiferSennheiserCommunications:
		return "KBluetoothCompanyIdentiferSennheiserCommunications"
	case KBluetoothCompanyIdentiferServerTechnology:
		return "KBluetoothCompanyIdentiferServerTechnology"
	case KBluetoothCompanyIdentiferShangHaiSuperSmartElectronics:
		return "KBluetoothCompanyIdentiferShangHaiSuperSmartElectronics"
	case KBluetoothCompanyIdentiferShenzhenExcelsecuDataTechnology:
		return "KBluetoothCompanyIdentiferShenzhenExcelsecuDataTechnology"
	case KBluetoothCompanyIdentiferSiRFTechnology:
		return "KBluetoothCompanyIdentiferSiRFTechnology"
	case KBluetoothCompanyIdentiferSigniaTechnologies:
		return "KBluetoothCompanyIdentiferSigniaTechnologies"
	case KBluetoothCompanyIdentiferSiliconWave:
		return "KBluetoothCompanyIdentiferSiliconWave"
	case KBluetoothCompanyIdentiferSmartifier:
		return "KBluetoothCompanyIdentiferSmartifier"
	case KBluetoothCompanyIdentiferSocketCommunications:
		return "KBluetoothCompanyIdentiferSocketCommunications"
	case KBluetoothCompanyIdentiferSonyEricssonMobileCommunications:
		return "KBluetoothCompanyIdentiferSonyEricssonMobileCommunications"
	case KBluetoothCompanyIdentiferSoundID:
		return "KBluetoothCompanyIdentiferSoundID"
	case KBluetoothCompanyIdentiferSportsTrackingTechnologies:
		return "KBluetoothCompanyIdentiferSportsTrackingTechnologies"
	case KBluetoothCompanyIdentiferStaccatoCommunications:
		return "KBluetoothCompanyIdentiferStaccatoCommunications"
	case KBluetoothCompanyIdentiferStalmartTechnology:
		return "KBluetoothCompanyIdentiferStalmartTechnology"
	case KBluetoothCompanyIdentiferStanleyBlackAndDecker:
		return "KBluetoothCompanyIdentiferStanleyBlackAndDecker"
	case KBluetoothCompanyIdentiferStarkeyLaboratories:
		return "KBluetoothCompanyIdentiferStarkeyLaboratories"
	case KBluetoothCompanyIdentiferStickNFind:
		return "KBluetoothCompanyIdentiferStickNFind"
	case KBluetoothCompanyIdentiferStonestreetOne:
		return "KBluetoothCompanyIdentiferStonestreetOne"
	case KBluetoothCompanyIdentiferSummitDataCommunications:
		return "KBluetoothCompanyIdentiferSummitDataCommunications"
	case KBluetoothCompanyIdentiferSuuntoOy:
		return "KBluetoothCompanyIdentiferSuuntoOy"
	case KBluetoothCompanyIdentiferSwirlNetworks:
		return "KBluetoothCompanyIdentiferSwirlNetworks"
	case KBluetoothCompanyIdentiferSymbolTechnologies:
		return "KBluetoothCompanyIdentiferSymbolTechnologies"
	case KBluetoothCompanyIdentiferSynopsys:
		return "KBluetoothCompanyIdentiferSynopsys"
	case KBluetoothCompanyIdentiferSystemsAndChips:
		return "KBluetoothCompanyIdentiferSystemsAndChips"
	case KBluetoothCompanyIdentiferTTPCom:
		return "KBluetoothCompanyIdentiferTTPCom"
	case KBluetoothCompanyIdentiferTZeroTechnologies:
		return "KBluetoothCompanyIdentiferTZeroTechnologies"
	case KBluetoothCompanyIdentiferTaixingbangTechnology:
		return "KBluetoothCompanyIdentiferTaixingbangTechnology"
	case KBluetoothCompanyIdentiferTelitWirelessSolutions:
		return "KBluetoothCompanyIdentiferTelitWirelessSolutions"
	case KBluetoothCompanyIdentiferTenovis:
		return "KBluetoothCompanyIdentiferTenovis"
	case KBluetoothCompanyIdentiferTerax:
		return "KBluetoothCompanyIdentiferTerax"
	case KBluetoothCompanyIdentiferTexasInstruments:
		return "KBluetoothCompanyIdentiferTexasInstruments"
	case KBluetoothCompanyIdentiferThinkOptics:
		return "KBluetoothCompanyIdentiferThinkOptics"
	case KBluetoothCompanyIdentiferTimeKeepingSystems:
		return "KBluetoothCompanyIdentiferTimeKeepingSystems"
	case KBluetoothCompanyIdentiferTimexGroup:
		return "KBluetoothCompanyIdentiferTimexGroup"
	case KBluetoothCompanyIdentiferTomTomInternational:
		return "KBluetoothCompanyIdentiferTomTomInternational"
	case KBluetoothCompanyIdentiferTopconPositioningSystems:
		return "KBluetoothCompanyIdentiferTopconPositioningSystems"
	case KBluetoothCompanyIdentiferToshiba:
		return "KBluetoothCompanyIdentiferToshiba"
	case KBluetoothCompanyIdentiferTransilica:
		return "KBluetoothCompanyIdentiferTransilica"
	case KBluetoothCompanyIdentiferTreLab:
		return "KBluetoothCompanyIdentiferTreLab"
	case KBluetoothCompanyIdentiferTypeProducts:
		return "KBluetoothCompanyIdentiferTypeProducts"
	case KBluetoothCompanyIdentiferUbiquitousComputingTechnology:
		return "KBluetoothCompanyIdentiferUbiquitousComputingTechnology"
	case KBluetoothCompanyIdentiferUniversalElectriconics:
		return "KBluetoothCompanyIdentiferUniversalElectriconics"
	case KBluetoothCompanyIdentiferVSNTechnologies:
		return "KBluetoothCompanyIdentiferVSNTechnologies"
	case KBluetoothCompanyIdentiferValenceTech:
		return "KBluetoothCompanyIdentiferValenceTech"
	case KBluetoothCompanyIdentiferVertu:
		return "KBluetoothCompanyIdentiferVertu"
	case KBluetoothCompanyIdentiferVisio:
		return "KBluetoothCompanyIdentiferVisio"
	case KBluetoothCompanyIdentiferVisteon:
		return "KBluetoothCompanyIdentiferVisteon"
	case KBluetoothCompanyIdentiferVoyetraTurtleBeach:
		return "KBluetoothCompanyIdentiferVoyetraTurtleBeach"
	case KBluetoothCompanyIdentiferVtrackSystems:
		return "KBluetoothCompanyIdentiferVtrackSystems"
	case KBluetoothCompanyIdentiferWavePlusTechnology:
		return "KBluetoothCompanyIdentiferWavePlusTechnology"
	case KBluetoothCompanyIdentiferWicentric:
		return "KBluetoothCompanyIdentiferWicentric"
	case KBluetoothCompanyIdentiferWidcomm:
		return "KBluetoothCompanyIdentiferWidcomm"
	case KBluetoothCompanyIdentiferWilliamDemantHolding:
		return "KBluetoothCompanyIdentiferWilliamDemantHolding"
	case KBluetoothCompanyIdentiferWitronTechnology:
		return "KBluetoothCompanyIdentiferWitronTechnology"
	case KBluetoothCompanyIdentiferWuXiVimicro:
		return "KBluetoothCompanyIdentiferWuXiVimicro"
	case KBluetoothCompanyIdentiferZeevo:
		return "KBluetoothCompanyIdentiferZeevo"
	case KBluetoothCompanyIdentiferZero1TV:
		return "KBluetoothCompanyIdentiferZero1TV"
	case KBluetoothCompanyIdentiferZomm:
		return "KBluetoothCompanyIdentiferZomm"
	case KBluetoothCompanyIdentiferZscanSoftware:
		return "KBluetoothCompanyIdentiferZscanSoftware"
	case KBluetoothCompanyIdentifertxtrGMBH:
		return "KBluetoothCompanyIdentifertxtrGMBH"
	default:
		return fmt.Sprintf("BluetoothCompanyIdentifers(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothFeatureBits
type BluetoothFeatureBits uint32

const (
	KBluetoothExtendedFeatureLEAndBREDRToSameDeviceHostMode     BluetoothFeatureBits = 4
	KBluetoothExtendedFeatureLESupportedHostMode                BluetoothFeatureBits = 2
	KBluetoothExtendedFeaturePing                               BluetoothFeatureBits = 2
	KBluetoothExtendedFeatureReserved                           BluetoothFeatureBits = 4
	KBluetoothExtendedFeatureSecureConnectionsControllerSupport BluetoothFeatureBits = 1
	KBluetoothExtendedFeatureSecureConnectionsHostMode          BluetoothFeatureBits = 8
	KBluetoothExtendedFeatureSimpleSecurePairingHostMode        BluetoothFeatureBits = 1
	KBluetoothExtendedFeatureSlotAvailabilityMask               BluetoothFeatureBits = 16
	KBluetoothExtendedFeatureTrainNudging                       BluetoothFeatureBits = 8
	KBluetoothFeature3SlotEnhancedDataRateACLPackets            BluetoothFeatureBits = 128
	KBluetoothFeature3SlotEnhancedDataRateeSCOPackets           BluetoothFeatureBits = 128
	KBluetoothFeature5SlotEnhancedDataRateACLPackets            BluetoothFeatureBits = 1
	KBluetoothFeatureAFHCapableMaster                           BluetoothFeatureBits = 8
	KBluetoothFeatureAFHCapablePeripheral                       BluetoothFeatureBits = 8
	KBluetoothFeatureAFHCapableSlave                            BluetoothFeatureBits = 8
	KBluetoothFeatureAFHClassificationMaster                    BluetoothFeatureBits = 16
	KBluetoothFeatureAFHClassificationPeripheral                BluetoothFeatureBits = 16
	KBluetoothFeatureAFHClassificationSlave                     BluetoothFeatureBits = 16
	KBluetoothFeatureALawLog                                    BluetoothFeatureBits = 128
	KBluetoothFeatureAbsenceMasks                               BluetoothFeatureBits = 4
	KBluetoothFeatureAliasAuhentication                         BluetoothFeatureBits = 32
	KBluetoothFeatureBroadcastEncryption                        BluetoothFeatureBits = 128
	KBluetoothFeatureCVSD                                       BluetoothFeatureBits = 1
	KBluetoothFeatureChannelQuality                             BluetoothFeatureBits = 4
	KBluetoothFeatureEV4Packets                                 BluetoothFeatureBits = 1
	KBluetoothFeatureEV5Packets                                 BluetoothFeatureBits = 2
	KBluetoothFeatureEncapsulatedPDU                            BluetoothFeatureBits = 16
	KBluetoothFeatureEncryption                                 BluetoothFeatureBits = 4
	KBluetoothFeatureEnhancedDataRateACL2MbpsMode               BluetoothFeatureBits = 2
	KBluetoothFeatureEnhancedDataRateACL3MbpsMode               BluetoothFeatureBits = 4
	KBluetoothFeatureEnhancedDataRateeSCO2MbpsMode              BluetoothFeatureBits = 32
	KBluetoothFeatureEnhancedDataRateeSCO3MbpsMode              BluetoothFeatureBits = 64
	KBluetoothFeatureEnhancedInquiryScan                        BluetoothFeatureBits = 8
	KBluetoothFeatureErroneousDataReporting                     BluetoothFeatureBits = 32
	KBluetoothFeatureExtendedFeatures                           BluetoothFeatureBits = 128
	KBluetoothFeatureExtendedInquiryResponse                    BluetoothFeatureBits = 1
	KBluetoothFeatureExtendedSCOLink                            BluetoothFeatureBits = 128
	KBluetoothFeatureFiveSlotPackets                            BluetoothFeatureBits = 2
	KBluetoothFeatureFlowControlLagBit0                         BluetoothFeatureBits = 16
	KBluetoothFeatureFlowControlLagBit1                         BluetoothFeatureBits = 32
	KBluetoothFeatureFlowControlLagBit2                         BluetoothFeatureBits = 64
	KBluetoothFeatureHV2Packets                                 BluetoothFeatureBits = 16
	KBluetoothFeatureHV3Packets                                 BluetoothFeatureBits = 32
	KBluetoothFeatureHoldMode                                   BluetoothFeatureBits = 64
	KBluetoothFeatureInquiryTransmissionPowerLevel              BluetoothFeatureBits = 2
	KBluetoothFeatureInterlacedInquiryScan                      BluetoothFeatureBits = 16
	KBluetoothFeatureInterlacedPageScan                         BluetoothFeatureBits = 32
	KBluetoothFeatureLESupportedController                      BluetoothFeatureBits = 64
	KBluetoothFeatureLinkSupervisionTimeoutChangedEvent         BluetoothFeatureBits = 1
	KBluetoothFeatureNonFlushablePacketBoundaryFlag             BluetoothFeatureBits = 64
	KBluetoothFeaturePagingScheme                               BluetoothFeatureBits = 2
	KBluetoothFeatureParkMode                                   BluetoothFeatureBits = 1
	KBluetoothFeaturePauseEncryption                            BluetoothFeatureBits = 4
	KBluetoothFeaturePowerControl                               BluetoothFeatureBits = 4
	KBluetoothFeaturePowerControlRequests                       BluetoothFeatureBits = 2
	KBluetoothFeatureRSSI                                       BluetoothFeatureBits = 2
	KBluetoothFeatureRSSIWithInquiryResult                      BluetoothFeatureBits = 64
	KBluetoothFeatureSCOLink                                    BluetoothFeatureBits = 8
	KBluetoothFeatureScatterMode                                BluetoothFeatureBits = 1
	KBluetoothFeatureSecureSimplePairing                        BluetoothFeatureBits = 8
	KBluetoothFeatureSlotOffset                                 BluetoothFeatureBits = 8
	KBluetoothFeatureSniffMode                                  BluetoothFeatureBits = 128
	KBluetoothFeatureSniffSubrating                             BluetoothFeatureBits = 2
	KBluetoothFeatureSwitchRoles                                BluetoothFeatureBits = 32
	KBluetoothFeatureThreeSlotPackets                           BluetoothFeatureBits = 1
	KBluetoothFeatureTimingAccuracy                             BluetoothFeatureBits = 16
	KBluetoothFeatureTransparentSCOData                         BluetoothFeatureBits = 8
	KBluetoothFeatureULawLog                                    BluetoothFeatureBits = 64
)

func (e BluetoothFeatureBits) String() string {
	switch e {
	case KBluetoothExtendedFeatureLEAndBREDRToSameDeviceHostMode:
		return "KBluetoothExtendedFeatureLEAndBREDRToSameDeviceHostMode"
	case KBluetoothExtendedFeatureLESupportedHostMode:
		return "KBluetoothExtendedFeatureLESupportedHostMode"
	case KBluetoothExtendedFeatureSecureConnectionsControllerSupport:
		return "KBluetoothExtendedFeatureSecureConnectionsControllerSupport"
	case KBluetoothExtendedFeatureSecureConnectionsHostMode:
		return "KBluetoothExtendedFeatureSecureConnectionsHostMode"
	case KBluetoothExtendedFeatureSlotAvailabilityMask:
		return "KBluetoothExtendedFeatureSlotAvailabilityMask"
	case KBluetoothFeature3SlotEnhancedDataRateACLPackets:
		return "KBluetoothFeature3SlotEnhancedDataRateACLPackets"
	case KBluetoothFeatureAliasAuhentication:
		return "KBluetoothFeatureAliasAuhentication"
	case KBluetoothFeatureEnhancedDataRateeSCO3MbpsMode:
		return "KBluetoothFeatureEnhancedDataRateeSCO3MbpsMode"
	default:
		return fmt.Sprintf("BluetoothFeatureBits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAFHChannelAssessmentModes
type BluetoothHCIAFHChannelAssessmentModes uint32

const (
	KAFHChannelAssessmentModeDisabled BluetoothHCIAFHChannelAssessmentModes = 0
	KAFHChannelAssessmentModeEnabled  BluetoothHCIAFHChannelAssessmentModes = 0x1
)

func (e BluetoothHCIAFHChannelAssessmentModes) String() string {
	switch e {
	case KAFHChannelAssessmentModeDisabled:
		return "KAFHChannelAssessmentModeDisabled"
	case KAFHChannelAssessmentModeEnabled:
		return "KAFHChannelAssessmentModeEnabled"
	default:
		return fmt.Sprintf("BluetoothHCIAFHChannelAssessmentModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAuthentionEnableModes
type BluetoothHCIAuthentionEnableModes uint32

const (
	KAuthenticationDisabled BluetoothHCIAuthentionEnableModes = 0
	KAuthenticationEnabled  BluetoothHCIAuthentionEnableModes = 0x1
)

func (e BluetoothHCIAuthentionEnableModes) String() string {
	switch e {
	case KAuthenticationDisabled:
		return "KAuthenticationDisabled"
	case KAuthenticationEnabled:
		return "KAuthenticationEnabled"
	default:
		return fmt.Sprintf("BluetoothHCIAuthentionEnableModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIConnectionModes
type BluetoothHCIConnectionModes uint32

const (
	KConnectionActiveMode               BluetoothHCIConnectionModes = 0
	KConnectionHoldMode                 BluetoothHCIConnectionModes = 1
	KConnectionModeReservedForFutureUse BluetoothHCIConnectionModes = 4
	KConnectionParkMode                 BluetoothHCIConnectionModes = 3
	KConnectionSniffMode                BluetoothHCIConnectionModes = 2
)

func (e BluetoothHCIConnectionModes) String() string {
	switch e {
	case KConnectionActiveMode:
		return "KConnectionActiveMode"
	case KConnectionHoldMode:
		return "KConnectionHoldMode"
	case KConnectionModeReservedForFutureUse:
		return "KConnectionModeReservedForFutureUse"
	case KConnectionParkMode:
		return "KConnectionParkMode"
	case KConnectionSniffMode:
		return "KConnectionSniffMode"
	default:
		return fmt.Sprintf("BluetoothHCIConnectionModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIDeleteStoredLinkKeyFlags
type BluetoothHCIDeleteStoredLinkKeyFlags uint32

const (
	KDeleteAllStoredLinkKeys         BluetoothHCIDeleteStoredLinkKeyFlags = 0x1
	KDeleteKeyForSpecifiedDeviceOnly BluetoothHCIDeleteStoredLinkKeyFlags = 0
)

func (e BluetoothHCIDeleteStoredLinkKeyFlags) String() string {
	switch e {
	case KDeleteAllStoredLinkKeys:
		return "KDeleteAllStoredLinkKeys"
	case KDeleteKeyForSpecifiedDeviceOnly:
		return "KDeleteKeyForSpecifiedDeviceOnly"
	default:
		return fmt.Sprintf("BluetoothHCIDeleteStoredLinkKeyFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEncryptionModes
type BluetoothHCIEncryptionModes uint32

const (
	KEncryptionDisabled                               BluetoothHCIEncryptionModes = 0
	KEncryptionForBothPointToPointAndBroadcastPackets BluetoothHCIEncryptionModes = 0x2
	KEncryptionOnlyForPointToPointPackets             BluetoothHCIEncryptionModes = 0x1
)

func (e BluetoothHCIEncryptionModes) String() string {
	switch e {
	case KEncryptionDisabled:
		return "KEncryptionDisabled"
	case KEncryptionForBothPointToPointAndBroadcastPackets:
		return "KEncryptionForBothPointToPointAndBroadcastPackets"
	case KEncryptionOnlyForPointToPointPackets:
		return "KEncryptionOnlyForPointToPointPackets"
	default:
		return fmt.Sprintf("BluetoothHCIEncryptionModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedInquiryResponseDataTypes
type BluetoothHCIExtendedInquiryResponseDataTypes uint32

const (
	KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsCompleteList      BluetoothHCIExtendedInquiryResponseDataTypes = 0x7
	KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsWithMoreAvailable BluetoothHCIExtendedInquiryResponseDataTypes = 0x6
	KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsCompleteList       BluetoothHCIExtendedInquiryResponseDataTypes = 0x3
	KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsWithMoreAvailable  BluetoothHCIExtendedInquiryResponseDataTypes = 0x2
	KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsCompleteList       BluetoothHCIExtendedInquiryResponseDataTypes = 0x5
	KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsWithMoreAvailable  BluetoothHCIExtendedInquiryResponseDataTypes = 0x4
	KBluetoothHCIExtendedInquiryResponseDataType3DInformationData                        BluetoothHCIExtendedInquiryResponseDataTypes = 0x3d
	KBluetoothHCIExtendedInquiryResponseDataTypeAdvertisingInterval                      BluetoothHCIExtendedInquiryResponseDataTypes = 0x1a
	KBluetoothHCIExtendedInquiryResponseDataTypeAppearance                               BluetoothHCIExtendedInquiryResponseDataTypes = 0x19
	KBluetoothHCIExtendedInquiryResponseDataTypeCompleteLocalName                        BluetoothHCIExtendedInquiryResponseDataTypes = 0x9
	KBluetoothHCIExtendedInquiryResponseDataTypeCsisRsiData                              BluetoothHCIExtendedInquiryResponseDataTypes = 0x2e
	KBluetoothHCIExtendedInquiryResponseDataTypeDeviceID                                 BluetoothHCIExtendedInquiryResponseDataTypes = 0x10
	KBluetoothHCIExtendedInquiryResponseDataTypeFlags                                    BluetoothHCIExtendedInquiryResponseDataTypes = 0x1
	KBluetoothHCIExtendedInquiryResponseDataTypeIndoorPositioning                        BluetoothHCIExtendedInquiryResponseDataTypes = 0x25
	KBluetoothHCIExtendedInquiryResponseDataTypeLEBluetoothDeviceAddress                 BluetoothHCIExtendedInquiryResponseDataTypes = 0x1b
	KBluetoothHCIExtendedInquiryResponseDataTypeLERole                                   BluetoothHCIExtendedInquiryResponseDataTypes = 0x1c
	KBluetoothHCIExtendedInquiryResponseDataTypeManufacturerSpecificData                 BluetoothHCIExtendedInquiryResponseDataTypes = 0xff
	KBluetoothHCIExtendedInquiryResponseDataTypePeripheralConnectionIntervalRange        BluetoothHCIExtendedInquiryResponseDataTypes = 0x12
	KBluetoothHCIExtendedInquiryResponseDataTypePublicTargetAddress                      BluetoothHCIExtendedInquiryResponseDataTypes = 0x17
	KBluetoothHCIExtendedInquiryResponseDataTypeRandomTargetAddress                      BluetoothHCIExtendedInquiryResponseDataTypes = 0x18
	KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBClassOfDevice                      BluetoothHCIExtendedInquiryResponseDataTypes = 0xd
	KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingHashC                 BluetoothHCIExtendedInquiryResponseDataTypes = 0xe
	KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingRandomizerR           BluetoothHCIExtendedInquiryResponseDataTypes = 0xf
	KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsConfirmationValue       BluetoothHCIExtendedInquiryResponseDataTypes = 0x22
	KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsRandomValue             BluetoothHCIExtendedInquiryResponseDataTypes = 0x23
	KBluetoothHCIExtendedInquiryResponseDataTypeSecurityManagerOOBFlags                  BluetoothHCIExtendedInquiryResponseDataTypes = 0x11
	KBluetoothHCIExtendedInquiryResponseDataTypeSecurityManagerTKValue                   BluetoothHCIExtendedInquiryResponseDataTypes = 0x10
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceData                              BluetoothHCIExtendedInquiryResponseDataTypes = 0x16
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceData128BitUUID                    BluetoothHCIExtendedInquiryResponseDataTypes = 0x21
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceData32BitUUID                     BluetoothHCIExtendedInquiryResponseDataTypes = 0x20
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation128BitUUIDs           BluetoothHCIExtendedInquiryResponseDataTypes = 0x15
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation16BitUUIDs            BluetoothHCIExtendedInquiryResponseDataTypes = 0x14
	KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation32BitUUIDs            BluetoothHCIExtendedInquiryResponseDataTypes = 0x1f
	KBluetoothHCIExtendedInquiryResponseDataTypeShortenedLocalName                       BluetoothHCIExtendedInquiryResponseDataTypes = 0x8
	KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingHash                        BluetoothHCIExtendedInquiryResponseDataTypes = 0x1d
	KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingRandomizer                  BluetoothHCIExtendedInquiryResponseDataTypes = 0x1e
	KBluetoothHCIExtendedInquiryResponseDataTypeSlaveConnectionIntervalRange             BluetoothHCIExtendedInquiryResponseDataTypes = 18
	KBluetoothHCIExtendedInquiryResponseDataTypeTransmitPowerLevel                       BluetoothHCIExtendedInquiryResponseDataTypes = 0xa
	KBluetoothHCIExtendedInquiryResponseDataTypeTransportDiscoveryData                   BluetoothHCIExtendedInquiryResponseDataTypes = 0x26
	KBluetoothHCIExtendedInquiryResponseDataTypeURI                                      BluetoothHCIExtendedInquiryResponseDataTypes = 0x24
)

func (e BluetoothHCIExtendedInquiryResponseDataTypes) String() string {
	switch e {
	case KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsCompleteList:
		return "KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsCompleteList"
	case KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsWithMoreAvailable:
		return "KBluetoothHCIExtendedInquiryResponseDataType128BitServiceClassUUIDsWithMoreAvailable"
	case KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsCompleteList:
		return "KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsCompleteList"
	case KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsWithMoreAvailable:
		return "KBluetoothHCIExtendedInquiryResponseDataType16BitServiceClassUUIDsWithMoreAvailable"
	case KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsCompleteList:
		return "KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsCompleteList"
	case KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsWithMoreAvailable:
		return "KBluetoothHCIExtendedInquiryResponseDataType32BitServiceClassUUIDsWithMoreAvailable"
	case KBluetoothHCIExtendedInquiryResponseDataType3DInformationData:
		return "KBluetoothHCIExtendedInquiryResponseDataType3DInformationData"
	case KBluetoothHCIExtendedInquiryResponseDataTypeAdvertisingInterval:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeAdvertisingInterval"
	case KBluetoothHCIExtendedInquiryResponseDataTypeAppearance:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeAppearance"
	case KBluetoothHCIExtendedInquiryResponseDataTypeCompleteLocalName:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeCompleteLocalName"
	case KBluetoothHCIExtendedInquiryResponseDataTypeCsisRsiData:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeCsisRsiData"
	case KBluetoothHCIExtendedInquiryResponseDataTypeDeviceID:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeDeviceID"
	case KBluetoothHCIExtendedInquiryResponseDataTypeFlags:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeFlags"
	case KBluetoothHCIExtendedInquiryResponseDataTypeIndoorPositioning:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeIndoorPositioning"
	case KBluetoothHCIExtendedInquiryResponseDataTypeLEBluetoothDeviceAddress:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeLEBluetoothDeviceAddress"
	case KBluetoothHCIExtendedInquiryResponseDataTypeLERole:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeLERole"
	case KBluetoothHCIExtendedInquiryResponseDataTypeManufacturerSpecificData:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeManufacturerSpecificData"
	case KBluetoothHCIExtendedInquiryResponseDataTypePeripheralConnectionIntervalRange:
		return "KBluetoothHCIExtendedInquiryResponseDataTypePeripheralConnectionIntervalRange"
	case KBluetoothHCIExtendedInquiryResponseDataTypePublicTargetAddress:
		return "KBluetoothHCIExtendedInquiryResponseDataTypePublicTargetAddress"
	case KBluetoothHCIExtendedInquiryResponseDataTypeRandomTargetAddress:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeRandomTargetAddress"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBClassOfDevice:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBClassOfDevice"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingHashC:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingHashC"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingRandomizerR:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSSPOOBSimplePairingRandomizerR"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsConfirmationValue:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsConfirmationValue"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsRandomValue:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSecureConnectionsRandomValue"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSecurityManagerOOBFlags:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSecurityManagerOOBFlags"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceData:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceData"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceData128BitUUID:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceData128BitUUID"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceData32BitUUID:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceData32BitUUID"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation128BitUUIDs:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation128BitUUIDs"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation16BitUUIDs:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation16BitUUIDs"
	case KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation32BitUUIDs:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeServiceSolicitation32BitUUIDs"
	case KBluetoothHCIExtendedInquiryResponseDataTypeShortenedLocalName:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeShortenedLocalName"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingHash:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingHash"
	case KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingRandomizer:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeSimplePairingRandomizer"
	case KBluetoothHCIExtendedInquiryResponseDataTypeTransmitPowerLevel:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeTransmitPowerLevel"
	case KBluetoothHCIExtendedInquiryResponseDataTypeTransportDiscoveryData:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeTransportDiscoveryData"
	case KBluetoothHCIExtendedInquiryResponseDataTypeURI:
		return "KBluetoothHCIExtendedInquiryResponseDataTypeURI"
	default:
		return fmt.Sprintf("BluetoothHCIExtendedInquiryResponseDataTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIFECRequiredValues
type BluetoothHCIFECRequiredValues uint32

const (
	KBluetoothHCIFECNotRequired BluetoothHCIFECRequiredValues = 0x1
	KBluetoothHCIFECRequired    BluetoothHCIFECRequiredValues = 0
)

func (e BluetoothHCIFECRequiredValues) String() string {
	switch e {
	case KBluetoothHCIFECNotRequired:
		return "KBluetoothHCIFECNotRequired"
	case KBluetoothHCIFECRequired:
		return "KBluetoothHCIFECRequired"
	default:
		return fmt.Sprintf("BluetoothHCIFECRequiredValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIGeneralFlowControlStates
type BluetoothHCIGeneralFlowControlStates uint32

const (
	KHCIACLDataPacketsOffHCISCODataPacketsOn BluetoothHCIGeneralFlowControlStates = 0x2
	KHCIACLDataPacketsOnHCISCODataPacketsOff BluetoothHCIGeneralFlowControlStates = 0x1
	KHCIACLDataPacketsOnHCISCODataPacketsOn  BluetoothHCIGeneralFlowControlStates = 0x3
	KHostControllerToHostFlowControlOff      BluetoothHCIGeneralFlowControlStates = 0
)

func (e BluetoothHCIGeneralFlowControlStates) String() string {
	switch e {
	case KHCIACLDataPacketsOffHCISCODataPacketsOn:
		return "KHCIACLDataPacketsOffHCISCODataPacketsOn"
	case KHCIACLDataPacketsOnHCISCODataPacketsOff:
		return "KHCIACLDataPacketsOnHCISCODataPacketsOff"
	case KHCIACLDataPacketsOnHCISCODataPacketsOn:
		return "KHCIACLDataPacketsOnHCISCODataPacketsOn"
	case KHostControllerToHostFlowControlOff:
		return "KHostControllerToHostFlowControlOff"
	default:
		return fmt.Sprintf("BluetoothHCIGeneralFlowControlStates(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIHoldModeActivityStates
type BluetoothHCIHoldModeActivityStates uint32

const (
	KMaintainCurrentPowerState BluetoothHCIHoldModeActivityStates = 0
	KSuspendInquiryScan        BluetoothHCIHoldModeActivityStates = 0x2
	KSuspendPageScan           BluetoothHCIHoldModeActivityStates = 0x1
	KSuspendPeriodicInquiries  BluetoothHCIHoldModeActivityStates = 0x3
)

func (e BluetoothHCIHoldModeActivityStates) String() string {
	switch e {
	case KMaintainCurrentPowerState:
		return "KMaintainCurrentPowerState"
	case KSuspendInquiryScan:
		return "KSuspendInquiryScan"
	case KSuspendPageScan:
		return "KSuspendPageScan"
	case KSuspendPeriodicInquiries:
		return "KSuspendPeriodicInquiries"
	default:
		return fmt.Sprintf("BluetoothHCIHoldModeActivityStates(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryModes
type BluetoothHCIInquiryModes uint32

const (
	KBluetoothHCIInquiryModeResultFormatStandard                              BluetoothHCIInquiryModes = 0
	KBluetoothHCIInquiryModeResultFormatWithRSSI                              BluetoothHCIInquiryModes = 0x1
	KBluetoothHCIInquiryModeResultFormatWithRSSIOrExtendedInquiryResultFormat BluetoothHCIInquiryModes = 0x2
)

func (e BluetoothHCIInquiryModes) String() string {
	switch e {
	case KBluetoothHCIInquiryModeResultFormatStandard:
		return "KBluetoothHCIInquiryModeResultFormatStandard"
	case KBluetoothHCIInquiryModeResultFormatWithRSSI:
		return "KBluetoothHCIInquiryModeResultFormatWithRSSI"
	case KBluetoothHCIInquiryModeResultFormatWithRSSIOrExtendedInquiryResultFormat:
		return "KBluetoothHCIInquiryModeResultFormatWithRSSIOrExtendedInquiryResultFormat"
	default:
		return fmt.Sprintf("BluetoothHCIInquiryModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryScanTypes
type BluetoothHCIInquiryScanTypes uint32

const (
	KBluetoothHCIInquiryScanTypeInterlaced    BluetoothHCIInquiryScanTypes = 0x1
	KBluetoothHCIInquiryScanTypeReservedEnd   BluetoothHCIInquiryScanTypes = 0xff
	KBluetoothHCIInquiryScanTypeReservedStart BluetoothHCIInquiryScanTypes = 0x2
	KBluetoothHCIInquiryScanTypeStandard      BluetoothHCIInquiryScanTypes = 0
)

func (e BluetoothHCIInquiryScanTypes) String() string {
	switch e {
	case KBluetoothHCIInquiryScanTypeInterlaced:
		return "KBluetoothHCIInquiryScanTypeInterlaced"
	case KBluetoothHCIInquiryScanTypeReservedEnd:
		return "KBluetoothHCIInquiryScanTypeReservedEnd"
	case KBluetoothHCIInquiryScanTypeReservedStart:
		return "KBluetoothHCIInquiryScanTypeReservedStart"
	case KBluetoothHCIInquiryScanTypeStandard:
		return "KBluetoothHCIInquiryScanTypeStandard"
	default:
		return fmt.Sprintf("BluetoothHCIInquiryScanTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkPolicySettingsValues
type BluetoothHCILinkPolicySettingsValues uint32

const (
	KDisableAllLMModes             BluetoothHCILinkPolicySettingsValues = 0
	KEnableCentralPeripheralSwitch BluetoothHCILinkPolicySettingsValues = 0x1
	KEnableHoldMode                BluetoothHCILinkPolicySettingsValues = 0x2
	KEnableMasterSlaveSwitch       BluetoothHCILinkPolicySettingsValues = 1
	KEnableParkMode                BluetoothHCILinkPolicySettingsValues = 0x8
	KEnableSniffMode               BluetoothHCILinkPolicySettingsValues = 0x4
	KReservedForFutureUse          BluetoothHCILinkPolicySettingsValues = 0x10
)

func (e BluetoothHCILinkPolicySettingsValues) String() string {
	switch e {
	case KDisableAllLMModes:
		return "KDisableAllLMModes"
	case KEnableCentralPeripheralSwitch:
		return "KEnableCentralPeripheralSwitch"
	case KEnableHoldMode:
		return "KEnableHoldMode"
	case KEnableParkMode:
		return "KEnableParkMode"
	case KEnableSniffMode:
		return "KEnableSniffMode"
	case KReservedForFutureUse:
		return "KReservedForFutureUse"
	default:
		return fmt.Sprintf("BluetoothHCILinkPolicySettingsValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIPageScanEnableStates
type BluetoothHCIPageScanEnableStates uint32

const (
	KInquiryScanDisabledPageScanEnabled BluetoothHCIPageScanEnableStates = 0x2
	KInquiryScanEnabledPageScanDisabled BluetoothHCIPageScanEnableStates = 0x1
	KInquiryScanEnabledPageScanEnabled  BluetoothHCIPageScanEnableStates = 0x3
	KNoScansEnabled                     BluetoothHCIPageScanEnableStates = 0
)

func (e BluetoothHCIPageScanEnableStates) String() string {
	switch e {
	case KInquiryScanDisabledPageScanEnabled:
		return "KInquiryScanDisabledPageScanEnabled"
	case KInquiryScanEnabledPageScanDisabled:
		return "KInquiryScanEnabledPageScanDisabled"
	case KInquiryScanEnabledPageScanEnabled:
		return "KInquiryScanEnabledPageScanEnabled"
	case KNoScansEnabled:
		return "KNoScansEnabled"
	default:
		return fmt.Sprintf("BluetoothHCIPageScanEnableStates(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIPageScanModes
type BluetoothHCIPageScanModes uint32

const (
	KMandatoryPageScanMode BluetoothHCIPageScanModes = 0
	KOptionalPageScanMode1 BluetoothHCIPageScanModes = 0x1
	KOptionalPageScanMode2 BluetoothHCIPageScanModes = 0x2
	KOptionalPageScanMode3 BluetoothHCIPageScanModes = 0x3
)

func (e BluetoothHCIPageScanModes) String() string {
	switch e {
	case KMandatoryPageScanMode:
		return "KMandatoryPageScanMode"
	case KOptionalPageScanMode1:
		return "KOptionalPageScanMode1"
	case KOptionalPageScanMode2:
		return "KOptionalPageScanMode2"
	case KOptionalPageScanMode3:
		return "KOptionalPageScanMode3"
	default:
		return fmt.Sprintf("BluetoothHCIPageScanModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIPageScanPeriodModes
type BluetoothHCIPageScanPeriodModes uint32

const (
	KP0Mode BluetoothHCIPageScanPeriodModes = 0
	KP1Mode BluetoothHCIPageScanPeriodModes = 0x1
	KP2Mode BluetoothHCIPageScanPeriodModes = 0x2
)

func (e BluetoothHCIPageScanPeriodModes) String() string {
	switch e {
	case KP0Mode:
		return "KP0Mode"
	case KP1Mode:
		return "KP1Mode"
	case KP2Mode:
		return "KP2Mode"
	default:
		return fmt.Sprintf("BluetoothHCIPageScanPeriodModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIPageScanTypes
type BluetoothHCIPageScanTypes uint32

const (
	KBluetoothHCIPageScanTypeInterlaced    BluetoothHCIPageScanTypes = 0x1
	KBluetoothHCIPageScanTypeReservedEnd   BluetoothHCIPageScanTypes = 0xff
	KBluetoothHCIPageScanTypeReservedStart BluetoothHCIPageScanTypes = 0x2
	KBluetoothHCIPageScanTypeStandard      BluetoothHCIPageScanTypes = 0
)

func (e BluetoothHCIPageScanTypes) String() string {
	switch e {
	case KBluetoothHCIPageScanTypeInterlaced:
		return "KBluetoothHCIPageScanTypeInterlaced"
	case KBluetoothHCIPageScanTypeReservedEnd:
		return "KBluetoothHCIPageScanTypeReservedEnd"
	case KBluetoothHCIPageScanTypeReservedStart:
		return "KBluetoothHCIPageScanTypeReservedStart"
	case KBluetoothHCIPageScanTypeStandard:
		return "KBluetoothHCIPageScanTypeStandard"
	default:
		return fmt.Sprintf("BluetoothHCIPageScanTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIPowerState
type BluetoothHCIPowerState uint32

const (
	KBluetoothHCIPowerStateOFF          BluetoothHCIPowerState = 0
	KBluetoothHCIPowerStateON           BluetoothHCIPowerState = 0x1
	KBluetoothHCIPowerStateUnintialized BluetoothHCIPowerState = 0xff
)

func (e BluetoothHCIPowerState) String() string {
	switch e {
	case KBluetoothHCIPowerStateOFF:
		return "KBluetoothHCIPowerStateOFF"
	case KBluetoothHCIPowerStateON:
		return "KBluetoothHCIPowerStateON"
	case KBluetoothHCIPowerStateUnintialized:
		return "KBluetoothHCIPowerStateUnintialized"
	default:
		return fmt.Sprintf("BluetoothHCIPowerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadStoredLinkKeysFlags
type BluetoothHCIReadStoredLinkKeysFlags uint32

const (
	KReadAllStoredLinkKeys               BluetoothHCIReadStoredLinkKeysFlags = 0x1
	KReturnLinkKeyForSpecifiedDeviceOnly BluetoothHCIReadStoredLinkKeysFlags = 0
)

func (e BluetoothHCIReadStoredLinkKeysFlags) String() string {
	switch e {
	case KReadAllStoredLinkKeys:
		return "KReadAllStoredLinkKeys"
	case KReturnLinkKeyForSpecifiedDeviceOnly:
		return "KReturnLinkKeyForSpecifiedDeviceOnly"
	default:
		return fmt.Sprintf("BluetoothHCIReadStoredLinkKeysFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRetransmissionEffortTypes
type BluetoothHCIRetransmissionEffortTypes uint32

const (
	KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeForPower    BluetoothHCIRetransmissionEffortTypes = 0x1
	KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeLinkQuality BluetoothHCIRetransmissionEffortTypes = 0x2
	KHCIRetransmissionEffortTypeDontCare                         BluetoothHCIRetransmissionEffortTypes = 0xff
	KHCIRetransmissionEffortTypeNone                             BluetoothHCIRetransmissionEffortTypes = 0
)

func (e BluetoothHCIRetransmissionEffortTypes) String() string {
	switch e {
	case KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeForPower:
		return "KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeForPower"
	case KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeLinkQuality:
		return "KHCIRetransmissionEffortTypeAtLeastOneAndOptimizeLinkQuality"
	case KHCIRetransmissionEffortTypeDontCare:
		return "KHCIRetransmissionEffortTypeDontCare"
	case KHCIRetransmissionEffortTypeNone:
		return "KHCIRetransmissionEffortTypeNone"
	default:
		return fmt.Sprintf("BluetoothHCIRetransmissionEffortTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRoles
type BluetoothHCIRoles uint32

const (
	KBluetoothHCICentralRole    BluetoothHCIRoles = 0
	KBluetoothHCIMasterRole     BluetoothHCIRoles = 0
	KBluetoothHCIPeripheralRole BluetoothHCIRoles = 0x1
	KBluetoothHCISlaveRole      BluetoothHCIRoles = 1
)

func (e BluetoothHCIRoles) String() string {
	switch e {
	case KBluetoothHCICentralRole:
		return "KBluetoothHCICentralRole"
	case KBluetoothHCIPeripheralRole:
		return "KBluetoothHCIPeripheralRole"
	default:
		return fmt.Sprintf("BluetoothHCIRoles(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISCOFlowControlStates
type BluetoothHCISCOFlowControlStates uint32

const (
	KSCOFlowControlDisabled BluetoothHCISCOFlowControlStates = 0
	KSCOFlowControlEnabled  BluetoothHCISCOFlowControlStates = 0x1
)

func (e BluetoothHCISCOFlowControlStates) String() string {
	switch e {
	case KSCOFlowControlDisabled:
		return "KSCOFlowControlDisabled"
	case KSCOFlowControlEnabled:
		return "KSCOFlowControlEnabled"
	default:
		return fmt.Sprintf("BluetoothHCISCOFlowControlStates(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISimplePairingModes
type BluetoothHCISimplePairingModes uint32

const (
	KBluetoothHCISimplePairingModeEnabled BluetoothHCISimplePairingModes = 0x1
	KBluetoothHCISimplePairingModeNotSet  BluetoothHCISimplePairingModes = 0
)

func (e BluetoothHCISimplePairingModes) String() string {
	switch e {
	case KBluetoothHCISimplePairingModeEnabled:
		return "KBluetoothHCISimplePairingModeEnabled"
	case KBluetoothHCISimplePairingModeNotSet:
		return "KBluetoothHCISimplePairingModeNotSet"
	default:
		return fmt.Sprintf("BluetoothHCISimplePairingModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCITimeoutValues
type BluetoothHCITimeoutValues uint32

const (
	KDefaultPageTimeout BluetoothHCITimeoutValues = 0x2710
)

func (e BluetoothHCITimeoutValues) String() string {
	switch e {
	case KDefaultPageTimeout:
		return "KDefaultPageTimeout"
	default:
		return fmt.Sprintf("BluetoothHCITimeoutValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCITransmitReadPowerLevelTypes
type BluetoothHCITransmitReadPowerLevelTypes uint32

const (
	KReadCurrentTransmitPowerLevel BluetoothHCITransmitReadPowerLevelTypes = 0
	KReadMaximumTransmitPowerLevel BluetoothHCITransmitReadPowerLevelTypes = 0x1
)

func (e BluetoothHCITransmitReadPowerLevelTypes) String() string {
	switch e {
	case KReadCurrentTransmitPowerLevel:
		return "KReadCurrentTransmitPowerLevel"
	case KReadMaximumTransmitPowerLevel:
		return "KReadMaximumTransmitPowerLevel"
	default:
		return fmt.Sprintf("BluetoothHCITransmitReadPowerLevelTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIVersions
type BluetoothHCIVersions uint32

const (
	KBluetoothHCIVersionCoreSpecification1_0b   BluetoothHCIVersions = 0
	KBluetoothHCIVersionCoreSpecification1_1    BluetoothHCIVersions = 0x1
	KBluetoothHCIVersionCoreSpecification1_2    BluetoothHCIVersions = 0x2
	KBluetoothHCIVersionCoreSpecification2_0EDR BluetoothHCIVersions = 0x3
	KBluetoothHCIVersionCoreSpecification2_1EDR BluetoothHCIVersions = 0x4
	KBluetoothHCIVersionCoreSpecification3_0HS  BluetoothHCIVersions = 0x5
	KBluetoothHCIVersionCoreSpecification4_0    BluetoothHCIVersions = 0x6
	KBluetoothHCIVersionCoreSpecification4_1    BluetoothHCIVersions = 0x7
	KBluetoothHCIVersionCoreSpecification4_2    BluetoothHCIVersions = 0x8
	KBluetoothHCIVersionCoreSpecification5_0    BluetoothHCIVersions = 0x9
	KBluetoothHCIVersionCoreSpecification5_1    BluetoothHCIVersions = 0xa
	KBluetoothHCIVersionCoreSpecification5_2    BluetoothHCIVersions = 0xb
	KBluetoothHCIVersionCoreSpecification5_3    BluetoothHCIVersions = 0xc
)

func (e BluetoothHCIVersions) String() string {
	switch e {
	case KBluetoothHCIVersionCoreSpecification1_0b:
		return "KBluetoothHCIVersionCoreSpecification1_0b"
	case KBluetoothHCIVersionCoreSpecification1_1:
		return "KBluetoothHCIVersionCoreSpecification1_1"
	case KBluetoothHCIVersionCoreSpecification1_2:
		return "KBluetoothHCIVersionCoreSpecification1_2"
	case KBluetoothHCIVersionCoreSpecification2_0EDR:
		return "KBluetoothHCIVersionCoreSpecification2_0EDR"
	case KBluetoothHCIVersionCoreSpecification2_1EDR:
		return "KBluetoothHCIVersionCoreSpecification2_1EDR"
	case KBluetoothHCIVersionCoreSpecification3_0HS:
		return "KBluetoothHCIVersionCoreSpecification3_0HS"
	case KBluetoothHCIVersionCoreSpecification4_0:
		return "KBluetoothHCIVersionCoreSpecification4_0"
	case KBluetoothHCIVersionCoreSpecification4_1:
		return "KBluetoothHCIVersionCoreSpecification4_1"
	case KBluetoothHCIVersionCoreSpecification4_2:
		return "KBluetoothHCIVersionCoreSpecification4_2"
	case KBluetoothHCIVersionCoreSpecification5_0:
		return "KBluetoothHCIVersionCoreSpecification5_0"
	case KBluetoothHCIVersionCoreSpecification5_1:
		return "KBluetoothHCIVersionCoreSpecification5_1"
	case KBluetoothHCIVersionCoreSpecification5_2:
		return "KBluetoothHCIVersionCoreSpecification5_2"
	case KBluetoothHCIVersionCoreSpecification5_3:
		return "KBluetoothHCIVersionCoreSpecification5_3"
	default:
		return fmt.Sprintf("BluetoothHCIVersions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothIOCapabilities
type BluetoothIOCapabilities uint32

const (
	KBluetoothCapabilityTypeDisplayOnly     BluetoothIOCapabilities = 0
	KBluetoothCapabilityTypeDisplayYesNo    BluetoothIOCapabilities = 0x1
	KBluetoothCapabilityTypeKeyboardOnly    BluetoothIOCapabilities = 0x2
	KBluetoothCapabilityTypeNoInputNoOutput BluetoothIOCapabilities = 0x3
)

func (e BluetoothIOCapabilities) String() string {
	switch e {
	case KBluetoothCapabilityTypeDisplayOnly:
		return "KBluetoothCapabilityTypeDisplayOnly"
	case KBluetoothCapabilityTypeDisplayYesNo:
		return "KBluetoothCapabilityTypeDisplayYesNo"
	case KBluetoothCapabilityTypeKeyboardOnly:
		return "KBluetoothCapabilityTypeKeyboardOnly"
	case KBluetoothCapabilityTypeNoInputNoOutput:
		return "KBluetoothCapabilityTypeNoInputNoOutput"
	default:
		return fmt.Sprintf("BluetoothIOCapabilities(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothKeypressNotificationTypes
type BluetoothKeypressNotificationTypes uint32

const (
	KBluetoothKeypressNotificationTypePasskeyCleared        BluetoothKeypressNotificationTypes = 3
	KBluetoothKeypressNotificationTypePasskeyDigitEntered   BluetoothKeypressNotificationTypes = 1
	KBluetoothKeypressNotificationTypePasskeyDigitErased    BluetoothKeypressNotificationTypes = 2
	KBluetoothKeypressNotificationTypePasskeyEntryCompleted BluetoothKeypressNotificationTypes = 4
	KBluetoothKeypressNotificationTypePasskeyEntryStarted   BluetoothKeypressNotificationTypes = 0
)

func (e BluetoothKeypressNotificationTypes) String() string {
	switch e {
	case KBluetoothKeypressNotificationTypePasskeyCleared:
		return "KBluetoothKeypressNotificationTypePasskeyCleared"
	case KBluetoothKeypressNotificationTypePasskeyDigitEntered:
		return "KBluetoothKeypressNotificationTypePasskeyDigitEntered"
	case KBluetoothKeypressNotificationTypePasskeyDigitErased:
		return "KBluetoothKeypressNotificationTypePasskeyDigitErased"
	case KBluetoothKeypressNotificationTypePasskeyEntryCompleted:
		return "KBluetoothKeypressNotificationTypePasskeyEntryCompleted"
	case KBluetoothKeypressNotificationTypePasskeyEntryStarted:
		return "KBluetoothKeypressNotificationTypePasskeyEntryStarted"
	default:
		return fmt.Sprintf("BluetoothKeypressNotificationTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPCommandCode
type BluetoothL2CAPCommandCode uint32

const (
	KBluetoothL2CAPCommandCodeCommandReject                     BluetoothL2CAPCommandCode = 0x1
	KBluetoothL2CAPCommandCodeConfigureRequest                  BluetoothL2CAPCommandCode = 0x4
	KBluetoothL2CAPCommandCodeConfigureResponse                 BluetoothL2CAPCommandCode = 0x5
	KBluetoothL2CAPCommandCodeConnectionParameterUpdateRequest  BluetoothL2CAPCommandCode = 0x12
	KBluetoothL2CAPCommandCodeConnectionParameterUpdateResponse BluetoothL2CAPCommandCode = 0x13
	KBluetoothL2CAPCommandCodeConnectionRequest                 BluetoothL2CAPCommandCode = 0x2
	KBluetoothL2CAPCommandCodeConnectionResponse                BluetoothL2CAPCommandCode = 0x3
	KBluetoothL2CAPCommandCodeCreateChannelRequest              BluetoothL2CAPCommandCode = 0xc
	KBluetoothL2CAPCommandCodeCreateChannelResponse             BluetoothL2CAPCommandCode = 0xd
	KBluetoothL2CAPCommandCodeDisconnectionRequest              BluetoothL2CAPCommandCode = 0x6
	KBluetoothL2CAPCommandCodeDisconnectionResponse             BluetoothL2CAPCommandCode = 0x7
	KBluetoothL2CAPCommandCodeEchoRequest                       BluetoothL2CAPCommandCode = 0x8
	KBluetoothL2CAPCommandCodeEchoResponse                      BluetoothL2CAPCommandCode = 0x9
	KBluetoothL2CAPCommandCodeInformationRequest                BluetoothL2CAPCommandCode = 0xa
	KBluetoothL2CAPCommandCodeInformationResponse               BluetoothL2CAPCommandCode = 0xb
	KBluetoothL2CAPCommandCodeLECreditBasedConnectionRequest    BluetoothL2CAPCommandCode = 0x14
	KBluetoothL2CAPCommandCodeLECreditBasedConnectionResponse   BluetoothL2CAPCommandCode = 0x15
	KBluetoothL2CAPCommandCodeLEFlowControlCredit               BluetoothL2CAPCommandCode = 0x16
	KBluetoothL2CAPCommandCodeMoveChannelConfirmation           BluetoothL2CAPCommandCode = 0x10
	KBluetoothL2CAPCommandCodeMoveChannelConfirmationResponse   BluetoothL2CAPCommandCode = 0x11
	KBluetoothL2CAPCommandCodeMoveChannelRequest                BluetoothL2CAPCommandCode = 0xe
	KBluetoothL2CAPCommandCodeMoveChannelResponse               BluetoothL2CAPCommandCode = 0xf
	KBluetoothL2CAPCommandCodeReserved                          BluetoothL2CAPCommandCode = 0
)

func (e BluetoothL2CAPCommandCode) String() string {
	switch e {
	case KBluetoothL2CAPCommandCodeCommandReject:
		return "KBluetoothL2CAPCommandCodeCommandReject"
	case KBluetoothL2CAPCommandCodeConfigureRequest:
		return "KBluetoothL2CAPCommandCodeConfigureRequest"
	case KBluetoothL2CAPCommandCodeConfigureResponse:
		return "KBluetoothL2CAPCommandCodeConfigureResponse"
	case KBluetoothL2CAPCommandCodeConnectionParameterUpdateRequest:
		return "KBluetoothL2CAPCommandCodeConnectionParameterUpdateRequest"
	case KBluetoothL2CAPCommandCodeConnectionParameterUpdateResponse:
		return "KBluetoothL2CAPCommandCodeConnectionParameterUpdateResponse"
	case KBluetoothL2CAPCommandCodeConnectionRequest:
		return "KBluetoothL2CAPCommandCodeConnectionRequest"
	case KBluetoothL2CAPCommandCodeConnectionResponse:
		return "KBluetoothL2CAPCommandCodeConnectionResponse"
	case KBluetoothL2CAPCommandCodeCreateChannelRequest:
		return "KBluetoothL2CAPCommandCodeCreateChannelRequest"
	case KBluetoothL2CAPCommandCodeCreateChannelResponse:
		return "KBluetoothL2CAPCommandCodeCreateChannelResponse"
	case KBluetoothL2CAPCommandCodeDisconnectionRequest:
		return "KBluetoothL2CAPCommandCodeDisconnectionRequest"
	case KBluetoothL2CAPCommandCodeDisconnectionResponse:
		return "KBluetoothL2CAPCommandCodeDisconnectionResponse"
	case KBluetoothL2CAPCommandCodeEchoRequest:
		return "KBluetoothL2CAPCommandCodeEchoRequest"
	case KBluetoothL2CAPCommandCodeEchoResponse:
		return "KBluetoothL2CAPCommandCodeEchoResponse"
	case KBluetoothL2CAPCommandCodeInformationRequest:
		return "KBluetoothL2CAPCommandCodeInformationRequest"
	case KBluetoothL2CAPCommandCodeInformationResponse:
		return "KBluetoothL2CAPCommandCodeInformationResponse"
	case KBluetoothL2CAPCommandCodeLECreditBasedConnectionRequest:
		return "KBluetoothL2CAPCommandCodeLECreditBasedConnectionRequest"
	case KBluetoothL2CAPCommandCodeLECreditBasedConnectionResponse:
		return "KBluetoothL2CAPCommandCodeLECreditBasedConnectionResponse"
	case KBluetoothL2CAPCommandCodeLEFlowControlCredit:
		return "KBluetoothL2CAPCommandCodeLEFlowControlCredit"
	case KBluetoothL2CAPCommandCodeMoveChannelConfirmation:
		return "KBluetoothL2CAPCommandCodeMoveChannelConfirmation"
	case KBluetoothL2CAPCommandCodeMoveChannelConfirmationResponse:
		return "KBluetoothL2CAPCommandCodeMoveChannelConfirmationResponse"
	case KBluetoothL2CAPCommandCodeMoveChannelRequest:
		return "KBluetoothL2CAPCommandCodeMoveChannelRequest"
	case KBluetoothL2CAPCommandCodeMoveChannelResponse:
		return "KBluetoothL2CAPCommandCodeMoveChannelResponse"
	case KBluetoothL2CAPCommandCodeReserved:
		return "KBluetoothL2CAPCommandCodeReserved"
	default:
		return fmt.Sprintf("BluetoothL2CAPCommandCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPCommandRejectReason
type BluetoothL2CAPCommandRejectReason uint32

const (
	KBluetoothL2CAPCommandRejectReasonCommandNotUnderstood  BluetoothL2CAPCommandRejectReason = 0
	KBluetoothL2CAPCommandRejectReasonInvalidCIDInRequest   BluetoothL2CAPCommandRejectReason = 0x2
	KBluetoothL2CAPCommandRejectReasonSignallingMTUExceeded BluetoothL2CAPCommandRejectReason = 0x1
)

func (e BluetoothL2CAPCommandRejectReason) String() string {
	switch e {
	case KBluetoothL2CAPCommandRejectReasonCommandNotUnderstood:
		return "KBluetoothL2CAPCommandRejectReasonCommandNotUnderstood"
	case KBluetoothL2CAPCommandRejectReasonInvalidCIDInRequest:
		return "KBluetoothL2CAPCommandRejectReasonInvalidCIDInRequest"
	case KBluetoothL2CAPCommandRejectReasonSignallingMTUExceeded:
		return "KBluetoothL2CAPCommandRejectReasonSignallingMTUExceeded"
	default:
		return fmt.Sprintf("BluetoothL2CAPCommandRejectReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPConfigurationOption
type BluetoothL2CAPConfigurationOption uint32

const (
	KBluetoothL2CAPConfigurationOptionExtendedFlowSpecification    BluetoothL2CAPConfigurationOption = 0x6
	KBluetoothL2CAPConfigurationOptionExtendedWindowSize           BluetoothL2CAPConfigurationOption = 0x7
	KBluetoothL2CAPConfigurationOptionFlushTimeout                 BluetoothL2CAPConfigurationOption = 0x2
	KBluetoothL2CAPConfigurationOptionFrameCheckSequence           BluetoothL2CAPConfigurationOption = 0x5
	KBluetoothL2CAPConfigurationOptionMTU                          BluetoothL2CAPConfigurationOption = 0x1
	KBluetoothL2CAPConfigurationOptionQoS                          BluetoothL2CAPConfigurationOption = 0x3
	KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControl BluetoothL2CAPConfigurationOption = 0x4
)

func (e BluetoothL2CAPConfigurationOption) String() string {
	switch e {
	case KBluetoothL2CAPConfigurationOptionExtendedFlowSpecification:
		return "KBluetoothL2CAPConfigurationOptionExtendedFlowSpecification"
	case KBluetoothL2CAPConfigurationOptionExtendedWindowSize:
		return "KBluetoothL2CAPConfigurationOptionExtendedWindowSize"
	case KBluetoothL2CAPConfigurationOptionFlushTimeout:
		return "KBluetoothL2CAPConfigurationOptionFlushTimeout"
	case KBluetoothL2CAPConfigurationOptionFrameCheckSequence:
		return "KBluetoothL2CAPConfigurationOptionFrameCheckSequence"
	case KBluetoothL2CAPConfigurationOptionMTU:
		return "KBluetoothL2CAPConfigurationOptionMTU"
	case KBluetoothL2CAPConfigurationOptionQoS:
		return "KBluetoothL2CAPConfigurationOptionQoS"
	case KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControl:
		return "KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControl"
	default:
		return fmt.Sprintf("BluetoothL2CAPConfigurationOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPConfigurationResult
type BluetoothL2CAPConfigurationResult uint32

const (
	KBluetoothL2CAPConfigurationResultRejected           BluetoothL2CAPConfigurationResult = 0x2
	KBluetoothL2CAPConfigurationResultSuccess            BluetoothL2CAPConfigurationResult = 0
	KBluetoothL2CAPConfigurationResultUnacceptableParams BluetoothL2CAPConfigurationResult = 0x1
	KBluetoothL2CAPConfigurationResultUnknownOptions     BluetoothL2CAPConfigurationResult = 0x3
)

func (e BluetoothL2CAPConfigurationResult) String() string {
	switch e {
	case KBluetoothL2CAPConfigurationResultRejected:
		return "KBluetoothL2CAPConfigurationResultRejected"
	case KBluetoothL2CAPConfigurationResultSuccess:
		return "KBluetoothL2CAPConfigurationResultSuccess"
	case KBluetoothL2CAPConfigurationResultUnacceptableParams:
		return "KBluetoothL2CAPConfigurationResultUnacceptableParams"
	case KBluetoothL2CAPConfigurationResultUnknownOptions:
		return "KBluetoothL2CAPConfigurationResultUnknownOptions"
	default:
		return fmt.Sprintf("BluetoothL2CAPConfigurationResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags
type BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags uint32

const (
	KBluetoothL2CAPConfigurationBasicL2CAPModeFlag         BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = 0
	KBluetoothL2CAPConfigurationEnhancedRetransmissionMode BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = 0x3
	KBluetoothL2CAPConfigurationFlowControlModeFlag        BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = 0x2
	KBluetoothL2CAPConfigurationRetransmissionModeFlag     BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = 0x1
	KBluetoothL2CAPConfigurationStreamingMode              BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = 0x4
)

func (e BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags) String() string {
	switch e {
	case KBluetoothL2CAPConfigurationBasicL2CAPModeFlag:
		return "KBluetoothL2CAPConfigurationBasicL2CAPModeFlag"
	case KBluetoothL2CAPConfigurationEnhancedRetransmissionMode:
		return "KBluetoothL2CAPConfigurationEnhancedRetransmissionMode"
	case KBluetoothL2CAPConfigurationFlowControlModeFlag:
		return "KBluetoothL2CAPConfigurationFlowControlModeFlag"
	case KBluetoothL2CAPConfigurationRetransmissionModeFlag:
		return "KBluetoothL2CAPConfigurationRetransmissionModeFlag"
	case KBluetoothL2CAPConfigurationStreamingMode:
		return "KBluetoothL2CAPConfigurationStreamingMode"
	default:
		return fmt.Sprintf("BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPConnectionResult
type BluetoothL2CAPConnectionResult uint32

const (
	KBluetoothL2CAPConnectionResultPending                          BluetoothL2CAPConnectionResult = 0x1
	KBluetoothL2CAPConnectionResultRefusedInvalidSourceCID          BluetoothL2CAPConnectionResult = 0x6
	KBluetoothL2CAPConnectionResultRefusedNoResources               BluetoothL2CAPConnectionResult = 0x4
	KBluetoothL2CAPConnectionResultRefusedPSMNotSupported           BluetoothL2CAPConnectionResult = 0x2
	KBluetoothL2CAPConnectionResultRefusedReserved                  BluetoothL2CAPConnectionResult = 0x5
	KBluetoothL2CAPConnectionResultRefusedSecurityBlock             BluetoothL2CAPConnectionResult = 0x3
	KBluetoothL2CAPConnectionResultRefusedSourceCIDAlreadyAllocated BluetoothL2CAPConnectionResult = 0x7
	KBluetoothL2CAPConnectionResultSuccessful                       BluetoothL2CAPConnectionResult = 0
)

func (e BluetoothL2CAPConnectionResult) String() string {
	switch e {
	case KBluetoothL2CAPConnectionResultPending:
		return "KBluetoothL2CAPConnectionResultPending"
	case KBluetoothL2CAPConnectionResultRefusedInvalidSourceCID:
		return "KBluetoothL2CAPConnectionResultRefusedInvalidSourceCID"
	case KBluetoothL2CAPConnectionResultRefusedNoResources:
		return "KBluetoothL2CAPConnectionResultRefusedNoResources"
	case KBluetoothL2CAPConnectionResultRefusedPSMNotSupported:
		return "KBluetoothL2CAPConnectionResultRefusedPSMNotSupported"
	case KBluetoothL2CAPConnectionResultRefusedReserved:
		return "KBluetoothL2CAPConnectionResultRefusedReserved"
	case KBluetoothL2CAPConnectionResultRefusedSecurityBlock:
		return "KBluetoothL2CAPConnectionResultRefusedSecurityBlock"
	case KBluetoothL2CAPConnectionResultRefusedSourceCIDAlreadyAllocated:
		return "KBluetoothL2CAPConnectionResultRefusedSourceCIDAlreadyAllocated"
	case KBluetoothL2CAPConnectionResultSuccessful:
		return "KBluetoothL2CAPConnectionResultSuccessful"
	default:
		return fmt.Sprintf("BluetoothL2CAPConnectionResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPConnectionStatus
type BluetoothL2CAPConnectionStatus uint32

const (
	KBluetoothL2CAPConnectionStatusAuthenticationPending BluetoothL2CAPConnectionStatus = 0x1
	KBluetoothL2CAPConnectionStatusAuthorizationPending  BluetoothL2CAPConnectionStatus = 0x2
	KBluetoothL2CAPConnectionStatusNoInfoAvailable       BluetoothL2CAPConnectionStatus = 0
)

func (e BluetoothL2CAPConnectionStatus) String() string {
	switch e {
	case KBluetoothL2CAPConnectionStatusAuthenticationPending:
		return "KBluetoothL2CAPConnectionStatusAuthenticationPending"
	case KBluetoothL2CAPConnectionStatusAuthorizationPending:
		return "KBluetoothL2CAPConnectionStatusAuthorizationPending"
	case KBluetoothL2CAPConnectionStatusNoInfoAvailable:
		return "KBluetoothL2CAPConnectionStatusNoInfoAvailable"
	default:
		return fmt.Sprintf("BluetoothL2CAPConnectionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPInformationExtendedFeaturesMask
type BluetoothL2CAPInformationExtendedFeaturesMask uint32

const (
	KBluetoothL2CAPInformationBidirectionalQoS           BluetoothL2CAPInformationExtendedFeaturesMask = 0x4
	KBluetoothL2CAPInformationEnhancedRetransmissionMode BluetoothL2CAPInformationExtendedFeaturesMask = 0x8
	KBluetoothL2CAPInformationExtendedFlowSpecification  BluetoothL2CAPInformationExtendedFeaturesMask = 0x40
	KBluetoothL2CAPInformationExtendedWindowSize         BluetoothL2CAPInformationExtendedFeaturesMask = 0x100
	KBluetoothL2CAPInformationFCSOption                  BluetoothL2CAPInformationExtendedFeaturesMask = 0x20
	KBluetoothL2CAPInformationFixedChannels              BluetoothL2CAPInformationExtendedFeaturesMask = 0x80
	KBluetoothL2CAPInformationFlowControlMode            BluetoothL2CAPInformationExtendedFeaturesMask = 0x1
	KBluetoothL2CAPInformationNoExtendedFeatures         BluetoothL2CAPInformationExtendedFeaturesMask = 0
	KBluetoothL2CAPInformationRetransmissionMode         BluetoothL2CAPInformationExtendedFeaturesMask = 0x2
	KBluetoothL2CAPInformationStreamingMode              BluetoothL2CAPInformationExtendedFeaturesMask = 0x10
	KBluetoothL2CAPUnicastConnectionlessDataReception    BluetoothL2CAPInformationExtendedFeaturesMask = 0x200
)

func (e BluetoothL2CAPInformationExtendedFeaturesMask) String() string {
	switch e {
	case KBluetoothL2CAPInformationBidirectionalQoS:
		return "KBluetoothL2CAPInformationBidirectionalQoS"
	case KBluetoothL2CAPInformationEnhancedRetransmissionMode:
		return "KBluetoothL2CAPInformationEnhancedRetransmissionMode"
	case KBluetoothL2CAPInformationExtendedFlowSpecification:
		return "KBluetoothL2CAPInformationExtendedFlowSpecification"
	case KBluetoothL2CAPInformationExtendedWindowSize:
		return "KBluetoothL2CAPInformationExtendedWindowSize"
	case KBluetoothL2CAPInformationFCSOption:
		return "KBluetoothL2CAPInformationFCSOption"
	case KBluetoothL2CAPInformationFixedChannels:
		return "KBluetoothL2CAPInformationFixedChannels"
	case KBluetoothL2CAPInformationFlowControlMode:
		return "KBluetoothL2CAPInformationFlowControlMode"
	case KBluetoothL2CAPInformationNoExtendedFeatures:
		return "KBluetoothL2CAPInformationNoExtendedFeatures"
	case KBluetoothL2CAPInformationRetransmissionMode:
		return "KBluetoothL2CAPInformationRetransmissionMode"
	case KBluetoothL2CAPInformationStreamingMode:
		return "KBluetoothL2CAPInformationStreamingMode"
	case KBluetoothL2CAPUnicastConnectionlessDataReception:
		return "KBluetoothL2CAPUnicastConnectionlessDataReception"
	default:
		return fmt.Sprintf("BluetoothL2CAPInformationExtendedFeaturesMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPInformationResult
type BluetoothL2CAPInformationResult uint32

const (
	KBluetoothL2CAPInformationResultNotSupported BluetoothL2CAPInformationResult = 0x1
	KBluetoothL2CAPInformationResultSuccess      BluetoothL2CAPInformationResult = 0
)

func (e BluetoothL2CAPInformationResult) String() string {
	switch e {
	case KBluetoothL2CAPInformationResultNotSupported:
		return "KBluetoothL2CAPInformationResultNotSupported"
	case KBluetoothL2CAPInformationResultSuccess:
		return "KBluetoothL2CAPInformationResultSuccess"
	default:
		return fmt.Sprintf("BluetoothL2CAPInformationResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPInformationType
type BluetoothL2CAPInformationType uint32

const (
	KBluetoothL2CAPInformationTypeConnectionlessMTU      BluetoothL2CAPInformationType = 0x1
	KBluetoothL2CAPInformationTypeExtendedFeatures       BluetoothL2CAPInformationType = 0x2
	KBluetoothL2CAPInformationTypeFixedChannelsSupported BluetoothL2CAPInformationType = 0x3
)

func (e BluetoothL2CAPInformationType) String() string {
	switch e {
	case KBluetoothL2CAPInformationTypeConnectionlessMTU:
		return "KBluetoothL2CAPInformationTypeConnectionlessMTU"
	case KBluetoothL2CAPInformationTypeExtendedFeatures:
		return "KBluetoothL2CAPInformationTypeExtendedFeatures"
	case KBluetoothL2CAPInformationTypeFixedChannelsSupported:
		return "KBluetoothL2CAPInformationTypeFixedChannelsSupported"
	default:
		return fmt.Sprintf("BluetoothL2CAPInformationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPQoSType
type BluetoothL2CAPQoSType uint32

const (
	KBluetoothL2CAPQoSTypeBestEffort BluetoothL2CAPQoSType = 0x1
	KBluetoothL2CAPQoSTypeGuaranteed BluetoothL2CAPQoSType = 0x2
	KBluetoothL2CAPQoSTypeNoTraffic  BluetoothL2CAPQoSType = 0
)

func (e BluetoothL2CAPQoSType) String() string {
	switch e {
	case KBluetoothL2CAPQoSTypeBestEffort:
		return "KBluetoothL2CAPQoSTypeBestEffort"
	case KBluetoothL2CAPQoSTypeGuaranteed:
		return "KBluetoothL2CAPQoSTypeGuaranteed"
	case KBluetoothL2CAPQoSTypeNoTraffic:
		return "KBluetoothL2CAPQoSTypeNoTraffic"
	default:
		return fmt.Sprintf("BluetoothL2CAPQoSType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPSegmentationAndReassembly
type BluetoothL2CAPSegmentationAndReassembly uint32

const (
	KBluetoothL2CAPSegmentationAndReassemblyContinuationOfSDU BluetoothL2CAPSegmentationAndReassembly = 0x3
	KBluetoothL2CAPSegmentationAndReassemblyEndOfSDU          BluetoothL2CAPSegmentationAndReassembly = 0x2
	KBluetoothL2CAPSegmentationAndReassemblyStartOfSDU        BluetoothL2CAPSegmentationAndReassembly = 0x1
	KBluetoothL2CAPSegmentationAndReassemblyUnsegmentedSDU    BluetoothL2CAPSegmentationAndReassembly = 0
)

func (e BluetoothL2CAPSegmentationAndReassembly) String() string {
	switch e {
	case KBluetoothL2CAPSegmentationAndReassemblyContinuationOfSDU:
		return "KBluetoothL2CAPSegmentationAndReassemblyContinuationOfSDU"
	case KBluetoothL2CAPSegmentationAndReassemblyEndOfSDU:
		return "KBluetoothL2CAPSegmentationAndReassemblyEndOfSDU"
	case KBluetoothL2CAPSegmentationAndReassemblyStartOfSDU:
		return "KBluetoothL2CAPSegmentationAndReassemblyStartOfSDU"
	case KBluetoothL2CAPSegmentationAndReassemblyUnsegmentedSDU:
		return "KBluetoothL2CAPSegmentationAndReassemblyUnsegmentedSDU"
	default:
		return fmt.Sprintf("BluetoothL2CAPSegmentationAndReassembly(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPSupervisoryFuctionType
type BluetoothL2CAPSupervisoryFuctionType uint32

const (
	KBluetoothL2CAPSupervisoryFuctionTypeReceiverNotReady BluetoothL2CAPSupervisoryFuctionType = 0x2
	KBluetoothL2CAPSupervisoryFuctionTypeReceiverReady    BluetoothL2CAPSupervisoryFuctionType = 0
	KBluetoothL2CAPSupervisoryFuctionTypeReject           BluetoothL2CAPSupervisoryFuctionType = 0x1
	KBluetoothL2CAPSupervisoryFuctionTypeSelectiveReject  BluetoothL2CAPSupervisoryFuctionType = 0x3
)

func (e BluetoothL2CAPSupervisoryFuctionType) String() string {
	switch e {
	case KBluetoothL2CAPSupervisoryFuctionTypeReceiverNotReady:
		return "KBluetoothL2CAPSupervisoryFuctionTypeReceiverNotReady"
	case KBluetoothL2CAPSupervisoryFuctionTypeReceiverReady:
		return "KBluetoothL2CAPSupervisoryFuctionTypeReceiverReady"
	case KBluetoothL2CAPSupervisoryFuctionTypeReject:
		return "KBluetoothL2CAPSupervisoryFuctionTypeReject"
	case KBluetoothL2CAPSupervisoryFuctionTypeSelectiveReject:
		return "KBluetoothL2CAPSupervisoryFuctionTypeSelectiveReject"
	default:
		return fmt.Sprintf("BluetoothL2CAPSupervisoryFuctionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEAddressType
type BluetoothLEAddressType uint32

const (
	BluetoothLEAddressTypePublic BluetoothLEAddressType = 0
	BluetoothLEAddressTypeRandom BluetoothLEAddressType = 0x1
)

func (e BluetoothLEAddressType) String() string {
	switch e {
	case BluetoothLEAddressTypePublic:
		return "BluetoothLEAddressTypePublic"
	case BluetoothLEAddressTypeRandom:
		return "BluetoothLEAddressTypeRandom"
	default:
		return fmt.Sprintf("BluetoothLEAddressType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEAdvertisingType
type BluetoothLEAdvertisingType uint32

const (
	BluetoothLEAdvertisingTypeConnectableDirected      BluetoothLEAdvertisingType = 0x1
	BluetoothLEAdvertisingTypeConnectableUndirected    BluetoothLEAdvertisingType = 0
	BluetoothLEAdvertisingTypeDiscoverableUndirected   BluetoothLEAdvertisingType = 0x2
	BluetoothLEAdvertisingTypeNonConnectableUndirected BluetoothLEAdvertisingType = 0x3
	BluetoothLEAdvertisingTypeScanResponse             BluetoothLEAdvertisingType = 0x4
)

func (e BluetoothLEAdvertisingType) String() string {
	switch e {
	case BluetoothLEAdvertisingTypeConnectableDirected:
		return "BluetoothLEAdvertisingTypeConnectableDirected"
	case BluetoothLEAdvertisingTypeConnectableUndirected:
		return "BluetoothLEAdvertisingTypeConnectableUndirected"
	case BluetoothLEAdvertisingTypeDiscoverableUndirected:
		return "BluetoothLEAdvertisingTypeDiscoverableUndirected"
	case BluetoothLEAdvertisingTypeNonConnectableUndirected:
		return "BluetoothLEAdvertisingTypeNonConnectableUndirected"
	case BluetoothLEAdvertisingTypeScanResponse:
		return "BluetoothLEAdvertisingTypeScanResponse"
	default:
		return fmt.Sprintf("BluetoothLEAdvertisingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEConnectionInterval
type BluetoothLEConnectionInterval uint32

const (
	BluetoothLEConnectionIntervalMax BluetoothLEConnectionInterval = 0xc80
	BluetoothLEConnectionIntervalMin BluetoothLEConnectionInterval = 0x6
)

func (e BluetoothLEConnectionInterval) String() string {
	switch e {
	case BluetoothLEConnectionIntervalMax:
		return "BluetoothLEConnectionIntervalMax"
	case BluetoothLEConnectionIntervalMin:
		return "BluetoothLEConnectionIntervalMin"
	default:
		return fmt.Sprintf("BluetoothLEConnectionInterval(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEFeatureBits
type BluetoothLEFeatureBits uint32

const (
	KBluetoothLEFeatureConnectionParamsRequestProcedure    BluetoothLEFeatureBits = 2
	KBluetoothLEFeatureExtendedRejectIndication            BluetoothLEFeatureBits = 4
	KBluetoothLEFeatureExtendedScannerFilterPolicies       BluetoothLEFeatureBits = 128
	KBluetoothLEFeatureLEDataPacketLengthExtension         BluetoothLEFeatureBits = 32
	KBluetoothLEFeatureLEEncryption                        BluetoothLEFeatureBits = 1
	KBluetoothLEFeatureLEPing                              BluetoothLEFeatureBits = 16
	KBluetoothLEFeatureLLPrivacy                           BluetoothLEFeatureBits = 64
	KBluetoothLEFeaturePeripheralInitiatedFeaturesExchange BluetoothLEFeatureBits = 8
	KBluetoothLEFeatureSlaveInitiatedFeaturesExchange      BluetoothLEFeatureBits = 8
)

func (e BluetoothLEFeatureBits) String() string {
	switch e {
	case KBluetoothLEFeatureConnectionParamsRequestProcedure:
		return "KBluetoothLEFeatureConnectionParamsRequestProcedure"
	case KBluetoothLEFeatureExtendedRejectIndication:
		return "KBluetoothLEFeatureExtendedRejectIndication"
	case KBluetoothLEFeatureExtendedScannerFilterPolicies:
		return "KBluetoothLEFeatureExtendedScannerFilterPolicies"
	case KBluetoothLEFeatureLEDataPacketLengthExtension:
		return "KBluetoothLEFeatureLEDataPacketLengthExtension"
	case KBluetoothLEFeatureLEEncryption:
		return "KBluetoothLEFeatureLEEncryption"
	case KBluetoothLEFeatureLEPing:
		return "KBluetoothLEFeatureLEPing"
	case KBluetoothLEFeatureLLPrivacy:
		return "KBluetoothLEFeatureLLPrivacy"
	case KBluetoothLEFeaturePeripheralInitiatedFeaturesExchange:
		return "KBluetoothLEFeaturePeripheralInitiatedFeaturesExchange"
	default:
		return fmt.Sprintf("BluetoothLEFeatureBits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEScan
type BluetoothLEScan uint32

const (
	BluetoothLEScanDisable BluetoothLEScan = 0
	BluetoothLEScanEnable  BluetoothLEScan = 0x1
)

func (e BluetoothLEScan) String() string {
	switch e {
	case BluetoothLEScanDisable:
		return "BluetoothLEScanDisable"
	case BluetoothLEScanEnable:
		return "BluetoothLEScanEnable"
	default:
		return fmt.Sprintf("BluetoothLEScan(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEScanDuplicateFilter
type BluetoothLEScanDuplicateFilter uint32

const (
	BluetoothLEScanDuplicateFilterDisable BluetoothLEScanDuplicateFilter = 0
	BluetoothLEScanDuplicateFilterEnable  BluetoothLEScanDuplicateFilter = 0x1
)

func (e BluetoothLEScanDuplicateFilter) String() string {
	switch e {
	case BluetoothLEScanDuplicateFilterDisable:
		return "BluetoothLEScanDuplicateFilterDisable"
	case BluetoothLEScanDuplicateFilterEnable:
		return "BluetoothLEScanDuplicateFilterEnable"
	default:
		return fmt.Sprintf("BluetoothLEScanDuplicateFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEScanFilter
type BluetoothLEScanFilter uint32

const (
	BluetoothLEScanFilterNone      BluetoothLEScanFilter = 0
	BluetoothLEScanFilterSafelist  BluetoothLEScanFilter = 0x1
	BluetoothLEScanFilterWhitelist BluetoothLEScanFilter = 1
)

func (e BluetoothLEScanFilter) String() string {
	switch e {
	case BluetoothLEScanFilterNone:
		return "BluetoothLEScanFilterNone"
	case BluetoothLEScanFilterSafelist:
		return "BluetoothLEScanFilterSafelist"
	default:
		return fmt.Sprintf("BluetoothLEScanFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLEScanType
type BluetoothLEScanType uint32

const (
	BluetoothLEScanTypeActive  BluetoothLEScanType = 0x1
	BluetoothLEScanTypePassive BluetoothLEScanType = 0
)

func (e BluetoothLEScanType) String() string {
	switch e {
	case BluetoothLEScanTypeActive:
		return "BluetoothLEScanTypeActive"
	case BluetoothLEScanTypePassive:
		return "BluetoothLEScanTypePassive"
	default:
		return fmt.Sprintf("BluetoothLEScanType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerCommandCode
type BluetoothLESecurityManagerCommandCode uint32

const (
	KBluetoothLESecurityManagerCommandCodeEncryptionInfo              BluetoothLESecurityManagerCommandCode = 0x6
	KBluetoothLESecurityManagerCommandCodeIdentityAddressInfo         BluetoothLESecurityManagerCommandCode = 0x9
	KBluetoothLESecurityManagerCommandCodeIdentityInfo                BluetoothLESecurityManagerCommandCode = 0x8
	KBluetoothLESecurityManagerCommandCodeMasterIdentification        BluetoothLESecurityManagerCommandCode = 0x7
	KBluetoothLESecurityManagerCommandCodePairingConfirm              BluetoothLESecurityManagerCommandCode = 0x3
	KBluetoothLESecurityManagerCommandCodePairingDHKeyCheck           BluetoothLESecurityManagerCommandCode = 0xd
	KBluetoothLESecurityManagerCommandCodePairingFailed               BluetoothLESecurityManagerCommandCode = 0x5
	KBluetoothLESecurityManagerCommandCodePairingKeypressNotification BluetoothLESecurityManagerCommandCode = 0xe
	KBluetoothLESecurityManagerCommandCodePairingPublicKey            BluetoothLESecurityManagerCommandCode = 0xc
	KBluetoothLESecurityManagerCommandCodePairingRandom               BluetoothLESecurityManagerCommandCode = 0x4
	KBluetoothLESecurityManagerCommandCodePairingRequest              BluetoothLESecurityManagerCommandCode = 0x1
	KBluetoothLESecurityManagerCommandCodePairingResponse             BluetoothLESecurityManagerCommandCode = 0x2
	KBluetoothLESecurityManagerCommandCodeReserved                    BluetoothLESecurityManagerCommandCode = 0
	KBluetoothLESecurityManagerCommandCodeReservedEnd                 BluetoothLESecurityManagerCommandCode = 0xff
	KBluetoothLESecurityManagerCommandCodeReservedStart               BluetoothLESecurityManagerCommandCode = 0xf
	KBluetoothLESecurityManagerCommandCodeSecurityRequest             BluetoothLESecurityManagerCommandCode = 0xb
	KBluetoothLESecurityManagerCommandCodeSigningInfo                 BluetoothLESecurityManagerCommandCode = 0xa
)

func (e BluetoothLESecurityManagerCommandCode) String() string {
	switch e {
	case KBluetoothLESecurityManagerCommandCodeEncryptionInfo:
		return "KBluetoothLESecurityManagerCommandCodeEncryptionInfo"
	case KBluetoothLESecurityManagerCommandCodeIdentityAddressInfo:
		return "KBluetoothLESecurityManagerCommandCodeIdentityAddressInfo"
	case KBluetoothLESecurityManagerCommandCodeIdentityInfo:
		return "KBluetoothLESecurityManagerCommandCodeIdentityInfo"
	case KBluetoothLESecurityManagerCommandCodeMasterIdentification:
		return "KBluetoothLESecurityManagerCommandCodeMasterIdentification"
	case KBluetoothLESecurityManagerCommandCodePairingConfirm:
		return "KBluetoothLESecurityManagerCommandCodePairingConfirm"
	case KBluetoothLESecurityManagerCommandCodePairingDHKeyCheck:
		return "KBluetoothLESecurityManagerCommandCodePairingDHKeyCheck"
	case KBluetoothLESecurityManagerCommandCodePairingFailed:
		return "KBluetoothLESecurityManagerCommandCodePairingFailed"
	case KBluetoothLESecurityManagerCommandCodePairingKeypressNotification:
		return "KBluetoothLESecurityManagerCommandCodePairingKeypressNotification"
	case KBluetoothLESecurityManagerCommandCodePairingPublicKey:
		return "KBluetoothLESecurityManagerCommandCodePairingPublicKey"
	case KBluetoothLESecurityManagerCommandCodePairingRandom:
		return "KBluetoothLESecurityManagerCommandCodePairingRandom"
	case KBluetoothLESecurityManagerCommandCodePairingRequest:
		return "KBluetoothLESecurityManagerCommandCodePairingRequest"
	case KBluetoothLESecurityManagerCommandCodePairingResponse:
		return "KBluetoothLESecurityManagerCommandCodePairingResponse"
	case KBluetoothLESecurityManagerCommandCodeReserved:
		return "KBluetoothLESecurityManagerCommandCodeReserved"
	case KBluetoothLESecurityManagerCommandCodeReservedEnd:
		return "KBluetoothLESecurityManagerCommandCodeReservedEnd"
	case KBluetoothLESecurityManagerCommandCodeReservedStart:
		return "KBluetoothLESecurityManagerCommandCodeReservedStart"
	case KBluetoothLESecurityManagerCommandCodeSecurityRequest:
		return "KBluetoothLESecurityManagerCommandCodeSecurityRequest"
	case KBluetoothLESecurityManagerCommandCodeSigningInfo:
		return "KBluetoothLESecurityManagerCommandCodeSigningInfo"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerCommandCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerIOCapability
type BluetoothLESecurityManagerIOCapability uint32

const (
	KBluetoothLESecurityManagerIOCapabilityDisplayOnly     BluetoothLESecurityManagerIOCapability = 0
	KBluetoothLESecurityManagerIOCapabilityDisplayYesNo    BluetoothLESecurityManagerIOCapability = 0x1
	KBluetoothLESecurityManagerIOCapabilityKeyboardDisplay BluetoothLESecurityManagerIOCapability = 0x4
	KBluetoothLESecurityManagerIOCapabilityKeyboardOnly    BluetoothLESecurityManagerIOCapability = 0x2
	KBluetoothLESecurityManagerIOCapabilityNoInputNoOutput BluetoothLESecurityManagerIOCapability = 0x3
	KBluetoothLESecurityManagerIOCapabilityReservedEnd     BluetoothLESecurityManagerIOCapability = 0xff
	KBluetoothLESecurityManagerIOCapabilityReservedStart   BluetoothLESecurityManagerIOCapability = 0x5
)

func (e BluetoothLESecurityManagerIOCapability) String() string {
	switch e {
	case KBluetoothLESecurityManagerIOCapabilityDisplayOnly:
		return "KBluetoothLESecurityManagerIOCapabilityDisplayOnly"
	case KBluetoothLESecurityManagerIOCapabilityDisplayYesNo:
		return "KBluetoothLESecurityManagerIOCapabilityDisplayYesNo"
	case KBluetoothLESecurityManagerIOCapabilityKeyboardDisplay:
		return "KBluetoothLESecurityManagerIOCapabilityKeyboardDisplay"
	case KBluetoothLESecurityManagerIOCapabilityKeyboardOnly:
		return "KBluetoothLESecurityManagerIOCapabilityKeyboardOnly"
	case KBluetoothLESecurityManagerIOCapabilityNoInputNoOutput:
		return "KBluetoothLESecurityManagerIOCapabilityNoInputNoOutput"
	case KBluetoothLESecurityManagerIOCapabilityReservedEnd:
		return "KBluetoothLESecurityManagerIOCapabilityReservedEnd"
	case KBluetoothLESecurityManagerIOCapabilityReservedStart:
		return "KBluetoothLESecurityManagerIOCapabilityReservedStart"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerIOCapability(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerKeyDistributionFormat
type BluetoothLESecurityManagerKeyDistributionFormat uint32

const (
	KBluetoothLESecurityManagerEncryptionKey BluetoothLESecurityManagerKeyDistributionFormat = 1
	KBluetoothLESecurityManagerIDKey         BluetoothLESecurityManagerKeyDistributionFormat = 2
	KBluetoothLESecurityManagerLinkKey       BluetoothLESecurityManagerKeyDistributionFormat = 8
	KBluetoothLESecurityManagerSignKey       BluetoothLESecurityManagerKeyDistributionFormat = 4
)

func (e BluetoothLESecurityManagerKeyDistributionFormat) String() string {
	switch e {
	case KBluetoothLESecurityManagerEncryptionKey:
		return "KBluetoothLESecurityManagerEncryptionKey"
	case KBluetoothLESecurityManagerIDKey:
		return "KBluetoothLESecurityManagerIDKey"
	case KBluetoothLESecurityManagerLinkKey:
		return "KBluetoothLESecurityManagerLinkKey"
	case KBluetoothLESecurityManagerSignKey:
		return "KBluetoothLESecurityManagerSignKey"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerKeyDistributionFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerKeypressNotificationType
type BluetoothLESecurityManagerKeypressNotificationType uint32

const (
	KBluetoothLESecurityManagerNotificationTypePasskeyCleared        BluetoothLESecurityManagerKeypressNotificationType = 3
	KBluetoothLESecurityManagerNotificationTypePasskeyDigitEntered   BluetoothLESecurityManagerKeypressNotificationType = 1
	KBluetoothLESecurityManagerNotificationTypePasskeyDigitErased    BluetoothLESecurityManagerKeypressNotificationType = 2
	KBluetoothLESecurityManagerNotificationTypePasskeyEntryCompleted BluetoothLESecurityManagerKeypressNotificationType = 4
	KBluetoothLESecurityManagerNotificationTypePasskeyEntryStarted   BluetoothLESecurityManagerKeypressNotificationType = 0
	KBluetoothLESecurityManagerNotificationTypeReservedEnd           BluetoothLESecurityManagerKeypressNotificationType = 255
	KBluetoothLESecurityManagerNotificationTypeReservedStart         BluetoothLESecurityManagerKeypressNotificationType = 5
)

func (e BluetoothLESecurityManagerKeypressNotificationType) String() string {
	switch e {
	case KBluetoothLESecurityManagerNotificationTypePasskeyCleared:
		return "KBluetoothLESecurityManagerNotificationTypePasskeyCleared"
	case KBluetoothLESecurityManagerNotificationTypePasskeyDigitEntered:
		return "KBluetoothLESecurityManagerNotificationTypePasskeyDigitEntered"
	case KBluetoothLESecurityManagerNotificationTypePasskeyDigitErased:
		return "KBluetoothLESecurityManagerNotificationTypePasskeyDigitErased"
	case KBluetoothLESecurityManagerNotificationTypePasskeyEntryCompleted:
		return "KBluetoothLESecurityManagerNotificationTypePasskeyEntryCompleted"
	case KBluetoothLESecurityManagerNotificationTypePasskeyEntryStarted:
		return "KBluetoothLESecurityManagerNotificationTypePasskeyEntryStarted"
	case KBluetoothLESecurityManagerNotificationTypeReservedEnd:
		return "KBluetoothLESecurityManagerNotificationTypeReservedEnd"
	case KBluetoothLESecurityManagerNotificationTypeReservedStart:
		return "KBluetoothLESecurityManagerNotificationTypeReservedStart"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerKeypressNotificationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerOOBData
type BluetoothLESecurityManagerOOBData uint32

const (
	KBluetoothLESecurityManagerOOBAuthenticationDataNotPresent BluetoothLESecurityManagerOOBData = 0
	KBluetoothLESecurityManagerOOBAuthenticationDataPresent    BluetoothLESecurityManagerOOBData = 0x1
	KBluetoothLESecurityManagerOOBDataReservedEnd              BluetoothLESecurityManagerOOBData = 0xff
	KBluetoothLESecurityManagerOOBDataReservedStart            BluetoothLESecurityManagerOOBData = 0x2
)

func (e BluetoothLESecurityManagerOOBData) String() string {
	switch e {
	case KBluetoothLESecurityManagerOOBAuthenticationDataNotPresent:
		return "KBluetoothLESecurityManagerOOBAuthenticationDataNotPresent"
	case KBluetoothLESecurityManagerOOBAuthenticationDataPresent:
		return "KBluetoothLESecurityManagerOOBAuthenticationDataPresent"
	case KBluetoothLESecurityManagerOOBDataReservedEnd:
		return "KBluetoothLESecurityManagerOOBDataReservedEnd"
	case KBluetoothLESecurityManagerOOBDataReservedStart:
		return "KBluetoothLESecurityManagerOOBDataReservedStart"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerOOBData(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerPairingFailedReasonCode
type BluetoothLESecurityManagerPairingFailedReasonCode uint32

const (
	KBluetoothLESecurityManagerReasonCodeAuthenticationRequirements                      BluetoothLESecurityManagerPairingFailedReasonCode = 0x3
	KBluetoothLESecurityManagerReasonCodeBREDRPairingInProgress                          BluetoothLESecurityManagerPairingFailedReasonCode = 0xd
	KBluetoothLESecurityManagerReasonCodeCommandNotSupported                             BluetoothLESecurityManagerPairingFailedReasonCode = 0x7
	KBluetoothLESecurityManagerReasonCodeConfirmValueFailed                              BluetoothLESecurityManagerPairingFailedReasonCode = 0x4
	KBluetoothLESecurityManagerReasonCodeCrossTransportKeyDerivationGenerationNotAllowed BluetoothLESecurityManagerPairingFailedReasonCode = 0xe
	KBluetoothLESecurityManagerReasonCodeDHKeyCheckFailed                                BluetoothLESecurityManagerPairingFailedReasonCode = 0xb
	KBluetoothLESecurityManagerReasonCodeEncryptionKeySize                               BluetoothLESecurityManagerPairingFailedReasonCode = 0x6
	KBluetoothLESecurityManagerReasonCodeInvalidParameters                               BluetoothLESecurityManagerPairingFailedReasonCode = 0xa
	KBluetoothLESecurityManagerReasonCodeNumericComparisonFailed                         BluetoothLESecurityManagerPairingFailedReasonCode = 0xc
	KBluetoothLESecurityManagerReasonCodeOOBNotAvailbale                                 BluetoothLESecurityManagerPairingFailedReasonCode = 0x2
	KBluetoothLESecurityManagerReasonCodePairingNotSupported                             BluetoothLESecurityManagerPairingFailedReasonCode = 0x5
	KBluetoothLESecurityManagerReasonCodePasskeyEntryFailed                              BluetoothLESecurityManagerPairingFailedReasonCode = 0x1
	KBluetoothLESecurityManagerReasonCodeRepeatedAttempts                                BluetoothLESecurityManagerPairingFailedReasonCode = 0x9
	KBluetoothLESecurityManagerReasonCodeReserved                                        BluetoothLESecurityManagerPairingFailedReasonCode = 0
	KBluetoothLESecurityManagerReasonCodeReservedEnd                                     BluetoothLESecurityManagerPairingFailedReasonCode = 0xff
	KBluetoothLESecurityManagerReasonCodeReservedStart                                   BluetoothLESecurityManagerPairingFailedReasonCode = 0xf
	KBluetoothLESecurityManagerReasonCodeUnspecifiedReason                               BluetoothLESecurityManagerPairingFailedReasonCode = 0x8
)

func (e BluetoothLESecurityManagerPairingFailedReasonCode) String() string {
	switch e {
	case KBluetoothLESecurityManagerReasonCodeAuthenticationRequirements:
		return "KBluetoothLESecurityManagerReasonCodeAuthenticationRequirements"
	case KBluetoothLESecurityManagerReasonCodeBREDRPairingInProgress:
		return "KBluetoothLESecurityManagerReasonCodeBREDRPairingInProgress"
	case KBluetoothLESecurityManagerReasonCodeCommandNotSupported:
		return "KBluetoothLESecurityManagerReasonCodeCommandNotSupported"
	case KBluetoothLESecurityManagerReasonCodeConfirmValueFailed:
		return "KBluetoothLESecurityManagerReasonCodeConfirmValueFailed"
	case KBluetoothLESecurityManagerReasonCodeCrossTransportKeyDerivationGenerationNotAllowed:
		return "KBluetoothLESecurityManagerReasonCodeCrossTransportKeyDerivationGenerationNotAllowed"
	case KBluetoothLESecurityManagerReasonCodeDHKeyCheckFailed:
		return "KBluetoothLESecurityManagerReasonCodeDHKeyCheckFailed"
	case KBluetoothLESecurityManagerReasonCodeEncryptionKeySize:
		return "KBluetoothLESecurityManagerReasonCodeEncryptionKeySize"
	case KBluetoothLESecurityManagerReasonCodeInvalidParameters:
		return "KBluetoothLESecurityManagerReasonCodeInvalidParameters"
	case KBluetoothLESecurityManagerReasonCodeNumericComparisonFailed:
		return "KBluetoothLESecurityManagerReasonCodeNumericComparisonFailed"
	case KBluetoothLESecurityManagerReasonCodeOOBNotAvailbale:
		return "KBluetoothLESecurityManagerReasonCodeOOBNotAvailbale"
	case KBluetoothLESecurityManagerReasonCodePairingNotSupported:
		return "KBluetoothLESecurityManagerReasonCodePairingNotSupported"
	case KBluetoothLESecurityManagerReasonCodePasskeyEntryFailed:
		return "KBluetoothLESecurityManagerReasonCodePasskeyEntryFailed"
	case KBluetoothLESecurityManagerReasonCodeRepeatedAttempts:
		return "KBluetoothLESecurityManagerReasonCodeRepeatedAttempts"
	case KBluetoothLESecurityManagerReasonCodeReserved:
		return "KBluetoothLESecurityManagerReasonCodeReserved"
	case KBluetoothLESecurityManagerReasonCodeReservedEnd:
		return "KBluetoothLESecurityManagerReasonCodeReservedEnd"
	case KBluetoothLESecurityManagerReasonCodeReservedStart:
		return "KBluetoothLESecurityManagerReasonCodeReservedStart"
	case KBluetoothLESecurityManagerReasonCodeUnspecifiedReason:
		return "KBluetoothLESecurityManagerReasonCodeUnspecifiedReason"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerPairingFailedReasonCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerUserInputCapability
type BluetoothLESecurityManagerUserInputCapability uint32

const (
	KBluetoothLESecurityManagerUserInputCapabilityKeyboard BluetoothLESecurityManagerUserInputCapability = 0x3
	KBluetoothLESecurityManagerUserInputCapabilityNoInput  BluetoothLESecurityManagerUserInputCapability = 0x1
	KBluetoothLESecurityManagerUserInputCapabilityYesNo    BluetoothLESecurityManagerUserInputCapability = 0x2
)

func (e BluetoothLESecurityManagerUserInputCapability) String() string {
	switch e {
	case KBluetoothLESecurityManagerUserInputCapabilityKeyboard:
		return "KBluetoothLESecurityManagerUserInputCapabilityKeyboard"
	case KBluetoothLESecurityManagerUserInputCapabilityNoInput:
		return "KBluetoothLESecurityManagerUserInputCapabilityNoInput"
	case KBluetoothLESecurityManagerUserInputCapabilityYesNo:
		return "KBluetoothLESecurityManagerUserInputCapabilityYesNo"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerUserInputCapability(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLESecurityManagerUserOutputCapability
type BluetoothLESecurityManagerUserOutputCapability uint32

const (
	KBluetoothLESecurityManagerUserOutputCapabilityNoOutput      BluetoothLESecurityManagerUserOutputCapability = 0x1
	KBluetoothLESecurityManagerUserOutputCapabilityNumericOutput BluetoothLESecurityManagerUserOutputCapability = 0x2
)

func (e BluetoothLESecurityManagerUserOutputCapability) String() string {
	switch e {
	case KBluetoothLESecurityManagerUserOutputCapabilityNoOutput:
		return "KBluetoothLESecurityManagerUserOutputCapabilityNoOutput"
	case KBluetoothLESecurityManagerUserOutputCapabilityNumericOutput:
		return "KBluetoothLESecurityManagerUserOutputCapabilityNumericOutput"
	default:
		return fmt.Sprintf("BluetoothLESecurityManagerUserOutputCapability(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLMPVersions
type BluetoothLMPVersions uint32

const (
	KBluetoothLMPVersionCoreSpecification1_0b   BluetoothLMPVersions = 0
	KBluetoothLMPVersionCoreSpecification1_1    BluetoothLMPVersions = 0x1
	KBluetoothLMPVersionCoreSpecification1_2    BluetoothLMPVersions = 0x2
	KBluetoothLMPVersionCoreSpecification2_0EDR BluetoothLMPVersions = 0x3
	KBluetoothLMPVersionCoreSpecification2_1EDR BluetoothLMPVersions = 0x4
	KBluetoothLMPVersionCoreSpecification3_0HS  BluetoothLMPVersions = 0x5
	KBluetoothLMPVersionCoreSpecification4_0    BluetoothLMPVersions = 0x6
	KBluetoothLMPVersionCoreSpecification4_1    BluetoothLMPVersions = 0x7
	KBluetoothLMPVersionCoreSpecification4_2    BluetoothLMPVersions = 0x8
	KBluetoothLMPVersionCoreSpecification5_0    BluetoothLMPVersions = 0x9
	KBluetoothLMPVersionCoreSpecification5_1    BluetoothLMPVersions = 0xa
	KBluetoothLMPVersionCoreSpecification5_2    BluetoothLMPVersions = 0xb
)

func (e BluetoothLMPVersions) String() string {
	switch e {
	case KBluetoothLMPVersionCoreSpecification1_0b:
		return "KBluetoothLMPVersionCoreSpecification1_0b"
	case KBluetoothLMPVersionCoreSpecification1_1:
		return "KBluetoothLMPVersionCoreSpecification1_1"
	case KBluetoothLMPVersionCoreSpecification1_2:
		return "KBluetoothLMPVersionCoreSpecification1_2"
	case KBluetoothLMPVersionCoreSpecification2_0EDR:
		return "KBluetoothLMPVersionCoreSpecification2_0EDR"
	case KBluetoothLMPVersionCoreSpecification2_1EDR:
		return "KBluetoothLMPVersionCoreSpecification2_1EDR"
	case KBluetoothLMPVersionCoreSpecification3_0HS:
		return "KBluetoothLMPVersionCoreSpecification3_0HS"
	case KBluetoothLMPVersionCoreSpecification4_0:
		return "KBluetoothLMPVersionCoreSpecification4_0"
	case KBluetoothLMPVersionCoreSpecification4_1:
		return "KBluetoothLMPVersionCoreSpecification4_1"
	case KBluetoothLMPVersionCoreSpecification4_2:
		return "KBluetoothLMPVersionCoreSpecification4_2"
	case KBluetoothLMPVersionCoreSpecification5_0:
		return "KBluetoothLMPVersionCoreSpecification5_0"
	case KBluetoothLMPVersionCoreSpecification5_1:
		return "KBluetoothLMPVersionCoreSpecification5_1"
	case KBluetoothLMPVersionCoreSpecification5_2:
		return "KBluetoothLMPVersionCoreSpecification5_2"
	default:
		return fmt.Sprintf("BluetoothLMPVersions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothLinkTypes
type BluetoothLinkTypes uint32

const (
	KBluetoothACLConnection  BluetoothLinkTypes = 1
	KBluetoothESCOConnection BluetoothLinkTypes = 2
	KBluetoothLinkTypeNone   BluetoothLinkTypes = 0xff
	KBluetoothSCOConnection  BluetoothLinkTypes = 0
)

func (e BluetoothLinkTypes) String() string {
	switch e {
	case KBluetoothACLConnection:
		return "KBluetoothACLConnection"
	case KBluetoothESCOConnection:
		return "KBluetoothESCOConnection"
	case KBluetoothLinkTypeNone:
		return "KBluetoothLinkTypeNone"
	case KBluetoothSCOConnection:
		return "KBluetoothSCOConnection"
	default:
		return fmt.Sprintf("BluetoothLinkTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothOOBDataPresenceValues
type BluetoothOOBDataPresenceValues uint32

const (
	KBluetoothOOBAuthenticationDataFromRemoteDevicePresent BluetoothOOBDataPresenceValues = 0x1
	KBluetoothOOBAuthenticationDataNotPresent              BluetoothOOBDataPresenceValues = 0
)

func (e BluetoothOOBDataPresenceValues) String() string {
	switch e {
	case KBluetoothOOBAuthenticationDataFromRemoteDevicePresent:
		return "KBluetoothOOBAuthenticationDataFromRemoteDevicePresent"
	case KBluetoothOOBAuthenticationDataNotPresent:
		return "KBluetoothOOBAuthenticationDataNotPresent"
	default:
		return fmt.Sprintf("BluetoothOOBDataPresenceValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothRFCOMMLineStatus
type BluetoothRFCOMMLineStatus uint32

const (
	BluetoothRFCOMMLineStatusFramingError BluetoothRFCOMMLineStatus = 3
	BluetoothRFCOMMLineStatusNoError      BluetoothRFCOMMLineStatus = 0
	BluetoothRFCOMMLineStatusOverrunError BluetoothRFCOMMLineStatus = 1
	BluetoothRFCOMMLineStatusParityError  BluetoothRFCOMMLineStatus = 2
)

func (e BluetoothRFCOMMLineStatus) String() string {
	switch e {
	case BluetoothRFCOMMLineStatusFramingError:
		return "BluetoothRFCOMMLineStatusFramingError"
	case BluetoothRFCOMMLineStatusNoError:
		return "BluetoothRFCOMMLineStatusNoError"
	case BluetoothRFCOMMLineStatusOverrunError:
		return "BluetoothRFCOMMLineStatusOverrunError"
	case BluetoothRFCOMMLineStatusParityError:
		return "BluetoothRFCOMMLineStatusParityError"
	default:
		return fmt.Sprintf("BluetoothRFCOMMLineStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothRFCOMMParityType
type BluetoothRFCOMMParityType uint32

const (
	KBluetoothRFCOMMParityTypeEvenParity BluetoothRFCOMMParityType = 2
	KBluetoothRFCOMMParityTypeMaxParity  BluetoothRFCOMMParityType = 3
	KBluetoothRFCOMMParityTypeNoParity   BluetoothRFCOMMParityType = 0
	KBluetoothRFCOMMParityTypeOddParity  BluetoothRFCOMMParityType = 1
)

func (e BluetoothRFCOMMParityType) String() string {
	switch e {
	case KBluetoothRFCOMMParityTypeEvenParity:
		return "KBluetoothRFCOMMParityTypeEvenParity"
	case KBluetoothRFCOMMParityTypeMaxParity:
		return "KBluetoothRFCOMMParityTypeMaxParity"
	case KBluetoothRFCOMMParityTypeNoParity:
		return "KBluetoothRFCOMMParityTypeNoParity"
	case KBluetoothRFCOMMParityTypeOddParity:
		return "KBluetoothRFCOMMParityTypeOddParity"
	default:
		return fmt.Sprintf("BluetoothRFCOMMParityType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothSimplePairingDebugModes
type BluetoothSimplePairingDebugModes uint32

const (
	KBluetoothHCISimplePairingDebugModeDisabled BluetoothSimplePairingDebugModes = 0
	KBluetoothHCISimplePairingDebugModeEnabled  BluetoothSimplePairingDebugModes = 0x1
)

func (e BluetoothSimplePairingDebugModes) String() string {
	switch e {
	case KBluetoothHCISimplePairingDebugModeDisabled:
		return "KBluetoothHCISimplePairingDebugModeDisabled"
	case KBluetoothHCISimplePairingDebugModeEnabled:
		return "KBluetoothHCISimplePairingDebugModeEnabled"
	default:
		return fmt.Sprintf("BluetoothSimplePairingDebugModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/BluetoothTransportTypes
type BluetoothTransportTypes uint32

const (
	KBluetoothTransportTypePCCard  BluetoothTransportTypes = 0x2
	KBluetoothTransportTypePCICard BluetoothTransportTypes = 0x3
	KBluetoothTransportTypePCIe    BluetoothTransportTypes = 0x5
	KBluetoothTransportTypeUART    BluetoothTransportTypes = 0x4
	KBluetoothTransportTypeUSB     BluetoothTransportTypes = 0x1
)

func (e BluetoothTransportTypes) String() string {
	switch e {
	case KBluetoothTransportTypePCCard:
		return "KBluetoothTransportTypePCCard"
	case KBluetoothTransportTypePCICard:
		return "KBluetoothTransportTypePCICard"
	case KBluetoothTransportTypePCIe:
		return "KBluetoothTransportTypePCIe"
	case KBluetoothTransportTypeUART:
		return "KBluetoothTransportTypeUART"
	case KBluetoothTransportTypeUSB:
		return "KBluetoothTransportTypeUSB"
	default:
		return fmt.Sprintf("BluetoothTransportTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/FTSFileType
type FTSFileType uint32

const (
	KFTSFileTypeFile   FTSFileType = 2
	KFTSFileTypeFolder FTSFileType = 1
)

func (e FTSFileType) String() string {
	switch e {
	case KFTSFileTypeFile:
		return "KFTSFileTypeFile"
	case KFTSFileTypeFolder:
		return "KFTSFileTypeFolder"
	default:
		return fmt.Sprintf("FTSFileType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchOptionsBits
type IOBluetoothDeviceSearchOptionsBits uint32

const (
	KSearchOptionsAlwaysStartInquiry   IOBluetoothDeviceSearchOptionsBits = 1
	KSearchOptionsDiscardCachedResults IOBluetoothDeviceSearchOptionsBits = 2
	KSearchOptionsNone                 IOBluetoothDeviceSearchOptionsBits = 0
)

func (e IOBluetoothDeviceSearchOptionsBits) String() string {
	switch e {
	case KSearchOptionsAlwaysStartInquiry:
		return "KSearchOptionsAlwaysStartInquiry"
	case KSearchOptionsDiscardCachedResults:
		return "KSearchOptionsDiscardCachedResults"
	case KSearchOptionsNone:
		return "KSearchOptionsNone"
	default:
		return fmt.Sprintf("IOBluetoothDeviceSearchOptionsBits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchTypesBits
type IOBluetoothDeviceSearchTypesBits uint32

const (
	KIOBluetoothDeviceSearchClassic IOBluetoothDeviceSearchTypesBits = 1
	KIOBluetoothDeviceSearchLE      IOBluetoothDeviceSearchTypesBits = 2
)

func (e IOBluetoothDeviceSearchTypesBits) String() string {
	switch e {
	case KIOBluetoothDeviceSearchClassic:
		return "KIOBluetoothDeviceSearchClassic"
	case KIOBluetoothDeviceSearchLE:
		return "KIOBluetoothDeviceSearchLE"
	default:
		return fmt.Sprintf("IOBluetoothDeviceSearchTypesBits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGatewayFeatures
type IOBluetoothHandsFreeAudioGatewayFeatures uint32

const (
	IOBluetoothHandsFreeAudioGatewayFeatureAttachedNumberToVoiceTag IOBluetoothHandsFreeAudioGatewayFeatures = 16
	IOBluetoothHandsFreeAudioGatewayFeatureCodecNegotiation         IOBluetoothHandsFreeAudioGatewayFeatures = 512
	IOBluetoothHandsFreeAudioGatewayFeatureECAndOrNRFunction        IOBluetoothHandsFreeAudioGatewayFeatures = 2
	IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallControl      IOBluetoothHandsFreeAudioGatewayFeatures = 128
	IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallStatus       IOBluetoothHandsFreeAudioGatewayFeatures = 64
	IOBluetoothHandsFreeAudioGatewayFeatureExtendedErrorResultCodes IOBluetoothHandsFreeAudioGatewayFeatures = 256
	IOBluetoothHandsFreeAudioGatewayFeatureInBandRingTone           IOBluetoothHandsFreeAudioGatewayFeatures = 8
	IOBluetoothHandsFreeAudioGatewayFeatureRejectCallCapability     IOBluetoothHandsFreeAudioGatewayFeatures = 32
	IOBluetoothHandsFreeAudioGatewayFeatureThreeWayCalling          IOBluetoothHandsFreeAudioGatewayFeatures = 1
	IOBluetoothHandsFreeAudioGatewayFeatureVoiceRecognition         IOBluetoothHandsFreeAudioGatewayFeatures = 4
)

func (e IOBluetoothHandsFreeAudioGatewayFeatures) String() string {
	switch e {
	case IOBluetoothHandsFreeAudioGatewayFeatureAttachedNumberToVoiceTag:
		return "IOBluetoothHandsFreeAudioGatewayFeatureAttachedNumberToVoiceTag"
	case IOBluetoothHandsFreeAudioGatewayFeatureCodecNegotiation:
		return "IOBluetoothHandsFreeAudioGatewayFeatureCodecNegotiation"
	case IOBluetoothHandsFreeAudioGatewayFeatureECAndOrNRFunction:
		return "IOBluetoothHandsFreeAudioGatewayFeatureECAndOrNRFunction"
	case IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallControl:
		return "IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallControl"
	case IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallStatus:
		return "IOBluetoothHandsFreeAudioGatewayFeatureEnhancedCallStatus"
	case IOBluetoothHandsFreeAudioGatewayFeatureExtendedErrorResultCodes:
		return "IOBluetoothHandsFreeAudioGatewayFeatureExtendedErrorResultCodes"
	case IOBluetoothHandsFreeAudioGatewayFeatureInBandRingTone:
		return "IOBluetoothHandsFreeAudioGatewayFeatureInBandRingTone"
	case IOBluetoothHandsFreeAudioGatewayFeatureRejectCallCapability:
		return "IOBluetoothHandsFreeAudioGatewayFeatureRejectCallCapability"
	case IOBluetoothHandsFreeAudioGatewayFeatureThreeWayCalling:
		return "IOBluetoothHandsFreeAudioGatewayFeatureThreeWayCalling"
	case IOBluetoothHandsFreeAudioGatewayFeatureVoiceRecognition:
		return "IOBluetoothHandsFreeAudioGatewayFeatureVoiceRecognition"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreeAudioGatewayFeatures(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallHoldModes
type IOBluetoothHandsFreeCallHoldModes uint

const (
	IOBluetoothHandsFreeCallHoldMode0    IOBluetoothHandsFreeCallHoldModes = 1
	IOBluetoothHandsFreeCallHoldMode1    IOBluetoothHandsFreeCallHoldModes = 2
	IOBluetoothHandsFreeCallHoldMode1idx IOBluetoothHandsFreeCallHoldModes = 4
	IOBluetoothHandsFreeCallHoldMode2    IOBluetoothHandsFreeCallHoldModes = 8
	IOBluetoothHandsFreeCallHoldMode2idx IOBluetoothHandsFreeCallHoldModes = 16
	IOBluetoothHandsFreeCallHoldMode3    IOBluetoothHandsFreeCallHoldModes = 32
	IOBluetoothHandsFreeCallHoldMode4    IOBluetoothHandsFreeCallHoldModes = 64
)

func (e IOBluetoothHandsFreeCallHoldModes) String() string {
	switch e {
	case IOBluetoothHandsFreeCallHoldMode0:
		return "IOBluetoothHandsFreeCallHoldMode0"
	case IOBluetoothHandsFreeCallHoldMode1:
		return "IOBluetoothHandsFreeCallHoldMode1"
	case IOBluetoothHandsFreeCallHoldMode1idx:
		return "IOBluetoothHandsFreeCallHoldMode1idx"
	case IOBluetoothHandsFreeCallHoldMode2:
		return "IOBluetoothHandsFreeCallHoldMode2"
	case IOBluetoothHandsFreeCallHoldMode2idx:
		return "IOBluetoothHandsFreeCallHoldMode2idx"
	case IOBluetoothHandsFreeCallHoldMode3:
		return "IOBluetoothHandsFreeCallHoldMode3"
	case IOBluetoothHandsFreeCallHoldMode4:
		return "IOBluetoothHandsFreeCallHoldMode4"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreeCallHoldModes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCodecID
type IOBluetoothHandsFreeCodecID uint8

const (
	IOBluetoothHandsFreeCodecIDAACELD IOBluetoothHandsFreeCodecID = 0x80
	IOBluetoothHandsFreeCodecIDCVSD   IOBluetoothHandsFreeCodecID = 0x1
	IOBluetoothHandsFreeCodecIDmSBC   IOBluetoothHandsFreeCodecID = 0x2
)

func (e IOBluetoothHandsFreeCodecID) String() string {
	switch e {
	case IOBluetoothHandsFreeCodecIDAACELD:
		return "IOBluetoothHandsFreeCodecIDAACELD"
	case IOBluetoothHandsFreeCodecIDCVSD:
		return "IOBluetoothHandsFreeCodecIDCVSD"
	case IOBluetoothHandsFreeCodecIDmSBC:
		return "IOBluetoothHandsFreeCodecIDmSBC"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreeCodecID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceFeatures
type IOBluetoothHandsFreeDeviceFeatures uint32

const (
	IOBluetoothHandsFreeDeviceFeatureCLIPresentation     IOBluetoothHandsFreeDeviceFeatures = 4
	IOBluetoothHandsFreeDeviceFeatureCodecNegotiation    IOBluetoothHandsFreeDeviceFeatures = 128
	IOBluetoothHandsFreeDeviceFeatureECAndOrNRFunction   IOBluetoothHandsFreeDeviceFeatures = 1
	IOBluetoothHandsFreeDeviceFeatureEnhancedCallControl IOBluetoothHandsFreeDeviceFeatures = 64
	IOBluetoothHandsFreeDeviceFeatureEnhancedCallStatus  IOBluetoothHandsFreeDeviceFeatures = 32
	IOBluetoothHandsFreeDeviceFeatureRemoteVolumeControl IOBluetoothHandsFreeDeviceFeatures = 16
	IOBluetoothHandsFreeDeviceFeatureThreeWayCalling     IOBluetoothHandsFreeDeviceFeatures = 2
	IOBluetoothHandsFreeDeviceFeatureVoiceRecognition    IOBluetoothHandsFreeDeviceFeatures = 8
)

func (e IOBluetoothHandsFreeDeviceFeatures) String() string {
	switch e {
	case IOBluetoothHandsFreeDeviceFeatureCLIPresentation:
		return "IOBluetoothHandsFreeDeviceFeatureCLIPresentation"
	case IOBluetoothHandsFreeDeviceFeatureCodecNegotiation:
		return "IOBluetoothHandsFreeDeviceFeatureCodecNegotiation"
	case IOBluetoothHandsFreeDeviceFeatureECAndOrNRFunction:
		return "IOBluetoothHandsFreeDeviceFeatureECAndOrNRFunction"
	case IOBluetoothHandsFreeDeviceFeatureEnhancedCallControl:
		return "IOBluetoothHandsFreeDeviceFeatureEnhancedCallControl"
	case IOBluetoothHandsFreeDeviceFeatureEnhancedCallStatus:
		return "IOBluetoothHandsFreeDeviceFeatureEnhancedCallStatus"
	case IOBluetoothHandsFreeDeviceFeatureRemoteVolumeControl:
		return "IOBluetoothHandsFreeDeviceFeatureRemoteVolumeControl"
	case IOBluetoothHandsFreeDeviceFeatureThreeWayCalling:
		return "IOBluetoothHandsFreeDeviceFeatureThreeWayCalling"
	case IOBluetoothHandsFreeDeviceFeatureVoiceRecognition:
		return "IOBluetoothHandsFreeDeviceFeatureVoiceRecognition"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreeDeviceFeatures(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreePDUMessageStatus
type IOBluetoothHandsFreePDUMessageStatus uint

const (
	IOBluetoothHandsFreePDUStatusAll       IOBluetoothHandsFreePDUMessageStatus = 4
	IOBluetoothHandsFreePDUStatusRecRead   IOBluetoothHandsFreePDUMessageStatus = 1
	IOBluetoothHandsFreePDUStatusRecUnread IOBluetoothHandsFreePDUMessageStatus = 0
	IOBluetoothHandsFreePDUStatusStoSent   IOBluetoothHandsFreePDUMessageStatus = 3
	IOBluetoothHandsFreePDUStatusStoUnsent IOBluetoothHandsFreePDUMessageStatus = 2
)

func (e IOBluetoothHandsFreePDUMessageStatus) String() string {
	switch e {
	case IOBluetoothHandsFreePDUStatusAll:
		return "IOBluetoothHandsFreePDUStatusAll"
	case IOBluetoothHandsFreePDUStatusRecRead:
		return "IOBluetoothHandsFreePDUStatusRecRead"
	case IOBluetoothHandsFreePDUStatusRecUnread:
		return "IOBluetoothHandsFreePDUStatusRecUnread"
	case IOBluetoothHandsFreePDUStatusStoSent:
		return "IOBluetoothHandsFreePDUStatusStoSent"
	case IOBluetoothHandsFreePDUStatusStoUnsent:
		return "IOBluetoothHandsFreePDUStatusStoUnsent"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreePDUMessageStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeSMSSupport
type IOBluetoothHandsFreeSMSSupport uint

const (
	IOBluetoothHandsFreeManufactureSpecificSMSSupport IOBluetoothHandsFreeSMSSupport = 4
	IOBluetoothHandsFreePhase2SMSSupport              IOBluetoothHandsFreeSMSSupport = 1
	IOBluetoothHandsFreePhase2pSMSSupport             IOBluetoothHandsFreeSMSSupport = 2
)

func (e IOBluetoothHandsFreeSMSSupport) String() string {
	switch e {
	case IOBluetoothHandsFreeManufactureSpecificSMSSupport:
		return "IOBluetoothHandsFreeManufactureSpecificSMSSupport"
	case IOBluetoothHandsFreePhase2SMSSupport:
		return "IOBluetoothHandsFreePhase2SMSSupport"
	case IOBluetoothHandsFreePhase2pSMSSupport:
		return "IOBluetoothHandsFreePhase2pSMSSupport"
	default:
		return fmt.Sprintf("IOBluetoothHandsFreeSMSSupport(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelEventType
type IOBluetoothL2CAPChannelEventType uint32

const (
	KIOBluetoothL2CAPChannelEventTypeClosed              IOBluetoothL2CAPChannelEventType = 0x3
	KIOBluetoothL2CAPChannelEventTypeData                IOBluetoothL2CAPChannelEventType = 0x1
	KIOBluetoothL2CAPChannelEventTypeOpenComplete        IOBluetoothL2CAPChannelEventType = 0x2
	KIOBluetoothL2CAPChannelEventTypeQueueSpaceAvailable IOBluetoothL2CAPChannelEventType = 0x6
	KIOBluetoothL2CAPChannelEventTypeReconfigured        IOBluetoothL2CAPChannelEventType = 0x4
	KIOBluetoothL2CAPChannelEventTypeWriteComplete       IOBluetoothL2CAPChannelEventType = 0x5
)

func (e IOBluetoothL2CAPChannelEventType) String() string {
	switch e {
	case KIOBluetoothL2CAPChannelEventTypeClosed:
		return "KIOBluetoothL2CAPChannelEventTypeClosed"
	case KIOBluetoothL2CAPChannelEventTypeData:
		return "KIOBluetoothL2CAPChannelEventTypeData"
	case KIOBluetoothL2CAPChannelEventTypeOpenComplete:
		return "KIOBluetoothL2CAPChannelEventTypeOpenComplete"
	case KIOBluetoothL2CAPChannelEventTypeQueueSpaceAvailable:
		return "KIOBluetoothL2CAPChannelEventTypeQueueSpaceAvailable"
	case KIOBluetoothL2CAPChannelEventTypeReconfigured:
		return "KIOBluetoothL2CAPChannelEventTypeReconfigured"
	case KIOBluetoothL2CAPChannelEventTypeWriteComplete:
		return "KIOBluetoothL2CAPChannelEventTypeWriteComplete"
	default:
		return fmt.Sprintf("IOBluetoothL2CAPChannelEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSMSMode
type IOBluetoothSMSMode uint

const (
	IOBluetoothSMSModePDU  IOBluetoothSMSMode = 0
	IOBluetoothSMSModeText IOBluetoothSMSMode = 1
)

func (e IOBluetoothSMSMode) String() string {
	switch e {
	case IOBluetoothSMSModePDU:
		return "IOBluetoothSMSModePDU"
	case IOBluetoothSMSModeText:
		return "IOBluetoothSMSModeText"
	default:
		return fmt.Sprintf("IOBluetoothSMSMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotificationChannelDirection
type IOBluetoothUserNotificationChannelDirection uint32

const (
	KIOBluetoothUserNotificationChannelDirectionAny      IOBluetoothUserNotificationChannelDirection = 0
	KIOBluetoothUserNotificationChannelDirectionIncoming IOBluetoothUserNotificationChannelDirection = 1
	KIOBluetoothUserNotificationChannelDirectionOutgoing IOBluetoothUserNotificationChannelDirection = 2
)

func (e IOBluetoothUserNotificationChannelDirection) String() string {
	switch e {
	case KIOBluetoothUserNotificationChannelDirectionAny:
		return "KIOBluetoothUserNotificationChannelDirectionAny"
	case KIOBluetoothUserNotificationChannelDirectionIncoming:
		return "KIOBluetoothUserNotificationChannelDirectionIncoming"
	case KIOBluetoothUserNotificationChannelDirectionOutgoing:
		return "KIOBluetoothUserNotificationChannelDirectionOutgoing"
	default:
		return fmt.Sprintf("IOBluetoothUserNotificationChannelDirection(%d)", e)
	}
}

type KBluetoothAirMode uint32

const (
	KBluetoothAirModeALawLog         KBluetoothAirMode = 0x1
	KBluetoothAirModeCVSD            KBluetoothAirMode = 0x2
	KBluetoothAirModeTransparentData KBluetoothAirMode = 0x3
	KBluetoothAirModeULawLog         KBluetoothAirMode = 0
)

func (e KBluetoothAirMode) String() string {
	switch e {
	case KBluetoothAirModeALawLog:
		return "KBluetoothAirModeALawLog"
	case KBluetoothAirModeCVSD:
		return "KBluetoothAirModeCVSD"
	case KBluetoothAirModeTransparentData:
		return "KBluetoothAirModeTransparentData"
	case KBluetoothAirModeULawLog:
		return "KBluetoothAirModeULawLog"
	default:
		return fmt.Sprintf("KBluetoothAirMode(%d)", e)
	}
}

type KBluetoothConnectionHandle uint32

const (
	KBluetoothConnectionHandleNone                 KBluetoothConnectionHandle = 0xffff
	KBluetoothConnectionHandleSerialDeviceReserved KBluetoothConnectionHandle = 0xfff
)

func (e KBluetoothConnectionHandle) String() string {
	switch e {
	case KBluetoothConnectionHandleNone:
		return "KBluetoothConnectionHandleNone"
	case KBluetoothConnectionHandleSerialDeviceReserved:
		return "KBluetoothConnectionHandleSerialDeviceReserved"
	default:
		return fmt.Sprintf("KBluetoothConnectionHandle(%d)", e)
	}
}

type KBluetoothDeviceClassMajor uint32

const (
	KBluetoothDeviceClassMajorAny            KBluetoothDeviceClassMajor = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KBluetoothDeviceClassMajorAudio          KBluetoothDeviceClassMajor = 0x4
	KBluetoothDeviceClassMajorComputer       KBluetoothDeviceClassMajor = 0x1
	KBluetoothDeviceClassMajorEnd            KBluetoothDeviceClassMajor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'f' // 'nonf'
	KBluetoothDeviceClassMajorHealth         KBluetoothDeviceClassMajor = 0x9
	KBluetoothDeviceClassMajorImaging        KBluetoothDeviceClassMajor = 0x6
	KBluetoothDeviceClassMajorLANAccessPoint KBluetoothDeviceClassMajor = 0x3
	KBluetoothDeviceClassMajorMiscellaneous  KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorNone           KBluetoothDeviceClassMajor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'none'
	KBluetoothDeviceClassMajorPeripheral     KBluetoothDeviceClassMajor = 0x5
	KBluetoothDeviceClassMajorPhone          KBluetoothDeviceClassMajor = 0x2
	KBluetoothDeviceClassMajorToy            KBluetoothDeviceClassMajor = 0x8
	KBluetoothDeviceClassMajorUnclassified   KBluetoothDeviceClassMajor = 0x1f
	KBluetoothDeviceClassMajorWearable       KBluetoothDeviceClassMajor = 0x7
)

func (e KBluetoothDeviceClassMajor) String() string {
	switch e {
	case KBluetoothDeviceClassMajorAny:
		return "KBluetoothDeviceClassMajorAny"
	case KBluetoothDeviceClassMajorAudio:
		return "KBluetoothDeviceClassMajorAudio"
	case KBluetoothDeviceClassMajorComputer:
		return "KBluetoothDeviceClassMajorComputer"
	case KBluetoothDeviceClassMajorEnd:
		return "KBluetoothDeviceClassMajorEnd"
	case KBluetoothDeviceClassMajorHealth:
		return "KBluetoothDeviceClassMajorHealth"
	case KBluetoothDeviceClassMajorImaging:
		return "KBluetoothDeviceClassMajorImaging"
	case KBluetoothDeviceClassMajorLANAccessPoint:
		return "KBluetoothDeviceClassMajorLANAccessPoint"
	case KBluetoothDeviceClassMajorMiscellaneous:
		return "KBluetoothDeviceClassMajorMiscellaneous"
	case KBluetoothDeviceClassMajorNone:
		return "KBluetoothDeviceClassMajorNone"
	case KBluetoothDeviceClassMajorPeripheral:
		return "KBluetoothDeviceClassMajorPeripheral"
	case KBluetoothDeviceClassMajorPhone:
		return "KBluetoothDeviceClassMajorPhone"
	case KBluetoothDeviceClassMajorToy:
		return "KBluetoothDeviceClassMajorToy"
	case KBluetoothDeviceClassMajorUnclassified:
		return "KBluetoothDeviceClassMajorUnclassified"
	case KBluetoothDeviceClassMajorWearable:
		return "KBluetoothDeviceClassMajorWearable"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMajor(%d)", e)
	}
}

type KBluetoothDeviceClassMinor uint32

const (
	KBluetoothDeviceClassMinorAny                             KBluetoothDeviceClassMinor = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KBluetoothDeviceClassMinorAudioCamcorder                  KBluetoothDeviceClassMinor = 0xd
	KBluetoothDeviceClassMinorAudioCar                        KBluetoothDeviceClassMinor = 0x8
	KBluetoothDeviceClassMinorAudioGamingToy                  KBluetoothDeviceClassMinor = 0x12
	KBluetoothDeviceClassMinorAudioHandsFree                  KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorAudioHeadphones                 KBluetoothDeviceClassMinor = 0x6
	KBluetoothDeviceClassMinorAudioHeadset                    KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorAudioHiFi                       KBluetoothDeviceClassMinor = 0xa
	KBluetoothDeviceClassMinorAudioLoudspeaker                KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorAudioMicrophone                 KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorAudioPortable                   KBluetoothDeviceClassMinor = 0x7
	KBluetoothDeviceClassMinorAudioReserved1                  KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorAudioReserved2                  KBluetoothDeviceClassMinor = 0x11
	KBluetoothDeviceClassMinorAudioSetTopBox                  KBluetoothDeviceClassMinor = 0x9
	KBluetoothDeviceClassMinorAudioUnclassified               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVCR                        KBluetoothDeviceClassMinor = 0xb
	KBluetoothDeviceClassMinorAudioVideoCamera                KBluetoothDeviceClassMinor = 0xc
	KBluetoothDeviceClassMinorAudioVideoConferencing          KBluetoothDeviceClassMinor = 0x10
	KBluetoothDeviceClassMinorAudioVideoDisplayAndLoudspeaker KBluetoothDeviceClassMinor = 0xf
	KBluetoothDeviceClassMinorAudioVideoMonitor               KBluetoothDeviceClassMinor = 0xe
	KBluetoothDeviceClassMinorComputerDesktopWorkstation      KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorComputerHandheld                KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorComputerLaptop                  KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorComputerPalmSized               KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorComputerServer                  KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorComputerUnclassified            KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerWearable                KBluetoothDeviceClassMinor = 0x6
	KBluetoothDeviceClassMinorEnd                             KBluetoothDeviceClassMinor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'f' // 'nonf'
	KBluetoothDeviceClassMinorHealthBloodPressureMonitor      KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorHealthDataDisplay               KBluetoothDeviceClassMinor = 0x7
	KBluetoothDeviceClassMinorHealthGlucoseMeter              KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorHealthHeartRateMonitor          KBluetoothDeviceClassMinor = 0x6
	KBluetoothDeviceClassMinorHealthPulseOximeter             KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorHealthScale                     KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorHealthThermometer               KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorHealthUndefined                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging1Camera                  KBluetoothDeviceClassMinor = 0x8
	KBluetoothDeviceClassMinorImaging1Display                 KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorImaging1Printer                 KBluetoothDeviceClassMinor = 0x20
	KBluetoothDeviceClassMinorImaging1Scanner                 KBluetoothDeviceClassMinor = 0x10
	KBluetoothDeviceClassMinorImaging2Unclassified            KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorNone                            KBluetoothDeviceClassMinor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'none'
	KBluetoothDeviceClassMinorPeripheral1Combo                KBluetoothDeviceClassMinor = 0x30
	KBluetoothDeviceClassMinorPeripheral1Keyboard             KBluetoothDeviceClassMinor = 0x10
	KBluetoothDeviceClassMinorPeripheral1Pointing             KBluetoothDeviceClassMinor = 0x20
	KBluetoothDeviceClassMinorPeripheral2AnyPointing          KBluetoothDeviceClassMinor = 'p'<<24 | 'o'<<16 | 'i'<<8 | 'n' // 'poin'
	KBluetoothDeviceClassMinorPeripheral2CardReader           KBluetoothDeviceClassMinor = 0x6
	KBluetoothDeviceClassMinorPeripheral2DigitalPen           KBluetoothDeviceClassMinor = 0x7
	KBluetoothDeviceClassMinorPeripheral2DigitizerTablet      KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorPeripheral2Gamepad              KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorPeripheral2GesturalInputDevice  KBluetoothDeviceClassMinor = 0x9
	KBluetoothDeviceClassMinorPeripheral2HandheldScanner      KBluetoothDeviceClassMinor = 0x8
	KBluetoothDeviceClassMinorPeripheral2Joystick             KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorPeripheral2RemoteControl        KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorPeripheral2SensingDevice        KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorPeripheral2Unclassified         KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneCellular                   KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorPhoneCommonISDNAccess           KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorPhoneCordless                   KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorPhoneSmartPhone                 KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorPhoneUnclassified               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneWiredModemOrVoiceGateway   KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorToyController                   KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorToyDollActionFigure             KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorToyGame                         KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorToyRobot                        KBluetoothDeviceClassMinor = 0x1
	KBluetoothDeviceClassMinorToyVehicle                      KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorWearableGlasses                 KBluetoothDeviceClassMinor = 0x5
	KBluetoothDeviceClassMinorWearableHelmet                  KBluetoothDeviceClassMinor = 0x4
	KBluetoothDeviceClassMinorWearableJacket                  KBluetoothDeviceClassMinor = 0x3
	KBluetoothDeviceClassMinorWearablePager                   KBluetoothDeviceClassMinor = 0x2
	KBluetoothDeviceClassMinorWearableWristWatch              KBluetoothDeviceClassMinor = 0x1
)

func (e KBluetoothDeviceClassMinor) String() string {
	switch e {
	case KBluetoothDeviceClassMinorAny:
		return "KBluetoothDeviceClassMinorAny"
	case KBluetoothDeviceClassMinorAudioCamcorder:
		return "KBluetoothDeviceClassMinorAudioCamcorder"
	case KBluetoothDeviceClassMinorAudioCar:
		return "KBluetoothDeviceClassMinorAudioCar"
	case KBluetoothDeviceClassMinorAudioGamingToy:
		return "KBluetoothDeviceClassMinorAudioGamingToy"
	case KBluetoothDeviceClassMinorAudioHandsFree:
		return "KBluetoothDeviceClassMinorAudioHandsFree"
	case KBluetoothDeviceClassMinorAudioHeadphones:
		return "KBluetoothDeviceClassMinorAudioHeadphones"
	case KBluetoothDeviceClassMinorAudioHeadset:
		return "KBluetoothDeviceClassMinorAudioHeadset"
	case KBluetoothDeviceClassMinorAudioHiFi:
		return "KBluetoothDeviceClassMinorAudioHiFi"
	case KBluetoothDeviceClassMinorAudioLoudspeaker:
		return "KBluetoothDeviceClassMinorAudioLoudspeaker"
	case KBluetoothDeviceClassMinorAudioMicrophone:
		return "KBluetoothDeviceClassMinorAudioMicrophone"
	case KBluetoothDeviceClassMinorAudioPortable:
		return "KBluetoothDeviceClassMinorAudioPortable"
	case KBluetoothDeviceClassMinorAudioReserved1:
		return "KBluetoothDeviceClassMinorAudioReserved1"
	case KBluetoothDeviceClassMinorAudioReserved2:
		return "KBluetoothDeviceClassMinorAudioReserved2"
	case KBluetoothDeviceClassMinorAudioSetTopBox:
		return "KBluetoothDeviceClassMinorAudioSetTopBox"
	case KBluetoothDeviceClassMinorAudioUnclassified:
		return "KBluetoothDeviceClassMinorAudioUnclassified"
	case KBluetoothDeviceClassMinorAudioVCR:
		return "KBluetoothDeviceClassMinorAudioVCR"
	case KBluetoothDeviceClassMinorAudioVideoCamera:
		return "KBluetoothDeviceClassMinorAudioVideoCamera"
	case KBluetoothDeviceClassMinorAudioVideoConferencing:
		return "KBluetoothDeviceClassMinorAudioVideoConferencing"
	case KBluetoothDeviceClassMinorAudioVideoDisplayAndLoudspeaker:
		return "KBluetoothDeviceClassMinorAudioVideoDisplayAndLoudspeaker"
	case KBluetoothDeviceClassMinorAudioVideoMonitor:
		return "KBluetoothDeviceClassMinorAudioVideoMonitor"
	case KBluetoothDeviceClassMinorEnd:
		return "KBluetoothDeviceClassMinorEnd"
	case KBluetoothDeviceClassMinorImaging1Printer:
		return "KBluetoothDeviceClassMinorImaging1Printer"
	case KBluetoothDeviceClassMinorNone:
		return "KBluetoothDeviceClassMinorNone"
	case KBluetoothDeviceClassMinorPeripheral1Combo:
		return "KBluetoothDeviceClassMinorPeripheral1Combo"
	case KBluetoothDeviceClassMinorPeripheral2AnyPointing:
		return "KBluetoothDeviceClassMinorPeripheral2AnyPointing"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMinor(%d)", e)
	}
}

const KBluetoothDeviceNameMaxLength uint32 = 248

type KBluetoothDontAllowRoleSwitch uint32

const (
	KBluetoothAllowRoleSwitch          KBluetoothDontAllowRoleSwitch = 0x1
	KBluetoothDontAllowRoleSwitchValue KBluetoothDontAllowRoleSwitch = 0
)

func (e KBluetoothDontAllowRoleSwitch) String() string {
	switch e {
	case KBluetoothAllowRoleSwitch:
		return "KBluetoothAllowRoleSwitch"
	case KBluetoothDontAllowRoleSwitchValue:
		return "KBluetoothDontAllowRoleSwitchValue"
	default:
		return fmt.Sprintf("KBluetoothDontAllowRoleSwitch(%d)", e)
	}
}

type KBluetoothEncryptionEnable uint32

const (
	KBluetoothEncryptionEnableBREDRAESCCM KBluetoothEncryptionEnable = 0x2
	KBluetoothEncryptionEnableBREDRE0     KBluetoothEncryptionEnable = 0x1
	KBluetoothEncryptionEnableLEAESCCM    KBluetoothEncryptionEnable = 0x1
	KBluetoothEncryptionEnableOff         KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableOn          KBluetoothEncryptionEnable = 0x1
)

func (e KBluetoothEncryptionEnable) String() string {
	switch e {
	case KBluetoothEncryptionEnableBREDRAESCCM:
		return "KBluetoothEncryptionEnableBREDRAESCCM"
	case KBluetoothEncryptionEnableBREDRE0:
		return "KBluetoothEncryptionEnableBREDRE0"
	case KBluetoothEncryptionEnableOff:
		return "KBluetoothEncryptionEnableOff"
	default:
		return fmt.Sprintf("KBluetoothEncryptionEnable(%d)", e)
	}
}

type KBluetoothGAPAppearance uint32

const (
	KBluetoothGAPAppearanceGenericBarcodeScanner               KBluetoothGAPAppearance = 704
	KBluetoothGAPAppearanceGenericBloodPressure                KBluetoothGAPAppearance = 896
	KBluetoothGAPAppearanceGenericClock                        KBluetoothGAPAppearance = 256
	KBluetoothGAPAppearanceGenericComputer                     KBluetoothGAPAppearance = 128
	KBluetoothGAPAppearanceGenericCycling                      KBluetoothGAPAppearance = 1152
	KBluetoothGAPAppearanceGenericDisplay                      KBluetoothGAPAppearance = 320
	KBluetoothGAPAppearanceGenericEyeGlasses                   KBluetoothGAPAppearance = 448
	KBluetoothGAPAppearanceGenericGlucoseMeter                 KBluetoothGAPAppearance = 1024
	KBluetoothGAPAppearanceGenericHeartrateSensor              KBluetoothGAPAppearance = 832
	KBluetoothGAPAppearanceGenericHumanInterfaceDevice         KBluetoothGAPAppearance = 960
	KBluetoothGAPAppearanceGenericKeyring                      KBluetoothGAPAppearance = 576
	KBluetoothGAPAppearanceGenericMediaPlayer                  KBluetoothGAPAppearance = 640
	KBluetoothGAPAppearanceGenericPhone                        KBluetoothGAPAppearance = 64
	KBluetoothGAPAppearanceGenericRemoteControl                KBluetoothGAPAppearance = 384
	KBluetoothGAPAppearanceGenericRunningWalkingSensor         KBluetoothGAPAppearance = 1088
	KBluetoothGAPAppearanceGenericTag                          KBluetoothGAPAppearance = 512
	KBluetoothGAPAppearanceGenericThermometer                  KBluetoothGAPAppearance = 768
	KBluetoothGAPAppearanceGenericWatch                        KBluetoothGAPAppearance = 192
	KBluetoothGAPAppearanceHumanInterfaceDeviceBarcodeScanner  KBluetoothGAPAppearance = 968
	KBluetoothGAPAppearanceHumanInterfaceDeviceCardReader      KBluetoothGAPAppearance = 966
	KBluetoothGAPAppearanceHumanInterfaceDeviceDigitalPen      KBluetoothGAPAppearance = 967
	KBluetoothGAPAppearanceHumanInterfaceDeviceDigitizerTablet KBluetoothGAPAppearance = 965
	KBluetoothGAPAppearanceHumanInterfaceDeviceGamepad         KBluetoothGAPAppearance = 964
	KBluetoothGAPAppearanceHumanInterfaceDeviceJoystick        KBluetoothGAPAppearance = 963
	KBluetoothGAPAppearanceHumanInterfaceDeviceKeyboard        KBluetoothGAPAppearance = 961
	KBluetoothGAPAppearanceHumanInterfaceDeviceMouse           KBluetoothGAPAppearance = 962
	KBluetoothGAPAppearanceUnknown                             KBluetoothGAPAppearance = 0
)

func (e KBluetoothGAPAppearance) String() string {
	switch e {
	case KBluetoothGAPAppearanceGenericBarcodeScanner:
		return "KBluetoothGAPAppearanceGenericBarcodeScanner"
	case KBluetoothGAPAppearanceGenericBloodPressure:
		return "KBluetoothGAPAppearanceGenericBloodPressure"
	case KBluetoothGAPAppearanceGenericClock:
		return "KBluetoothGAPAppearanceGenericClock"
	case KBluetoothGAPAppearanceGenericComputer:
		return "KBluetoothGAPAppearanceGenericComputer"
	case KBluetoothGAPAppearanceGenericCycling:
		return "KBluetoothGAPAppearanceGenericCycling"
	case KBluetoothGAPAppearanceGenericDisplay:
		return "KBluetoothGAPAppearanceGenericDisplay"
	case KBluetoothGAPAppearanceGenericEyeGlasses:
		return "KBluetoothGAPAppearanceGenericEyeGlasses"
	case KBluetoothGAPAppearanceGenericGlucoseMeter:
		return "KBluetoothGAPAppearanceGenericGlucoseMeter"
	case KBluetoothGAPAppearanceGenericHeartrateSensor:
		return "KBluetoothGAPAppearanceGenericHeartrateSensor"
	case KBluetoothGAPAppearanceGenericHumanInterfaceDevice:
		return "KBluetoothGAPAppearanceGenericHumanInterfaceDevice"
	case KBluetoothGAPAppearanceGenericKeyring:
		return "KBluetoothGAPAppearanceGenericKeyring"
	case KBluetoothGAPAppearanceGenericMediaPlayer:
		return "KBluetoothGAPAppearanceGenericMediaPlayer"
	case KBluetoothGAPAppearanceGenericPhone:
		return "KBluetoothGAPAppearanceGenericPhone"
	case KBluetoothGAPAppearanceGenericRemoteControl:
		return "KBluetoothGAPAppearanceGenericRemoteControl"
	case KBluetoothGAPAppearanceGenericRunningWalkingSensor:
		return "KBluetoothGAPAppearanceGenericRunningWalkingSensor"
	case KBluetoothGAPAppearanceGenericTag:
		return "KBluetoothGAPAppearanceGenericTag"
	case KBluetoothGAPAppearanceGenericThermometer:
		return "KBluetoothGAPAppearanceGenericThermometer"
	case KBluetoothGAPAppearanceGenericWatch:
		return "KBluetoothGAPAppearanceGenericWatch"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceBarcodeScanner:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceBarcodeScanner"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceCardReader:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceCardReader"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceDigitalPen:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceDigitalPen"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceDigitizerTablet:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceDigitizerTablet"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceGamepad:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceGamepad"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceJoystick:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceJoystick"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceKeyboard:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceKeyboard"
	case KBluetoothGAPAppearanceHumanInterfaceDeviceMouse:
		return "KBluetoothGAPAppearanceHumanInterfaceDeviceMouse"
	case KBluetoothGAPAppearanceUnknown:
		return "KBluetoothGAPAppearanceUnknown"
	default:
		return fmt.Sprintf("KBluetoothGAPAppearance(%d)", e)
	}
}

type KBluetoothGeneralInquiryAccessCodeIndex uint32

const (
	KBluetoothGeneralInquiryAccessCodeIndexValue KBluetoothGeneralInquiryAccessCodeIndex = 0
	KBluetoothGeneralInquiryAccessCodeLAPValue   KBluetoothGeneralInquiryAccessCodeIndex = 0x9e8b33
	KBluetoothLimitedInquiryAccessCodeEnd        KBluetoothGeneralInquiryAccessCodeIndex = 10390273
	KBluetoothLimitedInquiryAccessCodeIndex      KBluetoothGeneralInquiryAccessCodeIndex = 1
	KBluetoothLimitedInquiryAccessCodeLAPValue   KBluetoothGeneralInquiryAccessCodeIndex = 0x9e8b00
)

func (e KBluetoothGeneralInquiryAccessCodeIndex) String() string {
	switch e {
	case KBluetoothGeneralInquiryAccessCodeIndexValue:
		return "KBluetoothGeneralInquiryAccessCodeIndexValue"
	case KBluetoothGeneralInquiryAccessCodeLAPValue:
		return "KBluetoothGeneralInquiryAccessCodeLAPValue"
	case KBluetoothLimitedInquiryAccessCodeEnd:
		return "KBluetoothLimitedInquiryAccessCodeEnd"
	case KBluetoothLimitedInquiryAccessCodeIndex:
		return "KBluetoothLimitedInquiryAccessCodeIndex"
	case KBluetoothLimitedInquiryAccessCodeLAPValue:
		return "KBluetoothLimitedInquiryAccessCodeLAPValue"
	default:
		return fmt.Sprintf("KBluetoothGeneralInquiryAccessCodeIndex(%d)", e)
	}
}

type KBluetoothHCICommandPacketHeaderSize uint32

const (
	KBluetoothHCICommandPacketHeaderSizeValue KBluetoothHCICommandPacketHeaderSize = 3
	KBluetoothHCICommandPacketMaxDataSize     KBluetoothHCICommandPacketHeaderSize = 255
	KBluetoothHCIDataPacketHeaderSize         KBluetoothHCICommandPacketHeaderSize = 4
	KBluetoothHCIDataPacketMaxDataSize        KBluetoothHCICommandPacketHeaderSize = 65535
	KBluetoothHCIEventPacketHeaderSize        KBluetoothHCICommandPacketHeaderSize = 2
	KBluetoothHCIEventPacketMaxDataSize       KBluetoothHCICommandPacketHeaderSize = 255
	KBluetoothHCIMaxCommandPacketSize         KBluetoothHCICommandPacketHeaderSize = 258
	KBluetoothHCIMaxDataPacketSize            KBluetoothHCICommandPacketHeaderSize = 65539
	KBluetoothHCIMaxEventPacketSize           KBluetoothHCICommandPacketHeaderSize = 257
)

func (e KBluetoothHCICommandPacketHeaderSize) String() string {
	switch e {
	case KBluetoothHCICommandPacketHeaderSizeValue:
		return "KBluetoothHCICommandPacketHeaderSizeValue"
	case KBluetoothHCICommandPacketMaxDataSize:
		return "KBluetoothHCICommandPacketMaxDataSize"
	case KBluetoothHCIDataPacketHeaderSize:
		return "KBluetoothHCIDataPacketHeaderSize"
	case KBluetoothHCIDataPacketMaxDataSize:
		return "KBluetoothHCIDataPacketMaxDataSize"
	case KBluetoothHCIEventPacketHeaderSize:
		return "KBluetoothHCIEventPacketHeaderSize"
	case KBluetoothHCIMaxCommandPacketSize:
		return "KBluetoothHCIMaxCommandPacketSize"
	case KBluetoothHCIMaxDataPacketSize:
		return "KBluetoothHCIMaxDataPacketSize"
	case KBluetoothHCIMaxEventPacketSize:
		return "KBluetoothHCIMaxEventPacketSize"
	default:
		return fmt.Sprintf("KBluetoothHCICommandPacketHeaderSize(%d)", e)
	}
}

type KBluetoothHCIErroneousDataReporting uint32

const (
	KBluetoothHCIErroneousDataReportingDisabled      KBluetoothHCIErroneousDataReporting = 0
	KBluetoothHCIErroneousDataReportingEnabled       KBluetoothHCIErroneousDataReporting = 0x1
	KBluetoothHCIErroneousDataReportingReservedEnd   KBluetoothHCIErroneousDataReporting = 0xff
	KBluetoothHCIErroneousDataReportingReservedStart KBluetoothHCIErroneousDataReporting = 0x2
)

func (e KBluetoothHCIErroneousDataReporting) String() string {
	switch e {
	case KBluetoothHCIErroneousDataReportingDisabled:
		return "KBluetoothHCIErroneousDataReportingDisabled"
	case KBluetoothHCIErroneousDataReportingEnabled:
		return "KBluetoothHCIErroneousDataReportingEnabled"
	case KBluetoothHCIErroneousDataReportingReservedEnd:
		return "KBluetoothHCIErroneousDataReportingReservedEnd"
	case KBluetoothHCIErroneousDataReportingReservedStart:
		return "KBluetoothHCIErroneousDataReportingReservedStart"
	default:
		return fmt.Sprintf("KBluetoothHCIErroneousDataReporting(%d)", e)
	}
}

type KBluetoothHCIError uint32

const (
	KBluetoothHCIErrorACLConnectionAlreadyExists                    KBluetoothHCIError = 0xb
	KBluetoothHCIErrorAuthenticationFailure                         KBluetoothHCIError = 0x5
	KBluetoothHCIErrorChannelClassificationNotSupported             KBluetoothHCIError = 0x2e
	KBluetoothHCIErrorCoarseClockAdjustmentRejected                 KBluetoothHCIError = 0x40
	KBluetoothHCIErrorCommandDisallowed                             KBluetoothHCIError = 0xc
	KBluetoothHCIErrorConnectionFailedToBeEstablished               KBluetoothHCIError = 0x3e
	KBluetoothHCIErrorConnectionRejectedDueToNoSuitableChannelFound KBluetoothHCIError = 0x39
	KBluetoothHCIErrorConnectionTerminatedByLocalHost               KBluetoothHCIError = 0x16
	KBluetoothHCIErrorConnectionTerminatedDueToMICFailure           KBluetoothHCIError = 0x3d
	KBluetoothHCIErrorConnectionTimeout                             KBluetoothHCIError = 0x8
	KBluetoothHCIErrorControllerBusy                                KBluetoothHCIError = 0x3a
	KBluetoothHCIErrorDifferentTransactionCollision                 KBluetoothHCIError = 0x2a
	KBluetoothHCIErrorDirectedAdvertisingTimeout                    KBluetoothHCIError = 0x3c
	KBluetoothHCIErrorEncryptionModeNotAcceptable                   KBluetoothHCIError = 0x25
	KBluetoothHCIErrorExtendedInquiryResponseTooLarge               KBluetoothHCIError = 0x36
	KBluetoothHCIErrorHardwareFailure                               KBluetoothHCIError = 0x3
	KBluetoothHCIErrorHostBusyPairing                               KBluetoothHCIError = 0x38
	KBluetoothHCIErrorHostRejectedLimitedResources                  KBluetoothHCIError = 0xd
	KBluetoothHCIErrorHostRejectedRemoteDeviceIsPersonal            KBluetoothHCIError = 0xf
	KBluetoothHCIErrorHostRejectedSecurityReasons                   KBluetoothHCIError = 0xe
	KBluetoothHCIErrorHostRejectedUnacceptableDeviceAddress         KBluetoothHCIError = 0xf
	KBluetoothHCIErrorHostTimeout                                   KBluetoothHCIError = 0x10
	KBluetoothHCIErrorInstantPassed                                 KBluetoothHCIError = 0x28
	KBluetoothHCIErrorInsufficientSecurity                          KBluetoothHCIError = 0x2f
	KBluetoothHCIErrorInvalidHCICommandParameters                   KBluetoothHCIError = 0x12
	KBluetoothHCIErrorInvalidLMPParameters                          KBluetoothHCIError = 0x1e
	KBluetoothHCIErrorKeyMissing                                    KBluetoothHCIError = 0x6
	KBluetoothHCIErrorLMPErrorTransactionCollision                  KBluetoothHCIError = 0x23
	KBluetoothHCIErrorLMPPDUNotAllowed                              KBluetoothHCIError = 0x24
	KBluetoothHCIErrorLMPResponseTimeout                            KBluetoothHCIError = 0x22
	KBluetoothHCIErrorMACConnectionFailed                           KBluetoothHCIError = 0x3f
	KBluetoothHCIErrorMax                                           KBluetoothHCIError = 0x40
	KBluetoothHCIErrorMaxNumberOfConnections                        KBluetoothHCIError = 0x9
	KBluetoothHCIErrorMaxNumberOfSCOConnectionsToADevice            KBluetoothHCIError = 0xa
	KBluetoothHCIErrorMemoryFull                                    KBluetoothHCIError = 0x7
	KBluetoothHCIErrorNoConnection                                  KBluetoothHCIError = 0x2
	KBluetoothHCIErrorOtherEndTerminatedConnectionAboutToPowerOff   KBluetoothHCIError = 0x15
	KBluetoothHCIErrorOtherEndTerminatedConnectionLowResources      KBluetoothHCIError = 0x14
	KBluetoothHCIErrorOtherEndTerminatedConnectionUserEnded         KBluetoothHCIError = 0x13
	KBluetoothHCIErrorPageTimeout                                   KBluetoothHCIError = 0x4
	KBluetoothHCIErrorPairingNotAllowed                             KBluetoothHCIError = 0x18
	KBluetoothHCIErrorPairingWithUnitKeyNotSupported                KBluetoothHCIError = 0x29
	KBluetoothHCIErrorParameterOutOfMandatoryRange                  KBluetoothHCIError = 0x30
	KBluetoothHCIErrorQoSNotSupported                               KBluetoothHCIError = 0x27
	KBluetoothHCIErrorQoSRejected                                   KBluetoothHCIError = 0x2d
	KBluetoothHCIErrorQoSUnacceptableParameter                      KBluetoothHCIError = 0x2c
	KBluetoothHCIErrorRepeatedAttempts                              KBluetoothHCIError = 0x17
	KBluetoothHCIErrorReservedSlotViolation                         KBluetoothHCIError = 0x34
	KBluetoothHCIErrorRoleChangeNotAllowed                          KBluetoothHCIError = 0x21
	KBluetoothHCIErrorRoleSwitchFailed                              KBluetoothHCIError = 0x35
	KBluetoothHCIErrorRoleSwitchPending                             KBluetoothHCIError = 0x31
	KBluetoothHCIErrorSCOAirModeRejected                            KBluetoothHCIError = 0x1d
	KBluetoothHCIErrorSCOIntervalRejected                           KBluetoothHCIError = 0x1c
	KBluetoothHCIErrorSCOOffsetRejected                             KBluetoothHCIError = 0x1b
	KBluetoothHCIErrorSecureSimplePairingNotSupportedByHost         KBluetoothHCIError = 0x37
	KBluetoothHCIErrorSuccess                                       KBluetoothHCIError = 0
	KBluetoothHCIErrorUnacceptableConnectionInterval                KBluetoothHCIError = 0x3b
	KBluetoothHCIErrorUnitKeyUsed                                   KBluetoothHCIError = 0x26
	KBluetoothHCIErrorUnknownHCICommand                             KBluetoothHCIError = 0x1
	KBluetoothHCIErrorUnknownLMPPDU                                 KBluetoothHCIError = 0x19
	KBluetoothHCIErrorUnspecifiedError                              KBluetoothHCIError = 0x1f
	KBluetoothHCIErrorUnsupportedFeatureOrParameterValue            KBluetoothHCIError = 0x11
	KBluetoothHCIErrorUnsupportedLMPParameterValue                  KBluetoothHCIError = 0x20
	KBluetoothHCIErrorUnsupportedRemoteFeature                      KBluetoothHCIError = 0x1a
)

func (e KBluetoothHCIError) String() string {
	switch e {
	case KBluetoothHCIErrorACLConnectionAlreadyExists:
		return "KBluetoothHCIErrorACLConnectionAlreadyExists"
	case KBluetoothHCIErrorAuthenticationFailure:
		return "KBluetoothHCIErrorAuthenticationFailure"
	case KBluetoothHCIErrorChannelClassificationNotSupported:
		return "KBluetoothHCIErrorChannelClassificationNotSupported"
	case KBluetoothHCIErrorCoarseClockAdjustmentRejected:
		return "KBluetoothHCIErrorCoarseClockAdjustmentRejected"
	case KBluetoothHCIErrorCommandDisallowed:
		return "KBluetoothHCIErrorCommandDisallowed"
	case KBluetoothHCIErrorConnectionFailedToBeEstablished:
		return "KBluetoothHCIErrorConnectionFailedToBeEstablished"
	case KBluetoothHCIErrorConnectionRejectedDueToNoSuitableChannelFound:
		return "KBluetoothHCIErrorConnectionRejectedDueToNoSuitableChannelFound"
	case KBluetoothHCIErrorConnectionTerminatedByLocalHost:
		return "KBluetoothHCIErrorConnectionTerminatedByLocalHost"
	case KBluetoothHCIErrorConnectionTerminatedDueToMICFailure:
		return "KBluetoothHCIErrorConnectionTerminatedDueToMICFailure"
	case KBluetoothHCIErrorConnectionTimeout:
		return "KBluetoothHCIErrorConnectionTimeout"
	case KBluetoothHCIErrorControllerBusy:
		return "KBluetoothHCIErrorControllerBusy"
	case KBluetoothHCIErrorDifferentTransactionCollision:
		return "KBluetoothHCIErrorDifferentTransactionCollision"
	case KBluetoothHCIErrorDirectedAdvertisingTimeout:
		return "KBluetoothHCIErrorDirectedAdvertisingTimeout"
	case KBluetoothHCIErrorEncryptionModeNotAcceptable:
		return "KBluetoothHCIErrorEncryptionModeNotAcceptable"
	case KBluetoothHCIErrorExtendedInquiryResponseTooLarge:
		return "KBluetoothHCIErrorExtendedInquiryResponseTooLarge"
	case KBluetoothHCIErrorHardwareFailure:
		return "KBluetoothHCIErrorHardwareFailure"
	case KBluetoothHCIErrorHostBusyPairing:
		return "KBluetoothHCIErrorHostBusyPairing"
	case KBluetoothHCIErrorHostRejectedLimitedResources:
		return "KBluetoothHCIErrorHostRejectedLimitedResources"
	case KBluetoothHCIErrorHostRejectedRemoteDeviceIsPersonal:
		return "KBluetoothHCIErrorHostRejectedRemoteDeviceIsPersonal"
	case KBluetoothHCIErrorHostRejectedSecurityReasons:
		return "KBluetoothHCIErrorHostRejectedSecurityReasons"
	case KBluetoothHCIErrorHostTimeout:
		return "KBluetoothHCIErrorHostTimeout"
	case KBluetoothHCIErrorInstantPassed:
		return "KBluetoothHCIErrorInstantPassed"
	case KBluetoothHCIErrorInsufficientSecurity:
		return "KBluetoothHCIErrorInsufficientSecurity"
	case KBluetoothHCIErrorInvalidHCICommandParameters:
		return "KBluetoothHCIErrorInvalidHCICommandParameters"
	case KBluetoothHCIErrorInvalidLMPParameters:
		return "KBluetoothHCIErrorInvalidLMPParameters"
	case KBluetoothHCIErrorKeyMissing:
		return "KBluetoothHCIErrorKeyMissing"
	case KBluetoothHCIErrorLMPErrorTransactionCollision:
		return "KBluetoothHCIErrorLMPErrorTransactionCollision"
	case KBluetoothHCIErrorLMPPDUNotAllowed:
		return "KBluetoothHCIErrorLMPPDUNotAllowed"
	case KBluetoothHCIErrorLMPResponseTimeout:
		return "KBluetoothHCIErrorLMPResponseTimeout"
	case KBluetoothHCIErrorMACConnectionFailed:
		return "KBluetoothHCIErrorMACConnectionFailed"
	case KBluetoothHCIErrorMaxNumberOfConnections:
		return "KBluetoothHCIErrorMaxNumberOfConnections"
	case KBluetoothHCIErrorMaxNumberOfSCOConnectionsToADevice:
		return "KBluetoothHCIErrorMaxNumberOfSCOConnectionsToADevice"
	case KBluetoothHCIErrorMemoryFull:
		return "KBluetoothHCIErrorMemoryFull"
	case KBluetoothHCIErrorNoConnection:
		return "KBluetoothHCIErrorNoConnection"
	case KBluetoothHCIErrorOtherEndTerminatedConnectionAboutToPowerOff:
		return "KBluetoothHCIErrorOtherEndTerminatedConnectionAboutToPowerOff"
	case KBluetoothHCIErrorOtherEndTerminatedConnectionLowResources:
		return "KBluetoothHCIErrorOtherEndTerminatedConnectionLowResources"
	case KBluetoothHCIErrorOtherEndTerminatedConnectionUserEnded:
		return "KBluetoothHCIErrorOtherEndTerminatedConnectionUserEnded"
	case KBluetoothHCIErrorPageTimeout:
		return "KBluetoothHCIErrorPageTimeout"
	case KBluetoothHCIErrorPairingNotAllowed:
		return "KBluetoothHCIErrorPairingNotAllowed"
	case KBluetoothHCIErrorPairingWithUnitKeyNotSupported:
		return "KBluetoothHCIErrorPairingWithUnitKeyNotSupported"
	case KBluetoothHCIErrorParameterOutOfMandatoryRange:
		return "KBluetoothHCIErrorParameterOutOfMandatoryRange"
	case KBluetoothHCIErrorQoSNotSupported:
		return "KBluetoothHCIErrorQoSNotSupported"
	case KBluetoothHCIErrorQoSRejected:
		return "KBluetoothHCIErrorQoSRejected"
	case KBluetoothHCIErrorQoSUnacceptableParameter:
		return "KBluetoothHCIErrorQoSUnacceptableParameter"
	case KBluetoothHCIErrorRepeatedAttempts:
		return "KBluetoothHCIErrorRepeatedAttempts"
	case KBluetoothHCIErrorReservedSlotViolation:
		return "KBluetoothHCIErrorReservedSlotViolation"
	case KBluetoothHCIErrorRoleChangeNotAllowed:
		return "KBluetoothHCIErrorRoleChangeNotAllowed"
	case KBluetoothHCIErrorRoleSwitchFailed:
		return "KBluetoothHCIErrorRoleSwitchFailed"
	case KBluetoothHCIErrorRoleSwitchPending:
		return "KBluetoothHCIErrorRoleSwitchPending"
	case KBluetoothHCIErrorSCOAirModeRejected:
		return "KBluetoothHCIErrorSCOAirModeRejected"
	case KBluetoothHCIErrorSCOIntervalRejected:
		return "KBluetoothHCIErrorSCOIntervalRejected"
	case KBluetoothHCIErrorSCOOffsetRejected:
		return "KBluetoothHCIErrorSCOOffsetRejected"
	case KBluetoothHCIErrorSecureSimplePairingNotSupportedByHost:
		return "KBluetoothHCIErrorSecureSimplePairingNotSupportedByHost"
	case KBluetoothHCIErrorSuccess:
		return "KBluetoothHCIErrorSuccess"
	case KBluetoothHCIErrorUnacceptableConnectionInterval:
		return "KBluetoothHCIErrorUnacceptableConnectionInterval"
	case KBluetoothHCIErrorUnitKeyUsed:
		return "KBluetoothHCIErrorUnitKeyUsed"
	case KBluetoothHCIErrorUnknownHCICommand:
		return "KBluetoothHCIErrorUnknownHCICommand"
	case KBluetoothHCIErrorUnknownLMPPDU:
		return "KBluetoothHCIErrorUnknownLMPPDU"
	case KBluetoothHCIErrorUnspecifiedError:
		return "KBluetoothHCIErrorUnspecifiedError"
	case KBluetoothHCIErrorUnsupportedFeatureOrParameterValue:
		return "KBluetoothHCIErrorUnsupportedFeatureOrParameterValue"
	case KBluetoothHCIErrorUnsupportedLMPParameterValue:
		return "KBluetoothHCIErrorUnsupportedLMPParameterValue"
	case KBluetoothHCIErrorUnsupportedRemoteFeature:
		return "KBluetoothHCIErrorUnsupportedRemoteFeature"
	default:
		return fmt.Sprintf("KBluetoothHCIError(%d)", e)
	}
}

const KBluetoothHCIErrorPowerIsOFF uint32 = 65

type KBluetoothHCIEventInquiryComplete uint32

const (
	KBluetoothHCIEventAMPReceiverReport                       KBluetoothHCIEventInquiryComplete = 0x4b
	KBluetoothHCIEventAMPStartTest                            KBluetoothHCIEventInquiryComplete = 0x49
	KBluetoothHCIEventAMPStatusChange                         KBluetoothHCIEventInquiryComplete = 0x4d
	KBluetoothHCIEventAMPTestEnd                              KBluetoothHCIEventInquiryComplete = 0x4a
	KBluetoothHCIEventAuthenticationComplete                  KBluetoothHCIEventInquiryComplete = 0x6
	KBluetoothHCIEventChangeConnectionLinkKeyComplete         KBluetoothHCIEventInquiryComplete = 0x9
	KBluetoothHCIEventChannelSelected                         KBluetoothHCIEventInquiryComplete = 0x41
	KBluetoothHCIEventCommandComplete                         KBluetoothHCIEventInquiryComplete = 0xe
	KBluetoothHCIEventCommandStatus                           KBluetoothHCIEventInquiryComplete = 0xf
	KBluetoothHCIEventConnectionComplete                      KBluetoothHCIEventInquiryComplete = 0x3
	KBluetoothHCIEventConnectionPacketType                    KBluetoothHCIEventInquiryComplete = 0x1d
	KBluetoothHCIEventConnectionRequest                       KBluetoothHCIEventInquiryComplete = 0x4
	KBluetoothHCIEventDataBufferOverflow                      KBluetoothHCIEventInquiryComplete = 0x1a
	KBluetoothHCIEventDisconnectionComplete                   KBluetoothHCIEventInquiryComplete = 0x5
	KBluetoothHCIEventDisconnectionLogicalLinkComplete        KBluetoothHCIEventInquiryComplete = 0x46
	KBluetoothHCIEventDisconnectionPhysicalLinkComplete       KBluetoothHCIEventInquiryComplete = 0x42
	KBluetoothHCIEventEncryptionChange                        KBluetoothHCIEventInquiryComplete = 0x8
	KBluetoothHCIEventEncryptionKeyRefreshComplete            KBluetoothHCIEventInquiryComplete = 0x30
	KBluetoothHCIEventEnhancedFlushComplete                   KBluetoothHCIEventInquiryComplete = 0x39
	KBluetoothHCIEventExtendedInquiryResult                   KBluetoothHCIEventInquiryComplete = 0x2f
	KBluetoothHCIEventFlowSpecModifyComplete                  KBluetoothHCIEventInquiryComplete = 0x47
	KBluetoothHCIEventFlowSpecificationComplete               KBluetoothHCIEventInquiryComplete = 0x21
	KBluetoothHCIEventFlushOccurred                           KBluetoothHCIEventInquiryComplete = 0x11
	KBluetoothHCIEventHardwareError                           KBluetoothHCIEventInquiryComplete = 0x10
	KBluetoothHCIEventIOCapabilityRequest                     KBluetoothHCIEventInquiryComplete = 0x31
	KBluetoothHCIEventIOCapabilityResponse                    KBluetoothHCIEventInquiryComplete = 0x32
	KBluetoothHCIEventInquiryCompleteValue                    KBluetoothHCIEventInquiryComplete = 0x1
	KBluetoothHCIEventInquiryResult                           KBluetoothHCIEventInquiryComplete = 0x2
	KBluetoothHCIEventInquiryResultWithRSSI                   KBluetoothHCIEventInquiryComplete = 0x22
	KBluetoothHCIEventKeypressNotification                    KBluetoothHCIEventInquiryComplete = 0x3c
	KBluetoothHCIEventLEMetaEvent                             KBluetoothHCIEventInquiryComplete = 0x3e
	KBluetoothHCIEventLinkKeyNotification                     KBluetoothHCIEventInquiryComplete = 0x18
	KBluetoothHCIEventLinkKeyRequest                          KBluetoothHCIEventInquiryComplete = 0x17
	KBluetoothHCIEventLinkSupervisionTimeoutChanged           KBluetoothHCIEventInquiryComplete = 0x38
	KBluetoothHCIEventLogicalLinkComplete                     KBluetoothHCIEventInquiryComplete = 0x45
	KBluetoothHCIEventLogoTesting                             KBluetoothHCIEventInquiryComplete = 0xfe
	KBluetoothHCIEventLoopbackCommand                         KBluetoothHCIEventInquiryComplete = 0x19
	KBluetoothHCIEventMasterLinkKeyComplete                   KBluetoothHCIEventInquiryComplete = 0xa
	KBluetoothHCIEventMaxSlotsChange                          KBluetoothHCIEventInquiryComplete = 0x1b
	KBluetoothHCIEventModeChange                              KBluetoothHCIEventInquiryComplete = 0x14
	KBluetoothHCIEventNumberOfCompletedDataBlocks             KBluetoothHCIEventInquiryComplete = 0x48
	KBluetoothHCIEventNumberOfCompletedPackets                KBluetoothHCIEventInquiryComplete = 0x13
	KBluetoothHCIEventPINCodeRequest                          KBluetoothHCIEventInquiryComplete = 0x16
	KBluetoothHCIEventPageScanModeChange                      KBluetoothHCIEventInquiryComplete = 0x1f
	KBluetoothHCIEventPageScanRepetitionModeChange            KBluetoothHCIEventInquiryComplete = 0x20
	KBluetoothHCIEventPhysicalLinkComplete                    KBluetoothHCIEventInquiryComplete = 0x40
	KBluetoothHCIEventPhysicalLinkLossEarlyWarning            KBluetoothHCIEventInquiryComplete = 0x43
	KBluetoothHCIEventPhysicalLinkRecovery                    KBluetoothHCIEventInquiryComplete = 0x44
	KBluetoothHCIEventQoSSetupComplete                        KBluetoothHCIEventInquiryComplete = 0xd
	KBluetoothHCIEventQoSViolation                            KBluetoothHCIEventInquiryComplete = 0x1e
	KBluetoothHCIEventReadClockOffsetComplete                 KBluetoothHCIEventInquiryComplete = 0x1c
	KBluetoothHCIEventReadRemoteExtendedFeaturesComplete      KBluetoothHCIEventInquiryComplete = 0x23
	KBluetoothHCIEventReadRemoteSupportedFeaturesComplete     KBluetoothHCIEventInquiryComplete = 0xb
	KBluetoothHCIEventReadRemoteVersionInformationComplete    KBluetoothHCIEventInquiryComplete = 0xc
	KBluetoothHCIEventRemoteHostSupportedFeaturesNotification KBluetoothHCIEventInquiryComplete = 0x3d
	KBluetoothHCIEventRemoteNameRequestComplete               KBluetoothHCIEventInquiryComplete = 0x7
	KBluetoothHCIEventRemoteOOBDataRequest                    KBluetoothHCIEventInquiryComplete = 0x35
	KBluetoothHCIEventReturnLinkKeys                          KBluetoothHCIEventInquiryComplete = 0x15
	KBluetoothHCIEventRoleChange                              KBluetoothHCIEventInquiryComplete = 0x12
	KBluetoothHCIEventShortRangeModeChangeComplete            KBluetoothHCIEventInquiryComplete = 0x4c
	KBluetoothHCIEventSimplePairingComplete                   KBluetoothHCIEventInquiryComplete = 0x36
	KBluetoothHCIEventSniffSubrating                          KBluetoothHCIEventInquiryComplete = 0x2e
	KBluetoothHCIEventSynchronousConnectionChanged            KBluetoothHCIEventInquiryComplete = 0x2d
	KBluetoothHCIEventSynchronousConnectionComplete           KBluetoothHCIEventInquiryComplete = 0x2c
	KBluetoothHCIEventUserConfirmationRequest                 KBluetoothHCIEventInquiryComplete = 0x33
	KBluetoothHCIEventUserPasskeyNotification                 KBluetoothHCIEventInquiryComplete = 0x3b
	KBluetoothHCIEventUserPasskeyRequest                      KBluetoothHCIEventInquiryComplete = 0x34
	KBluetoothHCIEventVendorSpecific                          KBluetoothHCIEventInquiryComplete = 0xff
	KBluetoothHCISubEventLEAdvertisingReport                  KBluetoothHCIEventInquiryComplete = 0x2
	KBluetoothHCISubEventLEAdvertisingSetTerminated           KBluetoothHCIEventInquiryComplete = 0x12
	KBluetoothHCISubEventLEChannelSelectionAlgorithm          KBluetoothHCIEventInquiryComplete = 0x14
	KBluetoothHCISubEventLEConnectionComplete                 KBluetoothHCIEventInquiryComplete = 0x1
	KBluetoothHCISubEventLEConnectionUpdateComplete           KBluetoothHCIEventInquiryComplete = 0x3
	KBluetoothHCISubEventLEDataLengthChange                   KBluetoothHCIEventInquiryComplete = 0x7
	KBluetoothHCISubEventLEDirectAdvertisingReport            KBluetoothHCIEventInquiryComplete = 0xb
	KBluetoothHCISubEventLEEnhancedConnectionComplete         KBluetoothHCIEventInquiryComplete = 0xa
	KBluetoothHCISubEventLEExtendedAdvertising                KBluetoothHCIEventInquiryComplete = 0xd
	KBluetoothHCISubEventLEGenerateDHKeyComplete              KBluetoothHCIEventInquiryComplete = 0x9
	KBluetoothHCISubEventLELongTermKeyRequest                 KBluetoothHCIEventInquiryComplete = 0x5
	KBluetoothHCISubEventLEPeriodicAdvertisingReport          KBluetoothHCIEventInquiryComplete = 0xf
	KBluetoothHCISubEventLEPeriodicAdvertisingSyncEstablished KBluetoothHCIEventInquiryComplete = 0xe
	KBluetoothHCISubEventLEPeriodicAdvertisingSyncLost        KBluetoothHCIEventInquiryComplete = 0x10
	KBluetoothHCISubEventLEPhyUpdateComplete                  KBluetoothHCIEventInquiryComplete = 0xc
	KBluetoothHCISubEventLEReadLocalP256PublicKeyComplete     KBluetoothHCIEventInquiryComplete = 0x8
	KBluetoothHCISubEventLEReadRemoteUsedFeaturesComplete     KBluetoothHCIEventInquiryComplete = 0x4
	KBluetoothHCISubEventLERemoteConnectionParameterRequest   KBluetoothHCIEventInquiryComplete = 0x6
	KBluetoothHCISubEventLEScanRequestReceived                KBluetoothHCIEventInquiryComplete = 0x13
	KBluetoothHCISubEventLEScanTimeout                        KBluetoothHCIEventInquiryComplete = 0x11
)

func (e KBluetoothHCIEventInquiryComplete) String() string {
	switch e {
	case KBluetoothHCIEventAMPReceiverReport:
		return "KBluetoothHCIEventAMPReceiverReport"
	case KBluetoothHCIEventAMPStartTest:
		return "KBluetoothHCIEventAMPStartTest"
	case KBluetoothHCIEventAMPStatusChange:
		return "KBluetoothHCIEventAMPStatusChange"
	case KBluetoothHCIEventAMPTestEnd:
		return "KBluetoothHCIEventAMPTestEnd"
	case KBluetoothHCIEventAuthenticationComplete:
		return "KBluetoothHCIEventAuthenticationComplete"
	case KBluetoothHCIEventChangeConnectionLinkKeyComplete:
		return "KBluetoothHCIEventChangeConnectionLinkKeyComplete"
	case KBluetoothHCIEventChannelSelected:
		return "KBluetoothHCIEventChannelSelected"
	case KBluetoothHCIEventCommandComplete:
		return "KBluetoothHCIEventCommandComplete"
	case KBluetoothHCIEventCommandStatus:
		return "KBluetoothHCIEventCommandStatus"
	case KBluetoothHCIEventConnectionComplete:
		return "KBluetoothHCIEventConnectionComplete"
	case KBluetoothHCIEventConnectionPacketType:
		return "KBluetoothHCIEventConnectionPacketType"
	case KBluetoothHCIEventConnectionRequest:
		return "KBluetoothHCIEventConnectionRequest"
	case KBluetoothHCIEventDataBufferOverflow:
		return "KBluetoothHCIEventDataBufferOverflow"
	case KBluetoothHCIEventDisconnectionComplete:
		return "KBluetoothHCIEventDisconnectionComplete"
	case KBluetoothHCIEventDisconnectionLogicalLinkComplete:
		return "KBluetoothHCIEventDisconnectionLogicalLinkComplete"
	case KBluetoothHCIEventDisconnectionPhysicalLinkComplete:
		return "KBluetoothHCIEventDisconnectionPhysicalLinkComplete"
	case KBluetoothHCIEventEncryptionChange:
		return "KBluetoothHCIEventEncryptionChange"
	case KBluetoothHCIEventEncryptionKeyRefreshComplete:
		return "KBluetoothHCIEventEncryptionKeyRefreshComplete"
	case KBluetoothHCIEventEnhancedFlushComplete:
		return "KBluetoothHCIEventEnhancedFlushComplete"
	case KBluetoothHCIEventExtendedInquiryResult:
		return "KBluetoothHCIEventExtendedInquiryResult"
	case KBluetoothHCIEventFlowSpecModifyComplete:
		return "KBluetoothHCIEventFlowSpecModifyComplete"
	case KBluetoothHCIEventFlowSpecificationComplete:
		return "KBluetoothHCIEventFlowSpecificationComplete"
	case KBluetoothHCIEventFlushOccurred:
		return "KBluetoothHCIEventFlushOccurred"
	case KBluetoothHCIEventHardwareError:
		return "KBluetoothHCIEventHardwareError"
	case KBluetoothHCIEventIOCapabilityRequest:
		return "KBluetoothHCIEventIOCapabilityRequest"
	case KBluetoothHCIEventIOCapabilityResponse:
		return "KBluetoothHCIEventIOCapabilityResponse"
	case KBluetoothHCIEventInquiryCompleteValue:
		return "KBluetoothHCIEventInquiryCompleteValue"
	case KBluetoothHCIEventInquiryResult:
		return "KBluetoothHCIEventInquiryResult"
	case KBluetoothHCIEventInquiryResultWithRSSI:
		return "KBluetoothHCIEventInquiryResultWithRSSI"
	case KBluetoothHCIEventKeypressNotification:
		return "KBluetoothHCIEventKeypressNotification"
	case KBluetoothHCIEventLEMetaEvent:
		return "KBluetoothHCIEventLEMetaEvent"
	case KBluetoothHCIEventLinkKeyNotification:
		return "KBluetoothHCIEventLinkKeyNotification"
	case KBluetoothHCIEventLinkKeyRequest:
		return "KBluetoothHCIEventLinkKeyRequest"
	case KBluetoothHCIEventLinkSupervisionTimeoutChanged:
		return "KBluetoothHCIEventLinkSupervisionTimeoutChanged"
	case KBluetoothHCIEventLogicalLinkComplete:
		return "KBluetoothHCIEventLogicalLinkComplete"
	case KBluetoothHCIEventLogoTesting:
		return "KBluetoothHCIEventLogoTesting"
	case KBluetoothHCIEventLoopbackCommand:
		return "KBluetoothHCIEventLoopbackCommand"
	case KBluetoothHCIEventMasterLinkKeyComplete:
		return "KBluetoothHCIEventMasterLinkKeyComplete"
	case KBluetoothHCIEventMaxSlotsChange:
		return "KBluetoothHCIEventMaxSlotsChange"
	case KBluetoothHCIEventModeChange:
		return "KBluetoothHCIEventModeChange"
	case KBluetoothHCIEventNumberOfCompletedDataBlocks:
		return "KBluetoothHCIEventNumberOfCompletedDataBlocks"
	case KBluetoothHCIEventNumberOfCompletedPackets:
		return "KBluetoothHCIEventNumberOfCompletedPackets"
	case KBluetoothHCIEventPINCodeRequest:
		return "KBluetoothHCIEventPINCodeRequest"
	case KBluetoothHCIEventPageScanModeChange:
		return "KBluetoothHCIEventPageScanModeChange"
	case KBluetoothHCIEventPageScanRepetitionModeChange:
		return "KBluetoothHCIEventPageScanRepetitionModeChange"
	case KBluetoothHCIEventPhysicalLinkComplete:
		return "KBluetoothHCIEventPhysicalLinkComplete"
	case KBluetoothHCIEventPhysicalLinkLossEarlyWarning:
		return "KBluetoothHCIEventPhysicalLinkLossEarlyWarning"
	case KBluetoothHCIEventPhysicalLinkRecovery:
		return "KBluetoothHCIEventPhysicalLinkRecovery"
	case KBluetoothHCIEventQoSSetupComplete:
		return "KBluetoothHCIEventQoSSetupComplete"
	case KBluetoothHCIEventQoSViolation:
		return "KBluetoothHCIEventQoSViolation"
	case KBluetoothHCIEventReadClockOffsetComplete:
		return "KBluetoothHCIEventReadClockOffsetComplete"
	case KBluetoothHCIEventReadRemoteExtendedFeaturesComplete:
		return "KBluetoothHCIEventReadRemoteExtendedFeaturesComplete"
	case KBluetoothHCIEventReadRemoteSupportedFeaturesComplete:
		return "KBluetoothHCIEventReadRemoteSupportedFeaturesComplete"
	case KBluetoothHCIEventReadRemoteVersionInformationComplete:
		return "KBluetoothHCIEventReadRemoteVersionInformationComplete"
	case KBluetoothHCIEventRemoteHostSupportedFeaturesNotification:
		return "KBluetoothHCIEventRemoteHostSupportedFeaturesNotification"
	case KBluetoothHCIEventRemoteNameRequestComplete:
		return "KBluetoothHCIEventRemoteNameRequestComplete"
	case KBluetoothHCIEventRemoteOOBDataRequest:
		return "KBluetoothHCIEventRemoteOOBDataRequest"
	case KBluetoothHCIEventReturnLinkKeys:
		return "KBluetoothHCIEventReturnLinkKeys"
	case KBluetoothHCIEventRoleChange:
		return "KBluetoothHCIEventRoleChange"
	case KBluetoothHCIEventShortRangeModeChangeComplete:
		return "KBluetoothHCIEventShortRangeModeChangeComplete"
	case KBluetoothHCIEventSimplePairingComplete:
		return "KBluetoothHCIEventSimplePairingComplete"
	case KBluetoothHCIEventSniffSubrating:
		return "KBluetoothHCIEventSniffSubrating"
	case KBluetoothHCIEventSynchronousConnectionChanged:
		return "KBluetoothHCIEventSynchronousConnectionChanged"
	case KBluetoothHCIEventSynchronousConnectionComplete:
		return "KBluetoothHCIEventSynchronousConnectionComplete"
	case KBluetoothHCIEventUserConfirmationRequest:
		return "KBluetoothHCIEventUserConfirmationRequest"
	case KBluetoothHCIEventUserPasskeyNotification:
		return "KBluetoothHCIEventUserPasskeyNotification"
	case KBluetoothHCIEventUserPasskeyRequest:
		return "KBluetoothHCIEventUserPasskeyRequest"
	case KBluetoothHCIEventVendorSpecific:
		return "KBluetoothHCIEventVendorSpecific"
	default:
		return fmt.Sprintf("KBluetoothHCIEventInquiryComplete(%d)", e)
	}
}

type KBluetoothHCIEventMask uint32

const (
	KBluetoothHCIEventMaskAll                                  KBluetoothHCIEventMask = 0xffffffff
	KBluetoothHCIEventMaskAuthenticationComplete               KBluetoothHCIEventMask = 0x20
	KBluetoothHCIEventMaskChangeConnectionLinkKeyComplete      KBluetoothHCIEventMask = 0x100
	KBluetoothHCIEventMaskCommandComplete                      KBluetoothHCIEventMask = 0x2000
	KBluetoothHCIEventMaskCommandStatus                        KBluetoothHCIEventMask = 0x4000
	KBluetoothHCIEventMaskConnectionComplete                   KBluetoothHCIEventMask = 0x4
	KBluetoothHCIEventMaskConnectionPacketTypeChanged          KBluetoothHCIEventMask = 0x10000000
	KBluetoothHCIEventMaskConnectionRequest                    KBluetoothHCIEventMask = 0x8
	KBluetoothHCIEventMaskDataBufferOverflow                   KBluetoothHCIEventMask = 0x2000000
	KBluetoothHCIEventMaskDefault                              KBluetoothHCIEventMask = 4294967295
	KBluetoothHCIEventMaskDisconnectionComplete                KBluetoothHCIEventMask = 0x10
	KBluetoothHCIEventMaskEncryptionChange                     KBluetoothHCIEventMask = 0x80
	KBluetoothHCIEventMaskFlushOccurred                        KBluetoothHCIEventMask = 0x10000
	KBluetoothHCIEventMaskHardwareError                        KBluetoothHCIEventMask = 0x8000
	KBluetoothHCIEventMaskInquiryComplete                      KBluetoothHCIEventMask = 0x1
	KBluetoothHCIEventMaskInquiryResult                        KBluetoothHCIEventMask = 0x2
	KBluetoothHCIEventMaskLinkKeyNotification                  KBluetoothHCIEventMask = 0x800000
	KBluetoothHCIEventMaskLinkKeyRequest                       KBluetoothHCIEventMask = 0x400000
	KBluetoothHCIEventMaskLoopbackCommand                      KBluetoothHCIEventMask = 0x1000000
	KBluetoothHCIEventMaskMasterLinkKeyComplete                KBluetoothHCIEventMask = 0x200
	KBluetoothHCIEventMaskMaxSlotsChange                       KBluetoothHCIEventMask = 0x4000000
	KBluetoothHCIEventMaskModeChange                           KBluetoothHCIEventMask = 0x80000
	KBluetoothHCIEventMaskNone                                 KBluetoothHCIEventMask = 0
	KBluetoothHCIEventMaskNumberOfCompletedPackets             KBluetoothHCIEventMask = 0x40000
	KBluetoothHCIEventMaskPINCodeRequest                       KBluetoothHCIEventMask = 0x200000
	KBluetoothHCIEventMaskPageScanModeChange                   KBluetoothHCIEventMask = 0x40000000
	KBluetoothHCIEventMaskPageScanRepetitionModeChange         KBluetoothHCIEventMask = 0x80000000
	KBluetoothHCIEventMaskQoSSetupComplete                     KBluetoothHCIEventMask = 0x1000
	KBluetoothHCIEventMaskQoSViolation                         KBluetoothHCIEventMask = 0x20000000
	KBluetoothHCIEventMaskReadClockOffsetComplete              KBluetoothHCIEventMask = 0x8000000
	KBluetoothHCIEventMaskReadRemoteSupportedFeaturesComplete  KBluetoothHCIEventMask = 0x400
	KBluetoothHCIEventMaskReadRemoteVersionInformationComplete KBluetoothHCIEventMask = 0x800
	KBluetoothHCIEventMaskRemoteNameRequestComplete            KBluetoothHCIEventMask = 0x40
	KBluetoothHCIEventMaskReturnLinkKeys                       KBluetoothHCIEventMask = 0x100000
	KBluetoothHCIEventMaskRoleChange                           KBluetoothHCIEventMask = 0x20000
)

func (e KBluetoothHCIEventMask) String() string {
	switch e {
	case KBluetoothHCIEventMaskAll:
		return "KBluetoothHCIEventMaskAll"
	case KBluetoothHCIEventMaskAuthenticationComplete:
		return "KBluetoothHCIEventMaskAuthenticationComplete"
	case KBluetoothHCIEventMaskChangeConnectionLinkKeyComplete:
		return "KBluetoothHCIEventMaskChangeConnectionLinkKeyComplete"
	case KBluetoothHCIEventMaskCommandComplete:
		return "KBluetoothHCIEventMaskCommandComplete"
	case KBluetoothHCIEventMaskCommandStatus:
		return "KBluetoothHCIEventMaskCommandStatus"
	case KBluetoothHCIEventMaskConnectionComplete:
		return "KBluetoothHCIEventMaskConnectionComplete"
	case KBluetoothHCIEventMaskConnectionPacketTypeChanged:
		return "KBluetoothHCIEventMaskConnectionPacketTypeChanged"
	case KBluetoothHCIEventMaskConnectionRequest:
		return "KBluetoothHCIEventMaskConnectionRequest"
	case KBluetoothHCIEventMaskDataBufferOverflow:
		return "KBluetoothHCIEventMaskDataBufferOverflow"
	case KBluetoothHCIEventMaskDisconnectionComplete:
		return "KBluetoothHCIEventMaskDisconnectionComplete"
	case KBluetoothHCIEventMaskEncryptionChange:
		return "KBluetoothHCIEventMaskEncryptionChange"
	case KBluetoothHCIEventMaskFlushOccurred:
		return "KBluetoothHCIEventMaskFlushOccurred"
	case KBluetoothHCIEventMaskHardwareError:
		return "KBluetoothHCIEventMaskHardwareError"
	case KBluetoothHCIEventMaskInquiryComplete:
		return "KBluetoothHCIEventMaskInquiryComplete"
	case KBluetoothHCIEventMaskInquiryResult:
		return "KBluetoothHCIEventMaskInquiryResult"
	case KBluetoothHCIEventMaskLinkKeyNotification:
		return "KBluetoothHCIEventMaskLinkKeyNotification"
	case KBluetoothHCIEventMaskLinkKeyRequest:
		return "KBluetoothHCIEventMaskLinkKeyRequest"
	case KBluetoothHCIEventMaskLoopbackCommand:
		return "KBluetoothHCIEventMaskLoopbackCommand"
	case KBluetoothHCIEventMaskMasterLinkKeyComplete:
		return "KBluetoothHCIEventMaskMasterLinkKeyComplete"
	case KBluetoothHCIEventMaskMaxSlotsChange:
		return "KBluetoothHCIEventMaskMaxSlotsChange"
	case KBluetoothHCIEventMaskModeChange:
		return "KBluetoothHCIEventMaskModeChange"
	case KBluetoothHCIEventMaskNone:
		return "KBluetoothHCIEventMaskNone"
	case KBluetoothHCIEventMaskNumberOfCompletedPackets:
		return "KBluetoothHCIEventMaskNumberOfCompletedPackets"
	case KBluetoothHCIEventMaskPINCodeRequest:
		return "KBluetoothHCIEventMaskPINCodeRequest"
	case KBluetoothHCIEventMaskPageScanModeChange:
		return "KBluetoothHCIEventMaskPageScanModeChange"
	case KBluetoothHCIEventMaskPageScanRepetitionModeChange:
		return "KBluetoothHCIEventMaskPageScanRepetitionModeChange"
	case KBluetoothHCIEventMaskQoSSetupComplete:
		return "KBluetoothHCIEventMaskQoSSetupComplete"
	case KBluetoothHCIEventMaskQoSViolation:
		return "KBluetoothHCIEventMaskQoSViolation"
	case KBluetoothHCIEventMaskReadClockOffsetComplete:
		return "KBluetoothHCIEventMaskReadClockOffsetComplete"
	case KBluetoothHCIEventMaskReadRemoteSupportedFeaturesComplete:
		return "KBluetoothHCIEventMaskReadRemoteSupportedFeaturesComplete"
	case KBluetoothHCIEventMaskReadRemoteVersionInformationComplete:
		return "KBluetoothHCIEventMaskReadRemoteVersionInformationComplete"
	case KBluetoothHCIEventMaskRemoteNameRequestComplete:
		return "KBluetoothHCIEventMaskRemoteNameRequestComplete"
	case KBluetoothHCIEventMaskReturnLinkKeys:
		return "KBluetoothHCIEventMaskReturnLinkKeys"
	case KBluetoothHCIEventMaskRoleChange:
		return "KBluetoothHCIEventMaskRoleChange"
	default:
		return fmt.Sprintf("KBluetoothHCIEventMask(%d)", e)
	}
}

type KBluetoothHCILoopbackMode uint32

const (
	KBluetoothHCILoopbackModeLocal  KBluetoothHCILoopbackMode = 0x1
	KBluetoothHCILoopbackModeOff    KBluetoothHCILoopbackMode = 0
	KBluetoothHCILoopbackModeRemote KBluetoothHCILoopbackMode = 0x2
)

func (e KBluetoothHCILoopbackMode) String() string {
	switch e {
	case KBluetoothHCILoopbackModeLocal:
		return "KBluetoothHCILoopbackModeLocal"
	case KBluetoothHCILoopbackModeOff:
		return "KBluetoothHCILoopbackModeOff"
	case KBluetoothHCILoopbackModeRemote:
		return "KBluetoothHCILoopbackModeRemote"
	default:
		return fmt.Sprintf("KBluetoothHCILoopbackMode(%d)", e)
	}
}

type KBluetoothHCIOpCodeNoOp uint32

const (
	KBluetoothHCICommandAMPTest                                         KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandAMPTestEnd                                      KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandAcceptConnectionRequest                         KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandAcceptSniffRequest                              KBluetoothHCIOpCodeNoOp = 0x31
	KBluetoothHCICommandAcceptSynchronousConnectionRequest              KBluetoothHCIOpCodeNoOp = 0x29
	KBluetoothHCICommandAddSCOConnection                                KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandAuthenticationRequested                         KBluetoothHCIOpCodeNoOp = 0x11
	KBluetoothHCICommandChangeConnectionLinkKey                         KBluetoothHCIOpCodeNoOp = 0x15
	KBluetoothHCICommandChangeConnectionPacketType                      KBluetoothHCIOpCodeNoOp = 0xf
	KBluetoothHCICommandChangeLocalName                                 KBluetoothHCIOpCodeNoOp = 0x13
	KBluetoothHCICommandCreateConnection                                KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandCreateConnectionCancel                          KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandCreateNewUnitKey                                KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandDeleteReservedLTADDR                            KBluetoothHCIOpCodeNoOp = 0x75
	KBluetoothHCICommandDeleteStoredLinkKey                             KBluetoothHCIOpCodeNoOp = 0x12
	KBluetoothHCICommandDisconnect                                      KBluetoothHCIOpCodeNoOp = 0x6
	KBluetoothHCICommandEnableAMPReceiverReports                        KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandEnableDeviceUnderTestMode                       KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandEnhancedAcceptSynchronousConnectionRequest      KBluetoothHCIOpCodeNoOp = 0x3e
	KBluetoothHCICommandEnhancedFlush                                   KBluetoothHCIOpCodeNoOp = 0x5f
	KBluetoothHCICommandEnhancedSetupSynchronousConnection              KBluetoothHCIOpCodeNoOp = 0x3d
	KBluetoothHCICommandExitParkMode                                    KBluetoothHCIOpCodeNoOp = 0x6
	KBluetoothHCICommandExitPeriodicInquiryMode                         KBluetoothHCIOpCodeNoOp = 0x4
	KBluetoothHCICommandExitSniffMode                                   KBluetoothHCIOpCodeNoOp = 0x4
	KBluetoothHCICommandFlowSpecification                               KBluetoothHCIOpCodeNoOp = 0x10
	KBluetoothHCICommandFlush                                           KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandGetLinkQuality                                  KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandGetMWSTransportLayerConfiguration               KBluetoothHCIOpCodeNoOp = 0xc
	KBluetoothHCICommandGroupHostController                             KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandGroupInformational                              KBluetoothHCIOpCodeNoOp = 0x4
	KBluetoothHCICommandGroupLinkControl                                KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandGroupLinkPolicy                                 KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandGroupLogoTesting                                KBluetoothHCIOpCodeNoOp = 0x3e
	KBluetoothHCICommandGroupLowEnergy                                  KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandGroupMax                                        KBluetoothHCIOpCodeNoOp = 0x40
	KBluetoothHCICommandGroupNoOp                                       KBluetoothHCIOpCodeNoOp = 0
	KBluetoothHCICommandGroupStatus                                     KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandGroupTesting                                    KBluetoothHCIOpCodeNoOp = 0x6
	KBluetoothHCICommandGroupVendorSpecific                             KBluetoothHCIOpCodeNoOp = 0x3f
	KBluetoothHCICommandHoldMode                                        KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandHostBufferSize                                  KBluetoothHCIOpCodeNoOp = 0x33
	KBluetoothHCICommandHostNumberOfCompletedPackets                    KBluetoothHCIOpCodeNoOp = 0x35
	KBluetoothHCICommandIOCapabilityRequestNegativeReply                KBluetoothHCIOpCodeNoOp = 0x34
	KBluetoothHCICommandIOCapabilityRequestReply                        KBluetoothHCIOpCodeNoOp = 0x2b
	KBluetoothHCICommandInquiry                                         KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandInquiryCancel                                   KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandLEAddDeviceToPeriodicAdvertiserList             KBluetoothHCIOpCodeNoOp = 0x47
	KBluetoothHCICommandLEAddDeviceToResolvingList                      KBluetoothHCIOpCodeNoOp = 0x27
	KBluetoothHCICommandLEAddDeviceToWhiteList                          KBluetoothHCIOpCodeNoOp = 0x11
	KBluetoothHCICommandLEClearAdvertisingSets                          KBluetoothHCIOpCodeNoOp = 0x3d
	KBluetoothHCICommandLEClearPeriodicAdvertiserList                   KBluetoothHCIOpCodeNoOp = 0x49
	KBluetoothHCICommandLEClearResolvingList                            KBluetoothHCIOpCodeNoOp = 0x29
	KBluetoothHCICommandLEClearWhiteList                                KBluetoothHCIOpCodeNoOp = 0x10
	KBluetoothHCICommandLEConnectionUpdate                              KBluetoothHCIOpCodeNoOp = 0x13
	KBluetoothHCICommandLECreateConnection                              KBluetoothHCIOpCodeNoOp = 0xd
	KBluetoothHCICommandLECreateConnectionCancel                        KBluetoothHCIOpCodeNoOp = 0xe
	KBluetoothHCICommandLEEncrypt                                       KBluetoothHCIOpCodeNoOp = 0x17
	KBluetoothHCICommandLEEnhancedReceiverTest                          KBluetoothHCIOpCodeNoOp = 0x33
	KBluetoothHCICommandLEEnhancedTransmitterTest                       KBluetoothHCIOpCodeNoOp = 0x34
	KBluetoothHCICommandLEExtendedCreateConnection                      KBluetoothHCIOpCodeNoOp = 0x43
	KBluetoothHCICommandLEGenerateDHKey                                 KBluetoothHCIOpCodeNoOp = 0x26
	KBluetoothHCICommandLELongTermKeyRequestNegativeReply               KBluetoothHCIOpCodeNoOp = 0x1b
	KBluetoothHCICommandLELongTermKeyRequestReply                       KBluetoothHCIOpCodeNoOp = 0x1a
	KBluetoothHCICommandLEPeriodicAdvertisingCreateSync                 KBluetoothHCIOpCodeNoOp = 0x44
	KBluetoothHCICommandLEPeriodicAdvertisingCreateSyncCancel           KBluetoothHCIOpCodeNoOp = 0x45
	KBluetoothHCICommandLEPeriodicAdvertisingTerminateSync              KBluetoothHCIOpCodeNoOp = 0x46
	KBluetoothHCICommandLERand                                          KBluetoothHCIOpCodeNoOp = 0x18
	KBluetoothHCICommandLEReadAdvertisingChannelTxPower                 KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandLEReadBufferSize                                KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandLEReadChannelMap                                KBluetoothHCIOpCodeNoOp = 0x15
	KBluetoothHCICommandLEReadLocalP256PublicKey                        KBluetoothHCIOpCodeNoOp = 0x25
	KBluetoothHCICommandLEReadLocalResolvableAddress                    KBluetoothHCIOpCodeNoOp = 0x2c
	KBluetoothHCICommandLEReadLocalSupportedFeatures                    KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandLEReadMaximumDataLength                         KBluetoothHCIOpCodeNoOp = 0x2f
	KBluetoothHCICommandLEReadNumberofSupportedAdvertisingSets          KBluetoothHCIOpCodeNoOp = 0x3b
	KBluetoothHCICommandLEReadPeerResolvableAddress                     KBluetoothHCIOpCodeNoOp = 0x2b
	KBluetoothHCICommandLEReadPeriodicAdvertiserListSize                KBluetoothHCIOpCodeNoOp = 0x4a
	KBluetoothHCICommandLEReadPhy                                       KBluetoothHCIOpCodeNoOp = 0x30
	KBluetoothHCICommandLEReadRFPathCompensation                        KBluetoothHCIOpCodeNoOp = 0x4c
	KBluetoothHCICommandLEReadRemoteUsedFeatures                        KBluetoothHCIOpCodeNoOp = 0x16
	KBluetoothHCICommandLEReadResolvingListSize                         KBluetoothHCIOpCodeNoOp = 0x2a
	KBluetoothHCICommandLEReadSuggestedDefaultDataLength                KBluetoothHCIOpCodeNoOp = 0x23
	KBluetoothHCICommandLEReadSupportedStates                           KBluetoothHCIOpCodeNoOp = 0x1c
	KBluetoothHCICommandLEReadTransmitPower                             KBluetoothHCIOpCodeNoOp = 0x4b
	KBluetoothHCICommandLEReadWhiteListSize                             KBluetoothHCIOpCodeNoOp = 0xf
	KBluetoothHCICommandLEReceiverTest                                  KBluetoothHCIOpCodeNoOp = 0x1d
	KBluetoothHCICommandLERemoteConnectionParameterRequestNegativeReply KBluetoothHCIOpCodeNoOp = 0x21
	KBluetoothHCICommandLERemoteConnectionParameterRequestReply         KBluetoothHCIOpCodeNoOp = 0x20
	KBluetoothHCICommandLERemoveAdvertisingSet                          KBluetoothHCIOpCodeNoOp = 0x3c
	KBluetoothHCICommandLERemoveDeviceFromPeriodicAdvertiserList        KBluetoothHCIOpCodeNoOp = 0x48
	KBluetoothHCICommandLERemoveDeviceFromResolvingList                 KBluetoothHCIOpCodeNoOp = 0x28
	KBluetoothHCICommandLERemoveDeviceFromWhiteList                     KBluetoothHCIOpCodeNoOp = 0x12
	KBluetoothHCICommandLESetAddressResolutionEnable                    KBluetoothHCIOpCodeNoOp = 0x2d
	KBluetoothHCICommandLESetAdvertiseEnable                            KBluetoothHCIOpCodeNoOp = 0xa
	KBluetoothHCICommandLESetAdvertisingData                            KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandLESetAdvertisingParameters                      KBluetoothHCIOpCodeNoOp = 0x6
	KBluetoothHCICommandLESetAdvertisingSetRandomAddress                KBluetoothHCIOpCodeNoOp = 0x35
	KBluetoothHCICommandLESetDataLength                                 KBluetoothHCIOpCodeNoOp = 0x22
	KBluetoothHCICommandLESetDefaultPhy                                 KBluetoothHCIOpCodeNoOp = 0x31
	KBluetoothHCICommandLESetEventMask                                  KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandLESetExtendedAdvertisingData                    KBluetoothHCIOpCodeNoOp = 0x37
	KBluetoothHCICommandLESetExtendedAdvertisingEnableCommand           KBluetoothHCIOpCodeNoOp = 0x39
	KBluetoothHCICommandLESetExtendedAdvertisingParameters              KBluetoothHCIOpCodeNoOp = 0x36
	KBluetoothHCICommandLESetExtendedScanEnable                         KBluetoothHCIOpCodeNoOp = 0x42
	KBluetoothHCICommandLESetExtendedScanParameters                     KBluetoothHCIOpCodeNoOp = 0x41
	KBluetoothHCICommandLESetExtendedScanResponseData                   KBluetoothHCIOpCodeNoOp = 0x38
	KBluetoothHCICommandLESetHostChannelClassification                  KBluetoothHCIOpCodeNoOp = 0x14
	KBluetoothHCICommandLESetPeriodicAdvertisingData                    KBluetoothHCIOpCodeNoOp = 0x3f
	KBluetoothHCICommandLESetPeriodicAdvertisingEnable                  KBluetoothHCIOpCodeNoOp = 0x40
	KBluetoothHCICommandLESetPeriodicAdvertisingParameters              KBluetoothHCIOpCodeNoOp = 0x3e
	KBluetoothHCICommandLESetPhy                                        KBluetoothHCIOpCodeNoOp = 0x32
	KBluetoothHCICommandLESetPrivacyMode                                KBluetoothHCIOpCodeNoOp = 0x4e
	KBluetoothHCICommandLESetRandomAddress                              KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandLESetResolvablePrivateAddressTimeout            KBluetoothHCIOpCodeNoOp = 0x2e
	KBluetoothHCICommandLESetScanEnable                                 KBluetoothHCIOpCodeNoOp = 0xc
	KBluetoothHCICommandLESetScanParameters                             KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandLESetScanResponseData                           KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandLEStartEncryption                               KBluetoothHCIOpCodeNoOp = 0x19
	KBluetoothHCICommandLETestEnd                                       KBluetoothHCIOpCodeNoOp = 0x1f
	KBluetoothHCICommandLETransmitterTest                               KBluetoothHCIOpCodeNoOp = 0x1e
	KBluetoothHCICommandLEWriteRFPathCompensation                       KBluetoothHCIOpCodeNoOp = 0x4d
	KBluetoothHCICommandLEWriteSuggestedDefaultDataLength               KBluetoothHCIOpCodeNoOp = 0x24
	KBluetoothHCICommandLinkKeyRequestNegativeReply                     KBluetoothHCIOpCodeNoOp = 0xc
	KBluetoothHCICommandLinkKeyRequestReply                             KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandMasterLinkKey                                   KBluetoothHCIOpCodeNoOp = 0x17
	KBluetoothHCICommandMax                                             KBluetoothHCIOpCodeNoOp = 0x3ff
	KBluetoothHCICommandNoOp                                            KBluetoothHCIOpCodeNoOp = 0
	KBluetoothHCICommandPINCodeRequestNegativeReply                     KBluetoothHCIOpCodeNoOp = 0xe
	KBluetoothHCICommandPINCodeRequestReply                             KBluetoothHCIOpCodeNoOp = 0xd
	KBluetoothHCICommandParkMode                                        KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandPeriodicInquiryMode                             KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandQoSSetup                                        KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandReadAFHChannelAssessmentMode                    KBluetoothHCIOpCodeNoOp = 0x48
	KBluetoothHCICommandReadAFHMappings                                 KBluetoothHCIOpCodeNoOp = 0x6
	KBluetoothHCICommandReadAuthenticatedPayloadTimeout                 KBluetoothHCIOpCodeNoOp = 0x7b
	KBluetoothHCICommandReadAuthenticationEnable                        KBluetoothHCIOpCodeNoOp = 0x1f
	KBluetoothHCICommandReadAutomaticFlushTimeout                       KBluetoothHCIOpCodeNoOp = 0x27
	KBluetoothHCICommandReadBestEffortFlushTimeout                      KBluetoothHCIOpCodeNoOp = 0x69
	KBluetoothHCICommandReadBufferSize                                  KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandReadClassOfDevice                               KBluetoothHCIOpCodeNoOp = 0x23
	KBluetoothHCICommandReadClock                                       KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandReadClockOffset                                 KBluetoothHCIOpCodeNoOp = 0x1f
	KBluetoothHCICommandReadConnectionAcceptTimeout                     KBluetoothHCIOpCodeNoOp = 0x15
	KBluetoothHCICommandReadCountryCode                                 KBluetoothHCIOpCodeNoOp = 0x7
	KBluetoothHCICommandReadCurrentIACLAP                               KBluetoothHCIOpCodeNoOp = 0x39
	KBluetoothHCICommandReadDataBlockSize                               KBluetoothHCIOpCodeNoOp = 0xa
	KBluetoothHCICommandReadDefaultErroneousDataReporting               KBluetoothHCIOpCodeNoOp = 0x5a
	KBluetoothHCICommandReadDefaultLinkPolicySettings                   KBluetoothHCIOpCodeNoOp = 0xe
	KBluetoothHCICommandReadDeviceAddress                               KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandReadEncryptionKeySize                           KBluetoothHCIOpCodeNoOp = 0x8
	KBluetoothHCICommandReadEncryptionMode                              KBluetoothHCIOpCodeNoOp = 0x21
	KBluetoothHCICommandReadEnhancedTransmitPowerLevel                  KBluetoothHCIOpCodeNoOp = 0x68
	KBluetoothHCICommandReadExtendedInquiryLength                       KBluetoothHCIOpCodeNoOp = 0x80
	KBluetoothHCICommandReadExtendedInquiryResponse                     KBluetoothHCIOpCodeNoOp = 0x51
	KBluetoothHCICommandReadExtendedPageTimeout                         KBluetoothHCIOpCodeNoOp = 0x7e
	KBluetoothHCICommandReadFailedContactCounter                        KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandReadFlowControlMode                             KBluetoothHCIOpCodeNoOp = 0x66
	KBluetoothHCICommandReadHoldModeActivity                            KBluetoothHCIOpCodeNoOp = 0x2b
	KBluetoothHCICommandReadInquiryMode                                 KBluetoothHCIOpCodeNoOp = 0x44
	KBluetoothHCICommandReadInquiryResponseTransmitPower                KBluetoothHCIOpCodeNoOp = 0x58
	KBluetoothHCICommandReadInquiryScanActivity                         KBluetoothHCIOpCodeNoOp = 0x1d
	KBluetoothHCICommandReadInquiryScanType                             KBluetoothHCIOpCodeNoOp = 0x42
	KBluetoothHCICommandReadLEHostSupported                             KBluetoothHCIOpCodeNoOp = 0x6c
	KBluetoothHCICommandReadLMPHandle                                   KBluetoothHCIOpCodeNoOp = 0x20
	KBluetoothHCICommandReadLinkPolicySettings                          KBluetoothHCIOpCodeNoOp = 0xc
	KBluetoothHCICommandReadLinkSupervisionTimeout                      KBluetoothHCIOpCodeNoOp = 0x36
	KBluetoothHCICommandReadLocalAMPASSOC                               KBluetoothHCIOpCodeNoOp = 0xa
	KBluetoothHCICommandReadLocalAMPInfo                                KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandReadLocalExtendedFeatures                       KBluetoothHCIOpCodeNoOp = 0x4
	KBluetoothHCICommandReadLocalName                                   KBluetoothHCIOpCodeNoOp = 0x14
	KBluetoothHCICommandReadLocalOOBData                                KBluetoothHCIOpCodeNoOp = 0x57
	KBluetoothHCICommandReadLocalOOBExtendedData                        KBluetoothHCIOpCodeNoOp = 0x7d
	KBluetoothHCICommandReadLocalSupportedCodecs                        KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandReadLocalSupportedCommands                      KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandReadLocalSupportedFeatures                      KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandReadLocalVersionInformation                     KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandReadLocationData                                KBluetoothHCIOpCodeNoOp = 0x64
	KBluetoothHCICommandReadLogicalLinkAcceptTimeout                    KBluetoothHCIOpCodeNoOp = 0x61
	KBluetoothHCICommandReadLoopbackMode                                KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandReadNumberOfBroadcastRetransmissions            KBluetoothHCIOpCodeNoOp = 0x29
	KBluetoothHCICommandReadNumberOfSupportedIAC                        KBluetoothHCIOpCodeNoOp = 0x38
	KBluetoothHCICommandReadPINType                                     KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandReadPageScanActivity                            KBluetoothHCIOpCodeNoOp = 0x1b
	KBluetoothHCICommandReadPageScanMode                                KBluetoothHCIOpCodeNoOp = 0x3d
	KBluetoothHCICommandReadPageScanPeriodMode                          KBluetoothHCIOpCodeNoOp = 0x3b
	KBluetoothHCICommandReadPageScanType                                KBluetoothHCIOpCodeNoOp = 0x46
	KBluetoothHCICommandReadPageTimeout                                 KBluetoothHCIOpCodeNoOp = 0x17
	KBluetoothHCICommandReadRSSI                                        KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandReadRemoteExtendedFeatures                      KBluetoothHCIOpCodeNoOp = 0x1c
	KBluetoothHCICommandReadRemoteSupportedFeatures                     KBluetoothHCIOpCodeNoOp = 0x1b
	KBluetoothHCICommandReadRemoteVersionInformation                    KBluetoothHCIOpCodeNoOp = 0x1d
	KBluetoothHCICommandReadSCOFlowControlEnable                        KBluetoothHCIOpCodeNoOp = 0x2e
	KBluetoothHCICommandReadScanEnable                                  KBluetoothHCIOpCodeNoOp = 0x19
	KBluetoothHCICommandReadSecureConnectionsHostSupport                KBluetoothHCIOpCodeNoOp = 0x79
	KBluetoothHCICommandReadSimplePairingMode                           KBluetoothHCIOpCodeNoOp = 0x55
	KBluetoothHCICommandReadStoredLinkKey                               KBluetoothHCIOpCodeNoOp = 0xd
	KBluetoothHCICommandReadSynchronizationTrainParameters              KBluetoothHCIOpCodeNoOp = 0x77
	KBluetoothHCICommandReadTransmitPowerLevel                          KBluetoothHCIOpCodeNoOp = 0x2d
	KBluetoothHCICommandReadVoiceSetting                                KBluetoothHCIOpCodeNoOp = 0x25
	KBluetoothHCICommandReceiveSynchronizationTrain                     KBluetoothHCIOpCodeNoOp = 0x44
	KBluetoothHCICommandRefreshEncryptionKey                            KBluetoothHCIOpCodeNoOp = 0x53
	KBluetoothHCICommandRejectConnectionRequest                         KBluetoothHCIOpCodeNoOp = 0xa
	KBluetoothHCICommandRejectSniffRequest                              KBluetoothHCIOpCodeNoOp = 0x32
	KBluetoothHCICommandRejectSynchronousConnectionRequest              KBluetoothHCIOpCodeNoOp = 0x2a
	KBluetoothHCICommandRemoteNameRequest                               KBluetoothHCIOpCodeNoOp = 0x19
	KBluetoothHCICommandRemoteNameRequestCancel                         KBluetoothHCIOpCodeNoOp = 0x1a
	KBluetoothHCICommandRemoteOOBDataRequestNegativeReply               KBluetoothHCIOpCodeNoOp = 0x33
	KBluetoothHCICommandRemoteOOBDataRequestReply                       KBluetoothHCIOpCodeNoOp = 0x30
	KBluetoothHCICommandRemoteOOBExtendedDataRequestReply               KBluetoothHCIOpCodeNoOp = 0x45
	KBluetoothHCICommandReset                                           KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandResetFailedContactCounter                       KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandRoleDiscovery                                   KBluetoothHCIOpCodeNoOp = 0x9
	KBluetoothHCICommandSendKeypressNotification                        KBluetoothHCIOpCodeNoOp = 0x60
	KBluetoothHCICommandSetAFHClassification                            KBluetoothHCIOpCodeNoOp = 0x3f
	KBluetoothHCICommandSetConnectionEncryption                         KBluetoothHCIOpCodeNoOp = 0x13
	KBluetoothHCICommandSetConnectionlessPeripheralBroadcast            KBluetoothHCIOpCodeNoOp = 0x41
	KBluetoothHCICommandSetConnectionlessPeripheralBroadcastData        KBluetoothHCIOpCodeNoOp = 0x76
	KBluetoothHCICommandSetConnectionlessPeripheralBroadcastReceive     KBluetoothHCIOpCodeNoOp = 0x42
	KBluetoothHCICommandSetConnectionlessSlaveBroadcast                 KBluetoothHCIOpCodeNoOp = 65
	KBluetoothHCICommandSetConnectionlessSlaveBroadcastData             KBluetoothHCIOpCodeNoOp = 118
	KBluetoothHCICommandSetConnectionlessSlaveBroadcastReceive          KBluetoothHCIOpCodeNoOp = 66
	KBluetoothHCICommandSetEventFilter                                  KBluetoothHCIOpCodeNoOp = 0x5
	KBluetoothHCICommandSetEventMask                                    KBluetoothHCIOpCodeNoOp = 0x1
	KBluetoothHCICommandSetEventMaskPageTwo                             KBluetoothHCIOpCodeNoOp = 0x63
	KBluetoothHCICommandSetExternalFrameConfiguration                   KBluetoothHCIOpCodeNoOp = 0x6f
	KBluetoothHCICommandSetHostControllerToHostFlowControl              KBluetoothHCIOpCodeNoOp = 0x31
	KBluetoothHCICommandSetMWSChannelParameters                         KBluetoothHCIOpCodeNoOp = 0x6e
	KBluetoothHCICommandSetMWSPATTERNConfiguration                      KBluetoothHCIOpCodeNoOp = 0x73
	KBluetoothHCICommandSetMWSScanFrequencyTable                        KBluetoothHCIOpCodeNoOp = 0x72
	KBluetoothHCICommandSetMWSSignaling                                 KBluetoothHCIOpCodeNoOp = 0x70
	KBluetoothHCICommandSetMWSTransportLayer                            KBluetoothHCIOpCodeNoOp = 0x71
	KBluetoothHCICommandSetReservedLTADDR                               KBluetoothHCIOpCodeNoOp = 0x74
	KBluetoothHCICommandSetTriggeredClockCapture                        KBluetoothHCIOpCodeNoOp = 0xd
	KBluetoothHCICommandSetupSynchronousConnection                      KBluetoothHCIOpCodeNoOp = 0x28
	KBluetoothHCICommandShortRangeMode                                  KBluetoothHCIOpCodeNoOp = 0x6b
	KBluetoothHCICommandSniffMode                                       KBluetoothHCIOpCodeNoOp = 0x3
	KBluetoothHCICommandSniffSubrating                                  KBluetoothHCIOpCodeNoOp = 0x11
	KBluetoothHCICommandStartSynchronizationTrain                       KBluetoothHCIOpCodeNoOp = 0x43
	KBluetoothHCICommandSwitchRole                                      KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandTruncatedPage                                   KBluetoothHCIOpCodeNoOp = 0x3f
	KBluetoothHCICommandTruncatedPageCancel                             KBluetoothHCIOpCodeNoOp = 0x40
	KBluetoothHCICommandUserConfirmationRequestNegativeReply            KBluetoothHCIOpCodeNoOp = 0x2d
	KBluetoothHCICommandUserConfirmationRequestReply                    KBluetoothHCIOpCodeNoOp = 0x2c
	KBluetoothHCICommandUserPasskeyRequestNegativeReply                 KBluetoothHCIOpCodeNoOp = 0x2f
	KBluetoothHCICommandUserPasskeyRequestReply                         KBluetoothHCIOpCodeNoOp = 0x2e
	KBluetoothHCICommandWriteAFHChannelAssessmentMode                   KBluetoothHCIOpCodeNoOp = 0x49
	KBluetoothHCICommandWriteAuthenticatedPayloadTimeout                KBluetoothHCIOpCodeNoOp = 0x7c
	KBluetoothHCICommandWriteAuthenticationEnable                       KBluetoothHCIOpCodeNoOp = 0x20
	KBluetoothHCICommandWriteAutomaticFlushTimeout                      KBluetoothHCIOpCodeNoOp = 0x28
	KBluetoothHCICommandWriteBestEffortFlushTimeout                     KBluetoothHCIOpCodeNoOp = 0x6a
	KBluetoothHCICommandWriteClassOfDevice                              KBluetoothHCIOpCodeNoOp = 0x24
	KBluetoothHCICommandWriteConnectionAcceptTimeout                    KBluetoothHCIOpCodeNoOp = 0x16
	KBluetoothHCICommandWriteCurrentIACLAP                              KBluetoothHCIOpCodeNoOp = 0x3a
	KBluetoothHCICommandWriteDefaultErroneousDataReporting              KBluetoothHCIOpCodeNoOp = 0x5b
	KBluetoothHCICommandWriteDefaultLinkPolicySettings                  KBluetoothHCIOpCodeNoOp = 0xf
	KBluetoothHCICommandWriteEncryptionMode                             KBluetoothHCIOpCodeNoOp = 0x22
	KBluetoothHCICommandWriteExtendedInquiryLength                      KBluetoothHCIOpCodeNoOp = 0x81
	KBluetoothHCICommandWriteExtendedInquiryResponse                    KBluetoothHCIOpCodeNoOp = 0x52
	KBluetoothHCICommandWriteExtendedPageTimeout                        KBluetoothHCIOpCodeNoOp = 0x7f
	KBluetoothHCICommandWriteFlowControlMode                            KBluetoothHCIOpCodeNoOp = 0x67
	KBluetoothHCICommandWriteHoldModeActivity                           KBluetoothHCIOpCodeNoOp = 0x2c
	KBluetoothHCICommandWriteInquiryMode                                KBluetoothHCIOpCodeNoOp = 0x45
	KBluetoothHCICommandWriteInquiryResponseTransmitPower               KBluetoothHCIOpCodeNoOp = 0x59
	KBluetoothHCICommandWriteInquiryScanActivity                        KBluetoothHCIOpCodeNoOp = 0x1e
	KBluetoothHCICommandWriteInquiryScanType                            KBluetoothHCIOpCodeNoOp = 0x43
	KBluetoothHCICommandWriteLEHostSupported                            KBluetoothHCIOpCodeNoOp = 0x6d
	KBluetoothHCICommandWriteLinkPolicySettings                         KBluetoothHCIOpCodeNoOp = 0xd
	KBluetoothHCICommandWriteLinkSupervisionTimeout                     KBluetoothHCIOpCodeNoOp = 0x37
	KBluetoothHCICommandWriteLocationData                               KBluetoothHCIOpCodeNoOp = 0x65
	KBluetoothHCICommandWriteLogicalLinkAcceptTimeout                   KBluetoothHCIOpCodeNoOp = 0x62
	KBluetoothHCICommandWriteLoopbackMode                               KBluetoothHCIOpCodeNoOp = 0x2
	KBluetoothHCICommandWriteNumberOfBroadcastRetransmissions           KBluetoothHCIOpCodeNoOp = 0x2a
	KBluetoothHCICommandWritePINType                                    KBluetoothHCIOpCodeNoOp = 0xa
	KBluetoothHCICommandWritePageScanActivity                           KBluetoothHCIOpCodeNoOp = 0x1c
	KBluetoothHCICommandWritePageScanMode                               KBluetoothHCIOpCodeNoOp = 0x3e
	KBluetoothHCICommandWritePageScanPeriodMode                         KBluetoothHCIOpCodeNoOp = 0x3c
	KBluetoothHCICommandWritePageScanType                               KBluetoothHCIOpCodeNoOp = 0x47
	KBluetoothHCICommandWritePageTimeout                                KBluetoothHCIOpCodeNoOp = 0x18
	KBluetoothHCICommandWriteRemoteAMPASSOC                             KBluetoothHCIOpCodeNoOp = 0xb
	KBluetoothHCICommandWriteSCOFlowControlEnable                       KBluetoothHCIOpCodeNoOp = 0x2f
	KBluetoothHCICommandWriteScanEnable                                 KBluetoothHCIOpCodeNoOp = 0x1a
	KBluetoothHCICommandWriteSecureConnectionsHostSupport               KBluetoothHCIOpCodeNoOp = 0x7a
	KBluetoothHCICommandWriteSimplePairingDebugMode                     KBluetoothHCIOpCodeNoOp = 0x4
	KBluetoothHCICommandWriteSimplePairingMode                          KBluetoothHCIOpCodeNoOp = 0x56
	KBluetoothHCICommandWriteStoredLinkKey                              KBluetoothHCIOpCodeNoOp = 0x11
	KBluetoothHCICommandWriteSynchronizationTrainParameters             KBluetoothHCIOpCodeNoOp = 0x78
	KBluetoothHCICommandWriteVoiceSetting                               KBluetoothHCIOpCodeNoOp = 0x26
	KBluetoothHCIOpCodeNoOpValue                                        KBluetoothHCIOpCodeNoOp = 0
)

func (e KBluetoothHCIOpCodeNoOp) String() string {
	switch e {
	case KBluetoothHCICommandAMPTest:
		return "KBluetoothHCICommandAMPTest"
	case KBluetoothHCICommandAMPTestEnd:
		return "KBluetoothHCICommandAMPTestEnd"
	case KBluetoothHCICommandAcceptSniffRequest:
		return "KBluetoothHCICommandAcceptSniffRequest"
	case KBluetoothHCICommandAcceptSynchronousConnectionRequest:
		return "KBluetoothHCICommandAcceptSynchronousConnectionRequest"
	case KBluetoothHCICommandAddSCOConnection:
		return "KBluetoothHCICommandAddSCOConnection"
	case KBluetoothHCICommandAuthenticationRequested:
		return "KBluetoothHCICommandAuthenticationRequested"
	case KBluetoothHCICommandChangeConnectionLinkKey:
		return "KBluetoothHCICommandChangeConnectionLinkKey"
	case KBluetoothHCICommandChangeConnectionPacketType:
		return "KBluetoothHCICommandChangeConnectionPacketType"
	case KBluetoothHCICommandChangeLocalName:
		return "KBluetoothHCICommandChangeLocalName"
	case KBluetoothHCICommandCreateConnection:
		return "KBluetoothHCICommandCreateConnection"
	case KBluetoothHCICommandCreateNewUnitKey:
		return "KBluetoothHCICommandCreateNewUnitKey"
	case KBluetoothHCICommandDeleteReservedLTADDR:
		return "KBluetoothHCICommandDeleteReservedLTADDR"
	case KBluetoothHCICommandDeleteStoredLinkKey:
		return "KBluetoothHCICommandDeleteStoredLinkKey"
	case KBluetoothHCICommandDisconnect:
		return "KBluetoothHCICommandDisconnect"
	case KBluetoothHCICommandEnableDeviceUnderTestMode:
		return "KBluetoothHCICommandEnableDeviceUnderTestMode"
	case KBluetoothHCICommandEnhancedAcceptSynchronousConnectionRequest:
		return "KBluetoothHCICommandEnhancedAcceptSynchronousConnectionRequest"
	case KBluetoothHCICommandEnhancedFlush:
		return "KBluetoothHCICommandEnhancedFlush"
	case KBluetoothHCICommandEnhancedSetupSynchronousConnection:
		return "KBluetoothHCICommandEnhancedSetupSynchronousConnection"
	case KBluetoothHCICommandExitPeriodicInquiryMode:
		return "KBluetoothHCICommandExitPeriodicInquiryMode"
	case KBluetoothHCICommandFlowSpecification:
		return "KBluetoothHCICommandFlowSpecification"
	case KBluetoothHCICommandGetMWSTransportLayerConfiguration:
		return "KBluetoothHCICommandGetMWSTransportLayerConfiguration"
	case KBluetoothHCICommandGroupLinkControl:
		return "KBluetoothHCICommandGroupLinkControl"
	case KBluetoothHCICommandGroupLinkPolicy:
		return "KBluetoothHCICommandGroupLinkPolicy"
	case KBluetoothHCICommandGroupMax:
		return "KBluetoothHCICommandGroupMax"
	case KBluetoothHCICommandGroupNoOp:
		return "KBluetoothHCICommandGroupNoOp"
	case KBluetoothHCICommandGroupVendorSpecific:
		return "KBluetoothHCICommandGroupVendorSpecific"
	case KBluetoothHCICommandHostBufferSize:
		return "KBluetoothHCICommandHostBufferSize"
	case KBluetoothHCICommandHostNumberOfCompletedPackets:
		return "KBluetoothHCICommandHostNumberOfCompletedPackets"
	case KBluetoothHCICommandIOCapabilityRequestNegativeReply:
		return "KBluetoothHCICommandIOCapabilityRequestNegativeReply"
	case KBluetoothHCICommandIOCapabilityRequestReply:
		return "KBluetoothHCICommandIOCapabilityRequestReply"
	case KBluetoothHCICommandLEAddDeviceToPeriodicAdvertiserList:
		return "KBluetoothHCICommandLEAddDeviceToPeriodicAdvertiserList"
	case KBluetoothHCICommandLEAddDeviceToResolvingList:
		return "KBluetoothHCICommandLEAddDeviceToResolvingList"
	case KBluetoothHCICommandLEClearPeriodicAdvertiserList:
		return "KBluetoothHCICommandLEClearPeriodicAdvertiserList"
	case KBluetoothHCICommandLECreateConnection:
		return "KBluetoothHCICommandLECreateConnection"
	case KBluetoothHCICommandLECreateConnectionCancel:
		return "KBluetoothHCICommandLECreateConnectionCancel"
	case KBluetoothHCICommandLEEncrypt:
		return "KBluetoothHCICommandLEEncrypt"
	case KBluetoothHCICommandLEExtendedCreateConnection:
		return "KBluetoothHCICommandLEExtendedCreateConnection"
	case KBluetoothHCICommandLEGenerateDHKey:
		return "KBluetoothHCICommandLEGenerateDHKey"
	case KBluetoothHCICommandLELongTermKeyRequestNegativeReply:
		return "KBluetoothHCICommandLELongTermKeyRequestNegativeReply"
	case KBluetoothHCICommandLELongTermKeyRequestReply:
		return "KBluetoothHCICommandLELongTermKeyRequestReply"
	case KBluetoothHCICommandLEPeriodicAdvertisingCreateSync:
		return "KBluetoothHCICommandLEPeriodicAdvertisingCreateSync"
	case KBluetoothHCICommandLEPeriodicAdvertisingCreateSyncCancel:
		return "KBluetoothHCICommandLEPeriodicAdvertisingCreateSyncCancel"
	case KBluetoothHCICommandLEPeriodicAdvertisingTerminateSync:
		return "KBluetoothHCICommandLEPeriodicAdvertisingTerminateSync"
	case KBluetoothHCICommandLERand:
		return "KBluetoothHCICommandLERand"
	case KBluetoothHCICommandLEReadLocalP256PublicKey:
		return "KBluetoothHCICommandLEReadLocalP256PublicKey"
	case KBluetoothHCICommandLEReadLocalResolvableAddress:
		return "KBluetoothHCICommandLEReadLocalResolvableAddress"
	case KBluetoothHCICommandLEReadMaximumDataLength:
		return "KBluetoothHCICommandLEReadMaximumDataLength"
	case KBluetoothHCICommandLEReadNumberofSupportedAdvertisingSets:
		return "KBluetoothHCICommandLEReadNumberofSupportedAdvertisingSets"
	case KBluetoothHCICommandLEReadPeriodicAdvertiserListSize:
		return "KBluetoothHCICommandLEReadPeriodicAdvertiserListSize"
	case KBluetoothHCICommandLEReadPhy:
		return "KBluetoothHCICommandLEReadPhy"
	case KBluetoothHCICommandLEReadRFPathCompensation:
		return "KBluetoothHCICommandLEReadRFPathCompensation"
	case KBluetoothHCICommandLEReadRemoteUsedFeatures:
		return "KBluetoothHCICommandLEReadRemoteUsedFeatures"
	case KBluetoothHCICommandLEReadResolvingListSize:
		return "KBluetoothHCICommandLEReadResolvingListSize"
	case KBluetoothHCICommandLEReadSuggestedDefaultDataLength:
		return "KBluetoothHCICommandLEReadSuggestedDefaultDataLength"
	case KBluetoothHCICommandLEReadSupportedStates:
		return "KBluetoothHCICommandLEReadSupportedStates"
	case KBluetoothHCICommandLEReadTransmitPower:
		return "KBluetoothHCICommandLEReadTransmitPower"
	case KBluetoothHCICommandLEReceiverTest:
		return "KBluetoothHCICommandLEReceiverTest"
	case KBluetoothHCICommandLERemoteConnectionParameterRequestNegativeReply:
		return "KBluetoothHCICommandLERemoteConnectionParameterRequestNegativeReply"
	case KBluetoothHCICommandLERemoteConnectionParameterRequestReply:
		return "KBluetoothHCICommandLERemoteConnectionParameterRequestReply"
	case KBluetoothHCICommandLERemoveAdvertisingSet:
		return "KBluetoothHCICommandLERemoveAdvertisingSet"
	case KBluetoothHCICommandLERemoveDeviceFromPeriodicAdvertiserList:
		return "KBluetoothHCICommandLERemoveDeviceFromPeriodicAdvertiserList"
	case KBluetoothHCICommandLERemoveDeviceFromResolvingList:
		return "KBluetoothHCICommandLERemoveDeviceFromResolvingList"
	case KBluetoothHCICommandLESetAddressResolutionEnable:
		return "KBluetoothHCICommandLESetAddressResolutionEnable"
	case KBluetoothHCICommandLESetAdvertiseEnable:
		return "KBluetoothHCICommandLESetAdvertiseEnable"
	case KBluetoothHCICommandLESetDataLength:
		return "KBluetoothHCICommandLESetDataLength"
	case KBluetoothHCICommandLESetExtendedAdvertisingData:
		return "KBluetoothHCICommandLESetExtendedAdvertisingData"
	case KBluetoothHCICommandLESetExtendedAdvertisingEnableCommand:
		return "KBluetoothHCICommandLESetExtendedAdvertisingEnableCommand"
	case KBluetoothHCICommandLESetExtendedAdvertisingParameters:
		return "KBluetoothHCICommandLESetExtendedAdvertisingParameters"
	case KBluetoothHCICommandLESetExtendedScanEnable:
		return "KBluetoothHCICommandLESetExtendedScanEnable"
	case KBluetoothHCICommandLESetExtendedScanParameters:
		return "KBluetoothHCICommandLESetExtendedScanParameters"
	case KBluetoothHCICommandLESetExtendedScanResponseData:
		return "KBluetoothHCICommandLESetExtendedScanResponseData"
	case KBluetoothHCICommandLESetHostChannelClassification:
		return "KBluetoothHCICommandLESetHostChannelClassification"
	case KBluetoothHCICommandLESetPhy:
		return "KBluetoothHCICommandLESetPhy"
	case KBluetoothHCICommandLESetPrivacyMode:
		return "KBluetoothHCICommandLESetPrivacyMode"
	case KBluetoothHCICommandLESetResolvablePrivateAddressTimeout:
		return "KBluetoothHCICommandLESetResolvablePrivateAddressTimeout"
	case KBluetoothHCICommandLEStartEncryption:
		return "KBluetoothHCICommandLEStartEncryption"
	case KBluetoothHCICommandLETestEnd:
		return "KBluetoothHCICommandLETestEnd"
	case KBluetoothHCICommandLETransmitterTest:
		return "KBluetoothHCICommandLETransmitterTest"
	case KBluetoothHCICommandLEWriteRFPathCompensation:
		return "KBluetoothHCICommandLEWriteRFPathCompensation"
	case KBluetoothHCICommandLEWriteSuggestedDefaultDataLength:
		return "KBluetoothHCICommandLEWriteSuggestedDefaultDataLength"
	case KBluetoothHCICommandMax:
		return "KBluetoothHCICommandMax"
	case KBluetoothHCICommandReadAuthenticatedPayloadTimeout:
		return "KBluetoothHCICommandReadAuthenticatedPayloadTimeout"
	case KBluetoothHCICommandReadBestEffortFlushTimeout:
		return "KBluetoothHCICommandReadBestEffortFlushTimeout"
	case KBluetoothHCICommandReadDefaultErroneousDataReporting:
		return "KBluetoothHCICommandReadDefaultErroneousDataReporting"
	case KBluetoothHCICommandReadEnhancedTransmitPowerLevel:
		return "KBluetoothHCICommandReadEnhancedTransmitPowerLevel"
	case KBluetoothHCICommandReadExtendedInquiryLength:
		return "KBluetoothHCICommandReadExtendedInquiryLength"
	case KBluetoothHCICommandReadExtendedInquiryResponse:
		return "KBluetoothHCICommandReadExtendedInquiryResponse"
	case KBluetoothHCICommandReadExtendedPageTimeout:
		return "KBluetoothHCICommandReadExtendedPageTimeout"
	case KBluetoothHCICommandReadFlowControlMode:
		return "KBluetoothHCICommandReadFlowControlMode"
	case KBluetoothHCICommandReadInquiryResponseTransmitPower:
		return "KBluetoothHCICommandReadInquiryResponseTransmitPower"
	case KBluetoothHCICommandReadLEHostSupported:
		return "KBluetoothHCICommandReadLEHostSupported"
	case KBluetoothHCICommandReadLocalOOBData:
		return "KBluetoothHCICommandReadLocalOOBData"
	case KBluetoothHCICommandReadLocalOOBExtendedData:
		return "KBluetoothHCICommandReadLocalOOBExtendedData"
	case KBluetoothHCICommandReadLocationData:
		return "KBluetoothHCICommandReadLocationData"
	case KBluetoothHCICommandReadLogicalLinkAcceptTimeout:
		return "KBluetoothHCICommandReadLogicalLinkAcceptTimeout"
	case KBluetoothHCICommandReadSecureConnectionsHostSupport:
		return "KBluetoothHCICommandReadSecureConnectionsHostSupport"
	case KBluetoothHCICommandReadSimplePairingMode:
		return "KBluetoothHCICommandReadSimplePairingMode"
	case KBluetoothHCICommandReadSynchronizationTrainParameters:
		return "KBluetoothHCICommandReadSynchronizationTrainParameters"
	case KBluetoothHCICommandRefreshEncryptionKey:
		return "KBluetoothHCICommandRefreshEncryptionKey"
	case KBluetoothHCICommandSendKeypressNotification:
		return "KBluetoothHCICommandSendKeypressNotification"
	case KBluetoothHCICommandSetConnectionlessPeripheralBroadcastData:
		return "KBluetoothHCICommandSetConnectionlessPeripheralBroadcastData"
	case KBluetoothHCICommandSetEventMaskPageTwo:
		return "KBluetoothHCICommandSetEventMaskPageTwo"
	case KBluetoothHCICommandSetExternalFrameConfiguration:
		return "KBluetoothHCICommandSetExternalFrameConfiguration"
	case KBluetoothHCICommandSetMWSChannelParameters:
		return "KBluetoothHCICommandSetMWSChannelParameters"
	case KBluetoothHCICommandSetMWSPATTERNConfiguration:
		return "KBluetoothHCICommandSetMWSPATTERNConfiguration"
	case KBluetoothHCICommandSetMWSScanFrequencyTable:
		return "KBluetoothHCICommandSetMWSScanFrequencyTable"
	case KBluetoothHCICommandSetMWSSignaling:
		return "KBluetoothHCICommandSetMWSSignaling"
	case KBluetoothHCICommandSetMWSTransportLayer:
		return "KBluetoothHCICommandSetMWSTransportLayer"
	case KBluetoothHCICommandSetReservedLTADDR:
		return "KBluetoothHCICommandSetReservedLTADDR"
	case KBluetoothHCICommandShortRangeMode:
		return "KBluetoothHCICommandShortRangeMode"
	case KBluetoothHCICommandWriteAuthenticatedPayloadTimeout:
		return "KBluetoothHCICommandWriteAuthenticatedPayloadTimeout"
	case KBluetoothHCICommandWriteBestEffortFlushTimeout:
		return "KBluetoothHCICommandWriteBestEffortFlushTimeout"
	case KBluetoothHCICommandWriteCurrentIACLAP:
		return "KBluetoothHCICommandWriteCurrentIACLAP"
	case KBluetoothHCICommandWriteDefaultErroneousDataReporting:
		return "KBluetoothHCICommandWriteDefaultErroneousDataReporting"
	case KBluetoothHCICommandWriteExtendedInquiryLength:
		return "KBluetoothHCICommandWriteExtendedInquiryLength"
	case KBluetoothHCICommandWriteExtendedInquiryResponse:
		return "KBluetoothHCICommandWriteExtendedInquiryResponse"
	case KBluetoothHCICommandWriteExtendedPageTimeout:
		return "KBluetoothHCICommandWriteExtendedPageTimeout"
	case KBluetoothHCICommandWriteFlowControlMode:
		return "KBluetoothHCICommandWriteFlowControlMode"
	case KBluetoothHCICommandWriteInquiryResponseTransmitPower:
		return "KBluetoothHCICommandWriteInquiryResponseTransmitPower"
	case KBluetoothHCICommandWriteLEHostSupported:
		return "KBluetoothHCICommandWriteLEHostSupported"
	case KBluetoothHCICommandWriteLocationData:
		return "KBluetoothHCICommandWriteLocationData"
	case KBluetoothHCICommandWriteLogicalLinkAcceptTimeout:
		return "KBluetoothHCICommandWriteLogicalLinkAcceptTimeout"
	case KBluetoothHCICommandWriteSecureConnectionsHostSupport:
		return "KBluetoothHCICommandWriteSecureConnectionsHostSupport"
	case KBluetoothHCICommandWriteSimplePairingMode:
		return "KBluetoothHCICommandWriteSimplePairingMode"
	case KBluetoothHCICommandWriteSynchronizationTrainParameters:
		return "KBluetoothHCICommandWriteSynchronizationTrainParameters"
	default:
		return fmt.Sprintf("KBluetoothHCIOpCodeNoOp(%d)", e)
	}
}

type KBluetoothHCITransportUSB uint32

const (
	// KBluetoothHCITransportUSBClassCode: # Discussion
	KBluetoothHCITransportUSBClassCode KBluetoothHCITransportUSB = 0xe0
	// KBluetoothHCITransportUSBProtocolCode: # Discussion
	KBluetoothHCITransportUSBProtocolCode KBluetoothHCITransportUSB = 0x1
	// KBluetoothHCITransportUSBSubClassCode: # Discussion
	KBluetoothHCITransportUSBSubClassCode KBluetoothHCITransportUSB = 0x1
)

func (e KBluetoothHCITransportUSB) String() string {
	switch e {
	case KBluetoothHCITransportUSBClassCode:
		return "KBluetoothHCITransportUSBClassCode"
	case KBluetoothHCITransportUSBProtocolCode:
		return "KBluetoothHCITransportUSBProtocolCode"
	default:
		return fmt.Sprintf("KBluetoothHCITransportUSB(%d)", e)
	}
}

type KBluetoothKeyFlag uint32

const (
	KBluetoothKeyFlagSemiPermanent KBluetoothKeyFlag = 0
	KBluetoothKeyFlagTemporary     KBluetoothKeyFlag = 0x1
)

func (e KBluetoothKeyFlag) String() string {
	switch e {
	case KBluetoothKeyFlagSemiPermanent:
		return "KBluetoothKeyFlagSemiPermanent"
	case KBluetoothKeyFlagTemporary:
		return "KBluetoothKeyFlagTemporary"
	default:
		return fmt.Sprintf("KBluetoothKeyFlag(%d)", e)
	}
}

type KBluetoothKeyType uint32

const (
	KBluetoothKeyTypeAuthenticatedCombination       KBluetoothKeyType = 0x5
	KBluetoothKeyTypeAuthenticatedCombinationP256   KBluetoothKeyType = 0x8
	KBluetoothKeyTypeChangedCombination             KBluetoothKeyType = 0x6
	KBluetoothKeyTypeCombination                    KBluetoothKeyType = 0
	KBluetoothKeyTypeDebugCombination               KBluetoothKeyType = 0x3
	KBluetoothKeyTypeLocalUnit                      KBluetoothKeyType = 0x1
	KBluetoothKeyTypeRemoteUnit                     KBluetoothKeyType = 0x2
	KBluetoothKeyTypeUnauthenticatedCombination     KBluetoothKeyType = 0x4
	KBluetoothKeyTypeUnauthenticatedCombinationP256 KBluetoothKeyType = 0x7
)

func (e KBluetoothKeyType) String() string {
	switch e {
	case KBluetoothKeyTypeAuthenticatedCombination:
		return "KBluetoothKeyTypeAuthenticatedCombination"
	case KBluetoothKeyTypeAuthenticatedCombinationP256:
		return "KBluetoothKeyTypeAuthenticatedCombinationP256"
	case KBluetoothKeyTypeChangedCombination:
		return "KBluetoothKeyTypeChangedCombination"
	case KBluetoothKeyTypeCombination:
		return "KBluetoothKeyTypeCombination"
	case KBluetoothKeyTypeDebugCombination:
		return "KBluetoothKeyTypeDebugCombination"
	case KBluetoothKeyTypeLocalUnit:
		return "KBluetoothKeyTypeLocalUnit"
	case KBluetoothKeyTypeRemoteUnit:
		return "KBluetoothKeyTypeRemoteUnit"
	case KBluetoothKeyTypeUnauthenticatedCombination:
		return "KBluetoothKeyTypeUnauthenticatedCombination"
	case KBluetoothKeyTypeUnauthenticatedCombinationP256:
		return "KBluetoothKeyTypeUnauthenticatedCombinationP256"
	default:
		return fmt.Sprintf("KBluetoothKeyType(%d)", e)
	}
}

type KBluetoothL2CAPChannel uint32

const (
	KBluetoothL2CAPChannelAMPManagerProtocol   KBluetoothL2CAPChannel = 0x3
	KBluetoothL2CAPChannelAMPTestManager       KBluetoothL2CAPChannel = 0x3f
	KBluetoothL2CAPChannelAttributeProtocol    KBluetoothL2CAPChannel = 0x4
	KBluetoothL2CAPChannelBREDRSecurityManager KBluetoothL2CAPChannel = 0x7
	KBluetoothL2CAPChannelConnectionLessData   KBluetoothL2CAPChannel = 0x2
	KBluetoothL2CAPChannelDynamicEnd           KBluetoothL2CAPChannel = 0xffff
	KBluetoothL2CAPChannelDynamicStart         KBluetoothL2CAPChannel = 0x40
	KBluetoothL2CAPChannelEnd                  KBluetoothL2CAPChannel = 0xffff
	KBluetoothL2CAPChannelLEAP                 KBluetoothL2CAPChannel = 0x2a
	KBluetoothL2CAPChannelLEAS                 KBluetoothL2CAPChannel = 0x2b
	KBluetoothL2CAPChannelLESignalling         KBluetoothL2CAPChannel = 0x5
	KBluetoothL2CAPChannelMagicPairing         KBluetoothL2CAPChannel = 0x30
	KBluetoothL2CAPChannelMagnet               KBluetoothL2CAPChannel = 0x3a
	KBluetoothL2CAPChannelNull                 KBluetoothL2CAPChannel = 0
	KBluetoothL2CAPChannelReservedEnd          KBluetoothL2CAPChannel = 0x3e
	KBluetoothL2CAPChannelReservedStart        KBluetoothL2CAPChannel = 0x8
	KBluetoothL2CAPChannelSecurityManager      KBluetoothL2CAPChannel = 0x6
	KBluetoothL2CAPChannelSignalling           KBluetoothL2CAPChannel = 0x1
)

func (e KBluetoothL2CAPChannel) String() string {
	switch e {
	case KBluetoothL2CAPChannelAMPManagerProtocol:
		return "KBluetoothL2CAPChannelAMPManagerProtocol"
	case KBluetoothL2CAPChannelAMPTestManager:
		return "KBluetoothL2CAPChannelAMPTestManager"
	case KBluetoothL2CAPChannelAttributeProtocol:
		return "KBluetoothL2CAPChannelAttributeProtocol"
	case KBluetoothL2CAPChannelBREDRSecurityManager:
		return "KBluetoothL2CAPChannelBREDRSecurityManager"
	case KBluetoothL2CAPChannelConnectionLessData:
		return "KBluetoothL2CAPChannelConnectionLessData"
	case KBluetoothL2CAPChannelDynamicEnd:
		return "KBluetoothL2CAPChannelDynamicEnd"
	case KBluetoothL2CAPChannelDynamicStart:
		return "KBluetoothL2CAPChannelDynamicStart"
	case KBluetoothL2CAPChannelLEAP:
		return "KBluetoothL2CAPChannelLEAP"
	case KBluetoothL2CAPChannelLEAS:
		return "KBluetoothL2CAPChannelLEAS"
	case KBluetoothL2CAPChannelLESignalling:
		return "KBluetoothL2CAPChannelLESignalling"
	case KBluetoothL2CAPChannelMagicPairing:
		return "KBluetoothL2CAPChannelMagicPairing"
	case KBluetoothL2CAPChannelMagnet:
		return "KBluetoothL2CAPChannelMagnet"
	case KBluetoothL2CAPChannelNull:
		return "KBluetoothL2CAPChannelNull"
	case KBluetoothL2CAPChannelReservedEnd:
		return "KBluetoothL2CAPChannelReservedEnd"
	case KBluetoothL2CAPChannelReservedStart:
		return "KBluetoothL2CAPChannelReservedStart"
	case KBluetoothL2CAPChannelSecurityManager:
		return "KBluetoothL2CAPChannelSecurityManager"
	case KBluetoothL2CAPChannelSignalling:
		return "KBluetoothL2CAPChannelSignalling"
	default:
		return fmt.Sprintf("KBluetoothL2CAPChannel(%d)", e)
	}
}

type KBluetoothL2CAPConfigurationOptionMTULength uint32

const (
	KBluetoothL2CAPConfigurationOptionFlushTimeoutLength                 KBluetoothL2CAPConfigurationOptionMTULength = 2
	KBluetoothL2CAPConfigurationOptionMTULengthValue                     KBluetoothL2CAPConfigurationOptionMTULength = 2
	KBluetoothL2CAPConfigurationOptionQoSLength                          KBluetoothL2CAPConfigurationOptionMTULength = 22
	KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControlLength KBluetoothL2CAPConfigurationOptionMTULength = 9
)

func (e KBluetoothL2CAPConfigurationOptionMTULength) String() string {
	switch e {
	case KBluetoothL2CAPConfigurationOptionFlushTimeoutLength:
		return "KBluetoothL2CAPConfigurationOptionFlushTimeoutLength"
	case KBluetoothL2CAPConfigurationOptionQoSLength:
		return "KBluetoothL2CAPConfigurationOptionQoSLength"
	case KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControlLength:
		return "KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControlLength"
	default:
		return fmt.Sprintf("KBluetoothL2CAPConfigurationOptionMTULength(%d)", e)
	}
}

type KBluetoothL2CAPFlushTimeout uint32

const (
	KBluetoothL2CAPFlushTimeoutEnd         KBluetoothL2CAPFlushTimeout = 65536
	KBluetoothL2CAPFlushTimeoutForever     KBluetoothL2CAPFlushTimeout = 0xffff
	KBluetoothL2CAPFlushTimeoutImmediate   KBluetoothL2CAPFlushTimeout = 0x1
	KBluetoothL2CAPFlushTimeoutUseExisting KBluetoothL2CAPFlushTimeout = 0
)

func (e KBluetoothL2CAPFlushTimeout) String() string {
	switch e {
	case KBluetoothL2CAPFlushTimeoutEnd:
		return "KBluetoothL2CAPFlushTimeoutEnd"
	case KBluetoothL2CAPFlushTimeoutForever:
		return "KBluetoothL2CAPFlushTimeoutForever"
	case KBluetoothL2CAPFlushTimeoutImmediate:
		return "KBluetoothL2CAPFlushTimeoutImmediate"
	case KBluetoothL2CAPFlushTimeoutUseExisting:
		return "KBluetoothL2CAPFlushTimeoutUseExisting"
	default:
		return fmt.Sprintf("KBluetoothL2CAPFlushTimeout(%d)", e)
	}
}

const KBluetoothL2CAPInfoTypeMaxConnectionlessMTUSize uint32 = 0x1

type KBluetoothL2CAPMTULowEnergyDefault uint32

const (
	KBluetoothL2CAPFlushTimeoutDefault       KBluetoothL2CAPMTULowEnergyDefault = 65535
	KBluetoothL2CAPMTUDefault                KBluetoothL2CAPMTULowEnergyDefault = 0x3f9
	KBluetoothL2CAPMTULowEnergyDefaultValue  KBluetoothL2CAPMTULowEnergyDefault = 27
	KBluetoothL2CAPMTULowEnergyMax           KBluetoothL2CAPMTULowEnergyDefault = 251
	KBluetoothL2CAPMTUMaximum                KBluetoothL2CAPMTULowEnergyDefault = 0xffff
	KBluetoothL2CAPMTUMinimum                KBluetoothL2CAPMTULowEnergyDefault = 0x30
	KBluetoothL2CAPMTUSIG                    KBluetoothL2CAPMTULowEnergyDefault = 0x30
	KBluetoothL2CAPMTUStart                  KBluetoothL2CAPMTULowEnergyDefault = 0x7fff
	KBluetoothL2CAPQoSDelayVariationDefault  KBluetoothL2CAPMTULowEnergyDefault = 0xffffffff
	KBluetoothL2CAPQoSFlagsDefault           KBluetoothL2CAPMTULowEnergyDefault = 0
	KBluetoothL2CAPQoSLatencyDefault         KBluetoothL2CAPMTULowEnergyDefault = 0xffffffff
	KBluetoothL2CAPQoSPeakBandwidthDefault   KBluetoothL2CAPMTULowEnergyDefault = 0
	KBluetoothL2CAPQoSTokenBucketSizeDefault KBluetoothL2CAPMTULowEnergyDefault = 0
	KBluetoothL2CAPQoSTokenRateDefault       KBluetoothL2CAPMTULowEnergyDefault = 0
	KBluetoothL2CAPQoSTypeDefault            KBluetoothL2CAPMTULowEnergyDefault = 1
)

func (e KBluetoothL2CAPMTULowEnergyDefault) String() string {
	switch e {
	case KBluetoothL2CAPFlushTimeoutDefault:
		return "KBluetoothL2CAPFlushTimeoutDefault"
	case KBluetoothL2CAPMTUDefault:
		return "KBluetoothL2CAPMTUDefault"
	case KBluetoothL2CAPMTULowEnergyDefaultValue:
		return "KBluetoothL2CAPMTULowEnergyDefaultValue"
	case KBluetoothL2CAPMTULowEnergyMax:
		return "KBluetoothL2CAPMTULowEnergyMax"
	case KBluetoothL2CAPMTUMinimum:
		return "KBluetoothL2CAPMTUMinimum"
	case KBluetoothL2CAPMTUStart:
		return "KBluetoothL2CAPMTUStart"
	case KBluetoothL2CAPQoSDelayVariationDefault:
		return "KBluetoothL2CAPQoSDelayVariationDefault"
	case KBluetoothL2CAPQoSFlagsDefault:
		return "KBluetoothL2CAPQoSFlagsDefault"
	case KBluetoothL2CAPQoSTypeDefault:
		return "KBluetoothL2CAPQoSTypeDefault"
	default:
		return fmt.Sprintf("KBluetoothL2CAPMTULowEnergyDefault(%d)", e)
	}
}

type KBluetoothL2CAPMaxPacketSize uint32

const (
	KBluetoothACLLogicalChannelL2CAPContinue KBluetoothL2CAPMaxPacketSize = 1
	KBluetoothACLLogicalChannelL2CAPStart    KBluetoothL2CAPMaxPacketSize = 2
	KBluetoothACLLogicalChannelLMP           KBluetoothL2CAPMaxPacketSize = 3
	KBluetoothACLLogicalChannelReserved      KBluetoothL2CAPMaxPacketSize = 0
	KBluetoothL2CAPMaxPacketSizeValue        KBluetoothL2CAPMaxPacketSize = 65535
)

func (e KBluetoothL2CAPMaxPacketSize) String() string {
	switch e {
	case KBluetoothACLLogicalChannelL2CAPContinue:
		return "KBluetoothACLLogicalChannelL2CAPContinue"
	case KBluetoothACLLogicalChannelL2CAPStart:
		return "KBluetoothACLLogicalChannelL2CAPStart"
	case KBluetoothACLLogicalChannelLMP:
		return "KBluetoothACLLogicalChannelLMP"
	case KBluetoothACLLogicalChannelReserved:
		return "KBluetoothACLLogicalChannelReserved"
	case KBluetoothL2CAPMaxPacketSizeValue:
		return "KBluetoothL2CAPMaxPacketSizeValue"
	default:
		return fmt.Sprintf("KBluetoothL2CAPMaxPacketSize(%d)", e)
	}
}

type KBluetoothL2CAPPSM uint32

const (
	KBluetoothL2CAPPSMAACP             KBluetoothL2CAPPSM = 0x1001
	KBluetoothL2CAPPSMATT              KBluetoothL2CAPPSM = 0x1f
	KBluetoothL2CAPPSMAVCTP            KBluetoothL2CAPPSM = 0x17
	KBluetoothL2CAPPSMAVCTP_Browsing   KBluetoothL2CAPPSM = 0x1b
	KBluetoothL2CAPPSMAVDTP            KBluetoothL2CAPPSM = 0x19
	KBluetoothL2CAPPSMBNEP             KBluetoothL2CAPPSM = 0xf
	KBluetoothL2CAPPSMD2D              KBluetoothL2CAPPSM = 0x100f
	KBluetoothL2CAPPSMDynamicEnd       KBluetoothL2CAPPSM = 0xffff
	KBluetoothL2CAPPSMDynamicStart     KBluetoothL2CAPPSM = 0x1001
	KBluetoothL2CAPPSMHIDControl       KBluetoothL2CAPPSM = 0x11
	KBluetoothL2CAPPSMHIDInterrupt     KBluetoothL2CAPPSM = 0x13
	KBluetoothL2CAPPSMNone             KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMRFCOMM           KBluetoothL2CAPPSM = 0x3
	KBluetoothL2CAPPSMReservedEnd      KBluetoothL2CAPPSM = 0x1000
	KBluetoothL2CAPPSMReservedStart    KBluetoothL2CAPPSM = 0x1
	KBluetoothL2CAPPSMSDP              KBluetoothL2CAPPSM = 0x1
	KBluetoothL2CAPPSMTCS_BIN          KBluetoothL2CAPPSM = 0x5
	KBluetoothL2CAPPSMTCS_BIN_Cordless KBluetoothL2CAPPSM = 0x7
	KBluetoothL2CAPPSMUID_C_Plane      KBluetoothL2CAPPSM = 0x1d
)

func (e KBluetoothL2CAPPSM) String() string {
	switch e {
	case KBluetoothL2CAPPSMAACP:
		return "KBluetoothL2CAPPSMAACP"
	case KBluetoothL2CAPPSMATT:
		return "KBluetoothL2CAPPSMATT"
	case KBluetoothL2CAPPSMAVCTP:
		return "KBluetoothL2CAPPSMAVCTP"
	case KBluetoothL2CAPPSMAVCTP_Browsing:
		return "KBluetoothL2CAPPSMAVCTP_Browsing"
	case KBluetoothL2CAPPSMAVDTP:
		return "KBluetoothL2CAPPSMAVDTP"
	case KBluetoothL2CAPPSMBNEP:
		return "KBluetoothL2CAPPSMBNEP"
	case KBluetoothL2CAPPSMD2D:
		return "KBluetoothL2CAPPSMD2D"
	case KBluetoothL2CAPPSMDynamicEnd:
		return "KBluetoothL2CAPPSMDynamicEnd"
	case KBluetoothL2CAPPSMHIDControl:
		return "KBluetoothL2CAPPSMHIDControl"
	case KBluetoothL2CAPPSMHIDInterrupt:
		return "KBluetoothL2CAPPSMHIDInterrupt"
	case KBluetoothL2CAPPSMNone:
		return "KBluetoothL2CAPPSMNone"
	case KBluetoothL2CAPPSMRFCOMM:
		return "KBluetoothL2CAPPSMRFCOMM"
	case KBluetoothL2CAPPSMReservedEnd:
		return "KBluetoothL2CAPPSMReservedEnd"
	case KBluetoothL2CAPPSMReservedStart:
		return "KBluetoothL2CAPPSMReservedStart"
	case KBluetoothL2CAPPSMTCS_BIN:
		return "KBluetoothL2CAPPSMTCS_BIN"
	case KBluetoothL2CAPPSMTCS_BIN_Cordless:
		return "KBluetoothL2CAPPSMTCS_BIN_Cordless"
	case KBluetoothL2CAPPSMUID_C_Plane:
		return "KBluetoothL2CAPPSMUID_C_Plane"
	default:
		return fmt.Sprintf("KBluetoothL2CAPPSM(%d)", e)
	}
}

const KBluetoothL2CAPPacketHeaderSize uint32 = 4

type KBluetoothL2CAPTCICommand uint32

const (
	KBluetoothL2CAPTCICommandL2CA_ConfigReq         KBluetoothL2CAPTCICommand = 0x3
	KBluetoothL2CAPTCICommandL2CA_ConfigResp        KBluetoothL2CAPTCICommand = 0x13
	KBluetoothL2CAPTCICommandL2CA_ConnectReq        KBluetoothL2CAPTCICommand = 0x1
	KBluetoothL2CAPTCICommandL2CA_ConnectResp       KBluetoothL2CAPTCICommand = 0x11
	KBluetoothL2CAPTCICommandL2CA_DisableCLT        KBluetoothL2CAPTCICommand = 0x4
	KBluetoothL2CAPTCICommandL2CA_DisconnectReq     KBluetoothL2CAPTCICommand = 0x2
	KBluetoothL2CAPTCICommandL2CA_DisconnectResp    KBluetoothL2CAPTCICommand = 0x12
	KBluetoothL2CAPTCICommandL2CA_EnableCLT         KBluetoothL2CAPTCICommand = 0x5
	KBluetoothL2CAPTCICommandL2CA_GetInfo           KBluetoothL2CAPTCICommand = 0xe
	KBluetoothL2CAPTCICommandL2CA_GroupAddMember    KBluetoothL2CAPTCICommand = 0x8
	KBluetoothL2CAPTCICommandL2CA_GroupClose        KBluetoothL2CAPTCICommand = 0x7
	KBluetoothL2CAPTCICommandL2CA_GroupCreate       KBluetoothL2CAPTCICommand = 0x6
	KBluetoothL2CAPTCICommandL2CA_GroupMembership   KBluetoothL2CAPTCICommand = 0xa
	KBluetoothL2CAPTCICommandL2CA_GroupRemoveMember KBluetoothL2CAPTCICommand = 0x9
	KBluetoothL2CAPTCICommandL2CA_Ping              KBluetoothL2CAPTCICommand = 0xd
	KBluetoothL2CAPTCICommandL2CA_ReadData          KBluetoothL2CAPTCICommand = 0xc
	KBluetoothL2CAPTCICommandL2CA_Reserved1         KBluetoothL2CAPTCICommand = 0xf
	KBluetoothL2CAPTCICommandL2CA_Reserved2         KBluetoothL2CAPTCICommand = 0x10
	KBluetoothL2CAPTCICommandL2CA_WriteData         KBluetoothL2CAPTCICommand = 0xb
	KBluetoothL2CAPTCICommandReserved               KBluetoothL2CAPTCICommand = 0
)

func (e KBluetoothL2CAPTCICommand) String() string {
	switch e {
	case KBluetoothL2CAPTCICommandL2CA_ConfigReq:
		return "KBluetoothL2CAPTCICommandL2CA_ConfigReq"
	case KBluetoothL2CAPTCICommandL2CA_ConfigResp:
		return "KBluetoothL2CAPTCICommandL2CA_ConfigResp"
	case KBluetoothL2CAPTCICommandL2CA_ConnectReq:
		return "KBluetoothL2CAPTCICommandL2CA_ConnectReq"
	case KBluetoothL2CAPTCICommandL2CA_ConnectResp:
		return "KBluetoothL2CAPTCICommandL2CA_ConnectResp"
	case KBluetoothL2CAPTCICommandL2CA_DisableCLT:
		return "KBluetoothL2CAPTCICommandL2CA_DisableCLT"
	case KBluetoothL2CAPTCICommandL2CA_DisconnectReq:
		return "KBluetoothL2CAPTCICommandL2CA_DisconnectReq"
	case KBluetoothL2CAPTCICommandL2CA_DisconnectResp:
		return "KBluetoothL2CAPTCICommandL2CA_DisconnectResp"
	case KBluetoothL2CAPTCICommandL2CA_EnableCLT:
		return "KBluetoothL2CAPTCICommandL2CA_EnableCLT"
	case KBluetoothL2CAPTCICommandL2CA_GetInfo:
		return "KBluetoothL2CAPTCICommandL2CA_GetInfo"
	case KBluetoothL2CAPTCICommandL2CA_GroupAddMember:
		return "KBluetoothL2CAPTCICommandL2CA_GroupAddMember"
	case KBluetoothL2CAPTCICommandL2CA_GroupClose:
		return "KBluetoothL2CAPTCICommandL2CA_GroupClose"
	case KBluetoothL2CAPTCICommandL2CA_GroupCreate:
		return "KBluetoothL2CAPTCICommandL2CA_GroupCreate"
	case KBluetoothL2CAPTCICommandL2CA_GroupMembership:
		return "KBluetoothL2CAPTCICommandL2CA_GroupMembership"
	case KBluetoothL2CAPTCICommandL2CA_GroupRemoveMember:
		return "KBluetoothL2CAPTCICommandL2CA_GroupRemoveMember"
	case KBluetoothL2CAPTCICommandL2CA_Ping:
		return "KBluetoothL2CAPTCICommandL2CA_Ping"
	case KBluetoothL2CAPTCICommandL2CA_ReadData:
		return "KBluetoothL2CAPTCICommandL2CA_ReadData"
	case KBluetoothL2CAPTCICommandL2CA_Reserved1:
		return "KBluetoothL2CAPTCICommandL2CA_Reserved1"
	case KBluetoothL2CAPTCICommandL2CA_Reserved2:
		return "KBluetoothL2CAPTCICommandL2CA_Reserved2"
	case KBluetoothL2CAPTCICommandL2CA_WriteData:
		return "KBluetoothL2CAPTCICommandL2CA_WriteData"
	case KBluetoothL2CAPTCICommandReserved:
		return "KBluetoothL2CAPTCICommandReserved"
	default:
		return fmt.Sprintf("KBluetoothL2CAPTCICommand(%d)", e)
	}
}

type KBluetoothL2CAPTCIEventID uint32

const (
	KBluetoothL2CAPTCIEventIDL2CA_ConfigInd       KBluetoothL2CAPTCIEventID = 0x2
	KBluetoothL2CAPTCIEventIDL2CA_ConnectInd      KBluetoothL2CAPTCIEventID = 0x1
	KBluetoothL2CAPTCIEventIDL2CA_DisconnectInd   KBluetoothL2CAPTCIEventID = 0x3
	KBluetoothL2CAPTCIEventIDL2CA_QoSViolationInd KBluetoothL2CAPTCIEventID = 0x4
	KBluetoothL2CAPTCIEventIDL2CA_TimeOutInd      KBluetoothL2CAPTCIEventID = 0x5
	KBluetoothL2CAPTCIEventIDReserved             KBluetoothL2CAPTCIEventID = 0
)

func (e KBluetoothL2CAPTCIEventID) String() string {
	switch e {
	case KBluetoothL2CAPTCIEventIDL2CA_ConfigInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_ConfigInd"
	case KBluetoothL2CAPTCIEventIDL2CA_ConnectInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_ConnectInd"
	case KBluetoothL2CAPTCIEventIDL2CA_DisconnectInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_DisconnectInd"
	case KBluetoothL2CAPTCIEventIDL2CA_QoSViolationInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_QoSViolationInd"
	case KBluetoothL2CAPTCIEventIDL2CA_TimeOutInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_TimeOutInd"
	case KBluetoothL2CAPTCIEventIDReserved:
		return "KBluetoothL2CAPTCIEventIDReserved"
	default:
		return fmt.Sprintf("KBluetoothL2CAPTCIEventID(%d)", e)
	}
}

type KBluetoothLEMaxTX uint32

const (
	KBluetoothLEMaxTXOctetsDefault KBluetoothLEMaxTX = 0x80
	KBluetoothLEMaxTXOctetsMax     KBluetoothLEMaxTX = 0xfb
	KBluetoothLEMaxTXOctetsMin     KBluetoothLEMaxTX = 0x1b
	KBluetoothLEMaxTXTimeDefault   KBluetoothLEMaxTX = 27
	KBluetoothLEMaxTXTimeMax       KBluetoothLEMaxTX = 0x848
	KBluetoothLEMaxTXTimeMin       KBluetoothLEMaxTX = 0x148
)

func (e KBluetoothLEMaxTX) String() string {
	switch e {
	case KBluetoothLEMaxTXOctetsDefault:
		return "KBluetoothLEMaxTXOctetsDefault"
	case KBluetoothLEMaxTXOctetsMax:
		return "KBluetoothLEMaxTXOctetsMax"
	case KBluetoothLEMaxTXOctetsMin:
		return "KBluetoothLEMaxTXOctetsMin"
	case KBluetoothLEMaxTXTimeMax:
		return "KBluetoothLEMaxTXTimeMax"
	case KBluetoothLEMaxTXTimeMin:
		return "KBluetoothLEMaxTXTimeMin"
	default:
		return fmt.Sprintf("KBluetoothLEMaxTX(%d)", e)
	}
}

type KBluetoothLESecurityManagerNoBonding uint32

const (
	KBluetoothLESecurityManagerBonding        KBluetoothLESecurityManagerNoBonding = 1
	KBluetoothLESecurityManagerNoBondingValue KBluetoothLESecurityManagerNoBonding = 0
	KBluetoothLESecurityManagerReservedEnd    KBluetoothLESecurityManagerNoBonding = 3
	KBluetoothLESecurityManagerReservedStart  KBluetoothLESecurityManagerNoBonding = 2
)

func (e KBluetoothLESecurityManagerNoBonding) String() string {
	switch e {
	case KBluetoothLESecurityManagerBonding:
		return "KBluetoothLESecurityManagerBonding"
	case KBluetoothLESecurityManagerNoBondingValue:
		return "KBluetoothLESecurityManagerNoBondingValue"
	case KBluetoothLESecurityManagerReservedEnd:
		return "KBluetoothLESecurityManagerReservedEnd"
	case KBluetoothLESecurityManagerReservedStart:
		return "KBluetoothLESecurityManagerReservedStart"
	default:
		return fmt.Sprintf("KBluetoothLESecurityManagerNoBonding(%d)", e)
	}
}

type KBluetoothLETX uint32

const (
	KBluetoothLETXOctetsDefault KBluetoothLETX = 0x1b
	KBluetoothLETXOctetsMax     KBluetoothLETX = 0xfb
	KBluetoothLETXOctetsMin     KBluetoothLETX = 0x1b
	KBluetoothLETXTimeDefault   KBluetoothLETX = 0x148
	KBluetoothLETXTimeMax       KBluetoothLETX = 0x848
	KBluetoothLETXTimeMin       KBluetoothLETX = 0x148
)

func (e KBluetoothLETX) String() string {
	switch e {
	case KBluetoothLETXOctetsDefault:
		return "KBluetoothLETXOctetsDefault"
	case KBluetoothLETXOctetsMax:
		return "KBluetoothLETXOctetsMax"
	case KBluetoothLETXTimeDefault:
		return "KBluetoothLETXTimeDefault"
	case KBluetoothLETXTimeMax:
		return "KBluetoothLETXTimeMax"
	default:
		return fmt.Sprintf("KBluetoothLETX(%d)", e)
	}
}

type KBluetoothPacket uint32

const (
	KBluetoothPacketType2DH1Omit  KBluetoothPacket = 0x2
	KBluetoothPacketType2DH3Omit  KBluetoothPacket = 0x100
	KBluetoothPacketType2DH5Omit  KBluetoothPacket = 0x1000
	KBluetoothPacketType3DH1Omit  KBluetoothPacket = 0x4
	KBluetoothPacketType3DH3Omit  KBluetoothPacket = 0x200
	KBluetoothPacketType3DM5Omit  KBluetoothPacket = 0x2000
	KBluetoothPacketTypeAUX       KBluetoothPacket = 0x200
	KBluetoothPacketTypeDH1       KBluetoothPacket = 0x10
	KBluetoothPacketTypeDH3       KBluetoothPacket = 0x800
	KBluetoothPacketTypeDH5       KBluetoothPacket = 0x8000
	KBluetoothPacketTypeDM1       KBluetoothPacket = 0x8
	KBluetoothPacketTypeDM3       KBluetoothPacket = 0x400
	KBluetoothPacketTypeDM5       KBluetoothPacket = 0x4000
	KBluetoothPacketTypeDV        KBluetoothPacket = 0x100
	KBluetoothPacketTypeEnd       KBluetoothPacket = 32769
	KBluetoothPacketTypeHV1       KBluetoothPacket = 0x20
	KBluetoothPacketTypeHV2       KBluetoothPacket = 0x40
	KBluetoothPacketTypeHV3       KBluetoothPacket = 0x80
	KBluetoothPacketTypeReserved1 KBluetoothPacket = 0x1
)

func (e KBluetoothPacket) String() string {
	switch e {
	case KBluetoothPacketType2DH1Omit:
		return "KBluetoothPacketType2DH1Omit"
	case KBluetoothPacketType2DH3Omit:
		return "KBluetoothPacketType2DH3Omit"
	case KBluetoothPacketType2DH5Omit:
		return "KBluetoothPacketType2DH5Omit"
	case KBluetoothPacketType3DH1Omit:
		return "KBluetoothPacketType3DH1Omit"
	case KBluetoothPacketType3DH3Omit:
		return "KBluetoothPacketType3DH3Omit"
	case KBluetoothPacketType3DM5Omit:
		return "KBluetoothPacketType3DM5Omit"
	case KBluetoothPacketTypeDH1:
		return "KBluetoothPacketTypeDH1"
	case KBluetoothPacketTypeDH3:
		return "KBluetoothPacketTypeDH3"
	case KBluetoothPacketTypeDH5:
		return "KBluetoothPacketTypeDH5"
	case KBluetoothPacketTypeDM1:
		return "KBluetoothPacketTypeDM1"
	case KBluetoothPacketTypeDM3:
		return "KBluetoothPacketTypeDM3"
	case KBluetoothPacketTypeDM5:
		return "KBluetoothPacketTypeDM5"
	case KBluetoothPacketTypeEnd:
		return "KBluetoothPacketTypeEnd"
	case KBluetoothPacketTypeHV1:
		return "KBluetoothPacketTypeHV1"
	case KBluetoothPacketTypeHV2:
		return "KBluetoothPacketTypeHV2"
	case KBluetoothPacketTypeHV3:
		return "KBluetoothPacketTypeHV3"
	case KBluetoothPacketTypeReserved1:
		return "KBluetoothPacketTypeReserved1"
	default:
		return fmt.Sprintf("KBluetoothPacket(%d)", e)
	}
}

type KBluetoothPageScanMode uint32

const (
	KBluetoothPageScanModeMandatory KBluetoothPageScanMode = 0
	KBluetoothPageScanModeOptional1 KBluetoothPageScanMode = 0x1
	KBluetoothPageScanModeOptional2 KBluetoothPageScanMode = 0x2
	KBluetoothPageScanModeOptional3 KBluetoothPageScanMode = 0x3
)

func (e KBluetoothPageScanMode) String() string {
	switch e {
	case KBluetoothPageScanModeMandatory:
		return "KBluetoothPageScanModeMandatory"
	case KBluetoothPageScanModeOptional1:
		return "KBluetoothPageScanModeOptional1"
	case KBluetoothPageScanModeOptional2:
		return "KBluetoothPageScanModeOptional2"
	case KBluetoothPageScanModeOptional3:
		return "KBluetoothPageScanModeOptional3"
	default:
		return fmt.Sprintf("KBluetoothPageScanMode(%d)", e)
	}
}

type KBluetoothPageScanPeriodMode uint32

const (
	KBluetoothPageScanPeriodModeP0 KBluetoothPageScanPeriodMode = 0
	KBluetoothPageScanPeriodModeP1 KBluetoothPageScanPeriodMode = 0x1
	KBluetoothPageScanPeriodModeP2 KBluetoothPageScanPeriodMode = 0x2
)

func (e KBluetoothPageScanPeriodMode) String() string {
	switch e {
	case KBluetoothPageScanPeriodModeP0:
		return "KBluetoothPageScanPeriodModeP0"
	case KBluetoothPageScanPeriodModeP1:
		return "KBluetoothPageScanPeriodModeP1"
	case KBluetoothPageScanPeriodModeP2:
		return "KBluetoothPageScanPeriodModeP2"
	default:
		return fmt.Sprintf("KBluetoothPageScanPeriodMode(%d)", e)
	}
}

type KBluetoothPageScanRepetitionMode uint32

const (
	KBluetoothPageScanRepetitionModeR0 KBluetoothPageScanRepetitionMode = 0
	KBluetoothPageScanRepetitionModeR1 KBluetoothPageScanRepetitionMode = 0x1
	KBluetoothPageScanRepetitionModeR2 KBluetoothPageScanRepetitionMode = 0x2
)

func (e KBluetoothPageScanRepetitionMode) String() string {
	switch e {
	case KBluetoothPageScanRepetitionModeR0:
		return "KBluetoothPageScanRepetitionModeR0"
	case KBluetoothPageScanRepetitionModeR1:
		return "KBluetoothPageScanRepetitionModeR1"
	case KBluetoothPageScanRepetitionModeR2:
		return "KBluetoothPageScanRepetitionModeR2"
	default:
		return fmt.Sprintf("KBluetoothPageScanRepetitionMode(%d)", e)
	}
}

type KBluetoothRole uint32

const (
	KBluetoothRoleBecomeCentral    KBluetoothRole = 0
	KBluetoothRoleBecomeMaster     KBluetoothRole = 0
	KBluetoothRoleRemainPeripheral KBluetoothRole = 0x1
	KBluetoothRoleRemainSlave      KBluetoothRole = 1
)

func (e KBluetoothRole) String() string {
	switch e {
	case KBluetoothRoleBecomeCentral:
		return "KBluetoothRoleBecomeCentral"
	case KBluetoothRoleRemainPeripheral:
		return "KBluetoothRoleRemainPeripheral"
	default:
		return fmt.Sprintf("KBluetoothRole(%d)", e)
	}
}

type KBluetoothSDPDataElementType uint32

const (
	KBluetoothSDPDataElementTypeBoolean                KBluetoothSDPDataElementType = 5
	KBluetoothSDPDataElementTypeDataElementAlternative KBluetoothSDPDataElementType = 7
	KBluetoothSDPDataElementTypeDataElementSequence    KBluetoothSDPDataElementType = 6
	KBluetoothSDPDataElementTypeNil                    KBluetoothSDPDataElementType = 0
	KBluetoothSDPDataElementTypeReservedEnd            KBluetoothSDPDataElementType = 31
	KBluetoothSDPDataElementTypeReservedStart          KBluetoothSDPDataElementType = 9
	KBluetoothSDPDataElementTypeSignedInt              KBluetoothSDPDataElementType = 2
	KBluetoothSDPDataElementTypeString                 KBluetoothSDPDataElementType = 4
	KBluetoothSDPDataElementTypeURL                    KBluetoothSDPDataElementType = 8
	KBluetoothSDPDataElementTypeUUID                   KBluetoothSDPDataElementType = 3
	KBluetoothSDPDataElementTypeUnsignedInt            KBluetoothSDPDataElementType = 1
)

func (e KBluetoothSDPDataElementType) String() string {
	switch e {
	case KBluetoothSDPDataElementTypeBoolean:
		return "KBluetoothSDPDataElementTypeBoolean"
	case KBluetoothSDPDataElementTypeDataElementAlternative:
		return "KBluetoothSDPDataElementTypeDataElementAlternative"
	case KBluetoothSDPDataElementTypeDataElementSequence:
		return "KBluetoothSDPDataElementTypeDataElementSequence"
	case KBluetoothSDPDataElementTypeNil:
		return "KBluetoothSDPDataElementTypeNil"
	case KBluetoothSDPDataElementTypeReservedEnd:
		return "KBluetoothSDPDataElementTypeReservedEnd"
	case KBluetoothSDPDataElementTypeReservedStart:
		return "KBluetoothSDPDataElementTypeReservedStart"
	case KBluetoothSDPDataElementTypeSignedInt:
		return "KBluetoothSDPDataElementTypeSignedInt"
	case KBluetoothSDPDataElementTypeString:
		return "KBluetoothSDPDataElementTypeString"
	case KBluetoothSDPDataElementTypeURL:
		return "KBluetoothSDPDataElementTypeURL"
	case KBluetoothSDPDataElementTypeUUID:
		return "KBluetoothSDPDataElementTypeUUID"
	case KBluetoothSDPDataElementTypeUnsignedInt:
		return "KBluetoothSDPDataElementTypeUnsignedInt"
	default:
		return fmt.Sprintf("KBluetoothSDPDataElementType(%d)", e)
	}
}

type KBluetoothSDPErrorCode uint32

const (
	KBluetoothSDPErrorCodeInsufficientResources      KBluetoothSDPErrorCode = 0x6
	KBluetoothSDPErrorCodeInvalidContinuationState   KBluetoothSDPErrorCode = 0x5
	KBluetoothSDPErrorCodeInvalidPDUSize             KBluetoothSDPErrorCode = 0x4
	KBluetoothSDPErrorCodeInvalidRequestSyntax       KBluetoothSDPErrorCode = 0x3
	KBluetoothSDPErrorCodeInvalidSDPVersion          KBluetoothSDPErrorCode = 0x1
	KBluetoothSDPErrorCodeInvalidServiceRecordHandle KBluetoothSDPErrorCode = 0x2
	KBluetoothSDPErrorCodeReserved                   KBluetoothSDPErrorCode = 0
	KBluetoothSDPErrorCodeReservedEnd                KBluetoothSDPErrorCode = 0xffff
	KBluetoothSDPErrorCodeReservedStart              KBluetoothSDPErrorCode = 0x7
	KBluetoothSDPErrorCodeSuccess                    KBluetoothSDPErrorCode = 0
)

func (e KBluetoothSDPErrorCode) String() string {
	switch e {
	case KBluetoothSDPErrorCodeInsufficientResources:
		return "KBluetoothSDPErrorCodeInsufficientResources"
	case KBluetoothSDPErrorCodeInvalidContinuationState:
		return "KBluetoothSDPErrorCodeInvalidContinuationState"
	case KBluetoothSDPErrorCodeInvalidPDUSize:
		return "KBluetoothSDPErrorCodeInvalidPDUSize"
	case KBluetoothSDPErrorCodeInvalidRequestSyntax:
		return "KBluetoothSDPErrorCodeInvalidRequestSyntax"
	case KBluetoothSDPErrorCodeInvalidSDPVersion:
		return "KBluetoothSDPErrorCodeInvalidSDPVersion"
	case KBluetoothSDPErrorCodeInvalidServiceRecordHandle:
		return "KBluetoothSDPErrorCodeInvalidServiceRecordHandle"
	case KBluetoothSDPErrorCodeReserved:
		return "KBluetoothSDPErrorCodeReserved"
	case KBluetoothSDPErrorCodeReservedEnd:
		return "KBluetoothSDPErrorCodeReservedEnd"
	case KBluetoothSDPErrorCodeReservedStart:
		return "KBluetoothSDPErrorCodeReservedStart"
	default:
		return fmt.Sprintf("KBluetoothSDPErrorCode(%d)", e)
	}
}

type KBluetoothSDPPDUID uint32

const (
	KBluetoothSDPPDUIDErrorResponse                  KBluetoothSDPPDUID = 1
	KBluetoothSDPPDUIDReserved                       KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceAttributeRequest        KBluetoothSDPPDUID = 4
	KBluetoothSDPPDUIDServiceAttributeResponse       KBluetoothSDPPDUID = 5
	KBluetoothSDPPDUIDServiceSearchAttributeRequest  KBluetoothSDPPDUID = 6
	KBluetoothSDPPDUIDServiceSearchAttributeResponse KBluetoothSDPPDUID = 7
	KBluetoothSDPPDUIDServiceSearchRequest           KBluetoothSDPPDUID = 2
	KBluetoothSDPPDUIDServiceSearchResponse          KBluetoothSDPPDUID = 3
)

func (e KBluetoothSDPPDUID) String() string {
	switch e {
	case KBluetoothSDPPDUIDErrorResponse:
		return "KBluetoothSDPPDUIDErrorResponse"
	case KBluetoothSDPPDUIDReserved:
		return "KBluetoothSDPPDUIDReserved"
	case KBluetoothSDPPDUIDServiceAttributeRequest:
		return "KBluetoothSDPPDUIDServiceAttributeRequest"
	case KBluetoothSDPPDUIDServiceAttributeResponse:
		return "KBluetoothSDPPDUIDServiceAttributeResponse"
	case KBluetoothSDPPDUIDServiceSearchAttributeRequest:
		return "KBluetoothSDPPDUIDServiceSearchAttributeRequest"
	case KBluetoothSDPPDUIDServiceSearchAttributeResponse:
		return "KBluetoothSDPPDUIDServiceSearchAttributeResponse"
	case KBluetoothSDPPDUIDServiceSearchRequest:
		return "KBluetoothSDPPDUIDServiceSearchRequest"
	case KBluetoothSDPPDUIDServiceSearchResponse:
		return "KBluetoothSDPPDUIDServiceSearchResponse"
	default:
		return fmt.Sprintf("KBluetoothSDPPDUID(%d)", e)
	}
}

type KBluetoothSDPUUID16 uint32

const (
	KBluetoothSDPUUID16ATT                    KBluetoothSDPUUID16 = 0x7
	KBluetoothSDPUUID16AVCTP                  KBluetoothSDPUUID16 = 0x17
	KBluetoothSDPUUID16AVDTP                  KBluetoothSDPUUID16 = 0x19
	KBluetoothSDPUUID16BNEP                   KBluetoothSDPUUID16 = 0xf
	KBluetoothSDPUUID16Base                   KBluetoothSDPUUID16 = 0
	KBluetoothSDPUUID16CMPT                   KBluetoothSDPUUID16 = 0x1b
	KBluetoothSDPUUID16FTP                    KBluetoothSDPUUID16 = 0xa
	KBluetoothSDPUUID16HIDP                   KBluetoothSDPUUID16 = 0x11
	KBluetoothSDPUUID16HTTP                   KBluetoothSDPUUID16 = 0xc
	KBluetoothSDPUUID16HardcopyControlChannel KBluetoothSDPUUID16 = 0x12
	KBluetoothSDPUUID16HardcopyDataChannel    KBluetoothSDPUUID16 = 0x14
	KBluetoothSDPUUID16HardcopyNotification   KBluetoothSDPUUID16 = 0x16
	KBluetoothSDPUUID16IP                     KBluetoothSDPUUID16 = 0x9
	KBluetoothSDPUUID16L2CAP                  KBluetoothSDPUUID16 = 0x100
	KBluetoothSDPUUID16MCAPControlChannel     KBluetoothSDPUUID16 = 0x1e
	KBluetoothSDPUUID16MCAPDataChannel        KBluetoothSDPUUID16 = 0x1f
	KBluetoothSDPUUID16OBEX                   KBluetoothSDPUUID16 = 0x8
	KBluetoothSDPUUID16RFCOMM                 KBluetoothSDPUUID16 = 0x3
	KBluetoothSDPUUID16SDP                    KBluetoothSDPUUID16 = 0x1
	KBluetoothSDPUUID16TCP                    KBluetoothSDPUUID16 = 0x4
	KBluetoothSDPUUID16TCSAT                  KBluetoothSDPUUID16 = 0x6
	KBluetoothSDPUUID16TCSBIN                 KBluetoothSDPUUID16 = 0x5
	KBluetoothSDPUUID16UDI_C_Plane            KBluetoothSDPUUID16 = 0x1d
	KBluetoothSDPUUID16UDP                    KBluetoothSDPUUID16 = 0x2
	KBluetoothSDPUUID16UPNP                   KBluetoothSDPUUID16 = 0x10
	KBluetoothSDPUUID16WSP                    KBluetoothSDPUUID16 = 0xe
)

func (e KBluetoothSDPUUID16) String() string {
	switch e {
	case KBluetoothSDPUUID16ATT:
		return "KBluetoothSDPUUID16ATT"
	case KBluetoothSDPUUID16AVCTP:
		return "KBluetoothSDPUUID16AVCTP"
	case KBluetoothSDPUUID16AVDTP:
		return "KBluetoothSDPUUID16AVDTP"
	case KBluetoothSDPUUID16BNEP:
		return "KBluetoothSDPUUID16BNEP"
	case KBluetoothSDPUUID16Base:
		return "KBluetoothSDPUUID16Base"
	case KBluetoothSDPUUID16CMPT:
		return "KBluetoothSDPUUID16CMPT"
	case KBluetoothSDPUUID16FTP:
		return "KBluetoothSDPUUID16FTP"
	case KBluetoothSDPUUID16HIDP:
		return "KBluetoothSDPUUID16HIDP"
	case KBluetoothSDPUUID16HTTP:
		return "KBluetoothSDPUUID16HTTP"
	case KBluetoothSDPUUID16HardcopyControlChannel:
		return "KBluetoothSDPUUID16HardcopyControlChannel"
	case KBluetoothSDPUUID16HardcopyDataChannel:
		return "KBluetoothSDPUUID16HardcopyDataChannel"
	case KBluetoothSDPUUID16HardcopyNotification:
		return "KBluetoothSDPUUID16HardcopyNotification"
	case KBluetoothSDPUUID16IP:
		return "KBluetoothSDPUUID16IP"
	case KBluetoothSDPUUID16L2CAP:
		return "KBluetoothSDPUUID16L2CAP"
	case KBluetoothSDPUUID16MCAPControlChannel:
		return "KBluetoothSDPUUID16MCAPControlChannel"
	case KBluetoothSDPUUID16MCAPDataChannel:
		return "KBluetoothSDPUUID16MCAPDataChannel"
	case KBluetoothSDPUUID16OBEX:
		return "KBluetoothSDPUUID16OBEX"
	case KBluetoothSDPUUID16RFCOMM:
		return "KBluetoothSDPUUID16RFCOMM"
	case KBluetoothSDPUUID16SDP:
		return "KBluetoothSDPUUID16SDP"
	case KBluetoothSDPUUID16TCP:
		return "KBluetoothSDPUUID16TCP"
	case KBluetoothSDPUUID16TCSAT:
		return "KBluetoothSDPUUID16TCSAT"
	case KBluetoothSDPUUID16TCSBIN:
		return "KBluetoothSDPUUID16TCSBIN"
	case KBluetoothSDPUUID16UDI_C_Plane:
		return "KBluetoothSDPUUID16UDI_C_Plane"
	case KBluetoothSDPUUID16UDP:
		return "KBluetoothSDPUUID16UDP"
	case KBluetoothSDPUUID16UPNP:
		return "KBluetoothSDPUUID16UPNP"
	case KBluetoothSDPUUID16WSP:
		return "KBluetoothSDPUUID16WSP"
	default:
		return fmt.Sprintf("KBluetoothSDPUUID16(%d)", e)
	}
}

type KBluetoothServiceClassMajor uint32

const (
	KBluetoothServiceClassMajorAny                     KBluetoothServiceClassMajor = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	KBluetoothServiceClassMajorAudio                   KBluetoothServiceClassMajor = 0x100
	KBluetoothServiceClassMajorCapturing               KBluetoothServiceClassMajor = 0x40
	KBluetoothServiceClassMajorEnd                     KBluetoothServiceClassMajor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'f' // 'nonf'
	KBluetoothServiceClassMajorInformation             KBluetoothServiceClassMajor = 0x400
	KBluetoothServiceClassMajorLimitedDiscoverableMode KBluetoothServiceClassMajor = 0x1
	KBluetoothServiceClassMajorNetworking              KBluetoothServiceClassMajor = 0x10
	KBluetoothServiceClassMajorNone                    KBluetoothServiceClassMajor = 'n'<<24 | 'o'<<16 | 'n'<<8 | 'e' // 'none'
	KBluetoothServiceClassMajorObjectTransfer          KBluetoothServiceClassMajor = 0x80
	KBluetoothServiceClassMajorPositioning             KBluetoothServiceClassMajor = 0x8
	KBluetoothServiceClassMajorRendering               KBluetoothServiceClassMajor = 0x20
	KBluetoothServiceClassMajorReserved1               KBluetoothServiceClassMajor = 0x2
	KBluetoothServiceClassMajorReserved2               KBluetoothServiceClassMajor = 0x4
	KBluetoothServiceClassMajorTelephony               KBluetoothServiceClassMajor = 0x200
)

func (e KBluetoothServiceClassMajor) String() string {
	switch e {
	case KBluetoothServiceClassMajorAny:
		return "KBluetoothServiceClassMajorAny"
	case KBluetoothServiceClassMajorAudio:
		return "KBluetoothServiceClassMajorAudio"
	case KBluetoothServiceClassMajorCapturing:
		return "KBluetoothServiceClassMajorCapturing"
	case KBluetoothServiceClassMajorEnd:
		return "KBluetoothServiceClassMajorEnd"
	case KBluetoothServiceClassMajorInformation:
		return "KBluetoothServiceClassMajorInformation"
	case KBluetoothServiceClassMajorLimitedDiscoverableMode:
		return "KBluetoothServiceClassMajorLimitedDiscoverableMode"
	case KBluetoothServiceClassMajorNetworking:
		return "KBluetoothServiceClassMajorNetworking"
	case KBluetoothServiceClassMajorNone:
		return "KBluetoothServiceClassMajorNone"
	case KBluetoothServiceClassMajorObjectTransfer:
		return "KBluetoothServiceClassMajorObjectTransfer"
	case KBluetoothServiceClassMajorPositioning:
		return "KBluetoothServiceClassMajorPositioning"
	case KBluetoothServiceClassMajorRendering:
		return "KBluetoothServiceClassMajorRendering"
	case KBluetoothServiceClassMajorReserved1:
		return "KBluetoothServiceClassMajorReserved1"
	case KBluetoothServiceClassMajorReserved2:
		return "KBluetoothServiceClassMajorReserved2"
	case KBluetoothServiceClassMajorTelephony:
		return "KBluetoothServiceClassMajorTelephony"
	default:
		return fmt.Sprintf("KBluetoothServiceClassMajor(%d)", e)
	}
}

type KBluetoothSynchronousConnectionPacket uint32

const (
	KBluetoothSynchronousConnectionPacketType2EV3Omit  KBluetoothSynchronousConnectionPacket = 0x40
	KBluetoothSynchronousConnectionPacketType2EV5Omit  KBluetoothSynchronousConnectionPacket = 0x100
	KBluetoothSynchronousConnectionPacketType3EV3Omit  KBluetoothSynchronousConnectionPacket = 0x80
	KBluetoothSynchronousConnectionPacketType3EV5Omit  KBluetoothSynchronousConnectionPacket = 0x200
	KBluetoothSynchronousConnectionPacketTypeAll       KBluetoothSynchronousConnectionPacket = 0xffff
	KBluetoothSynchronousConnectionPacketTypeEV3       KBluetoothSynchronousConnectionPacket = 0x8
	KBluetoothSynchronousConnectionPacketTypeEV4       KBluetoothSynchronousConnectionPacket = 0x10
	KBluetoothSynchronousConnectionPacketTypeEV5       KBluetoothSynchronousConnectionPacket = 0x20
	KBluetoothSynchronousConnectionPacketTypeEnd       KBluetoothSynchronousConnectionPacket = 65536
	KBluetoothSynchronousConnectionPacketTypeFutureUse KBluetoothSynchronousConnectionPacket = 0xfc00
	KBluetoothSynchronousConnectionPacketTypeHV1       KBluetoothSynchronousConnectionPacket = 0x1
	KBluetoothSynchronousConnectionPacketTypeHV2       KBluetoothSynchronousConnectionPacket = 0x2
	KBluetoothSynchronousConnectionPacketTypeHV3       KBluetoothSynchronousConnectionPacket = 0x4
	KBluetoothSynchronousConnectionPacketTypeNone      KBluetoothSynchronousConnectionPacket = 0
)

func (e KBluetoothSynchronousConnectionPacket) String() string {
	switch e {
	case KBluetoothSynchronousConnectionPacketType2EV3Omit:
		return "KBluetoothSynchronousConnectionPacketType2EV3Omit"
	case KBluetoothSynchronousConnectionPacketType2EV5Omit:
		return "KBluetoothSynchronousConnectionPacketType2EV5Omit"
	case KBluetoothSynchronousConnectionPacketType3EV3Omit:
		return "KBluetoothSynchronousConnectionPacketType3EV3Omit"
	case KBluetoothSynchronousConnectionPacketType3EV5Omit:
		return "KBluetoothSynchronousConnectionPacketType3EV5Omit"
	case KBluetoothSynchronousConnectionPacketTypeAll:
		return "KBluetoothSynchronousConnectionPacketTypeAll"
	case KBluetoothSynchronousConnectionPacketTypeEV3:
		return "KBluetoothSynchronousConnectionPacketTypeEV3"
	case KBluetoothSynchronousConnectionPacketTypeEV4:
		return "KBluetoothSynchronousConnectionPacketTypeEV4"
	case KBluetoothSynchronousConnectionPacketTypeEV5:
		return "KBluetoothSynchronousConnectionPacketTypeEV5"
	case KBluetoothSynchronousConnectionPacketTypeEnd:
		return "KBluetoothSynchronousConnectionPacketTypeEnd"
	case KBluetoothSynchronousConnectionPacketTypeFutureUse:
		return "KBluetoothSynchronousConnectionPacketTypeFutureUse"
	case KBluetoothSynchronousConnectionPacketTypeHV1:
		return "KBluetoothSynchronousConnectionPacketTypeHV1"
	case KBluetoothSynchronousConnectionPacketTypeHV2:
		return "KBluetoothSynchronousConnectionPacketTypeHV2"
	case KBluetoothSynchronousConnectionPacketTypeHV3:
		return "KBluetoothSynchronousConnectionPacketTypeHV3"
	case KBluetoothSynchronousConnectionPacketTypeNone:
		return "KBluetoothSynchronousConnectionPacketTypeNone"
	default:
		return fmt.Sprintf("KBluetoothSynchronousConnectionPacket(%d)", e)
	}
}

type KBluetoothVoiceSettingAirCodingFormat uint32

const (
	KBluetoothVoiceSettingAirCodingFormatALaw            KBluetoothVoiceSettingAirCodingFormat = 0x2
	KBluetoothVoiceSettingAirCodingFormatCVSD            KBluetoothVoiceSettingAirCodingFormat = 0
	KBluetoothVoiceSettingAirCodingFormatMask            KBluetoothVoiceSettingAirCodingFormat = 0x3
	KBluetoothVoiceSettingAirCodingFormatTransparentData KBluetoothVoiceSettingAirCodingFormat = 0x3
	KBluetoothVoiceSettingAirCodingFormatULaw            KBluetoothVoiceSettingAirCodingFormat = 0x1
)

func (e KBluetoothVoiceSettingAirCodingFormat) String() string {
	switch e {
	case KBluetoothVoiceSettingAirCodingFormatALaw:
		return "KBluetoothVoiceSettingAirCodingFormatALaw"
	case KBluetoothVoiceSettingAirCodingFormatCVSD:
		return "KBluetoothVoiceSettingAirCodingFormatCVSD"
	case KBluetoothVoiceSettingAirCodingFormatMask:
		return "KBluetoothVoiceSettingAirCodingFormatMask"
	case KBluetoothVoiceSettingAirCodingFormatULaw:
		return "KBluetoothVoiceSettingAirCodingFormatULaw"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingAirCodingFormat(%d)", e)
	}
}

type KBluetoothVoiceSettingInputCoding uint32

const (
	KBluetoothVoiceSettingInputCodingALawInputCoding   KBluetoothVoiceSettingInputCoding = 0x200
	KBluetoothVoiceSettingInputCodingLinearInputCoding KBluetoothVoiceSettingInputCoding = 0
	KBluetoothVoiceSettingInputCodingMask              KBluetoothVoiceSettingInputCoding = 0x300
	KBluetoothVoiceSettingInputCodingULawInputCoding   KBluetoothVoiceSettingInputCoding = 0x100
)

func (e KBluetoothVoiceSettingInputCoding) String() string {
	switch e {
	case KBluetoothVoiceSettingInputCodingALawInputCoding:
		return "KBluetoothVoiceSettingInputCodingALawInputCoding"
	case KBluetoothVoiceSettingInputCodingLinearInputCoding:
		return "KBluetoothVoiceSettingInputCodingLinearInputCoding"
	case KBluetoothVoiceSettingInputCodingMask:
		return "KBluetoothVoiceSettingInputCodingMask"
	case KBluetoothVoiceSettingInputCodingULawInputCoding:
		return "KBluetoothVoiceSettingInputCodingULawInputCoding"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputCoding(%d)", e)
	}
}

type KBluetoothVoiceSettingInputData uint32

const (
	KBluetoothVoiceSettingInputDataFormat1sComplement  KBluetoothVoiceSettingInputData = 0
	KBluetoothVoiceSettingInputDataFormat2sComplement  KBluetoothVoiceSettingInputData = 0x40
	KBluetoothVoiceSettingInputDataFormatMask          KBluetoothVoiceSettingInputData = 0xc0
	KBluetoothVoiceSettingInputDataFormatSignMagnitude KBluetoothVoiceSettingInputData = 0x80
	KBluetoothVoiceSettingInputDataFormatUnsigned      KBluetoothVoiceSettingInputData = 0xc0
)

func (e KBluetoothVoiceSettingInputData) String() string {
	switch e {
	case KBluetoothVoiceSettingInputDataFormat1sComplement:
		return "KBluetoothVoiceSettingInputDataFormat1sComplement"
	case KBluetoothVoiceSettingInputDataFormat2sComplement:
		return "KBluetoothVoiceSettingInputDataFormat2sComplement"
	case KBluetoothVoiceSettingInputDataFormatMask:
		return "KBluetoothVoiceSettingInputDataFormatMask"
	case KBluetoothVoiceSettingInputDataFormatSignMagnitude:
		return "KBluetoothVoiceSettingInputDataFormatSignMagnitude"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputData(%d)", e)
	}
}

type KBluetoothVoiceSettingInputSample uint32

const (
	KBluetoothVoiceSettingInputSampleSize16Bit KBluetoothVoiceSettingInputSample = 0x20
	KBluetoothVoiceSettingInputSampleSize8Bit  KBluetoothVoiceSettingInputSample = 0
	KBluetoothVoiceSettingInputSampleSizeMask  KBluetoothVoiceSettingInputSample = 0x20
)

func (e KBluetoothVoiceSettingInputSample) String() string {
	switch e {
	case KBluetoothVoiceSettingInputSampleSize16Bit:
		return "KBluetoothVoiceSettingInputSampleSize16Bit"
	case KBluetoothVoiceSettingInputSampleSize8Bit:
		return "KBluetoothVoiceSettingInputSampleSize8Bit"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputSample(%d)", e)
	}
}

const KBluetoothVoiceSettingPCMBitPositionMask uint32 = 0x1c

const KMaximumNumberOfInquiryAccessCodes uint32 = 0x40

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXConnectFlagValues
type OBEXConnectFlagValues uint32

const (
	KOBEXConnectFlag1Reserved                       OBEXConnectFlagValues = 2
	KOBEXConnectFlag2Reserved                       OBEXConnectFlagValues = 4
	KOBEXConnectFlag3Reserved                       OBEXConnectFlagValues = 8
	KOBEXConnectFlag4Reserved                       OBEXConnectFlagValues = 16
	KOBEXConnectFlag5Reserved                       OBEXConnectFlagValues = 32
	KOBEXConnectFlag6Reserved                       OBEXConnectFlagValues = 64
	KOBEXConnectFlag7Reserved                       OBEXConnectFlagValues = 128
	KOBEXConnectFlagNone                            OBEXConnectFlagValues = 0
	KOBEXConnectFlagSupportMultipleItLMPConnections OBEXConnectFlagValues = 1
)

func (e OBEXConnectFlagValues) String() string {
	switch e {
	case KOBEXConnectFlag1Reserved:
		return "KOBEXConnectFlag1Reserved"
	case KOBEXConnectFlag2Reserved:
		return "KOBEXConnectFlag2Reserved"
	case KOBEXConnectFlag3Reserved:
		return "KOBEXConnectFlag3Reserved"
	case KOBEXConnectFlag4Reserved:
		return "KOBEXConnectFlag4Reserved"
	case KOBEXConnectFlag5Reserved:
		return "KOBEXConnectFlag5Reserved"
	case KOBEXConnectFlag6Reserved:
		return "KOBEXConnectFlag6Reserved"
	case KOBEXConnectFlag7Reserved:
		return "KOBEXConnectFlag7Reserved"
	case KOBEXConnectFlagNone:
		return "KOBEXConnectFlagNone"
	case KOBEXConnectFlagSupportMultipleItLMPConnections:
		return "KOBEXConnectFlagSupportMultipleItLMPConnections"
	default:
		return fmt.Sprintf("OBEXConnectFlagValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXErrorCodes
type OBEXErrorCodes int32

const (
	KOBEXBadArgumentError OBEXErrorCodes = -21854
	KOBEXBadRequestError  OBEXErrorCodes = -21856
	KOBEXCancelledError   OBEXErrorCodes = -21857
	KOBEXConflictError    OBEXErrorCodes = -21861
	// KOBEXErrorRangeMax: # Discussion
	KOBEXErrorRangeMax OBEXErrorCodes = -21899
	// KOBEXErrorRangeMin: # Discussion
	KOBEXErrorRangeMin                OBEXErrorCodes = -21850
	KOBEXForbiddenError               OBEXErrorCodes = -21858
	KOBEXGeneralError                 OBEXErrorCodes = -21850
	KOBEXInternalError                OBEXErrorCodes = -21853
	KOBEXMethodNotAllowedError        OBEXErrorCodes = -21862
	KOBEXNoResourcesError             OBEXErrorCodes = -21851
	KOBEXNotAcceptableError           OBEXErrorCodes = -21860
	KOBEXNotFoundError                OBEXErrorCodes = -21863
	KOBEXNotImplementedError          OBEXErrorCodes = -21864
	KOBEXPreconditionFailedError      OBEXErrorCodes = -21865
	KOBEXSessionAlreadyConnectedError OBEXErrorCodes = -21882
	KOBEXSessionBadRequestError       OBEXErrorCodes = -21877
	KOBEXSessionBadResponseError      OBEXErrorCodes = -21878
	KOBEXSessionBusyError             OBEXErrorCodes = -21875
	KOBEXSessionNoTransportError      OBEXErrorCodes = -21879
	KOBEXSessionNotConnectedError     OBEXErrorCodes = -21876
	KOBEXSessionTimeoutError          OBEXErrorCodes = -21881
	KOBEXSessionTransportDiedError    OBEXErrorCodes = -21880
	KOBEXSuccess                      OBEXErrorCodes = 0
	KOBEXTimeoutError                 OBEXErrorCodes = -21855
	KOBEXUnauthorizedError            OBEXErrorCodes = -21859
	KOBEXUnsupportedError             OBEXErrorCodes = -21852
)

func (e OBEXErrorCodes) String() string {
	switch e {
	case KOBEXBadArgumentError:
		return "KOBEXBadArgumentError"
	case KOBEXBadRequestError:
		return "KOBEXBadRequestError"
	case KOBEXCancelledError:
		return "KOBEXCancelledError"
	case KOBEXConflictError:
		return "KOBEXConflictError"
	case KOBEXErrorRangeMax:
		return "KOBEXErrorRangeMax"
	case KOBEXErrorRangeMin:
		return "KOBEXErrorRangeMin"
	case KOBEXForbiddenError:
		return "KOBEXForbiddenError"
	case KOBEXInternalError:
		return "KOBEXInternalError"
	case KOBEXMethodNotAllowedError:
		return "KOBEXMethodNotAllowedError"
	case KOBEXNoResourcesError:
		return "KOBEXNoResourcesError"
	case KOBEXNotAcceptableError:
		return "KOBEXNotAcceptableError"
	case KOBEXNotFoundError:
		return "KOBEXNotFoundError"
	case KOBEXNotImplementedError:
		return "KOBEXNotImplementedError"
	case KOBEXPreconditionFailedError:
		return "KOBEXPreconditionFailedError"
	case KOBEXSessionAlreadyConnectedError:
		return "KOBEXSessionAlreadyConnectedError"
	case KOBEXSessionBadRequestError:
		return "KOBEXSessionBadRequestError"
	case KOBEXSessionBadResponseError:
		return "KOBEXSessionBadResponseError"
	case KOBEXSessionBusyError:
		return "KOBEXSessionBusyError"
	case KOBEXSessionNoTransportError:
		return "KOBEXSessionNoTransportError"
	case KOBEXSessionNotConnectedError:
		return "KOBEXSessionNotConnectedError"
	case KOBEXSessionTimeoutError:
		return "KOBEXSessionTimeoutError"
	case KOBEXSessionTransportDiedError:
		return "KOBEXSessionTransportDiedError"
	case KOBEXSuccess:
		return "KOBEXSuccess"
	case KOBEXTimeoutError:
		return "KOBEXTimeoutError"
	case KOBEXUnauthorizedError:
		return "KOBEXUnauthorizedError"
	case KOBEXUnsupportedError:
		return "KOBEXUnsupportedError"
	default:
		return fmt.Sprintf("OBEXErrorCodes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXHeaderIdentifiers
type OBEXHeaderIdentifiers uint32

const (
	// KOBEXHeaderIDAppParameters: # Discussion
	KOBEXHeaderIDAppParameters OBEXHeaderIdentifiers = 0x4c
	// KOBEXHeaderIDAuthorizationChallenge: # Discussion
	KOBEXHeaderIDAuthorizationChallenge OBEXHeaderIdentifiers = 0x4d
	// KOBEXHeaderIDAuthorizationResponse: # Discussion
	KOBEXHeaderIDAuthorizationResponse OBEXHeaderIdentifiers = 0x4e
	// KOBEXHeaderIDBody: # Discussion
	KOBEXHeaderIDBody OBEXHeaderIdentifiers = 0x48
	// KOBEXHeaderIDConnectionID: # Discussion
	KOBEXHeaderIDConnectionID OBEXHeaderIdentifiers = 0xcb
	// KOBEXHeaderIDCount: # Discussion
	KOBEXHeaderIDCount OBEXHeaderIdentifiers = 0xc0
	// KOBEXHeaderIDDescription: # Discussion
	KOBEXHeaderIDDescription OBEXHeaderIdentifiers = 0x5
	// KOBEXHeaderIDEndOfBody: # Discussion
	KOBEXHeaderIDEndOfBody OBEXHeaderIdentifiers = 0x49
	// KOBEXHeaderIDHTTP: # Discussion
	KOBEXHeaderIDHTTP OBEXHeaderIdentifiers = 0x47
	// KOBEXHeaderIDLength: # Discussion
	KOBEXHeaderIDLength OBEXHeaderIdentifiers = 0xc3
	// KOBEXHeaderIDName: # Discussion
	KOBEXHeaderIDName OBEXHeaderIdentifiers = 0x1
	// KOBEXHeaderIDOBEX13CreatorID: # Discussion
	KOBEXHeaderIDOBEX13CreatorID OBEXHeaderIdentifiers = 0xcf
	// KOBEXHeaderIDOBEX13ObjectClass: # Discussion
	KOBEXHeaderIDOBEX13ObjectClass OBEXHeaderIdentifiers = 0x51
	// KOBEXHeaderIDOBEX13SessionParameters: # Discussion
	KOBEXHeaderIDOBEX13SessionParameters OBEXHeaderIdentifiers = 0x52
	// KOBEXHeaderIDOBEX13SessionSequenceNumber: # Discussion
	KOBEXHeaderIDOBEX13SessionSequenceNumber OBEXHeaderIdentifiers = 0x93
	// KOBEXHeaderIDOBEX13WANUUID: # Discussion
	KOBEXHeaderIDOBEX13WANUUID OBEXHeaderIdentifiers = 0x50
	// KOBEXHeaderIDObjectClass: # Discussion
	KOBEXHeaderIDObjectClass OBEXHeaderIdentifiers = 0x4f
	// KOBEXHeaderIDReservedRangeEnd: # Discussion
	KOBEXHeaderIDReservedRangeEnd OBEXHeaderIdentifiers = 0x2f
	// KOBEXHeaderIDReservedRangeStart: # Discussion
	KOBEXHeaderIDReservedRangeStart OBEXHeaderIdentifiers = 0x10
	// KOBEXHeaderIDTarget: # Discussion
	KOBEXHeaderIDTarget OBEXHeaderIdentifiers = 0x46
	// KOBEXHeaderIDTime4Byte: # Discussion
	KOBEXHeaderIDTime4Byte OBEXHeaderIdentifiers = 0xc4
	// KOBEXHeaderIDTimeISO: # Discussion
	KOBEXHeaderIDTimeISO OBEXHeaderIdentifiers = 0x44
	// KOBEXHeaderIDType: # Discussion
	KOBEXHeaderIDType OBEXHeaderIdentifiers = 0x42
	// KOBEXHeaderIDUserDefinedRangeEnd: # Discussion
	KOBEXHeaderIDUserDefinedRangeEnd OBEXHeaderIdentifiers = 0x3f
	// KOBEXHeaderIDUserDefinedRangeStart: # Discussion
	KOBEXHeaderIDUserDefinedRangeStart OBEXHeaderIdentifiers = 0x30
	// KOBEXHeaderIDWho: # Discussion
	KOBEXHeaderIDWho OBEXHeaderIdentifiers = 0x4a
)

func (e OBEXHeaderIdentifiers) String() string {
	switch e {
	case KOBEXHeaderIDAppParameters:
		return "KOBEXHeaderIDAppParameters"
	case KOBEXHeaderIDAuthorizationChallenge:
		return "KOBEXHeaderIDAuthorizationChallenge"
	case KOBEXHeaderIDAuthorizationResponse:
		return "KOBEXHeaderIDAuthorizationResponse"
	case KOBEXHeaderIDBody:
		return "KOBEXHeaderIDBody"
	case KOBEXHeaderIDConnectionID:
		return "KOBEXHeaderIDConnectionID"
	case KOBEXHeaderIDCount:
		return "KOBEXHeaderIDCount"
	case KOBEXHeaderIDDescription:
		return "KOBEXHeaderIDDescription"
	case KOBEXHeaderIDEndOfBody:
		return "KOBEXHeaderIDEndOfBody"
	case KOBEXHeaderIDHTTP:
		return "KOBEXHeaderIDHTTP"
	case KOBEXHeaderIDLength:
		return "KOBEXHeaderIDLength"
	case KOBEXHeaderIDName:
		return "KOBEXHeaderIDName"
	case KOBEXHeaderIDOBEX13CreatorID:
		return "KOBEXHeaderIDOBEX13CreatorID"
	case KOBEXHeaderIDOBEX13ObjectClass:
		return "KOBEXHeaderIDOBEX13ObjectClass"
	case KOBEXHeaderIDOBEX13SessionParameters:
		return "KOBEXHeaderIDOBEX13SessionParameters"
	case KOBEXHeaderIDOBEX13SessionSequenceNumber:
		return "KOBEXHeaderIDOBEX13SessionSequenceNumber"
	case KOBEXHeaderIDOBEX13WANUUID:
		return "KOBEXHeaderIDOBEX13WANUUID"
	case KOBEXHeaderIDObjectClass:
		return "KOBEXHeaderIDObjectClass"
	case KOBEXHeaderIDReservedRangeEnd:
		return "KOBEXHeaderIDReservedRangeEnd"
	case KOBEXHeaderIDReservedRangeStart:
		return "KOBEXHeaderIDReservedRangeStart"
	case KOBEXHeaderIDTarget:
		return "KOBEXHeaderIDTarget"
	case KOBEXHeaderIDTime4Byte:
		return "KOBEXHeaderIDTime4Byte"
	case KOBEXHeaderIDTimeISO:
		return "KOBEXHeaderIDTimeISO"
	case KOBEXHeaderIDType:
		return "KOBEXHeaderIDType"
	case KOBEXHeaderIDUserDefinedRangeEnd:
		return "KOBEXHeaderIDUserDefinedRangeEnd"
	case KOBEXHeaderIDUserDefinedRangeStart:
		return "KOBEXHeaderIDUserDefinedRangeStart"
	case KOBEXHeaderIDWho:
		return "KOBEXHeaderIDWho"
	default:
		return fmt.Sprintf("OBEXHeaderIdentifiers(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXNonceFlagValues
type OBEXNonceFlagValues uint32

const (
	KOBEXNonceFlag2Reserved            OBEXNonceFlagValues = 4
	KOBEXNonceFlag3Reserved            OBEXNonceFlagValues = 8
	KOBEXNonceFlag4Reserved            OBEXNonceFlagValues = 16
	KOBEXNonceFlag5Reserved            OBEXNonceFlagValues = 32
	KOBEXNonceFlag6Reserved            OBEXNonceFlagValues = 64
	KOBEXNonceFlag7Reserved            OBEXNonceFlagValues = 128
	KOBEXNonceFlagAccessModeReadOnly   OBEXNonceFlagValues = 2
	KOBEXNonceFlagNone                 OBEXNonceFlagValues = 0
	KOBEXNonceFlagSendUserIDInResponse OBEXNonceFlagValues = 1
)

func (e OBEXNonceFlagValues) String() string {
	switch e {
	case KOBEXNonceFlag2Reserved:
		return "KOBEXNonceFlag2Reserved"
	case KOBEXNonceFlag3Reserved:
		return "KOBEXNonceFlag3Reserved"
	case KOBEXNonceFlag4Reserved:
		return "KOBEXNonceFlag4Reserved"
	case KOBEXNonceFlag5Reserved:
		return "KOBEXNonceFlag5Reserved"
	case KOBEXNonceFlag6Reserved:
		return "KOBEXNonceFlag6Reserved"
	case KOBEXNonceFlag7Reserved:
		return "KOBEXNonceFlag7Reserved"
	case KOBEXNonceFlagAccessModeReadOnly:
		return "KOBEXNonceFlagAccessModeReadOnly"
	case KOBEXNonceFlagNone:
		return "KOBEXNonceFlagNone"
	case KOBEXNonceFlagSendUserIDInResponse:
		return "KOBEXNonceFlagSendUserIDInResponse"
	default:
		return fmt.Sprintf("OBEXNonceFlagValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXOpCodeCommandValues
type OBEXOpCodeCommandValues uint32

const (
	KOBEXOpCodeAbort                  OBEXOpCodeCommandValues = 0xff
	KOBEXOpCodeConnect                OBEXOpCodeCommandValues = 0x80
	KOBEXOpCodeDisconnect             OBEXOpCodeCommandValues = 0x81
	KOBEXOpCodeGet                    OBEXOpCodeCommandValues = 0x3
	KOBEXOpCodeGetWithHighBitSet      OBEXOpCodeCommandValues = 0x83
	KOBEXOpCodePut                    OBEXOpCodeCommandValues = 0x2
	KOBEXOpCodePutWithHighBitSet      OBEXOpCodeCommandValues = 0x82
	KOBEXOpCodeReserved               OBEXOpCodeCommandValues = 0x4
	KOBEXOpCodeReservedRangeEnd       OBEXOpCodeCommandValues = 0xf
	KOBEXOpCodeReservedRangeStart     OBEXOpCodeCommandValues = 0x6
	KOBEXOpCodeReservedWithHighBitSet OBEXOpCodeCommandValues = 0x84
	KOBEXOpCodeSetPath                OBEXOpCodeCommandValues = 0x85
	KOBEXOpCodeUserDefinedEnd         OBEXOpCodeCommandValues = 0x1f
	KOBEXOpCodeUserDefinedStart       OBEXOpCodeCommandValues = 0x10
)

func (e OBEXOpCodeCommandValues) String() string {
	switch e {
	case KOBEXOpCodeAbort:
		return "KOBEXOpCodeAbort"
	case KOBEXOpCodeConnect:
		return "KOBEXOpCodeConnect"
	case KOBEXOpCodeDisconnect:
		return "KOBEXOpCodeDisconnect"
	case KOBEXOpCodeGet:
		return "KOBEXOpCodeGet"
	case KOBEXOpCodeGetWithHighBitSet:
		return "KOBEXOpCodeGetWithHighBitSet"
	case KOBEXOpCodePut:
		return "KOBEXOpCodePut"
	case KOBEXOpCodePutWithHighBitSet:
		return "KOBEXOpCodePutWithHighBitSet"
	case KOBEXOpCodeReserved:
		return "KOBEXOpCodeReserved"
	case KOBEXOpCodeReservedRangeEnd:
		return "KOBEXOpCodeReservedRangeEnd"
	case KOBEXOpCodeReservedRangeStart:
		return "KOBEXOpCodeReservedRangeStart"
	case KOBEXOpCodeReservedWithHighBitSet:
		return "KOBEXOpCodeReservedWithHighBitSet"
	case KOBEXOpCodeSetPath:
		return "KOBEXOpCodeSetPath"
	case KOBEXOpCodeUserDefinedEnd:
		return "KOBEXOpCodeUserDefinedEnd"
	case KOBEXOpCodeUserDefinedStart:
		return "KOBEXOpCodeUserDefinedStart"
	default:
		return fmt.Sprintf("OBEXOpCodeCommandValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXOpCodeResponseValues
type OBEXOpCodeResponseValues uint32

const (
	KOBEXResponseCodeAccepted                                OBEXOpCodeResponseValues = 0x22
	KOBEXResponseCodeAcceptedWithFinalBit                    OBEXOpCodeResponseValues = 0xa2
	KOBEXResponseCodeBadGateway                              OBEXOpCodeResponseValues = 0x52
	KOBEXResponseCodeBadGatewayWithFinalBit                  OBEXOpCodeResponseValues = 0xd2
	KOBEXResponseCodeBadRequest                              OBEXOpCodeResponseValues = 0x40
	KOBEXResponseCodeBadRequestWithFinalBit                  OBEXOpCodeResponseValues = 0xc0
	KOBEXResponseCodeConflict                                OBEXOpCodeResponseValues = 0x49
	KOBEXResponseCodeConflictWithFinalBit                    OBEXOpCodeResponseValues = 0xc9
	KOBEXResponseCodeContinue                                OBEXOpCodeResponseValues = 0x10
	KOBEXResponseCodeContinueWithFinalBit                    OBEXOpCodeResponseValues = 0x90
	KOBEXResponseCodeCreated                                 OBEXOpCodeResponseValues = 0x21
	KOBEXResponseCodeCreatedWithFinalBit                     OBEXOpCodeResponseValues = 0xa1
	KOBEXResponseCodeDatabaseFull                            OBEXOpCodeResponseValues = 0x60
	KOBEXResponseCodeDatabaseFullWithFinalBit                OBEXOpCodeResponseValues = 0xe0
	KOBEXResponseCodeDatabaseLocked                          OBEXOpCodeResponseValues = 0x61
	KOBEXResponseCodeDatabaseLockedWithFinalBit              OBEXOpCodeResponseValues = 0xe1
	KOBEXResponseCodeForbidden                               OBEXOpCodeResponseValues = 0x43
	KOBEXResponseCodeForbiddenWithFinalBit                   OBEXOpCodeResponseValues = 0xc3
	KOBEXResponseCodeGatewayTimeout                          OBEXOpCodeResponseValues = 0x54
	KOBEXResponseCodeGatewayTimeoutWithFinalBit              OBEXOpCodeResponseValues = 0xd4
	KOBEXResponseCodeGone                                    OBEXOpCodeResponseValues = 0x4a
	KOBEXResponseCodeGoneWithFinalBit                        OBEXOpCodeResponseValues = 0xca
	KOBEXResponseCodeHTTPVersionNotSupported                 OBEXOpCodeResponseValues = 0x55
	KOBEXResponseCodeHTTPVersionNotSupportedWithFinalBit     OBEXOpCodeResponseValues = 0xd5
	KOBEXResponseCodeInternalServerError                     OBEXOpCodeResponseValues = 0x50
	KOBEXResponseCodeInternalServerErrorWithFinalBit         OBEXOpCodeResponseValues = 0xd0
	KOBEXResponseCodeLengthRequired                          OBEXOpCodeResponseValues = 0x4b
	KOBEXResponseCodeLengthRequiredFinalBit                  OBEXOpCodeResponseValues = 0xcb
	KOBEXResponseCodeMethodNotAllowed                        OBEXOpCodeResponseValues = 0x45
	KOBEXResponseCodeMethodNotAllowedWithFinalBit            OBEXOpCodeResponseValues = 0xc5
	KOBEXResponseCodeMovedPermanently                        OBEXOpCodeResponseValues = 0x31
	KOBEXResponseCodeMovedPermanentlyWithFinalBit            OBEXOpCodeResponseValues = 0xb1
	KOBEXResponseCodeMovedTemporarily                        OBEXOpCodeResponseValues = 0x32
	KOBEXResponseCodeMovedTemporarilyWithFinalBit            OBEXOpCodeResponseValues = 0xb2
	KOBEXResponseCodeMultipleChoices                         OBEXOpCodeResponseValues = 0x30
	KOBEXResponseCodeMultipleChoicesWithFinalBit             OBEXOpCodeResponseValues = 0xb0
	KOBEXResponseCodeNoContent                               OBEXOpCodeResponseValues = 0x24
	KOBEXResponseCodeNoContentWithFinalBit                   OBEXOpCodeResponseValues = 0xa4
	KOBEXResponseCodeNonAuthoritativeInfo                    OBEXOpCodeResponseValues = 0x23
	KOBEXResponseCodeNonAuthoritativeInfoWithFinalBit        OBEXOpCodeResponseValues = 0xa3
	KOBEXResponseCodeNotAcceptable                           OBEXOpCodeResponseValues = 0x46
	KOBEXResponseCodeNotAcceptableWithFinalBit               OBEXOpCodeResponseValues = 0xc6
	KOBEXResponseCodeNotFound                                OBEXOpCodeResponseValues = 0x44
	KOBEXResponseCodeNotFoundWithFinalBit                    OBEXOpCodeResponseValues = 0xc4
	KOBEXResponseCodeNotImplemented                          OBEXOpCodeResponseValues = 0x51
	KOBEXResponseCodeNotImplementedWithFinalBit              OBEXOpCodeResponseValues = 0xd1
	KOBEXResponseCodeNotModified                             OBEXOpCodeResponseValues = 0x34
	KOBEXResponseCodeNotModifiedWithFinalBit                 OBEXOpCodeResponseValues = 0xb4
	KOBEXResponseCodePartialContent                          OBEXOpCodeResponseValues = 0x26
	KOBEXResponseCodePartialContentWithFinalBit              OBEXOpCodeResponseValues = 0xa6
	KOBEXResponseCodePaymentRequired                         OBEXOpCodeResponseValues = 0x42
	KOBEXResponseCodePaymentRequiredWithFinalBit             OBEXOpCodeResponseValues = 0xc2
	KOBEXResponseCodePreconditionFailed                      OBEXOpCodeResponseValues = 0x4c
	KOBEXResponseCodePreconditionFailedWithFinalBit          OBEXOpCodeResponseValues = 0xcc
	KOBEXResponseCodeProxyAuthenticationRequired             OBEXOpCodeResponseValues = 0x47
	KOBEXResponseCodeProxyAuthenticationRequiredWithFinalBit OBEXOpCodeResponseValues = 0xc7
	KOBEXResponseCodeRequestTimeOut                          OBEXOpCodeResponseValues = 0x48
	KOBEXResponseCodeRequestTimeOutWithFinalBit              OBEXOpCodeResponseValues = 0xc8
	KOBEXResponseCodeRequestURLTooLarge                      OBEXOpCodeResponseValues = 0x4e
	KOBEXResponseCodeRequestURLTooLargeWithFinalBit          OBEXOpCodeResponseValues = 0xce
	KOBEXResponseCodeRequestedEntityTooLarge                 OBEXOpCodeResponseValues = 0x4d
	KOBEXResponseCodeRequestedEntityTooLargeWithFinalBit     OBEXOpCodeResponseValues = 0xcd
	KOBEXResponseCodeReservedRangeEnd                        OBEXOpCodeResponseValues = 0xf
	KOBEXResponseCodeReservedRangeStart                      OBEXOpCodeResponseValues = 0
	KOBEXResponseCodeResetContent                            OBEXOpCodeResponseValues = 0x25
	KOBEXResponseCodeResetContentWithFinalBit                OBEXOpCodeResponseValues = 0xa5
	KOBEXResponseCodeSeeOther                                OBEXOpCodeResponseValues = 0x33
	KOBEXResponseCodeSeeOtherWithFinalBit                    OBEXOpCodeResponseValues = 0xb3
	KOBEXResponseCodeServiceUnavailable                      OBEXOpCodeResponseValues = 0x53
	KOBEXResponseCodeServiceUnavailableWithFinalBit          OBEXOpCodeResponseValues = 0xd3
	KOBEXResponseCodeSuccess                                 OBEXOpCodeResponseValues = 0x20
	KOBEXResponseCodeSuccessWithFinalBit                     OBEXOpCodeResponseValues = 0xa0
	KOBEXResponseCodeUnauthorized                            OBEXOpCodeResponseValues = 0x41
	KOBEXResponseCodeUnauthorizedWithFinalBit                OBEXOpCodeResponseValues = 0xc1
	KOBEXResponseCodeUnsupportedMediaType                    OBEXOpCodeResponseValues = 0x4f
	KOBEXResponseCodeUnsupportedMediaTypeWithFinalBit        OBEXOpCodeResponseValues = 0xcf
	KOBEXResponseCodeUseProxy                                OBEXOpCodeResponseValues = 0x35
	KOBEXResponseCodeUseProxyWithFinalBit                    OBEXOpCodeResponseValues = 0xb5
)

func (e OBEXOpCodeResponseValues) String() string {
	switch e {
	case KOBEXResponseCodeAccepted:
		return "KOBEXResponseCodeAccepted"
	case KOBEXResponseCodeAcceptedWithFinalBit:
		return "KOBEXResponseCodeAcceptedWithFinalBit"
	case KOBEXResponseCodeBadGateway:
		return "KOBEXResponseCodeBadGateway"
	case KOBEXResponseCodeBadGatewayWithFinalBit:
		return "KOBEXResponseCodeBadGatewayWithFinalBit"
	case KOBEXResponseCodeBadRequest:
		return "KOBEXResponseCodeBadRequest"
	case KOBEXResponseCodeBadRequestWithFinalBit:
		return "KOBEXResponseCodeBadRequestWithFinalBit"
	case KOBEXResponseCodeConflict:
		return "KOBEXResponseCodeConflict"
	case KOBEXResponseCodeConflictWithFinalBit:
		return "KOBEXResponseCodeConflictWithFinalBit"
	case KOBEXResponseCodeContinue:
		return "KOBEXResponseCodeContinue"
	case KOBEXResponseCodeContinueWithFinalBit:
		return "KOBEXResponseCodeContinueWithFinalBit"
	case KOBEXResponseCodeCreated:
		return "KOBEXResponseCodeCreated"
	case KOBEXResponseCodeCreatedWithFinalBit:
		return "KOBEXResponseCodeCreatedWithFinalBit"
	case KOBEXResponseCodeDatabaseFull:
		return "KOBEXResponseCodeDatabaseFull"
	case KOBEXResponseCodeDatabaseFullWithFinalBit:
		return "KOBEXResponseCodeDatabaseFullWithFinalBit"
	case KOBEXResponseCodeDatabaseLocked:
		return "KOBEXResponseCodeDatabaseLocked"
	case KOBEXResponseCodeDatabaseLockedWithFinalBit:
		return "KOBEXResponseCodeDatabaseLockedWithFinalBit"
	case KOBEXResponseCodeForbidden:
		return "KOBEXResponseCodeForbidden"
	case KOBEXResponseCodeForbiddenWithFinalBit:
		return "KOBEXResponseCodeForbiddenWithFinalBit"
	case KOBEXResponseCodeGatewayTimeout:
		return "KOBEXResponseCodeGatewayTimeout"
	case KOBEXResponseCodeGatewayTimeoutWithFinalBit:
		return "KOBEXResponseCodeGatewayTimeoutWithFinalBit"
	case KOBEXResponseCodeGone:
		return "KOBEXResponseCodeGone"
	case KOBEXResponseCodeGoneWithFinalBit:
		return "KOBEXResponseCodeGoneWithFinalBit"
	case KOBEXResponseCodeHTTPVersionNotSupported:
		return "KOBEXResponseCodeHTTPVersionNotSupported"
	case KOBEXResponseCodeHTTPVersionNotSupportedWithFinalBit:
		return "KOBEXResponseCodeHTTPVersionNotSupportedWithFinalBit"
	case KOBEXResponseCodeInternalServerError:
		return "KOBEXResponseCodeInternalServerError"
	case KOBEXResponseCodeInternalServerErrorWithFinalBit:
		return "KOBEXResponseCodeInternalServerErrorWithFinalBit"
	case KOBEXResponseCodeLengthRequired:
		return "KOBEXResponseCodeLengthRequired"
	case KOBEXResponseCodeLengthRequiredFinalBit:
		return "KOBEXResponseCodeLengthRequiredFinalBit"
	case KOBEXResponseCodeMethodNotAllowed:
		return "KOBEXResponseCodeMethodNotAllowed"
	case KOBEXResponseCodeMethodNotAllowedWithFinalBit:
		return "KOBEXResponseCodeMethodNotAllowedWithFinalBit"
	case KOBEXResponseCodeMovedPermanently:
		return "KOBEXResponseCodeMovedPermanently"
	case KOBEXResponseCodeMovedPermanentlyWithFinalBit:
		return "KOBEXResponseCodeMovedPermanentlyWithFinalBit"
	case KOBEXResponseCodeMovedTemporarily:
		return "KOBEXResponseCodeMovedTemporarily"
	case KOBEXResponseCodeMovedTemporarilyWithFinalBit:
		return "KOBEXResponseCodeMovedTemporarilyWithFinalBit"
	case KOBEXResponseCodeMultipleChoices:
		return "KOBEXResponseCodeMultipleChoices"
	case KOBEXResponseCodeMultipleChoicesWithFinalBit:
		return "KOBEXResponseCodeMultipleChoicesWithFinalBit"
	case KOBEXResponseCodeNoContent:
		return "KOBEXResponseCodeNoContent"
	case KOBEXResponseCodeNoContentWithFinalBit:
		return "KOBEXResponseCodeNoContentWithFinalBit"
	case KOBEXResponseCodeNonAuthoritativeInfo:
		return "KOBEXResponseCodeNonAuthoritativeInfo"
	case KOBEXResponseCodeNonAuthoritativeInfoWithFinalBit:
		return "KOBEXResponseCodeNonAuthoritativeInfoWithFinalBit"
	case KOBEXResponseCodeNotAcceptable:
		return "KOBEXResponseCodeNotAcceptable"
	case KOBEXResponseCodeNotAcceptableWithFinalBit:
		return "KOBEXResponseCodeNotAcceptableWithFinalBit"
	case KOBEXResponseCodeNotFound:
		return "KOBEXResponseCodeNotFound"
	case KOBEXResponseCodeNotFoundWithFinalBit:
		return "KOBEXResponseCodeNotFoundWithFinalBit"
	case KOBEXResponseCodeNotImplemented:
		return "KOBEXResponseCodeNotImplemented"
	case KOBEXResponseCodeNotImplementedWithFinalBit:
		return "KOBEXResponseCodeNotImplementedWithFinalBit"
	case KOBEXResponseCodeNotModified:
		return "KOBEXResponseCodeNotModified"
	case KOBEXResponseCodeNotModifiedWithFinalBit:
		return "KOBEXResponseCodeNotModifiedWithFinalBit"
	case KOBEXResponseCodePartialContent:
		return "KOBEXResponseCodePartialContent"
	case KOBEXResponseCodePartialContentWithFinalBit:
		return "KOBEXResponseCodePartialContentWithFinalBit"
	case KOBEXResponseCodePaymentRequired:
		return "KOBEXResponseCodePaymentRequired"
	case KOBEXResponseCodePaymentRequiredWithFinalBit:
		return "KOBEXResponseCodePaymentRequiredWithFinalBit"
	case KOBEXResponseCodePreconditionFailed:
		return "KOBEXResponseCodePreconditionFailed"
	case KOBEXResponseCodePreconditionFailedWithFinalBit:
		return "KOBEXResponseCodePreconditionFailedWithFinalBit"
	case KOBEXResponseCodeProxyAuthenticationRequired:
		return "KOBEXResponseCodeProxyAuthenticationRequired"
	case KOBEXResponseCodeProxyAuthenticationRequiredWithFinalBit:
		return "KOBEXResponseCodeProxyAuthenticationRequiredWithFinalBit"
	case KOBEXResponseCodeRequestTimeOut:
		return "KOBEXResponseCodeRequestTimeOut"
	case KOBEXResponseCodeRequestTimeOutWithFinalBit:
		return "KOBEXResponseCodeRequestTimeOutWithFinalBit"
	case KOBEXResponseCodeRequestURLTooLarge:
		return "KOBEXResponseCodeRequestURLTooLarge"
	case KOBEXResponseCodeRequestURLTooLargeWithFinalBit:
		return "KOBEXResponseCodeRequestURLTooLargeWithFinalBit"
	case KOBEXResponseCodeRequestedEntityTooLarge:
		return "KOBEXResponseCodeRequestedEntityTooLarge"
	case KOBEXResponseCodeRequestedEntityTooLargeWithFinalBit:
		return "KOBEXResponseCodeRequestedEntityTooLargeWithFinalBit"
	case KOBEXResponseCodeReservedRangeEnd:
		return "KOBEXResponseCodeReservedRangeEnd"
	case KOBEXResponseCodeReservedRangeStart:
		return "KOBEXResponseCodeReservedRangeStart"
	case KOBEXResponseCodeResetContent:
		return "KOBEXResponseCodeResetContent"
	case KOBEXResponseCodeResetContentWithFinalBit:
		return "KOBEXResponseCodeResetContentWithFinalBit"
	case KOBEXResponseCodeSeeOther:
		return "KOBEXResponseCodeSeeOther"
	case KOBEXResponseCodeSeeOtherWithFinalBit:
		return "KOBEXResponseCodeSeeOtherWithFinalBit"
	case KOBEXResponseCodeServiceUnavailable:
		return "KOBEXResponseCodeServiceUnavailable"
	case KOBEXResponseCodeServiceUnavailableWithFinalBit:
		return "KOBEXResponseCodeServiceUnavailableWithFinalBit"
	case KOBEXResponseCodeSuccess:
		return "KOBEXResponseCodeSuccess"
	case KOBEXResponseCodeSuccessWithFinalBit:
		return "KOBEXResponseCodeSuccessWithFinalBit"
	case KOBEXResponseCodeUnauthorized:
		return "KOBEXResponseCodeUnauthorized"
	case KOBEXResponseCodeUnauthorizedWithFinalBit:
		return "KOBEXResponseCodeUnauthorizedWithFinalBit"
	case KOBEXResponseCodeUnsupportedMediaType:
		return "KOBEXResponseCodeUnsupportedMediaType"
	case KOBEXResponseCodeUnsupportedMediaTypeWithFinalBit:
		return "KOBEXResponseCodeUnsupportedMediaTypeWithFinalBit"
	case KOBEXResponseCodeUseProxy:
		return "KOBEXResponseCodeUseProxy"
	case KOBEXResponseCodeUseProxyWithFinalBit:
		return "KOBEXResponseCodeUseProxyWithFinalBit"
	default:
		return fmt.Sprintf("OBEXOpCodeResponseValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXOpCodeSessionValues
type OBEXOpCodeSessionValues uint32

const (
	KOBEXOpCodeCloseSession   OBEXOpCodeSessionValues = 0x1
	KOBEXOpCodeCreateSession  OBEXOpCodeSessionValues = 0
	KOBEXOpCodeResumeSession  OBEXOpCodeSessionValues = 0x3
	KOBEXOpCodeSetTimeout     OBEXOpCodeSessionValues = 0x4
	KOBEXOpCodeSuspendSession OBEXOpCodeSessionValues = 0x2
)

func (e OBEXOpCodeSessionValues) String() string {
	switch e {
	case KOBEXOpCodeCloseSession:
		return "KOBEXOpCodeCloseSession"
	case KOBEXOpCodeCreateSession:
		return "KOBEXOpCodeCreateSession"
	case KOBEXOpCodeResumeSession:
		return "KOBEXOpCodeResumeSession"
	case KOBEXOpCodeSetTimeout:
		return "KOBEXOpCodeSetTimeout"
	case KOBEXOpCodeSuspendSession:
		return "KOBEXOpCodeSuspendSession"
	default:
		return fmt.Sprintf("OBEXOpCodeSessionValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXPutFlagValues
type OBEXPutFlagValues uint32

const (
	KOBEXPutFlag2Reserved           OBEXPutFlagValues = 4
	KOBEXPutFlag3Reserved           OBEXPutFlagValues = 8
	KOBEXPutFlag4Reserved           OBEXPutFlagValues = 16
	KOBEXPutFlag5Reserved           OBEXPutFlagValues = 32
	KOBEXPutFlag6Reserved           OBEXPutFlagValues = 64
	KOBEXPutFlag7Reserved           OBEXPutFlagValues = 128
	KOBEXPutFlagDontCreateDirectory OBEXPutFlagValues = 2
	KOBEXPutFlagGoToParentDirFirst  OBEXPutFlagValues = 1
	KOBEXPutFlagNone                OBEXPutFlagValues = 0
)

func (e OBEXPutFlagValues) String() string {
	switch e {
	case KOBEXPutFlag2Reserved:
		return "KOBEXPutFlag2Reserved"
	case KOBEXPutFlag3Reserved:
		return "KOBEXPutFlag3Reserved"
	case KOBEXPutFlag4Reserved:
		return "KOBEXPutFlag4Reserved"
	case KOBEXPutFlag5Reserved:
		return "KOBEXPutFlag5Reserved"
	case KOBEXPutFlag6Reserved:
		return "KOBEXPutFlag6Reserved"
	case KOBEXPutFlag7Reserved:
		return "KOBEXPutFlag7Reserved"
	case KOBEXPutFlagDontCreateDirectory:
		return "KOBEXPutFlagDontCreateDirectory"
	case KOBEXPutFlagGoToParentDirFirst:
		return "KOBEXPutFlagGoToParentDirFirst"
	case KOBEXPutFlagNone:
		return "KOBEXPutFlagNone"
	default:
		return fmt.Sprintf("OBEXPutFlagValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXRealmValues
type OBEXRealmValues uint32

const (
	KOBEXRealmASCII    OBEXRealmValues = 0
	KOBEXRealmISO88591 OBEXRealmValues = 0x1
	KOBEXRealmISO88592 OBEXRealmValues = 0x2
	KOBEXRealmISO88593 OBEXRealmValues = 0x3
	KOBEXRealmISO88594 OBEXRealmValues = 0x4
	KOBEXRealmISO88595 OBEXRealmValues = 0x5
	KOBEXRealmISO88596 OBEXRealmValues = 0x6
	KOBEXRealmISO88597 OBEXRealmValues = 0x7
	KOBEXRealmISO88598 OBEXRealmValues = 0x8
	KOBEXRealmISO88599 OBEXRealmValues = 0x9
	KOBEXRealmUNICODE  OBEXRealmValues = 0xff
)

func (e OBEXRealmValues) String() string {
	switch e {
	case KOBEXRealmASCII:
		return "KOBEXRealmASCII"
	case KOBEXRealmISO88591:
		return "KOBEXRealmISO88591"
	case KOBEXRealmISO88592:
		return "KOBEXRealmISO88592"
	case KOBEXRealmISO88593:
		return "KOBEXRealmISO88593"
	case KOBEXRealmISO88594:
		return "KOBEXRealmISO88594"
	case KOBEXRealmISO88595:
		return "KOBEXRealmISO88595"
	case KOBEXRealmISO88596:
		return "KOBEXRealmISO88596"
	case KOBEXRealmISO88597:
		return "KOBEXRealmISO88597"
	case KOBEXRealmISO88598:
		return "KOBEXRealmISO88598"
	case KOBEXRealmISO88599:
		return "KOBEXRealmISO88599"
	case KOBEXRealmUNICODE:
		return "KOBEXRealmUNICODE"
	default:
		return fmt.Sprintf("OBEXRealmValues(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionEventTypes
type OBEXSessionEventTypes uint32

const (
	KOBEXSessionEventTypeAbortCommandReceived              OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'A' // 'OSEA'
	KOBEXSessionEventTypeAbortCommandResponseReceived      OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'A' // 'OCEA'
	KOBEXSessionEventTypeConnectCommandReceived            OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'C' // 'OSEC'
	KOBEXSessionEventTypeConnectCommandResponseReceived    OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'C' // 'OCEC'
	KOBEXSessionEventTypeDisconnectCommandReceived         OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'D' // 'OSED'
	KOBEXSessionEventTypeDisconnectCommandResponseReceived OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'D' // 'OCED'
	KOBEXSessionEventTypeError                             OBEXSessionEventTypes = 'O'<<24 | 'G'<<16 | 'E'<<8 | 'E' // 'OGEE'
	KOBEXSessionEventTypeGetCommandReceived                OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'G' // 'OSEG'
	KOBEXSessionEventTypeGetCommandResponseReceived        OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'G' // 'OCEG'
	KOBEXSessionEventTypePutCommandReceived                OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'P' // 'OSEP'
	KOBEXSessionEventTypePutCommandResponseReceived        OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'P' // 'OCEP'
	KOBEXSessionEventTypeSetPathCommandReceived            OBEXSessionEventTypes = 'O'<<24 | 'S'<<16 | 'E'<<8 | 'S' // 'OSES'
	KOBEXSessionEventTypeSetPathCommandResponseReceived    OBEXSessionEventTypes = 'O'<<24 | 'C'<<16 | 'E'<<8 | 'S' // 'OCES'
)

func (e OBEXSessionEventTypes) String() string {
	switch e {
	case KOBEXSessionEventTypeAbortCommandReceived:
		return "KOBEXSessionEventTypeAbortCommandReceived"
	case KOBEXSessionEventTypeAbortCommandResponseReceived:
		return "KOBEXSessionEventTypeAbortCommandResponseReceived"
	case KOBEXSessionEventTypeConnectCommandReceived:
		return "KOBEXSessionEventTypeConnectCommandReceived"
	case KOBEXSessionEventTypeConnectCommandResponseReceived:
		return "KOBEXSessionEventTypeConnectCommandResponseReceived"
	case KOBEXSessionEventTypeDisconnectCommandReceived:
		return "KOBEXSessionEventTypeDisconnectCommandReceived"
	case KOBEXSessionEventTypeDisconnectCommandResponseReceived:
		return "KOBEXSessionEventTypeDisconnectCommandResponseReceived"
	case KOBEXSessionEventTypeError:
		return "KOBEXSessionEventTypeError"
	case KOBEXSessionEventTypeGetCommandReceived:
		return "KOBEXSessionEventTypeGetCommandReceived"
	case KOBEXSessionEventTypeGetCommandResponseReceived:
		return "KOBEXSessionEventTypeGetCommandResponseReceived"
	case KOBEXSessionEventTypePutCommandReceived:
		return "KOBEXSessionEventTypePutCommandReceived"
	case KOBEXSessionEventTypePutCommandResponseReceived:
		return "KOBEXSessionEventTypePutCommandResponseReceived"
	case KOBEXSessionEventTypeSetPathCommandReceived:
		return "KOBEXSessionEventTypeSetPathCommandReceived"
	case KOBEXSessionEventTypeSetPathCommandResponseReceived:
		return "KOBEXSessionEventTypeSetPathCommandResponseReceived"
	default:
		return fmt.Sprintf("OBEXSessionEventTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionParameterTags
type OBEXSessionParameterTags uint32

const (
	KOBEXSessionParameterTagDeviceAddress      OBEXSessionParameterTags = 0
	KOBEXSessionParameterTagNextSequenceNumber OBEXSessionParameterTags = 0x3
	KOBEXSessionParameterTagNonce              OBEXSessionParameterTags = 0x1
	KOBEXSessionParameterTagSessionID          OBEXSessionParameterTags = 0x2
	KOBEXSessionParameterTagSessionOpcode      OBEXSessionParameterTags = 0x5
	KOBEXSessionParameterTagTimeout            OBEXSessionParameterTags = 0x4
)

func (e OBEXSessionParameterTags) String() string {
	switch e {
	case KOBEXSessionParameterTagDeviceAddress:
		return "KOBEXSessionParameterTagDeviceAddress"
	case KOBEXSessionParameterTagNextSequenceNumber:
		return "KOBEXSessionParameterTagNextSequenceNumber"
	case KOBEXSessionParameterTagNonce:
		return "KOBEXSessionParameterTagNonce"
	case KOBEXSessionParameterTagSessionID:
		return "KOBEXSessionParameterTagSessionID"
	case KOBEXSessionParameterTagSessionOpcode:
		return "KOBEXSessionParameterTagSessionOpcode"
	case KOBEXSessionParameterTagTimeout:
		return "KOBEXSessionParameterTagTimeout"
	default:
		return fmt.Sprintf("OBEXSessionParameterTags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXTransportEventTypes
type OBEXTransportEventTypes uint32

const (
	KOBEXTransportEventTypeDataReceived OBEXTransportEventTypes = 'D'<<24 | 'a'<<16 | 't'<<8 | 'A' // 'DatA'
	KOBEXTransportEventTypeStatus       OBEXTransportEventTypes = 'S'<<24 | 't'<<16 | 'a'<<8 | 'T' // 'StaT'
)

func (e OBEXTransportEventTypes) String() string {
	switch e {
	case KOBEXTransportEventTypeDataReceived:
		return "KOBEXTransportEventTypeDataReceived"
	case KOBEXTransportEventTypeStatus:
		return "KOBEXTransportEventTypeStatus"
	default:
		return fmt.Sprintf("OBEXTransportEventTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXVersions
type OBEXVersions uint32

const (
	KOBEXVersion10 OBEXVersions = 0x10
)

func (e OBEXVersions) String() string {
	switch e {
	case KOBEXVersion10:
		return "KOBEXVersion10"
	default:
		return fmt.Sprintf("OBEXVersions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/ProtocolParameters
type ProtocolParameters uint32

const (
	KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList ProtocolParameters = 2
	KBluetoothSDPProtocolParameterBNEPVersion                        ProtocolParameters = 1
	KBluetoothSDPProtocolParameterL2CAPPSM                           ProtocolParameters = 1
	KBluetoothSDPProtocolParameterRFCOMMChannel                      ProtocolParameters = 1
	KBluetoothSDPProtocolParameterTCPPort                            ProtocolParameters = 1
	KBluetoothSDPProtocolParameterUDPPort                            ProtocolParameters = 1
)

func (e ProtocolParameters) String() string {
	switch e {
	case KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList:
		return "KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList"
	case KBluetoothSDPProtocolParameterBNEPVersion:
		return "KBluetoothSDPProtocolParameterBNEPVersion"
	default:
		return fmt.Sprintf("ProtocolParameters(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/SDPAttributeDeviceIdentificationRecord
type SDPAttributeDeviceIdentificationRecord uint32

const (
	KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL SDPAttributeDeviceIdentificationRecord = 0xb
	KBluetoothSDPAttributeDeviceIdentifierDocumentationURL    SDPAttributeDeviceIdentificationRecord = 0xa
	KBluetoothSDPAttributeDeviceIdentifierPrimaryRecord       SDPAttributeDeviceIdentificationRecord = 0x204
	KBluetoothSDPAttributeDeviceIdentifierProductID           SDPAttributeDeviceIdentificationRecord = 0x202
	KBluetoothSDPAttributeDeviceIdentifierReservedRangeEnd    SDPAttributeDeviceIdentificationRecord = 0x2ff
	KBluetoothSDPAttributeDeviceIdentifierReservedRangeStart  SDPAttributeDeviceIdentificationRecord = 0x206
	KBluetoothSDPAttributeDeviceIdentifierServiceDescription  SDPAttributeDeviceIdentificationRecord = 0x1
	KBluetoothSDPAttributeDeviceIdentifierSpecificationID     SDPAttributeDeviceIdentificationRecord = 0x200
	KBluetoothSDPAttributeDeviceIdentifierVendorID            SDPAttributeDeviceIdentificationRecord = 0x201
	KBluetoothSDPAttributeDeviceIdentifierVendorIDSource      SDPAttributeDeviceIdentificationRecord = 0x205
	KBluetoothSDPAttributeDeviceIdentifierVersion             SDPAttributeDeviceIdentificationRecord = 0x203
)

func (e SDPAttributeDeviceIdentificationRecord) String() string {
	switch e {
	case KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL:
		return "KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL"
	case KBluetoothSDPAttributeDeviceIdentifierDocumentationURL:
		return "KBluetoothSDPAttributeDeviceIdentifierDocumentationURL"
	case KBluetoothSDPAttributeDeviceIdentifierPrimaryRecord:
		return "KBluetoothSDPAttributeDeviceIdentifierPrimaryRecord"
	case KBluetoothSDPAttributeDeviceIdentifierProductID:
		return "KBluetoothSDPAttributeDeviceIdentifierProductID"
	case KBluetoothSDPAttributeDeviceIdentifierReservedRangeEnd:
		return "KBluetoothSDPAttributeDeviceIdentifierReservedRangeEnd"
	case KBluetoothSDPAttributeDeviceIdentifierReservedRangeStart:
		return "KBluetoothSDPAttributeDeviceIdentifierReservedRangeStart"
	case KBluetoothSDPAttributeDeviceIdentifierServiceDescription:
		return "KBluetoothSDPAttributeDeviceIdentifierServiceDescription"
	case KBluetoothSDPAttributeDeviceIdentifierSpecificationID:
		return "KBluetoothSDPAttributeDeviceIdentifierSpecificationID"
	case KBluetoothSDPAttributeDeviceIdentifierVendorID:
		return "KBluetoothSDPAttributeDeviceIdentifierVendorID"
	case KBluetoothSDPAttributeDeviceIdentifierVendorIDSource:
		return "KBluetoothSDPAttributeDeviceIdentifierVendorIDSource"
	case KBluetoothSDPAttributeDeviceIdentifierVersion:
		return "KBluetoothSDPAttributeDeviceIdentifierVersion"
	default:
		return fmt.Sprintf("SDPAttributeDeviceIdentificationRecord(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/SDPAttributeIdentifierCodes
type SDPAttributeIdentifierCodes uint32

const (
	KBluetoothSDPAttributeIdentifierAdditionalProtocolsDescriptorList SDPAttributeIdentifierCodes = 0xd
	KBluetoothSDPAttributeIdentifierAudioFeedbackSupport              SDPAttributeIdentifierCodes = 0x305
	KBluetoothSDPAttributeIdentifierBluetoothProfileDescriptorList    SDPAttributeIdentifierCodes = 0x9
	KBluetoothSDPAttributeIdentifierBrowseGroupList                   SDPAttributeIdentifierCodes = 0x5
	KBluetoothSDPAttributeIdentifierClientExecutableURL               SDPAttributeIdentifierCodes = 0xb
	KBluetoothSDPAttributeIdentifierDocumentationURL                  SDPAttributeIdentifierCodes = 0xa
	KBluetoothSDPAttributeIdentifierExternalNetwork                   SDPAttributeIdentifierCodes = 0x301
	KBluetoothSDPAttributeIdentifierFaxClass1Support                  SDPAttributeIdentifierCodes = 0x302
	KBluetoothSDPAttributeIdentifierFaxClass2Support                  SDPAttributeIdentifierCodes = 0x304
	KBluetoothSDPAttributeIdentifierFaxClass2_0Support                SDPAttributeIdentifierCodes = 0x303
	KBluetoothSDPAttributeIdentifierGroupID                           SDPAttributeIdentifierCodes = 0x200
	KBluetoothSDPAttributeIdentifierHIDBatteryPower                   SDPAttributeIdentifierCodes = 0x209
	KBluetoothSDPAttributeIdentifierHIDBootDevice                     SDPAttributeIdentifierCodes = 0x20e
	KBluetoothSDPAttributeIdentifierHIDCountryCode                    SDPAttributeIdentifierCodes = 0x203
	KBluetoothSDPAttributeIdentifierHIDDescriptorList                 SDPAttributeIdentifierCodes = 0x206
	KBluetoothSDPAttributeIdentifierHIDDeviceSubclass                 SDPAttributeIdentifierCodes = 0x202
	KBluetoothSDPAttributeIdentifierHIDLangIDBaseList                 SDPAttributeIdentifierCodes = 0x207
	KBluetoothSDPAttributeIdentifierHIDNormallyConnectable            SDPAttributeIdentifierCodes = 0x20d
	KBluetoothSDPAttributeIdentifierHIDParserVersion                  SDPAttributeIdentifierCodes = 0x201
	KBluetoothSDPAttributeIdentifierHIDProfileVersion                 SDPAttributeIdentifierCodes = 0x20b
	KBluetoothSDPAttributeIdentifierHIDReconnectInitiate              SDPAttributeIdentifierCodes = 0x205
	KBluetoothSDPAttributeIdentifierHIDReleaseNumber                  SDPAttributeIdentifierCodes = 0x200
	KBluetoothSDPAttributeIdentifierHIDRemoteWake                     SDPAttributeIdentifierCodes = 0x20a
	KBluetoothSDPAttributeIdentifierHIDSDPDisable                     SDPAttributeIdentifierCodes = 0x208
	KBluetoothSDPAttributeIdentifierHIDSSRHostMaxLatency              SDPAttributeIdentifierCodes = 0x20f
	KBluetoothSDPAttributeIdentifierHIDSSRHostMinTimeout              SDPAttributeIdentifierCodes = 0x210
	KBluetoothSDPAttributeIdentifierHIDSupervisionTimeout             SDPAttributeIdentifierCodes = 0x20c
	KBluetoothSDPAttributeIdentifierHIDVirtualCable                   SDPAttributeIdentifierCodes = 0x204
	KBluetoothSDPAttributeIdentifierHomepageURL                       SDPAttributeIdentifierCodes = 0x308
	KBluetoothSDPAttributeIdentifierIPSubnet                          SDPAttributeIdentifierCodes = 0x200
	KBluetoothSDPAttributeIdentifierIconURL                           SDPAttributeIdentifierCodes = 0xc
	KBluetoothSDPAttributeIdentifierLanguageBaseAttributeIDList       SDPAttributeIdentifierCodes = 0x6
	KBluetoothSDPAttributeIdentifierMaxNetAccessRate                  SDPAttributeIdentifierCodes = 0x30c
	KBluetoothSDPAttributeIdentifierNetAccessType                     SDPAttributeIdentifierCodes = 0x30b
	KBluetoothSDPAttributeIdentifierNetwork                           SDPAttributeIdentifierCodes = 0x301
	KBluetoothSDPAttributeIdentifierNetworkAddress                    SDPAttributeIdentifierCodes = 0x306
	KBluetoothSDPAttributeIdentifierProtocolDescriptorList            SDPAttributeIdentifierCodes = 0x4
	KBluetoothSDPAttributeIdentifierProviderName                      SDPAttributeIdentifierCodes = 0x2
	KBluetoothSDPAttributeIdentifierRemoteAudioVolumeControl          SDPAttributeIdentifierCodes = 0x302
	KBluetoothSDPAttributeIdentifierSecurityDescription               SDPAttributeIdentifierCodes = 0x30a
	KBluetoothSDPAttributeIdentifierServiceAvailability               SDPAttributeIdentifierCodes = 0x8
	KBluetoothSDPAttributeIdentifierServiceClassIDList                SDPAttributeIdentifierCodes = 0x1
	KBluetoothSDPAttributeIdentifierServiceDatabaseState              SDPAttributeIdentifierCodes = 0x201
	KBluetoothSDPAttributeIdentifierServiceDescription                SDPAttributeIdentifierCodes = 0x1
	KBluetoothSDPAttributeIdentifierServiceID                         SDPAttributeIdentifierCodes = 0x3
	KBluetoothSDPAttributeIdentifierServiceInfoTimeToLive             SDPAttributeIdentifierCodes = 0x7
	KBluetoothSDPAttributeIdentifierServiceName                       SDPAttributeIdentifierCodes = 0
	KBluetoothSDPAttributeIdentifierServiceRecordHandle               SDPAttributeIdentifierCodes = 0
	KBluetoothSDPAttributeIdentifierServiceRecordState                SDPAttributeIdentifierCodes = 0x2
	KBluetoothSDPAttributeIdentifierServiceVersion                    SDPAttributeIdentifierCodes = 0x300
	KBluetoothSDPAttributeIdentifierSupportedCapabilities             SDPAttributeIdentifierCodes = 0x310
	KBluetoothSDPAttributeIdentifierSupportedDataStoresList           SDPAttributeIdentifierCodes = 0x301
	KBluetoothSDPAttributeIdentifierSupportedFeatures                 SDPAttributeIdentifierCodes = 0x311
	KBluetoothSDPAttributeIdentifierSupportedFunctions                SDPAttributeIdentifierCodes = 0x312
	KBluetoothSDPAttributeIdentifierSupporterFormatsList              SDPAttributeIdentifierCodes = 0x303
	KBluetoothSDPAttributeIdentifierTotalImagingDataCapacity          SDPAttributeIdentifierCodes = 0x313
	KBluetoothSDPAttributeIdentifierVersionNumberList                 SDPAttributeIdentifierCodes = 0x200
	KBluetoothSDPAttributeIdentifierWAPGateway                        SDPAttributeIdentifierCodes = 0x307
	KBluetoothSDPAttributeIdentifierWAPStackType                      SDPAttributeIdentifierCodes = 0x309
)

func (e SDPAttributeIdentifierCodes) String() string {
	switch e {
	case KBluetoothSDPAttributeIdentifierAdditionalProtocolsDescriptorList:
		return "KBluetoothSDPAttributeIdentifierAdditionalProtocolsDescriptorList"
	case KBluetoothSDPAttributeIdentifierAudioFeedbackSupport:
		return "KBluetoothSDPAttributeIdentifierAudioFeedbackSupport"
	case KBluetoothSDPAttributeIdentifierBluetoothProfileDescriptorList:
		return "KBluetoothSDPAttributeIdentifierBluetoothProfileDescriptorList"
	case KBluetoothSDPAttributeIdentifierBrowseGroupList:
		return "KBluetoothSDPAttributeIdentifierBrowseGroupList"
	case KBluetoothSDPAttributeIdentifierClientExecutableURL:
		return "KBluetoothSDPAttributeIdentifierClientExecutableURL"
	case KBluetoothSDPAttributeIdentifierDocumentationURL:
		return "KBluetoothSDPAttributeIdentifierDocumentationURL"
	case KBluetoothSDPAttributeIdentifierExternalNetwork:
		return "KBluetoothSDPAttributeIdentifierExternalNetwork"
	case KBluetoothSDPAttributeIdentifierFaxClass1Support:
		return "KBluetoothSDPAttributeIdentifierFaxClass1Support"
	case KBluetoothSDPAttributeIdentifierFaxClass2Support:
		return "KBluetoothSDPAttributeIdentifierFaxClass2Support"
	case KBluetoothSDPAttributeIdentifierFaxClass2_0Support:
		return "KBluetoothSDPAttributeIdentifierFaxClass2_0Support"
	case KBluetoothSDPAttributeIdentifierGroupID:
		return "KBluetoothSDPAttributeIdentifierGroupID"
	case KBluetoothSDPAttributeIdentifierHIDBatteryPower:
		return "KBluetoothSDPAttributeIdentifierHIDBatteryPower"
	case KBluetoothSDPAttributeIdentifierHIDBootDevice:
		return "KBluetoothSDPAttributeIdentifierHIDBootDevice"
	case KBluetoothSDPAttributeIdentifierHIDCountryCode:
		return "KBluetoothSDPAttributeIdentifierHIDCountryCode"
	case KBluetoothSDPAttributeIdentifierHIDDescriptorList:
		return "KBluetoothSDPAttributeIdentifierHIDDescriptorList"
	case KBluetoothSDPAttributeIdentifierHIDDeviceSubclass:
		return "KBluetoothSDPAttributeIdentifierHIDDeviceSubclass"
	case KBluetoothSDPAttributeIdentifierHIDLangIDBaseList:
		return "KBluetoothSDPAttributeIdentifierHIDLangIDBaseList"
	case KBluetoothSDPAttributeIdentifierHIDNormallyConnectable:
		return "KBluetoothSDPAttributeIdentifierHIDNormallyConnectable"
	case KBluetoothSDPAttributeIdentifierHIDParserVersion:
		return "KBluetoothSDPAttributeIdentifierHIDParserVersion"
	case KBluetoothSDPAttributeIdentifierHIDProfileVersion:
		return "KBluetoothSDPAttributeIdentifierHIDProfileVersion"
	case KBluetoothSDPAttributeIdentifierHIDReconnectInitiate:
		return "KBluetoothSDPAttributeIdentifierHIDReconnectInitiate"
	case KBluetoothSDPAttributeIdentifierHIDRemoteWake:
		return "KBluetoothSDPAttributeIdentifierHIDRemoteWake"
	case KBluetoothSDPAttributeIdentifierHIDSDPDisable:
		return "KBluetoothSDPAttributeIdentifierHIDSDPDisable"
	case KBluetoothSDPAttributeIdentifierHIDSSRHostMaxLatency:
		return "KBluetoothSDPAttributeIdentifierHIDSSRHostMaxLatency"
	case KBluetoothSDPAttributeIdentifierHIDSSRHostMinTimeout:
		return "KBluetoothSDPAttributeIdentifierHIDSSRHostMinTimeout"
	case KBluetoothSDPAttributeIdentifierHIDSupervisionTimeout:
		return "KBluetoothSDPAttributeIdentifierHIDSupervisionTimeout"
	case KBluetoothSDPAttributeIdentifierHIDVirtualCable:
		return "KBluetoothSDPAttributeIdentifierHIDVirtualCable"
	case KBluetoothSDPAttributeIdentifierHomepageURL:
		return "KBluetoothSDPAttributeIdentifierHomepageURL"
	case KBluetoothSDPAttributeIdentifierIconURL:
		return "KBluetoothSDPAttributeIdentifierIconURL"
	case KBluetoothSDPAttributeIdentifierLanguageBaseAttributeIDList:
		return "KBluetoothSDPAttributeIdentifierLanguageBaseAttributeIDList"
	case KBluetoothSDPAttributeIdentifierMaxNetAccessRate:
		return "KBluetoothSDPAttributeIdentifierMaxNetAccessRate"
	case KBluetoothSDPAttributeIdentifierNetAccessType:
		return "KBluetoothSDPAttributeIdentifierNetAccessType"
	case KBluetoothSDPAttributeIdentifierNetworkAddress:
		return "KBluetoothSDPAttributeIdentifierNetworkAddress"
	case KBluetoothSDPAttributeIdentifierProtocolDescriptorList:
		return "KBluetoothSDPAttributeIdentifierProtocolDescriptorList"
	case KBluetoothSDPAttributeIdentifierProviderName:
		return "KBluetoothSDPAttributeIdentifierProviderName"
	case KBluetoothSDPAttributeIdentifierSecurityDescription:
		return "KBluetoothSDPAttributeIdentifierSecurityDescription"
	case KBluetoothSDPAttributeIdentifierServiceAvailability:
		return "KBluetoothSDPAttributeIdentifierServiceAvailability"
	case KBluetoothSDPAttributeIdentifierServiceClassIDList:
		return "KBluetoothSDPAttributeIdentifierServiceClassIDList"
	case KBluetoothSDPAttributeIdentifierServiceID:
		return "KBluetoothSDPAttributeIdentifierServiceID"
	case KBluetoothSDPAttributeIdentifierServiceInfoTimeToLive:
		return "KBluetoothSDPAttributeIdentifierServiceInfoTimeToLive"
	case KBluetoothSDPAttributeIdentifierServiceName:
		return "KBluetoothSDPAttributeIdentifierServiceName"
	case KBluetoothSDPAttributeIdentifierServiceVersion:
		return "KBluetoothSDPAttributeIdentifierServiceVersion"
	case KBluetoothSDPAttributeIdentifierSupportedCapabilities:
		return "KBluetoothSDPAttributeIdentifierSupportedCapabilities"
	case KBluetoothSDPAttributeIdentifierSupportedFeatures:
		return "KBluetoothSDPAttributeIdentifierSupportedFeatures"
	case KBluetoothSDPAttributeIdentifierSupportedFunctions:
		return "KBluetoothSDPAttributeIdentifierSupportedFunctions"
	case KBluetoothSDPAttributeIdentifierTotalImagingDataCapacity:
		return "KBluetoothSDPAttributeIdentifierTotalImagingDataCapacity"
	case KBluetoothSDPAttributeIdentifierWAPGateway:
		return "KBluetoothSDPAttributeIdentifierWAPGateway"
	case KBluetoothSDPAttributeIdentifierWAPStackType:
		return "KBluetoothSDPAttributeIdentifierWAPStackType"
	default:
		return fmt.Sprintf("SDPAttributeIdentifierCodes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/SDPServiceClasses
type SDPServiceClasses uint32

const (
	KBluetoothSDPUUID16ServiceClassAVRemoteControl                       SDPServiceClasses = 0x110e
	KBluetoothSDPUUID16ServiceClassAVRemoteControlController             SDPServiceClasses = 0x110f
	KBluetoothSDPUUID16ServiceClassAVRemoteControlTarget                 SDPServiceClasses = 0x110c
	KBluetoothSDPUUID16ServiceClassAdvancedAudioDistribution             SDPServiceClasses = 0x110d
	KBluetoothSDPUUID16ServiceClassAudioSink                             SDPServiceClasses = 0x110b
	KBluetoothSDPUUID16ServiceClassAudioSource                           SDPServiceClasses = 0x110a
	KBluetoothSDPUUID16ServiceClassAudioVideo                            SDPServiceClasses = 0x112c
	KBluetoothSDPUUID16ServiceClassBasicPrinting                         SDPServiceClasses = 0x1122
	KBluetoothSDPUUID16ServiceClassBrowseGroupDescriptor                 SDPServiceClasses = 0x1001
	KBluetoothSDPUUID16ServiceClassCommonISDNAccess                      SDPServiceClasses = 0x1128
	KBluetoothSDPUUID16ServiceClassCordlessTelephony                     SDPServiceClasses = 0x1109
	KBluetoothSDPUUID16ServiceClassDialupNetworking                      SDPServiceClasses = 0x1103
	KBluetoothSDPUUID16ServiceClassDirectPrinting                        SDPServiceClasses = 0x1118
	KBluetoothSDPUUID16ServiceClassDirectPrintingReferenceObjectsService SDPServiceClasses = 0x1120
	KBluetoothSDPUUID16ServiceClassFax                                   SDPServiceClasses = 0x1111
	KBluetoothSDPUUID16ServiceClassGATT                                  SDPServiceClasses = 0x1801
	KBluetoothSDPUUID16ServiceClassGN                                    SDPServiceClasses = 0x1117
	KBluetoothSDPUUID16ServiceClassGenericAudio                          SDPServiceClasses = 0x1203
	KBluetoothSDPUUID16ServiceClassGenericFileTransfer                   SDPServiceClasses = 0x1202
	KBluetoothSDPUUID16ServiceClassGenericNetworking                     SDPServiceClasses = 0x1201
	KBluetoothSDPUUID16ServiceClassGenericTelephony                      SDPServiceClasses = 0x1204
	KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystem       SDPServiceClasses = 0x1135
	KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystemServer SDPServiceClasses = 0x1136
	KBluetoothSDPUUID16ServiceClassHCR_Print                             SDPServiceClasses = 0x1126
	KBluetoothSDPUUID16ServiceClassHCR_Scan                              SDPServiceClasses = 0x1127
	KBluetoothSDPUUID16ServiceClassHandsFree                             SDPServiceClasses = 0x111e
	KBluetoothSDPUUID16ServiceClassHandsFreeAudioGateway                 SDPServiceClasses = 0x111f
	KBluetoothSDPUUID16ServiceClassHardcopyCableReplacement              SDPServiceClasses = 0x1125
	KBluetoothSDPUUID16ServiceClassHeadset                               SDPServiceClasses = 0x1108
	KBluetoothSDPUUID16ServiceClassHeadsetAudioGateway                   SDPServiceClasses = 0x1112
	KBluetoothSDPUUID16ServiceClassHeadset_HS                            SDPServiceClasses = 0x1131
	KBluetoothSDPUUID16ServiceClassHealthDevice                          SDPServiceClasses = 0x1400
	KBluetoothSDPUUID16ServiceClassHealthDeviceSink                      SDPServiceClasses = 0x1402
	KBluetoothSDPUUID16ServiceClassHealthDeviceSource                    SDPServiceClasses = 0x1401
	KBluetoothSDPUUID16ServiceClassHumanInterfaceDeviceService           SDPServiceClasses = 0x1124
	KBluetoothSDPUUID16ServiceClassImaging                               SDPServiceClasses = 0x111a
	KBluetoothSDPUUID16ServiceClassImagingAutomaticArchive               SDPServiceClasses = 0x111c
	KBluetoothSDPUUID16ServiceClassImagingReferencedObjects              SDPServiceClasses = 0x111d
	KBluetoothSDPUUID16ServiceClassImagingResponder                      SDPServiceClasses = 0x111b
	KBluetoothSDPUUID16ServiceClassIntercom                              SDPServiceClasses = 0x1110
	KBluetoothSDPUUID16ServiceClassIrMCSync                              SDPServiceClasses = 0x1104
	KBluetoothSDPUUID16ServiceClassIrMCSyncCommand                       SDPServiceClasses = 0x1107
	KBluetoothSDPUUID16ServiceClassLANAccessUsingPPP                     SDPServiceClasses = 0x1102
	KBluetoothSDPUUID16ServiceClassMessageAccessProfile                  SDPServiceClasses = 0x1134
	KBluetoothSDPUUID16ServiceClassMessageAccessServer                   SDPServiceClasses = 0x1132
	KBluetoothSDPUUID16ServiceClassMessageNotificationServer             SDPServiceClasses = 0x1133
	KBluetoothSDPUUID16ServiceClassNAP                                   SDPServiceClasses = 0x1116
	KBluetoothSDPUUID16ServiceClassOBEXFileTransfer                      SDPServiceClasses = 0x1106
	KBluetoothSDPUUID16ServiceClassOBEXObjectPush                        SDPServiceClasses = 0x1105
	KBluetoothSDPUUID16ServiceClassPANU                                  SDPServiceClasses = 0x1115
	KBluetoothSDPUUID16ServiceClassPhonebookAccess                       SDPServiceClasses = 0x1130
	KBluetoothSDPUUID16ServiceClassPhonebookAccess_PCE                   SDPServiceClasses = 0x112e
	KBluetoothSDPUUID16ServiceClassPhonebookAccess_PSE                   SDPServiceClasses = 0x112f
	KBluetoothSDPUUID16ServiceClassPnPInformation                        SDPServiceClasses = 0x1200
	KBluetoothSDPUUID16ServiceClassPrintingStatus                        SDPServiceClasses = 0x1123
	KBluetoothSDPUUID16ServiceClassPublicBrowseGroup                     SDPServiceClasses = 0x1002
	KBluetoothSDPUUID16ServiceClassReferencePrinting                     SDPServiceClasses = 0x1119
	KBluetoothSDPUUID16ServiceClassReflectedUI                           SDPServiceClasses = 0x1121
	KBluetoothSDPUUID16ServiceClassSIM_Access                            SDPServiceClasses = 0x112d
	KBluetoothSDPUUID16ServiceClassSerialPort                            SDPServiceClasses = 0x1101
	KBluetoothSDPUUID16ServiceClassServiceDiscoveryServer                SDPServiceClasses = 0x1000
	KBluetoothSDPUUID16ServiceClassUDI_MT                                SDPServiceClasses = 0x112a
	KBluetoothSDPUUID16ServiceClassUDI_TA                                SDPServiceClasses = 0x112b
	KBluetoothSDPUUID16ServiceClassVideoConferencingGW                   SDPServiceClasses = 0x1129
	KBluetoothSDPUUID16ServiceClassVideoDistribution                     SDPServiceClasses = 0x1305
	KBluetoothSDPUUID16ServiceClassVideoSink                             SDPServiceClasses = 0x1304
	KBluetoothSDPUUID16ServiceClassVideoSource                           SDPServiceClasses = 0x1303
	KBluetoothSDPUUID16ServiceClassWAP                                   SDPServiceClasses = 0x1113
	KBluetoothSDPUUID16ServiceClassWAPClient                             SDPServiceClasses = 0x1114
)

func (e SDPServiceClasses) String() string {
	switch e {
	case KBluetoothSDPUUID16ServiceClassAVRemoteControl:
		return "KBluetoothSDPUUID16ServiceClassAVRemoteControl"
	case KBluetoothSDPUUID16ServiceClassAVRemoteControlController:
		return "KBluetoothSDPUUID16ServiceClassAVRemoteControlController"
	case KBluetoothSDPUUID16ServiceClassAVRemoteControlTarget:
		return "KBluetoothSDPUUID16ServiceClassAVRemoteControlTarget"
	case KBluetoothSDPUUID16ServiceClassAdvancedAudioDistribution:
		return "KBluetoothSDPUUID16ServiceClassAdvancedAudioDistribution"
	case KBluetoothSDPUUID16ServiceClassAudioSink:
		return "KBluetoothSDPUUID16ServiceClassAudioSink"
	case KBluetoothSDPUUID16ServiceClassAudioSource:
		return "KBluetoothSDPUUID16ServiceClassAudioSource"
	case KBluetoothSDPUUID16ServiceClassAudioVideo:
		return "KBluetoothSDPUUID16ServiceClassAudioVideo"
	case KBluetoothSDPUUID16ServiceClassBasicPrinting:
		return "KBluetoothSDPUUID16ServiceClassBasicPrinting"
	case KBluetoothSDPUUID16ServiceClassBrowseGroupDescriptor:
		return "KBluetoothSDPUUID16ServiceClassBrowseGroupDescriptor"
	case KBluetoothSDPUUID16ServiceClassCommonISDNAccess:
		return "KBluetoothSDPUUID16ServiceClassCommonISDNAccess"
	case KBluetoothSDPUUID16ServiceClassCordlessTelephony:
		return "KBluetoothSDPUUID16ServiceClassCordlessTelephony"
	case KBluetoothSDPUUID16ServiceClassDialupNetworking:
		return "KBluetoothSDPUUID16ServiceClassDialupNetworking"
	case KBluetoothSDPUUID16ServiceClassDirectPrinting:
		return "KBluetoothSDPUUID16ServiceClassDirectPrinting"
	case KBluetoothSDPUUID16ServiceClassDirectPrintingReferenceObjectsService:
		return "KBluetoothSDPUUID16ServiceClassDirectPrintingReferenceObjectsService"
	case KBluetoothSDPUUID16ServiceClassFax:
		return "KBluetoothSDPUUID16ServiceClassFax"
	case KBluetoothSDPUUID16ServiceClassGATT:
		return "KBluetoothSDPUUID16ServiceClassGATT"
	case KBluetoothSDPUUID16ServiceClassGN:
		return "KBluetoothSDPUUID16ServiceClassGN"
	case KBluetoothSDPUUID16ServiceClassGenericAudio:
		return "KBluetoothSDPUUID16ServiceClassGenericAudio"
	case KBluetoothSDPUUID16ServiceClassGenericFileTransfer:
		return "KBluetoothSDPUUID16ServiceClassGenericFileTransfer"
	case KBluetoothSDPUUID16ServiceClassGenericNetworking:
		return "KBluetoothSDPUUID16ServiceClassGenericNetworking"
	case KBluetoothSDPUUID16ServiceClassGenericTelephony:
		return "KBluetoothSDPUUID16ServiceClassGenericTelephony"
	case KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystem:
		return "KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystem"
	case KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystemServer:
		return "KBluetoothSDPUUID16ServiceClassGlobalNavigationSatelliteSystemServer"
	case KBluetoothSDPUUID16ServiceClassHCR_Print:
		return "KBluetoothSDPUUID16ServiceClassHCR_Print"
	case KBluetoothSDPUUID16ServiceClassHCR_Scan:
		return "KBluetoothSDPUUID16ServiceClassHCR_Scan"
	case KBluetoothSDPUUID16ServiceClassHandsFree:
		return "KBluetoothSDPUUID16ServiceClassHandsFree"
	case KBluetoothSDPUUID16ServiceClassHandsFreeAudioGateway:
		return "KBluetoothSDPUUID16ServiceClassHandsFreeAudioGateway"
	case KBluetoothSDPUUID16ServiceClassHardcopyCableReplacement:
		return "KBluetoothSDPUUID16ServiceClassHardcopyCableReplacement"
	case KBluetoothSDPUUID16ServiceClassHeadset:
		return "KBluetoothSDPUUID16ServiceClassHeadset"
	case KBluetoothSDPUUID16ServiceClassHeadsetAudioGateway:
		return "KBluetoothSDPUUID16ServiceClassHeadsetAudioGateway"
	case KBluetoothSDPUUID16ServiceClassHeadset_HS:
		return "KBluetoothSDPUUID16ServiceClassHeadset_HS"
	case KBluetoothSDPUUID16ServiceClassHealthDevice:
		return "KBluetoothSDPUUID16ServiceClassHealthDevice"
	case KBluetoothSDPUUID16ServiceClassHealthDeviceSink:
		return "KBluetoothSDPUUID16ServiceClassHealthDeviceSink"
	case KBluetoothSDPUUID16ServiceClassHealthDeviceSource:
		return "KBluetoothSDPUUID16ServiceClassHealthDeviceSource"
	case KBluetoothSDPUUID16ServiceClassHumanInterfaceDeviceService:
		return "KBluetoothSDPUUID16ServiceClassHumanInterfaceDeviceService"
	case KBluetoothSDPUUID16ServiceClassImaging:
		return "KBluetoothSDPUUID16ServiceClassImaging"
	case KBluetoothSDPUUID16ServiceClassImagingAutomaticArchive:
		return "KBluetoothSDPUUID16ServiceClassImagingAutomaticArchive"
	case KBluetoothSDPUUID16ServiceClassImagingReferencedObjects:
		return "KBluetoothSDPUUID16ServiceClassImagingReferencedObjects"
	case KBluetoothSDPUUID16ServiceClassImagingResponder:
		return "KBluetoothSDPUUID16ServiceClassImagingResponder"
	case KBluetoothSDPUUID16ServiceClassIntercom:
		return "KBluetoothSDPUUID16ServiceClassIntercom"
	case KBluetoothSDPUUID16ServiceClassIrMCSync:
		return "KBluetoothSDPUUID16ServiceClassIrMCSync"
	case KBluetoothSDPUUID16ServiceClassIrMCSyncCommand:
		return "KBluetoothSDPUUID16ServiceClassIrMCSyncCommand"
	case KBluetoothSDPUUID16ServiceClassLANAccessUsingPPP:
		return "KBluetoothSDPUUID16ServiceClassLANAccessUsingPPP"
	case KBluetoothSDPUUID16ServiceClassMessageAccessProfile:
		return "KBluetoothSDPUUID16ServiceClassMessageAccessProfile"
	case KBluetoothSDPUUID16ServiceClassMessageAccessServer:
		return "KBluetoothSDPUUID16ServiceClassMessageAccessServer"
	case KBluetoothSDPUUID16ServiceClassMessageNotificationServer:
		return "KBluetoothSDPUUID16ServiceClassMessageNotificationServer"
	case KBluetoothSDPUUID16ServiceClassNAP:
		return "KBluetoothSDPUUID16ServiceClassNAP"
	case KBluetoothSDPUUID16ServiceClassOBEXFileTransfer:
		return "KBluetoothSDPUUID16ServiceClassOBEXFileTransfer"
	case KBluetoothSDPUUID16ServiceClassOBEXObjectPush:
		return "KBluetoothSDPUUID16ServiceClassOBEXObjectPush"
	case KBluetoothSDPUUID16ServiceClassPANU:
		return "KBluetoothSDPUUID16ServiceClassPANU"
	case KBluetoothSDPUUID16ServiceClassPhonebookAccess:
		return "KBluetoothSDPUUID16ServiceClassPhonebookAccess"
	case KBluetoothSDPUUID16ServiceClassPhonebookAccess_PCE:
		return "KBluetoothSDPUUID16ServiceClassPhonebookAccess_PCE"
	case KBluetoothSDPUUID16ServiceClassPhonebookAccess_PSE:
		return "KBluetoothSDPUUID16ServiceClassPhonebookAccess_PSE"
	case KBluetoothSDPUUID16ServiceClassPnPInformation:
		return "KBluetoothSDPUUID16ServiceClassPnPInformation"
	case KBluetoothSDPUUID16ServiceClassPrintingStatus:
		return "KBluetoothSDPUUID16ServiceClassPrintingStatus"
	case KBluetoothSDPUUID16ServiceClassPublicBrowseGroup:
		return "KBluetoothSDPUUID16ServiceClassPublicBrowseGroup"
	case KBluetoothSDPUUID16ServiceClassReferencePrinting:
		return "KBluetoothSDPUUID16ServiceClassReferencePrinting"
	case KBluetoothSDPUUID16ServiceClassReflectedUI:
		return "KBluetoothSDPUUID16ServiceClassReflectedUI"
	case KBluetoothSDPUUID16ServiceClassSIM_Access:
		return "KBluetoothSDPUUID16ServiceClassSIM_Access"
	case KBluetoothSDPUUID16ServiceClassSerialPort:
		return "KBluetoothSDPUUID16ServiceClassSerialPort"
	case KBluetoothSDPUUID16ServiceClassServiceDiscoveryServer:
		return "KBluetoothSDPUUID16ServiceClassServiceDiscoveryServer"
	case KBluetoothSDPUUID16ServiceClassUDI_MT:
		return "KBluetoothSDPUUID16ServiceClassUDI_MT"
	case KBluetoothSDPUUID16ServiceClassUDI_TA:
		return "KBluetoothSDPUUID16ServiceClassUDI_TA"
	case KBluetoothSDPUUID16ServiceClassVideoConferencingGW:
		return "KBluetoothSDPUUID16ServiceClassVideoConferencingGW"
	case KBluetoothSDPUUID16ServiceClassVideoDistribution:
		return "KBluetoothSDPUUID16ServiceClassVideoDistribution"
	case KBluetoothSDPUUID16ServiceClassVideoSink:
		return "KBluetoothSDPUUID16ServiceClassVideoSink"
	case KBluetoothSDPUUID16ServiceClassVideoSource:
		return "KBluetoothSDPUUID16ServiceClassVideoSource"
	case KBluetoothSDPUUID16ServiceClassWAP:
		return "KBluetoothSDPUUID16ServiceClassWAP"
	case KBluetoothSDPUUID16ServiceClassWAPClient:
		return "KBluetoothSDPUUID16ServiceClassWAPClient"
	default:
		return fmt.Sprintf("SDPServiceClasses(%d)", e)
	}
}
