package dao

//
//func TestAlarmProject_Create(t *testing.T) {
//	dao := NewAlarmProjectDao()
//	create, err := dao.Create(&model.AlarmProject{
//		ProjectId: 1,
//		Name:      "告警111",
//		Type:      "js错误",
//		Silence:   false,
//	})
//	if err != nil {
//		panic(err)
//	}
//	utils.PrettyString(create)
//}
//
//func TestAlarmProject_AppendRules(t *testing.T) {
//	dao := NewAlarmProjectDao()
//	err := dao.AppendRules(1, []model.AlarmRule{model.AlarmRule{ID: 2}})
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestAlarmProject_AppendTargets(t *testing.T) {
//	dao := NewAlarmProjectDao()
//	err := dao.AppendTargets(1, []model.AlarmTarget{model.AlarmTarget{ID: 2}})
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestAlarmProject_List(t *testing.T) {
//	dao := NewAlarmProjectDao()
//	list, err := dao.List()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	utils.PrettyPrint(list)
//}
