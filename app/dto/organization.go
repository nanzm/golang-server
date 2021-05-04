package dto

// 创建组织参数
type CreateOrganization struct {
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Type         string `json:"type" binding:"required"`
}

// 获取组织详情参数
type QueryOrganizationDetail struct {
	OrganizationId uint `form:"organization_id"  binding:"required"`
}

// 获取组织成员参数
type GetOrganizationMembers struct {
	OrganizationId uint `form:"organization_id"  binding:"required"`
}

// 添加组织成员参数
type AddOrganizationMembers struct {
	OrganizationId uint   `json:"organization_id"  binding:"required"`
	UserIds         []uint `json:"user_ids"  binding:"required"`
}

// 移除组织成员参数
type RemoveOrganizationMembers struct {
	OrganizationId uint `json:"organization_id"  binding:"required"`
	UserId         uint `json:"user_id"  binding:"required"`
}
