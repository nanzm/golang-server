package model

func Tables() []interface{} {
	return []interface{}{
		new(SysLog),

		// 项目 用户
		new(Project),
		new(Role),
		new(User),
		new(UserSetting),

		// 监控业务
		new(Issue),
		new(IssueUserStatus),
		new(SourceMap),

		// 制品 项目备份
		new(Artifact),

		// 告警
		new(AlarmLog),
	}
}
