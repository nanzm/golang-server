package model

func Tables() []interface{} {
	return []interface{}{
		new(Organization),
		new(Project),
		new(Role),
		new(SysLog),
		new(User),
		new(UserSetting),

		// 监控业务
		new(Issue),
		new(IssueUserStatus),
		new(SourceMap),

		// 制品 项目备份
		new(Artifact),

		// 告警
		new(AlarmProject),
		new(AlarmRule),
		new(AlarmLog),
	}
}
