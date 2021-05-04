package slslogComponent

//func Test_cleanLogs(t *testing.T) {
//	log1 := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"agg_md5": "{}", "content2": fmt.Sprintf("%v", "1212312312")})
//	log2 := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"agg_md5": "2", "content2": fmt.Sprintf("%v", "1212312312")})
//	log3 := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"agg_md5": "123", "content2": fmt.Sprintf("%v", "1212312312")})
//
//	var logs []*sls.Log
//	logs = append(logs, log1, log2, log3)
//
//	i := cleanLogs(logs)
//	utils.PrettyPrint(i)
//}
//
//func Test_getKeyMd5Val(t *testing.T) {
//	md5 := "agg_md5"
//	d := "10086"
//	a := []*sls.LogContent{{
//		Key:   &md5,
//		Value: &d,
//	}, {
//		Key:   &d,
//		Value: &d,
//	}}
//
//	val := getAggMd5Val(a)
//	if val != "10086" {
//		t.Fatalf("failed go val: %#v \n", val)
//	}
//	t.Logf("success got val : %#v", val)
//}
