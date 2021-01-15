package paragraph

import "testing"

func TestSpacing_SetBefore(t *testing.T) {
    s := Spacing{}
    if true == s.isSet {
        t.Fatal("验证失败")
    }

    if nil != s.before {
        t.Fatal("验证失败")
    }

    s.SetBefore(30)
    if false == s.isSet {
        t.Fatal("验证失败")
    }

    if nil == s.before {
        t.Fatal("验证失败")
    }

    if 30 != *s.before {
        t.Fatal("验证失败")
    }
}

func TestSpacing_SetAfter(t *testing.T) {
    s := Spacing{}
    if true == s.isSet {
        t.Fatal("验证失败")
    }

    if nil != s.after {
        t.Fatal("验证失败")
    }

    s.SetAfter(30)
    if false == s.isSet {
        t.Fatal("验证失败")
    }

    if nil == s.after {
        t.Fatal("验证失败")
    }

    if 30 != *s.after {
        t.Fatal("验证失败")
    }
}

func TestSpacing_SetSpace(t *testing.T) {
    s := Spacing{}
    if true == s.isSet {
        t.Fatal("验证失败")
    }

    if nil != s.before {
        t.Fatal("验证失败")
    }
    if nil != s.after {
        t.Fatal("验证失败")
    }

    s.SetSpace(30)
    if false == s.isSet {
        t.Fatal("验证失败")
    }

    if nil == s.before {
        t.Fatal("验证失败")
    }
    if nil == s.after {
        t.Fatal("验证失败")
    }

    if 30 != *s.before {
        t.Fatal("验证失败")
    }
    if 30 != *s.after {
        t.Fatal("验证失败")
    }
}

func TestSpacing_SetLine(t *testing.T) {
    // TODO
}

func TestSpacing_SetLineRule(t *testing.T) {
    // TODO
}

func TestSpacing_SetBeforeAutoSpacing(t *testing.T) {
    // TODO
}

func TestSpacing_SetAfterAutoSpacing(t *testing.T) {
    // TODO
}

func TestSpacing_GetXmlBytes(t *testing.T) {
    // TODO
}
