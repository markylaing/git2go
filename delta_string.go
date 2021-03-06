// Code generated by "stringer -type Delta -trimprefix Delta -tags static"; DO NOT EDIT.

package git

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeltaUnmodified-0]
	_ = x[DeltaAdded-1]
	_ = x[DeltaDeleted-2]
	_ = x[DeltaModified-3]
	_ = x[DeltaRenamed-4]
	_ = x[DeltaCopied-5]
	_ = x[DeltaIgnored-6]
	_ = x[DeltaUntracked-7]
	_ = x[DeltaTypeChange-8]
	_ = x[DeltaUnreadable-9]
	_ = x[DeltaConflicted-10]
}

const _Delta_name = "UnmodifiedAddedDeletedModifiedRenamedCopiedIgnoredUntrackedTypeChangeUnreadableConflicted"

var _Delta_index = [...]uint8{0, 10, 15, 22, 30, 37, 43, 50, 59, 69, 79, 89}

func (i Delta) String() string {
	if i < 0 || i >= Delta(len(_Delta_index)-1) {
		return "Delta(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Delta_name[_Delta_index[i]:_Delta_index[i+1]]
}
