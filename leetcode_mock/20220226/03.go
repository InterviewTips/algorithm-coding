package _0220226

//https://leetcode-cn.com/problems/snapshot-array/

type SnapshotArray struct {
	revision int                 // 当前版本
	data     map[int]map[int]int // map[int]int 表示 snap_id val
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		data: make(map[int]map[int]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	_, ok := this.data[index]
	if !ok {
		this.data[index] = make(map[int]int)
	}
	this.data[index][this.revision] = val
	return
}

func (this *SnapshotArray) Snap() int {
	snapId := this.revision
	this.revision++
	return snapId
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	value, ok := this.data[index]
	if !ok {
		return 0 // 直接返回
	}
	// fmt.Println("index is", index, "snap_id is", snap_id, "value is", value)

	for snap_id >= 0 {
		_, ok = value[snap_id]
		if ok {
			return value[snap_id]
		}
		snap_id--
	}

	return value[snap_id]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
