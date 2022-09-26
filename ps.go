package xutil

import (
	"math"
	"net"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	_ = 1 << (10 * iota)
	_KB
	_MB
	_GB
	_TB
)

func Bytes2GB(n uint64) float64 {
	return float64(n) / _GB
}

func Bytes2MB(n uint64) float64 {
	return float64(n) / _MB
}

// firstOrDefaultIntArgs
func firstOrDefaultIntArgs(dft int, args ...int) (val int) {
	val = dft
	if len(args) > 0 {
		val = args[0]
	}
	return val
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision ...int) float64 {
	p := firstOrDefaultIntArgs(2, precision...)
	output := math.Pow(10, float64(p))
	return float64(round(num*output)) / output
}

type PsMetric struct {
	// host info
	Uptime   uint64 `json:"uptime,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	OS       string `json:"os,omitempty"`
	HostId   string `json:"host_id"`

	// mem
	Total       float64 `json:"total,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty"`

	// cpu
	Cores     int32  `json:"cores,omitempty"`
	ModelName string `json:"model_name,omitempty"`

	// load
	Load1  float64 `json:"load1,omitempty"`
	Load5  float64 `json:"load5,omitempty"`
	Load15 float64 `json:"load15,omitempty"`
}

func panicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

func NewPsMetric() *PsMetric {
	p := &PsMetric{}
	p.spawn()
	return p
}

func (p *PsMetric) spawn() {
	p.Mem()
	p.CPU()
	p.OsInfo()
	p.LoadAvg()
}

func RealTimeInfo() map[string]interface{} {
	p := NewPsMetric()
	dat := make(map[string]interface{})
	dat["total"] = p.Total
	dat["used_percent"] = p.UsedPercent
	dat["load1"] = p.Load1
	return dat
}

func (p *PsMetric) Mem() *mem.VirtualMemoryStat {
	v, e := mem.VirtualMemory()
	panicIfErr(e)
	p.Total = toFixed(Bytes2MB(v.Total))
	p.UsedPercent = toFixed(v.UsedPercent)
	return v
}

func (p *PsMetric) OsInfo() *host.InfoStat {
	v, e := host.Info()
	panicIfErr(e)
	p.Uptime = v.Uptime
	p.Hostname = v.Hostname
	p.OS = v.OS
	p.HostId = v.HostID
	return v
}

func (p *PsMetric) CPU() {
	v, e := cpu.Info()
	panicIfErr(e)
	cpu := v[0]

	p.Cores = cpu.Cores
	p.ModelName = cpu.ModelName
}

func (p *PsMetric) LoadAvg() {
	v, e := load.Avg()
	panicIfErr(e)

	p.Load1 = toFixed(v.Load1)
	p.Load5 = toFixed(v.Load5)
	p.Load15 = toFixed(v.Load15)
}

func HostUUID() string {
	v, e := host.HostID()
	panicIfErr(e)
	return v
}

// GetLocalIpAddr
//
//	@return ip:port
func GetLocalIpAddr() (ip, port string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panicIfErr(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	lst := strings.Split(localAddr, ":")
	ip, port = lst[0], lst[1]
	return
}
