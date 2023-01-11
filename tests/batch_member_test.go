package tests

import (
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/msq"
	"github.com/ixre/go2o/tests/ti"
	"strconv"
	"testing"
	"time"
)

/**
 * Copyright 2009-2019 @ 56x.net
 * name : batch_member_test.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2019-06-05 01:02
 * description :
 * history :
 */
func init() {
	// 初始化producer
	//msq.Configure(msq.KAFKA, []string{"127.0.0.1:9092"})
}

func TestBatchPushMember(t *testing.T) {
	defer msq.Close()
	orm := ti.Factory.GetOrm()

	var members []member.Member
	err := orm.SelectByQuery(&members, "select * FROM mm_member where id > 0 LIMIT 100 OFFSET 0")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, v := range members {
		id := int(v.Id)
		msq.Push(msq.MemberUpdated, "update|"+strconv.Itoa(id))
		msq.PushDelay(msq.MemberProfileUpdated, strconv.Itoa(id), 1000)
		msq.PushDelay(msq.MemberRelationUpdated, strconv.Itoa(id), 1000)
		t.Log("notify ", id)
	}
	t.Log("finished")
	time.Sleep(5 * time.Minute)
}
