// Code generated by "stringer -type=fetcherState"; DO NOT EDIT.

package row

import "strconv"

const _fetcherState_name = "stateInvalidstateInitFetchstateDecodeFirstKVOfRowstateSeekPrefixstateFetchNextKVWithUnfinishedRowstateFinalizeRowstateEmitLastBatchstateFinished"

var _fetcherState_index = [...]uint8{0, 12, 26, 49, 64, 97, 113, 131, 144}

func (i fetcherState) String() string {
	if i < 0 || i >= fetcherState(len(_fetcherState_index)-1) {
		return "fetcherState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _fetcherState_name[_fetcherState_index[i]:_fetcherState_index[i+1]]
}
