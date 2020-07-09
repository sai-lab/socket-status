package status

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/host"
)

func GetServerStat() ServerStat {
	var wg sync.WaitGroup
	var ss ServerStat

	wg.Add(2)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetHostStat()
	}(&ss)
	go func(d *ServerStat) {
		defer wg.Done()
		//		d.GetSocketStatNet()
		//		d.GetSocketStatSs()
		d.GetSocketStatSock()
	}(&ss)

	wg.Wait()
	ss.GetTime()

	return ss
}

func (s *ServerStat) GetHostStat() {
	h, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}
	s.HostName = h.Hostname
	s.HostID = h.HostID
	s.VirtualizationSystem = h.VirtualizationSystem
}

func (s *ServerStat) GetSocketStatNet(h string) {
	result := 0
	out, err := exec.Command("sudo", "netstat", "-na").Output()
	if err != nil {
		log.Fatal(err)
	}
	lin := strings.Split(string(out), "\n")
	for _, v := range lin {
		words := strings.Fields(v)
		if len(words) == 0 {
			continue
		} else if words[0] == "tcp" {
			switch words[len(words)-1] {
			case "ESTABLISHED":
				result++
			case "SYN_SENT":
				result++
			case "SYN_RECEIVED":
				result++
			default:
			}
		}
	}
	s.Socket = string(result)
}

func (s *ServerStat) GetSocketStatSs(h string) {
	result := 0
	out, err := exec.Command("sudo", "ss", "-nta").Output()
	if err != nil {
		log.Fatal(err)
	}

	lin := strings.Split(string(out), "\n")
	for _, v := range lin {
		words := strings.Fields(v)
		if len(words) == 0 {
			continue
		} else if words[0] == "State" {
			continue
		}
		switch words[0] {
		case "ESTAB":
			result++
		case "SYN-SENT":
			result++
		case "SYN-RECEIVED":
			result++
		default:
		}
	}
	s.Socket = string(result)
}

func (s *ServerStat) GetSocketStatSock() {
	result := 0
	flag := false
	out, err := exec.Command("sudo", "cat", "/proc/net/sockstat").Output()
	if err != nil {
		log.Fatal(err)
	}

	lin := strings.Split(string(out), "\n")
	for _, v := range lin {
		words := strings.Fields(v)
		if len(words) == 0 {
			continue
		}
		if words[0] == "TCP:" {
			for i := 1; i < len(words)-1; i++ {
				if words[i] == "inuse" {
					flag = true
					continue
				}
				if flag {
					result, _ = strconv.Atoi(string(words[i]))
					break
				}
			}
		}
		if flag {
			flag = false
			break
		}

		s.Socket = strconv.Itoa(result)
	}
}

func (s *ServerStat) GetTime() {
	now := time.Now()
	s.Time = now.String()
}
