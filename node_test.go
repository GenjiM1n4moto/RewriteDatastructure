package rewritedatastructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOutput(t *testing.T) {
	testNodeList1 := &Node{
		val: 1,
		next: &Node{
			val:  "test",
			next: nil,
		},
	}
	testNodeList2 := &Node{}
	testNodeList1Array, err := testNodeList1.output()
	require.NoError(t, err)
	require.Equal(t, []interface{}{1, "test"}, testNodeList1Array)
	testNodeList2Array, err := testNodeList2.output()
	require.Nil(t, testNodeList2Array)
	require.Error(t, err)
}

func TestInsert(t *testing.T) {
	testNodeList1 := &Node{
		val: 1,
		next: &Node{
			val:  "test",
			next: nil,
		},
	}
	err := testNodeList1.insert("test2")
	require.NoError(t, err)
	testNodeList1Array, err := testNodeList1.output()
	require.NoError(t, err)
	require.Equal(t, testNodeList1Array, []interface{}{1, "test", "test2"})

	testNodeList2 := &Node{}
	err = testNodeList2.insert(1)
	require.Error(t, err)
}

func TestDeleteIndex(t *testing.T) {
	testNodeList := &Node{
		val:  1,
		next: nil,
	}
	_, err := testNodeList.deleteIndex(1)
	require.Error(t, err)
	_, err = testNodeList.deleteIndex(0)
	require.Error(t, err)

	testNodeList = &Node{
		val: 1,
		next: &Node{
			val: "test",
			next: &Node{
				val:  "test2",
				next: nil,
			},
		},
	}
	head, err := testNodeList.deleteIndex(1)
	require.NoError(t, err)
	outputhead, err := head.output()
	require.NoError(t, err)
	require.Equal(t, []interface{}{1, "test2"}, outputhead)

	testNodeList = &Node{
		val: 1,
		next: &Node{
			val: "test",
			next: &Node{
				val: "test2",
				next: &Node{
					val:  "test3",
					next: nil,
				},
			},
		},
	}
	head, err = testNodeList.deleteIndex(3)
	require.NoError(t, err)
	outputhead, err = head.output()
	require.NoError(t, err)
	require.Equal(t, []interface{}{1, "test", "test2"}, outputhead)

	testNodeList = &Node{}
	_, err = testNodeList.deleteIndex(0)
	require.Error(t, err)

	testNodeList = &Node{
		val: 1,
		next: &Node{
			val: "test",
			next: &Node{
				val: "test2",
				next: &Node{
					val:  "test3",
					next: nil,
				},
			},
		},
	}
	head, err = testNodeList.deleteIndex(0)
	require.NoError(t, err)
	outputhead, err = head.output()
	require.NoError(t, err)
	require.Equal(t, []interface{}{"test", "test2", "test3"}, outputhead)

	testNodeList = &Node{
		val: 1,
		next: &Node{
			val: "test",
			next: &Node{
				val: "test2",
				next: &Node{
					val:  "test3",
					next: nil,
				},
			},
		},
	}
	_, err = testNodeList.deleteIndex(10)
	require.Error(t, err)
}

func TestReverse(t *testing.T) {
	testNodeList := &Node{
		val: 1,
		next: &Node{
			val: "test",
			next: &Node{
				val: "test2",
				next: &Node{
					val:  "test3",
					next: nil,
				},
			},
		},
	}
	head := testNodeList.reverse()
	testNodeListArray, err := head.output()
	require.NoError(t, err)
	require.Equal(t, []interface{}{"test3", "test2", "test", 1}, testNodeListArray)
}
