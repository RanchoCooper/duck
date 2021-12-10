package duck

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/11
 */

func TestCopyFields(t *testing.T) {
    type A struct {
        ID   int
        Name string
    }
    type B struct {
        ID   int
        Name string
    }
    a := A{
        ID:   1,
        Name: "rancho",
    }
    var b B
    err := CopyFields(&a, &b)
    assert.Nil(t, err)
    assert.Equal(t, a.ID, b.ID)
    assert.Equal(t, a.Name, b.Name)
}
