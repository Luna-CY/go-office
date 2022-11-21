package xlsx1

import "testing"

func TestRootRelationship_FilePath(t *testing.T) {
	rs := RootRelationship{}

	if "_rels/.rels" != rs.Filepath() {
		t.Fatal("测试失败")
	}
}

func TestRootRelationship_Marshal(t *testing.T) {
	rs := RootRelationship{}
	rs.Relationships = append(rs.Relationships, Relationship{Id: "rId1", Type: "test-type", Target: "test-target"})

	content, err := rs.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Relationships xmlns=""><Relationship Id="rId1" Type="test-type" Target="test-target"></Relationship></Relationships>` != string(content) {
		t.Fatal("测试失败")
	}
}
