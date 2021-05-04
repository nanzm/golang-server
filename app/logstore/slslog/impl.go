package slslogComponent

import (
	store "dora/app/logstore/core"
	"dora/app/logstore/response"
	"dora/pkg"
	"dora/pkg/logger"
	"fmt"
)

type slsQuery struct {
}

func (s slsQuery) GetLogByMd5(from, to int64, md5 string) (*response.LogsResponse, error) {
	queryExp := fmt.Sprintf("* and agg_md5: %s", md5)

	r, err := baseQueryLogs(from, to, queryExp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	result := &response.LogsResponse{
		Count: r.Count,
		Logs:  r.Logs,
	}
	return result, err
}

func (s slsQuery) LogCountByMd5(from, to int64, md5 string) (*response.LogCountByMd5Res, error) {
	queryExp := fmt.Sprintf(`* and agg_md5: %s | SELECT COUNT(*) as count, approx_distinct("_uuid") as effect_user`, md5)

	slsRes, err := baseQueryLogs(from, to, queryExp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.LogCountByMd5Res{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error) {
	exp, err := buildQueryExp(appId, PvUvTotal)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.PvUvTotalRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PvUvTrend(appId string, from, to, interval int64) (*response.PvUvTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, PvUvTrend)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.PvUvTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.PvUvTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.PvUvTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) SdkVersionCount(appId string, from, to int64) (*response.SdkVersionCountRes, error) {
	exp, err := buildQueryExp(appId, SdkVersionCount)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.SdkVersionCountRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		list := make([]*response.SdkVersionItem, 0)
		for _, log := range slsRes.Logs {
			item := &response.SdkVersionItem{}
			err := pkg.WeekDecode(log, item)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			list = append(list, item)
		}

		result.List = list
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) CategoryCount(appId string, from, to int64) (*response.CategoryCountRes, error) {
	exp, err := buildQueryExp(appId, CategoryCount)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.CategoryCountRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		list := make([]*response.CategoryCountItem, 0)
		for _, log := range slsRes.Logs {
			item := &response.CategoryCountItem{}
			err := pkg.WeekDecode(log, item)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			list = append(list, item)
		}

		result.List = list
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PagesCount(appId string, from, to int64) (*response.PageTotalRes, error) {
	exp, err := buildQueryExp(appId, PageTotal)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.PageTotalRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		list := make([]*response.PageTotalItemRes, 0)
		for _, log := range slsRes.Logs {
			item := &response.PageTotalItemRes{}
			err := pkg.WeekDecode(log, item)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			list = append(list, item)
		}

		result.List = list
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ErrorCount(appId string, from, to int64) (*response.ErrorCountRes, error) {
	exp, err := buildQueryExp(appId, ErrorCount)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.ErrorCountRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ErrorCountTrend(appId string, from, to, interval int64) (*response.ErrorCountTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, ErrorCountTrend)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ErrorCountTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ErrorCountTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.ErrorCountTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ApiErrorCount(appId string, from, to int64) (*response.ApiErrorCountRes, error) {
	exp, err := buildQueryExp(appId, ApiErrorCount)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.ApiErrorCountRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ApiErrorTrend(appId string, from, to int64, interval int64) (*response.ApiErrorTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, ApiErrorTrend)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ApiErrorTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ApiErrorTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.ApiErrorTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ApiErrorList(appId string, from, to int64) (*response.ApiErrorListRes, error) {
	exp, err := buildQueryExp(appId, ApiErrorList)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ApiErrorListRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ApiErrorItem, 0)
		for i, log := range slsRes.Logs {
			trendItem := &response.ApiErrorItem{Id: i}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfNavigationTimingTrend(appId string, from, to int64, interval int64) (*response.PerfNavigationTimingTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, PerfNavigationTimingTrend)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.PerfNavigationTimingTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.NavigationTimingTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.NavigationTimingTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfNavigationTimingValues(appId string, from, to int64) (*response.PerfNavigationTimingValuesRes, error) {
	exp, err := buildQueryExp(appId, PerfNavigationTimingValues)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.PerfNavigationTimingValuesRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfDataConsumptionTrend(appId string, from, to int64, interval int64) (*response.PerfDataConsumptionTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, PerfDataConsumption)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.PerfDataConsumptionTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.PerfDataConsumptionTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.PerfDataConsumptionTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfDataConsumptionValues(appId string, from, to int64) (*response.PerfDataConsumptionValuesRes, error) {
	exp, err := buildQueryExp(appId, PerfDataConsumptionValues)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.PerfDataConsumptionValuesRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfMetricsTrend(appId string, from, to int64, interval int64) (*response.PerfMetricsTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, PerfMetrics)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.PerfMetricsTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.PerfMetricsTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.PerfMetricsTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) PerfMetricsValues(appId string, from, to int64) (*response.PerfMetricsValuesRes, error) {
	exp, err := buildQueryExp(appId, PerfMetricsValues)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.PerfMetricsValuesRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ResLoadFailTotalTrend(appId string, from, to, interval int64) (*response.ResLoadFailTotalTrendRes, error) {
	exp, err := buildQueryTrendExp(appId, interval, ResLoadFailTotalTrend)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ResLoadFailTotalTrendRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ResLoadFailTotalTrendItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.ResLoadFailTotalTrendItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error) {
	exp, err := buildQueryExp(appId, ResLoadFailTotal)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.ResLoadFailTotalRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ResLoadFailList(appId string, from, to int64) (*response.ResLoadFailListRes, error) {
	exp, err := buildQueryExp(appId, ResLoadFailList)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ResLoadFailListRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ResLoadFailItemRes, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.ResLoadFailItemRes{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectIpToCountry(appId string, from, to int64) (*response.ProjectIpToCountryRes, error) {
	panic("implement me")
}

func (s slsQuery) ProjectIpToProvince(appId string, from, to int64) (*response.ProjectIpToProvinceRes, error) {
	panic("implement me")
}

func (s slsQuery) ProjectIpToCity(appId string, from, to int64) (*response.ProjectIpToCityRes, error) {
	panic("implement me")
}

func (s slsQuery) ProjectEventCount(appId string, from, to int64) (*response.ProjectEventCountRes, error) {
	exp, err := buildQueryExp(appId, ProjectEventCount)
	if err != nil {
		return nil, err
	}

	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		input := slsRes.Logs[0]
		result := &response.ProjectEventCountRes{}
		err := pkg.WeekDecode(input, result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectSendMode(appId string, from, to int64) (*response.ProjectSendModeRes, error) {
	exp, err := buildQueryExp(appId, ProjectSendMode)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ProjectSendModeRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.SendModeItem, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.SendModeItem{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectVersion(appId string, from, to int64) (*response.ProjectVersionRes, error) {
	exp, err := buildQueryExp(appId, ProjectVersion)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ProjectVersionRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.VersionItem, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.VersionItem{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectUserScreen(appId string, from, to int64) (*response.ProjectUserScreenRes, error) {
	exp, err := buildQueryExp(appId, projectUserScreen)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ProjectUserScreenRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.ScreenItem, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.ScreenItem{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectCategory(appId string, from, to int64) (*response.ProjectCategoryRes, error) {
	exp, err := buildQueryExp(appId, projectCategory)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ProjectCategoryRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.CategoryItem, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.CategoryItem{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func (s slsQuery) ProjectEnv(appId string, from, to int64) (*response.ProjectEnvRes, error) {
	exp, err := buildQueryExp(appId, ProjectEnv)
	if err != nil {
		return nil, err
	}
	slsRes, err := baseQueryLogs(from, to, exp)
	if err != nil {
		logger.Printf("query log err : %v", err)
		return nil, err
	}

	if len(slsRes.Logs) > 0 {
		result := &response.ProjectEnvRes{
			Total: len(slsRes.Logs),
		}

		// 遍历
		trendList := make([]*response.EnvItem, 0)
		for _, log := range slsRes.Logs {
			trendItem := &response.EnvItem{}
			err := pkg.WeekDecode(log, trendItem)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			trendList = append(trendList, trendItem)
		}

		result.List = trendList
		return result, nil
	}
	return nil, nil
}

func NewSlsQuery() store.QueryMethods {
	return &slsQuery{}
}
