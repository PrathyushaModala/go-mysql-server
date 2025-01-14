// Code generated by "stringer -type=RuleId -linecomment"; DO NOT EDIT.

package analyzer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[applyDefaultSelectLimitId-0]
	_ = x[validateOffsetAndLimitId-1]
	_ = x[validateStarExpressionsId-2]
	_ = x[validateCreateTableId-3]
	_ = x[validateExprSemId-4]
	_ = x[resolveVariablesId-5]
	_ = x[resolveNamedWindowsId-6]
	_ = x[resolveSetVariablesId-7]
	_ = x[resolveViewsId-8]
	_ = x[liftCtesId-9]
	_ = x[resolveCtesId-10]
	_ = x[liftRecursiveCtesId-11]
	_ = x[resolveDatabasesId-12]
	_ = x[resolveTablesId-13]
	_ = x[loadStoredProceduresId-14]
	_ = x[validateDropTablesId-15]
	_ = x[pruneDropTablesId-16]
	_ = x[setTargetSchemasId-17]
	_ = x[resolveCreateLikeId-18]
	_ = x[parseColumnDefaultsId-19]
	_ = x[resolveDropConstraintId-20]
	_ = x[validateDropConstraintId-21]
	_ = x[loadCheckConstraintsId-22]
	_ = x[assignCatalogId-23]
	_ = x[resolveAnalyzeTablesId-24]
	_ = x[resolveCreateSelectId-25]
	_ = x[resolveSubqueriesId-26]
	_ = x[setViewTargetSchemaId-27]
	_ = x[resolveUnionsId-28]
	_ = x[resolveDescribeQueryId-29]
	_ = x[checkUniqueTableNamesId-30]
	_ = x[disambiguateTableFunctionsId-31]
	_ = x[resolveTableFunctionsId-32]
	_ = x[resolveDeclarationsId-33]
	_ = x[resolveColumnDefaultsId-34]
	_ = x[validateColumnDefaultsId-35]
	_ = x[validateCreateTriggerId-36]
	_ = x[validateCreateProcedureId-37]
	_ = x[resolveCreateProcedureId-38]
	_ = x[loadInfoSchemaId-39]
	_ = x[validateReadOnlyDatabaseId-40]
	_ = x[validateReadOnlyTransactionId-41]
	_ = x[validateDatabaseSetId-42]
	_ = x[validatePrivilegesId-43]
	_ = x[reresolveTablesId-44]
	_ = x[setInsertColumnsId-45]
	_ = x[validateJoinComplexityId-46]
	_ = x[applyBinlogReplicaControllerId-47]
	_ = x[resolveNaturalJoinsId-48]
	_ = x[resolveOrderbyLiteralsId-49]
	_ = x[resolveFunctionsId-50]
	_ = x[flattenTableAliasesId-51]
	_ = x[pushdownSortId-52]
	_ = x[pushdownGroupbyAliasesId-53]
	_ = x[pushdownSubqueryAliasFiltersId-54]
	_ = x[qualifyColumnsId-55]
	_ = x[resolveColumnsId-56]
	_ = x[validateCheckConstraintId-57]
	_ = x[resolveBarewordSetVariablesId-58]
	_ = x[replaceCountStarId-59]
	_ = x[expandStarsId-60]
	_ = x[transposeRightJoinsId-61]
	_ = x[resolveHavingId-62]
	_ = x[mergeUnionSchemasId-63]
	_ = x[flattenAggregationExprsId-64]
	_ = x[reorderProjectionId-65]
	_ = x[resolveSubqueryExprsId-66]
	_ = x[replaceCrossJoinsId-67]
	_ = x[moveJoinCondsToFilterId-68]
	_ = x[evalFilterId-69]
	_ = x[optimizeDistinctId-70]
	_ = x[hoistOutOfScopeFiltersId-71]
	_ = x[transformJoinApplyId-72]
	_ = x[hoistSelectExistsId-73]
	_ = x[finalizeSubqueriesId-74]
	_ = x[finalizeUnionsId-75]
	_ = x[loadTriggersId-76]
	_ = x[loadEventsId-77]
	_ = x[processTruncateId-78]
	_ = x[resolveAlterColumnId-79]
	_ = x[resolveGeneratorsId-80]
	_ = x[removeUnnecessaryConvertsId-81]
	_ = x[pruneColumnsId-82]
	_ = x[stripTableNameInDefaultsId-83]
	_ = x[foldEmptyJoinsId-84]
	_ = x[optimizeJoinsId-85]
	_ = x[pushdownFiltersId-86]
	_ = x[subqueryIndexesId-87]
	_ = x[pruneTablesId-88]
	_ = x[setJoinScopeLenId-89]
	_ = x[eraseProjectionId-90]
	_ = x[replaceSortPkId-91]
	_ = x[insertTopNId-92]
	_ = x[applyHashInId-93]
	_ = x[resolveInsertRowsId-94]
	_ = x[resolvePreparedInsertId-95]
	_ = x[applyTriggersId-96]
	_ = x[applyProceduresId-97]
	_ = x[assignRoutinesId-98]
	_ = x[modifyUpdateExprsForJoinId-99]
	_ = x[applyRowUpdateAccumulatorsId-100]
	_ = x[wrapWithRollbackId-101]
	_ = x[applyFKsId-102]
	_ = x[validateResolvedId-103]
	_ = x[validateOrderById-104]
	_ = x[validateGroupById-105]
	_ = x[validateSchemaSourceId-106]
	_ = x[validateIndexCreationId-107]
	_ = x[validateOperandsId-108]
	_ = x[validateCaseResultTypesId-109]
	_ = x[validateIntervalUsageId-110]
	_ = x[validateExplodeUsageId-111]
	_ = x[validateSubqueryColumnsId-112]
	_ = x[validateUnionSchemasMatchId-113]
	_ = x[validateAggregationsId-114]
	_ = x[validateDeleteFromId-115]
	_ = x[cacheSubqueryResultsId-116]
	_ = x[cacheSubqueryAliasesInJoinsId-117]
	_ = x[AutocommitId-118]
	_ = x[TrackProcessId-119]
	_ = x[parallelizeId-120]
	_ = x[clearWarningsId-121]
}

const _RuleId_name = "applyDefaultSelectLimitvalidateOffsetAndLimitvalidateStarExpressionsvalidateCreateTablevalidateExprSemresolveVariablesresolveNamedWindowsresolveSetVariablesresolveViewsliftCtesresolveCtesliftRecursiveCtesresolveDatabasesresolveTablesloadStoredProceduresvalidateDropTablespruneDropTablessetTargetSchemasresolveCreateLikeparseColumnDefaultsresolveDropConstraintvalidateDropConstraintloadCheckConstraintsassignCatalogresolveAnalyzeTablesresolveCreateSelectresolveSubqueriessetViewTargetSchemaresolveUnionsresolveDescribeQuerycheckUniqueTableNamesdisambiguateTableFunctionsresolveTableFunctionsresolveDeclarationsresolveColumnDefaultsvalidateColumnDefaultsvalidateCreateTriggervalidateCreateProcedureresolveCreateProcedureloadInfoSchemavalidateReadOnlyDatabasevalidateReadOnlyTransactionvalidateDatabaseSetvalidatePrivilegesreresolveTablessetInsertColumnsvalidateJoinComplexityapplyBinlogReplicaControllerresolveNaturalJoinsresolveOrderbyLiteralsresolveFunctionsflattenTableAliasespushdownSortpushdownGroupbyAliasespushdownSubqueryAliasFiltersqualifyColumnsresolveColumnsvalidateCheckConstraintresolveBarewordSetVariablesreplaceCountStarexpandStarstransposeRightJoinsresolveHavingmergeUnionSchemasflattenAggregationExprsreorderProjectionresolveSubqueryExprsreplaceCrossJoinsmoveJoinCondsToFilterevalFilteroptimizeDistincthoistOutOfScopeFilterstransformJoinApplyhoistSelectExistsfinalizeSubqueriesfinalizeUnionsloadTriggersloadEventsprocessTruncateresolveAlterColumnresolveGeneratorsremoveUnnecessaryConvertspruneColumnsstripTableNamesFromColumnDefaultsfoldEmptyJoinsoptimizeJoinspushdownFilterssubqueryIndexespruneTablessetJoinScopeLeneraseProjectionreplaceSortPkinsertTopNapplyHashInresolveInsertRowsresolvePreparedInsertapplyTriggersapplyProceduresassignRoutinesmodifyUpdateExprsForJoinapplyRowUpdateAccumulatorsrollback triggersapplyFKsvalidateResolvedvalidateOrderByvalidateGroupByvalidateSchemaSourcevalidateIndexCreationvalidateOperandsvalidateCaseResultTypesvalidateIntervalUsagevalidateExplodeUsagevalidateSubqueryColumnsvalidateUnionSchemasMatchvalidateAggregationsvalidateDeleteFromcacheSubqueryResultscacheSubqueryAliasesInJoinsaddAutocommitNodetrackProcessparallelizeclearWarnings"

var _RuleId_index = [...]uint16{0, 23, 45, 68, 87, 102, 118, 137, 156, 168, 176, 187, 204, 220, 233, 253, 271, 286, 302, 319, 338, 359, 381, 401, 414, 434, 453, 470, 489, 502, 522, 543, 569, 590, 609, 630, 652, 673, 696, 718, 732, 756, 783, 802, 820, 835, 851, 873, 901, 920, 942, 958, 977, 989, 1011, 1039, 1053, 1067, 1090, 1117, 1133, 1144, 1163, 1176, 1193, 1216, 1233, 1253, 1270, 1291, 1301, 1317, 1339, 1357, 1374, 1392, 1406, 1418, 1428, 1443, 1461, 1478, 1503, 1515, 1548, 1562, 1575, 1590, 1605, 1616, 1631, 1646, 1659, 1669, 1680, 1697, 1718, 1731, 1746, 1760, 1784, 1810, 1827, 1835, 1851, 1866, 1881, 1901, 1922, 1938, 1961, 1982, 2002, 2025, 2050, 2070, 2088, 2108, 2135, 2152, 2164, 2175, 2188}

func (i RuleId) String() string {
	if i < 0 || i >= RuleId(len(_RuleId_index)-1) {
		return "RuleId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleId_name[_RuleId_index[i]:_RuleId_index[i+1]]
}
