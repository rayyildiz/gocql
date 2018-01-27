// +build appengine

package murmur

import "encoding/binary"

func getBlock(data []byte, n int) (int64, int64) {
	k1 := binary.LittleEndian.Uint64(data[n*16:])
	k2 := binary.LittleEndian.Uint64(data[(n*16)+8:])
	return int64(k1), int64(k2)
}
