package reader

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/infraboard/keyauth/pkg/ip2region"
)

const (
	INDEX_BLOCK_LENGTH  = 12
	TOTAL_HEADER_LENGTH = 8192
)

// New todo
func New(db io.Reader) (*IPReader, error) {
	dbBinStr, err := ioutil.ReadAll(db)
	if err != nil {
		return nil, fmt.Errorf("load db file to memory error")
	}

	return &IPReader{
		dbBinStr:      dbBinStr,
		firstIndexPtr: getLong(dbBinStr, 0),
		lastIndexPtr:  getLong(dbBinStr, 4),
	}, nil
}

// IPReader tood
type IPReader struct {
	// super block index info
	firstIndexPtr int64
	lastIndexPtr  int64
	dbBinStr      []byte
}

// TotalBlocks todo
func (r *IPReader) TotalBlocks() int64 {
	return (r.lastIndexPtr-r.firstIndexPtr)/INDEX_BLOCK_LENGTH + 1
}

// MemorySearch todo
func (r *IPReader) MemorySearch(ipStr string) (*ip2region.IPInfo, error) {
	ipInfo := ip2region.IPInfo{}

	ip, err := ip2long(ipStr)
	if err != nil {
		return nil, err
	}

	h := r.TotalBlocks()
	var dataPtr, l int64
	for l <= h {
		m := (l + h) >> 1
		p := r.firstIndexPtr + m*INDEX_BLOCK_LENGTH
		sip := getLong(r.dbBinStr, p)
		if ip < sip {
			h = m - 1
		} else {
			eip := getLong(r.dbBinStr, p+4)
			if ip > eip {
				l = m + 1
			} else {
				dataPtr = getLong(r.dbBinStr, p+8)
				break
			}
		}
	}
	if dataPtr == 0 {
		return &ipInfo, errors.New("not found")
	}

	dataLen := ((dataPtr >> 24) & 0xFF)
	dataPtr = (dataPtr & 0x00FFFFFF)
	ipInfo = getIPInfo(getLong(r.dbBinStr, dataPtr), r.dbBinStr[(dataPtr)+4:dataPtr+dataLen])
	return &ipInfo, nil
}

func getLong(b []byte, offset int64) int64 {

	val := (int64(b[offset]) |
		int64(b[offset+1])<<8 |
		int64(b[offset+2])<<16 |
		int64(b[offset+3])<<24)

	return val

}

func ip2long(IPStr string) (int64, error) {
	if IPStr == "" {
		return 0, fmt.Errorf("ip ins \"\"")
	}

	bits := strings.Split(IPStr, ".")
	if len(bits) != 4 {
		return 0, fmt.Errorf("ip [%s] format error", IPStr)
	}

	var sum int64
	for i, n := range bits {
		bit, _ := strconv.ParseInt(n, 10, 64)
		sum += bit << uint(24-8*i)
	}

	return sum, nil
}

func getIPInfo(cityID int64, line []byte) ip2region.IPInfo {
	lineSlice := strings.Split(string(line), "|")
	ipInfo := ip2region.IPInfo{}
	length := len(lineSlice)
	ipInfo.CityID = cityID
	if length < 5 {
		for i := 0; i <= 5-length; i++ {
			lineSlice = append(lineSlice, "")
		}
	}

	ipInfo.Country = lineSlice[0]
	ipInfo.Region = lineSlice[1]
	ipInfo.Province = lineSlice[2]
	ipInfo.City = lineSlice[3]
	ipInfo.ISP = lineSlice[4]
	return ipInfo
}
