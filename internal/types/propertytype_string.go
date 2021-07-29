// Code generated by "stringer -type=PropertyType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AckedTransitions-0]
	_ = x[AckRequired-1]
	_ = x[Action-2]
	_ = x[ActionText-3]
	_ = x[ActiveText-4]
	_ = x[ActiveVtSessions-5]
	_ = x[AlarmValue-6]
	_ = x[AlarmValues-7]
	_ = x[All-8]
	_ = x[AllWritesSuccessful-9]
	_ = x[ApduSegmentTimeout-10]
	_ = x[ApduTimeout-11]
	_ = x[ApplicationSoftwareVersion-12]
	_ = x[Archive-13]
	_ = x[Bias-14]
	_ = x[ChangeOfStateCount-15]
	_ = x[ChangeOfStateTime-16]
	_ = x[NotificationClassProp-17]
	_ = x[ControlledVariableReference-19]
	_ = x[ControlledVariableUnits-20]
	_ = x[ControlledVariableValue-21]
	_ = x[CovIncrement-22]
	_ = x[DateList-23]
	_ = x[DaylightSavingsStatus-24]
	_ = x[Deadband-25]
	_ = x[DerivativeAnt-26]
	_ = x[DerivativeAntUnits-27]
	_ = x[Description-28]
	_ = x[DescriptionOfHalt-29]
	_ = x[DeviceAddressBinding-30]
	_ = x[DeviceType-31]
	_ = x[EffectivePeriod-32]
	_ = x[ElapsedActiveTime-33]
	_ = x[ErrorLimit-34]
	_ = x[EventEnable-35]
	_ = x[EventState-36]
	_ = x[EventType-37]
	_ = x[ExceptionSchedule-38]
	_ = x[FaultValues-39]
	_ = x[FeedbackValue-40]
	_ = x[FileAccessMethod-41]
	_ = x[FileSize-42]
	_ = x[FileType-43]
	_ = x[FirmwareRevision-44]
	_ = x[HighLimit-45]
	_ = x[InactiveText-46]
	_ = x[InProcess-47]
	_ = x[InstanceOf-48]
	_ = x[IntegralAnt-49]
	_ = x[IntegralAntUnits-50]
	_ = x[IssueConfirmedNotifications-51]
	_ = x[LimitEnable-52]
	_ = x[ListOfGroupMembers-53]
	_ = x[ListOfObjectPropertyReferences-54]
	_ = x[ListOfSessionKeys-55]
	_ = x[LocalDate-56]
	_ = x[LocalTime-57]
	_ = x[Location-58]
	_ = x[LowLimit-59]
	_ = x[ManipulatedVariableReference-60]
	_ = x[MaximumOutput-61]
	_ = x[MaxApduLengthAccepted-62]
	_ = x[MaxInfoFrames-63]
	_ = x[MaxMaster-64]
	_ = x[MaxPresValue-65]
	_ = x[MinimumOffTime-66]
	_ = x[MinimumOnTime-67]
	_ = x[MinimumOutput-68]
	_ = x[MinPresValue-69]
	_ = x[ModelName-70]
	_ = x[ModificationDate-71]
	_ = x[NotifyType-72]
	_ = x[NumberOfApduRetries-73]
	_ = x[NumberOfStates-74]
	_ = x[ObjectIdentifier-75]
	_ = x[ObjectList-76]
	_ = x[ObjectName-77]
	_ = x[ObjectPropertyReference-78]
	_ = x[ObjectTypeProp-79]
	_ = x[Optional-80]
	_ = x[OutOfService-81]
	_ = x[OutputUnits-82]
	_ = x[EventParameters-83]
	_ = x[Polarity-84]
	_ = x[PresentValue-85]
	_ = x[Priority-86]
	_ = x[PriorityArray-87]
	_ = x[PriorityForWriting-88]
	_ = x[ProcessIdentifier-89]
	_ = x[ProgramChange-90]
	_ = x[ProgramLocation-91]
	_ = x[ProgramState-92]
	_ = x[ProportionalAnt-93]
	_ = x[ProportionalAntUnits-94]
	_ = x[ProtocolConformanceClass-95]
	_ = x[ProtocolObjectTypesSupported-96]
	_ = x[ProtocolServicesSupported-97]
	_ = x[ProtocolVersion-98]
	_ = x[ReadOnly-99]
	_ = x[ReasonForHalt-100]
	_ = x[Recipient-101]
	_ = x[RecipientList-102]
	_ = x[Reliability-103]
	_ = x[RelinquishDefault-104]
	_ = x[Required-105]
	_ = x[Resolution-106]
	_ = x[SegmentationSupported-107]
	_ = x[Setpoint-108]
	_ = x[SetpointReference-109]
	_ = x[StateText-110]
	_ = x[StatusFlags-111]
	_ = x[SystemStatus-112]
	_ = x[TimeDelay-113]
	_ = x[TimeOfActiveTimeReset-114]
	_ = x[TimeOfStateCountReset-115]
	_ = x[TimeSynchronizationRecipients-116]
	_ = x[Units-117]
	_ = x[UpdateInterval-118]
	_ = x[UtcOffset-119]
	_ = x[VendorIdentifier-120]
	_ = x[VendorName-121]
	_ = x[VtClassesSupported-122]
	_ = x[WeeklySchedule-123]
	_ = x[AttemptedSamples-124]
	_ = x[AverageValue-125]
	_ = x[BufferSize-126]
	_ = x[ClientCovIncrement-127]
	_ = x[CovResubscriptionInterval-128]
	_ = x[CurrentNotifyTime-129]
	_ = x[EventTimeStamps-130]
	_ = x[LogBuffer-131]
	_ = x[LogDeviceObjectProperty-132]
	_ = x[Enable-133]
	_ = x[LogInterval-134]
	_ = x[MaximumValue-135]
	_ = x[MinimumValue-136]
	_ = x[NotificationThreshold-137]
	_ = x[PreviousNotifyTime-138]
	_ = x[ProtocolRevision-139]
	_ = x[RecordsSinceNotification-140]
	_ = x[RecordCount-141]
	_ = x[StartTime-142]
	_ = x[StopTime-143]
	_ = x[StopWhenFull-144]
	_ = x[TotalRecordCount-145]
	_ = x[ValidSamples-146]
	_ = x[WindowInterval-147]
	_ = x[WindowSamples-148]
	_ = x[MaximumValueTimestamp-149]
	_ = x[MinimumValueTimestamp-150]
	_ = x[VarianceValue-151]
	_ = x[ActiveCovSubscriptions-152]
	_ = x[BackupFailureTimeout-153]
	_ = x[ConfigurationFiles-154]
	_ = x[DatabaseRevision-155]
	_ = x[DirectReading-156]
	_ = x[LastRestoreTime-157]
	_ = x[MaintenanceRequired-158]
	_ = x[MemberOf-159]
	_ = x[Mode-160]
	_ = x[OperationExpected-161]
	_ = x[Setting-162]
	_ = x[Silenced-163]
	_ = x[TrackingValue-164]
	_ = x[ZoneMembers-165]
	_ = x[LifeSafetyAlarmValues-166]
	_ = x[MaxSegmentsAccepted-167]
	_ = x[ProfileName-168]
	_ = x[AutoSlaveDiscovery-169]
	_ = x[ManualSlaveAddressBinding-170]
	_ = x[SlaveAddressBinding-171]
	_ = x[SlaveProxyEnable-172]
	_ = x[LastNotifyRecord-173]
	_ = x[ScheduleDefault-174]
	_ = x[AcceptedModes-175]
	_ = x[AdjustValue-176]
	_ = x[Count-177]
	_ = x[CountBeforeChange-178]
	_ = x[CountChangeTime-179]
	_ = x[CovPeriod-180]
	_ = x[InputReference-181]
	_ = x[LimitMonitoringInterval-182]
	_ = x[LoggingObject-183]
	_ = x[LoggingRecord-184]
	_ = x[Prescale-185]
	_ = x[PulseRate-186]
	_ = x[Scale-187]
	_ = x[ScaleFactor-188]
	_ = x[UpdateTime-189]
	_ = x[ValueBeforeChange-190]
	_ = x[ValueSet-191]
	_ = x[ValueChangeTime-192]
	_ = x[AlignIntervals-193]
	_ = x[IntervalOffset-195]
	_ = x[LastRestartReason-196]
	_ = x[LoggingType-197]
	_ = x[RestartNotificationRecipients-202]
	_ = x[TimeOfDeviceRestart-203]
	_ = x[TimeSynchronizationInterval-204]
	_ = x[Trigger-205]
	_ = x[UTCTimeSynchronizationRecipients-206]
	_ = x[NodeSubtype-207]
	_ = x[NodeType-208]
	_ = x[StructuredObjectList-209]
	_ = x[SubordinateAnnotations-210]
	_ = x[SubordinateList-211]
	_ = x[ActualShedLevel-212]
	_ = x[DutyWindow-213]
	_ = x[ExpectedShedLevel-214]
	_ = x[FullDutyBaseline-215]
	_ = x[RequestedShedLevel-218]
	_ = x[ShedDuration-219]
	_ = x[ShedLevelDescriptions-220]
	_ = x[ShedLevels-221]
	_ = x[StateDescription-222]
	_ = x[DoorAlarmState-226]
	_ = x[DoorExtendedPulseTime-227]
	_ = x[DoorMembers-228]
	_ = x[DoorOpenTooLongTime-229]
	_ = x[DoorPulseTime-230]
	_ = x[DoorStatus-231]
	_ = x[DoorUnlockDelayTime-232]
	_ = x[LockStatus-233]
	_ = x[MaskedAlarmValues-234]
	_ = x[SecuredStatus-235]
	_ = x[AbsenteeLimit-244]
	_ = x[AccessAlarmEvents-245]
	_ = x[AccessDoors-246]
	_ = x[AccessEvent-247]
	_ = x[AccessEventAuthenticationFactor-248]
	_ = x[AccessEventCredential-249]
	_ = x[AccessEventTime-250]
	_ = x[AccessTransactionEvents-251]
	_ = x[Accompaniment-252]
	_ = x[AccompanimentTime-253]
	_ = x[ActivationTime-254]
	_ = x[ActiveAuthenticationPolicy-255]
	_ = x[AssignedAccessRights-256]
	_ = x[AuthenticationFactors-257]
	_ = x[AuthenticationPolicyList-258]
	_ = x[AuthenticationPolicyNames-259]
	_ = x[AuthenticationStatus-260]
	_ = x[AuthorizationMode-261]
	_ = x[BelongsTo-262]
	_ = x[CredentialDisable-263]
	_ = x[CredentialStatus-264]
	_ = x[Credentials-265]
	_ = x[CredentialsInZone-266]
	_ = x[DaysRemaining-267]
	_ = x[EntryPoints-268]
	_ = x[ExitPoints-269]
	_ = x[ExpiryTime-270]
	_ = x[ExtendedTimeEnable-271]
	_ = x[FailedAttemptEvents-272]
	_ = x[FailedAttempts-273]
	_ = x[FailedAttemptsTime-274]
	_ = x[LastAccessEvent-275]
	_ = x[LastAccessPoint-276]
	_ = x[LastCredentialAdded-277]
	_ = x[LastCredentialAddedTime-278]
	_ = x[LastCredentialRemoved-279]
	_ = x[LastCredentialRemovedTime-280]
	_ = x[LastUseTime-281]
	_ = x[Lockout-282]
	_ = x[LockoutRelinquishTime-283]
	_ = x[MasterExemption-284]
	_ = x[MaxFailedAttempts-285]
	_ = x[Members-286]
	_ = x[MusterPoint-287]
	_ = x[NegativeAccessRules-288]
	_ = x[NumberOfAuthenticationPolicies-289]
	_ = x[OccupancyCount-290]
	_ = x[OccupancyCountAdjust-291]
	_ = x[OccupancyCountEnable-292]
	_ = x[OccupancyExemption-293]
	_ = x[OccupancyLowerLimit-294]
	_ = x[OccupancyLowerLimitEnforced-295]
	_ = x[OccupancyState-296]
	_ = x[OccupancyUpperLimit-297]
	_ = x[OccupancyUpperLimitEnforced-298]
	_ = x[PassbackExemption-299]
	_ = x[PassbackMode-300]
	_ = x[PassbackTimeout-301]
	_ = x[PositiveAccessRules-302]
	_ = x[ReasonForDisable-303]
	_ = x[SupportedFormats-304]
	_ = x[SupportedFormatClasses-305]
	_ = x[ThreatAuthority-306]
	_ = x[ThreatLevel-307]
	_ = x[TraceFlag-308]
	_ = x[TransactionNotificationClass-309]
	_ = x[UserExternalIdentifier-310]
	_ = x[UserInformationReference-311]
	_ = x[UserName-317]
	_ = x[UserType-318]
	_ = x[UsesRemaining-319]
	_ = x[ZoneFrom-320]
	_ = x[ZoneTo-321]
	_ = x[AccessEventTag-322]
	_ = x[GlobalIdentifier-323]
	_ = x[VerificationTime-326]
	_ = x[BaseDeviceSecurityPolicy-327]
	_ = x[DistributionKeyRevision-328]
	_ = x[DoNotHide-329]
	_ = x[KeySets-330]
	_ = x[LastKeyServer-331]
	_ = x[NetworkAccessSecurityPolicies-332]
	_ = x[PacketReorderTime-333]
	_ = x[SecurityPduTimeout-334]
	_ = x[SecurityTimeWindow-335]
	_ = x[SupportedSecurityAlgorithm-336]
	_ = x[UpdateKeySetTimeout-337]
	_ = x[BackupAndRestoreState-338]
	_ = x[BackupPreparationTime-339]
	_ = x[RestoreCompletionTime-340]
	_ = x[RestorePreparationTime-341]
	_ = x[BitMask-342]
	_ = x[BitText-343]
	_ = x[IsUTC-344]
	_ = x[GroupMembers-345]
	_ = x[GroupMemberNames-346]
	_ = x[MemberStatusFlags-347]
	_ = x[RequestedUpdateInterval-348]
	_ = x[CovuPeriod-349]
	_ = x[CovuRecipients-350]
	_ = x[EventMessageTexts-351]
	_ = x[EventMessageTextsConfig-352]
	_ = x[EventDetectionEnable-353]
	_ = x[EventAlgorithmInhibit-354]
	_ = x[EventAlgorithmInhibitRef-355]
	_ = x[TimeDelayNormal-356]
	_ = x[ReliabilityEvaluationInhibit-357]
	_ = x[FaultParameters-358]
	_ = x[FaultType-359]
	_ = x[LocalForwardingOnly-360]
	_ = x[ProcessIdentifierFilter-361]
	_ = x[SubscribedRecipients-362]
	_ = x[PortFilter-363]
	_ = x[AuthorizationExemptions-364]
	_ = x[AllowGroupDelayInhibit-365]
	_ = x[ChannelNumber-366]
	_ = x[ControlGroups-367]
	_ = x[ExecutionDelay-368]
	_ = x[LastPriority-369]
	_ = x[WriteStatus-370]
	_ = x[PropertyList-371]
	_ = x[SerialNumber-372]
	_ = x[BlinkWarnEnable-373]
	_ = x[DefaultFadeTime-374]
	_ = x[DefaultRampRate-375]
	_ = x[DefaultStepIncrement-376]
	_ = x[EgressTime-377]
	_ = x[InProgress-378]
	_ = x[InstantaneousPower-379]
	_ = x[LightingCommand-380]
	_ = x[LightingCommandDefaultPriority-381]
	_ = x[MaxActualValue-382]
	_ = x[MinActualValue-383]
	_ = x[Power-384]
	_ = x[Transition-385]
	_ = x[EgressActive-386]
}

const (
	_PropertyType_name_0 = "AckedTransitionsAckRequiredActionActionTextActiveTextActiveVtSessionsAlarmValueAlarmValuesAllAllWritesSuccessfulApduSegmentTimeoutApduTimeoutApplicationSoftwareVersionArchiveBiasChangeOfStateCountChangeOfStateTimeNotificationClassProp"
	_PropertyType_name_1 = "ControlledVariableReferenceControlledVariableUnitsControlledVariableValueCovIncrementDateListDaylightSavingsStatusDeadbandDerivativeAntDerivativeAntUnitsDescriptionDescriptionOfHaltDeviceAddressBindingDeviceTypeEffectivePeriodElapsedActiveTimeErrorLimitEventEnableEventStateEventTypeExceptionScheduleFaultValuesFeedbackValueFileAccessMethodFileSizeFileTypeFirmwareRevisionHighLimitInactiveTextInProcessInstanceOfIntegralAntIntegralAntUnitsIssueConfirmedNotificationsLimitEnableListOfGroupMembersListOfObjectPropertyReferencesListOfSessionKeysLocalDateLocalTimeLocationLowLimitManipulatedVariableReferenceMaximumOutputMaxApduLengthAcceptedMaxInfoFramesMaxMasterMaxPresValueMinimumOffTimeMinimumOnTimeMinimumOutputMinPresValueModelNameModificationDateNotifyTypeNumberOfApduRetriesNumberOfStatesObjectIdentifierObjectListObjectNameObjectPropertyReferenceObjectTypePropOptionalOutOfServiceOutputUnitsEventParametersPolarityPresentValuePriorityPriorityArrayPriorityForWritingProcessIdentifierProgramChangeProgramLocationProgramStateProportionalAntProportionalAntUnitsProtocolConformanceClassProtocolObjectTypesSupportedProtocolServicesSupportedProtocolVersionReadOnlyReasonForHaltRecipientRecipientListReliabilityRelinquishDefaultRequiredResolutionSegmentationSupportedSetpointSetpointReferenceStateTextStatusFlagsSystemStatusTimeDelayTimeOfActiveTimeResetTimeOfStateCountResetTimeSynchronizationRecipientsUnitsUpdateIntervalUtcOffsetVendorIdentifierVendorNameVtClassesSupportedWeeklyScheduleAttemptedSamplesAverageValueBufferSizeClientCovIncrementCovResubscriptionIntervalCurrentNotifyTimeEventTimeStampsLogBufferLogDeviceObjectPropertyEnableLogIntervalMaximumValueMinimumValueNotificationThresholdPreviousNotifyTimeProtocolRevisionRecordsSinceNotificationRecordCountStartTimeStopTimeStopWhenFullTotalRecordCountValidSamplesWindowIntervalWindowSamplesMaximumValueTimestampMinimumValueTimestampVarianceValueActiveCovSubscriptionsBackupFailureTimeoutConfigurationFilesDatabaseRevisionDirectReadingLastRestoreTimeMaintenanceRequiredMemberOfModeOperationExpectedSettingSilencedTrackingValueZoneMembersLifeSafetyAlarmValuesMaxSegmentsAcceptedProfileNameAutoSlaveDiscoveryManualSlaveAddressBindingSlaveAddressBindingSlaveProxyEnableLastNotifyRecordScheduleDefaultAcceptedModesAdjustValueCountCountBeforeChangeCountChangeTimeCovPeriodInputReferenceLimitMonitoringIntervalLoggingObjectLoggingRecordPrescalePulseRateScaleScaleFactorUpdateTimeValueBeforeChangeValueSetValueChangeTimeAlignIntervals"
	_PropertyType_name_2 = "IntervalOffsetLastRestartReasonLoggingType"
	_PropertyType_name_3 = "RestartNotificationRecipientsTimeOfDeviceRestartTimeSynchronizationIntervalTriggerUTCTimeSynchronizationRecipientsNodeSubtypeNodeTypeStructuredObjectListSubordinateAnnotationsSubordinateListActualShedLevelDutyWindowExpectedShedLevelFullDutyBaseline"
	_PropertyType_name_4 = "RequestedShedLevelShedDurationShedLevelDescriptionsShedLevelsStateDescription"
	_PropertyType_name_5 = "DoorAlarmStateDoorExtendedPulseTimeDoorMembersDoorOpenTooLongTimeDoorPulseTimeDoorStatusDoorUnlockDelayTimeLockStatusMaskedAlarmValuesSecuredStatus"
	_PropertyType_name_6 = "AbsenteeLimitAccessAlarmEventsAccessDoorsAccessEventAccessEventAuthenticationFactorAccessEventCredentialAccessEventTimeAccessTransactionEventsAccompanimentAccompanimentTimeActivationTimeActiveAuthenticationPolicyAssignedAccessRightsAuthenticationFactorsAuthenticationPolicyListAuthenticationPolicyNamesAuthenticationStatusAuthorizationModeBelongsToCredentialDisableCredentialStatusCredentialsCredentialsInZoneDaysRemainingEntryPointsExitPointsExpiryTimeExtendedTimeEnableFailedAttemptEventsFailedAttemptsFailedAttemptsTimeLastAccessEventLastAccessPointLastCredentialAddedLastCredentialAddedTimeLastCredentialRemovedLastCredentialRemovedTimeLastUseTimeLockoutLockoutRelinquishTimeMasterExemptionMaxFailedAttemptsMembersMusterPointNegativeAccessRulesNumberOfAuthenticationPoliciesOccupancyCountOccupancyCountAdjustOccupancyCountEnableOccupancyExemptionOccupancyLowerLimitOccupancyLowerLimitEnforcedOccupancyStateOccupancyUpperLimitOccupancyUpperLimitEnforcedPassbackExemptionPassbackModePassbackTimeoutPositiveAccessRulesReasonForDisableSupportedFormatsSupportedFormatClassesThreatAuthorityThreatLevelTraceFlagTransactionNotificationClassUserExternalIdentifierUserInformationReference"
	_PropertyType_name_7 = "UserNameUserTypeUsesRemainingZoneFromZoneToAccessEventTagGlobalIdentifier"
	_PropertyType_name_8 = "VerificationTimeBaseDeviceSecurityPolicyDistributionKeyRevisionDoNotHideKeySetsLastKeyServerNetworkAccessSecurityPoliciesPacketReorderTimeSecurityPduTimeoutSecurityTimeWindowSupportedSecurityAlgorithmUpdateKeySetTimeoutBackupAndRestoreStateBackupPreparationTimeRestoreCompletionTimeRestorePreparationTimeBitMaskBitTextIsUTCGroupMembersGroupMemberNamesMemberStatusFlagsRequestedUpdateIntervalCovuPeriodCovuRecipientsEventMessageTextsEventMessageTextsConfigEventDetectionEnableEventAlgorithmInhibitEventAlgorithmInhibitRefTimeDelayNormalReliabilityEvaluationInhibitFaultParametersFaultTypeLocalForwardingOnlyProcessIdentifierFilterSubscribedRecipientsPortFilterAuthorizationExemptionsAllowGroupDelayInhibitChannelNumberControlGroupsExecutionDelayLastPriorityWriteStatusPropertyListSerialNumberBlinkWarnEnableDefaultFadeTimeDefaultRampRateDefaultStepIncrementEgressTimeInProgressInstantaneousPowerLightingCommandLightingCommandDefaultPriorityMaxActualValueMinActualValuePowerTransitionEgressActive"
)

var (
	_PropertyType_index_0 = [...]uint8{0, 16, 27, 33, 43, 53, 69, 79, 90, 93, 112, 130, 141, 167, 174, 178, 196, 213, 234}
	_PropertyType_index_1 = [...]uint16{0, 27, 50, 73, 85, 93, 114, 122, 135, 153, 164, 181, 201, 211, 226, 243, 253, 264, 274, 283, 300, 311, 324, 340, 348, 356, 372, 381, 393, 402, 412, 423, 439, 466, 477, 495, 525, 542, 551, 560, 568, 576, 604, 617, 638, 651, 660, 672, 686, 699, 712, 724, 733, 749, 759, 778, 792, 808, 818, 828, 851, 865, 873, 885, 896, 911, 919, 931, 939, 952, 970, 987, 1000, 1015, 1027, 1042, 1062, 1086, 1114, 1139, 1154, 1162, 1175, 1184, 1197, 1208, 1225, 1233, 1243, 1264, 1272, 1289, 1298, 1309, 1321, 1330, 1351, 1372, 1401, 1406, 1420, 1429, 1445, 1455, 1473, 1487, 1503, 1515, 1525, 1543, 1568, 1585, 1600, 1609, 1632, 1638, 1649, 1661, 1673, 1694, 1712, 1728, 1752, 1763, 1772, 1780, 1792, 1808, 1820, 1834, 1847, 1868, 1889, 1902, 1924, 1944, 1962, 1978, 1991, 2006, 2025, 2033, 2037, 2054, 2061, 2069, 2082, 2093, 2114, 2133, 2144, 2162, 2187, 2206, 2222, 2238, 2253, 2266, 2277, 2282, 2299, 2314, 2323, 2337, 2360, 2373, 2386, 2394, 2403, 2408, 2419, 2429, 2446, 2454, 2469, 2483}
	_PropertyType_index_2 = [...]uint8{0, 14, 31, 42}
	_PropertyType_index_3 = [...]uint8{0, 29, 48, 75, 82, 114, 125, 133, 153, 175, 190, 205, 215, 232, 248}
	_PropertyType_index_4 = [...]uint8{0, 18, 30, 51, 61, 77}
	_PropertyType_index_5 = [...]uint8{0, 14, 35, 46, 65, 78, 88, 107, 117, 134, 147}
	_PropertyType_index_6 = [...]uint16{0, 13, 30, 41, 52, 83, 104, 119, 142, 155, 172, 186, 212, 232, 253, 277, 302, 322, 339, 348, 365, 381, 392, 409, 422, 433, 443, 453, 471, 490, 504, 522, 537, 552, 571, 594, 615, 640, 651, 658, 679, 694, 711, 718, 729, 748, 778, 792, 812, 832, 850, 869, 896, 910, 929, 956, 973, 985, 1000, 1019, 1035, 1051, 1073, 1088, 1099, 1108, 1136, 1158, 1182}
	_PropertyType_index_7 = [...]uint8{0, 8, 16, 29, 37, 43, 57, 73}
	_PropertyType_index_8 = [...]uint16{0, 16, 40, 63, 72, 79, 92, 121, 138, 156, 174, 200, 219, 240, 261, 282, 304, 311, 318, 323, 335, 351, 368, 391, 401, 415, 432, 455, 475, 496, 520, 535, 563, 578, 587, 606, 629, 649, 659, 682, 704, 717, 730, 744, 756, 767, 779, 791, 806, 821, 836, 856, 866, 876, 894, 909, 939, 953, 967, 972, 982, 994}
)

func (i PropertyType) String() string {
	switch {
	case i <= 17:
		return _PropertyType_name_0[_PropertyType_index_0[i]:_PropertyType_index_0[i+1]]
	case 19 <= i && i <= 193:
		i -= 19
		return _PropertyType_name_1[_PropertyType_index_1[i]:_PropertyType_index_1[i+1]]
	case 195 <= i && i <= 197:
		i -= 195
		return _PropertyType_name_2[_PropertyType_index_2[i]:_PropertyType_index_2[i+1]]
	case 202 <= i && i <= 215:
		i -= 202
		return _PropertyType_name_3[_PropertyType_index_3[i]:_PropertyType_index_3[i+1]]
	case 218 <= i && i <= 222:
		i -= 218
		return _PropertyType_name_4[_PropertyType_index_4[i]:_PropertyType_index_4[i+1]]
	case 226 <= i && i <= 235:
		i -= 226
		return _PropertyType_name_5[_PropertyType_index_5[i]:_PropertyType_index_5[i+1]]
	case 244 <= i && i <= 311:
		i -= 244
		return _PropertyType_name_6[_PropertyType_index_6[i]:_PropertyType_index_6[i+1]]
	case 317 <= i && i <= 323:
		i -= 317
		return _PropertyType_name_7[_PropertyType_index_7[i]:_PropertyType_index_7[i+1]]
	case 326 <= i && i <= 386:
		i -= 326
		return _PropertyType_name_8[_PropertyType_index_8[i]:_PropertyType_index_8[i+1]]
	default:
		return "PropertyType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
