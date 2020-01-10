package ospf

import "net"

const (
	p1LSUpdateHex = "00000001000b02010afc27120afc27128001c2ae4fd50078020000080afc2712ffffffff0300000a0a6406fe0a6406640200000a0afc23010afc27120100000a0afc2301ffffffff0300000a0afc27010afc27120100000a0afc2701ffffffff0300000a0afc27110afc27120100000a0afc2711ffffffff0300000a"
	p1HeaderHex   = "020400940afc279400000000afe600000000000000000000"
	p1LsaHex      = "000b02010afc27120afc27128001c2ae4fd50078020000080afc2712ffffffff0300000a0a6406fe0a6406640200000a0afc23010afc27120100000a0afc2301ffffffff0300000a0afc27010afc27120100000a0afc2701ffffffff0300000a0afc27110afc27120100000a0afc2711ffffffff0300000a"

	p2LSUpdateHex = "00000002000a22010aff00170aff00178001d20c49ac0054020000050afe0108fffffffc030000320aff000c0afe0217010000010aff000b0afe0217010000010afe0217ffffffff030000000a6417fe0a6417fe02000001000a22020a6417fe0aff00178000707b5cf90024ffffff000aff00170afa00010afc1d04"
	p21LsaHex     = "000a22010aff00170aff00178001d20c49ac0054020000050afe0108fffffffc030000320aff000c0afe0217010000010aff000b0afe0217010000010afe0217ffffffff030000000a6417fe0a6417fe02000001"
	p22LsaHex     = "000a22020a6417fe0aff00178000707b5cf90024ffffff000aff00170afa00010afc1d04"
)

var (
	p1Header = Header{
		Version:  v2,
		Type:     lsUpdate,
		Length:   148,
		RouterID: net.ParseIP("10.252.39.148").To4(),
		AreaID:   net.ParseIP("0.0.0.0").To4(),
	}

	p1LSA = LSA{
		LSAge:       11,
		Options:     2,
		LSType:      LSTypeRouter,
		LinkStateID: net.ParseIP("10.252.39.18").To4(),
		AdvRouter:   net.ParseIP("10.252.39.18").To4(),
		SeqNumber:   2147599022,
		Length:      120,
		Checksum:    20437,
	}

	p21LSA = LSA{
		LSAge:       10,
		Options:     0x22,
		LSType:      LSTypeRouter,
		LinkStateID: net.ParseIP("10.255.0.23").To4(),
		AdvRouter:   net.ParseIP("10.255.0.23").To4(),
		SeqNumber:   0x8001d20c,
		Length:      84,
		Checksum:    0x000049ac,
	}

	p22LSA = LSA{
		LSAge:       10,
		Options:     0x22,
		LSType:      LSTypeNetwork,
		LinkStateID: net.ParseIP("10.100.23.254").To4(),
		AdvRouter:   net.ParseIP("10.255.0.23").To4(),
		SeqNumber:   0x8000707b,
		Length:      36,
		Checksum:    0x00005cf9,
	}

	p1LS1 = LS1{
		LSA: p1LSA,
		LS1Flags: LS1Flags{
			VirtualLink: false,
			ASBR:        true,
			ABR:         false,
		},
		NumOfLinks: 8,
		Links: []LS1Link{
			{
				LinkID:     net.ParseIP("10.252.39.18").To4(),
				Data:       net.ParseIP("255.255.255.255").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.100.6.254").To4(),
				Data:       net.ParseIP("10.100.6.100").To4(),
				Type:       Transit,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.35.1").To4(),
				Data:       net.ParseIP("10.252.39.18").To4(),
				Type:       PTP,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.35.1").To4(),
				Data:       net.ParseIP("255.255.255.255").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.39.1").To4(),
				Data:       net.ParseIP("10.252.39.18").To4(),
				Type:       PTP,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.39.1").To4(),
				Data:       net.ParseIP("255.255.255.255").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.39.17").To4(),
				Data:       net.ParseIP("10.252.39.18").To4(),
				Type:       PTP,
				MetricsNum: 0,
				Metric:     10,
			},
			{
				LinkID:     net.ParseIP("10.252.39.17").To4(),
				Data:       net.ParseIP("255.255.255.255").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     10,
			},
		},
	}

	p21LS1 = LS1{
		LSA: p21LSA,
		LS1Flags: LS1Flags{
			VirtualLink: false,
			ASBR:        true,
			ABR:         false,
		},
		NumOfLinks: 5,
		Links: []LS1Link{
			{
				LinkID:     net.ParseIP("10.254.1.8").To4(),
				Data:       net.ParseIP("255.255.255.252").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     50,
			},
			{
				LinkID:     net.ParseIP("10.255.0.12").To4(),
				Data:       net.ParseIP("10.254.2.23").To4(),
				Type:       PTP,
				MetricsNum: 0,
				Metric:     1,
			},
			{
				LinkID:     net.ParseIP("10.255.0.11").To4(),
				Data:       net.ParseIP("10.254.2.23").To4(),
				Type:       PTP,
				MetricsNum: 0,
				Metric:     1,
			},
			{
				LinkID:     net.ParseIP("10.254.2.23").To4(),
				Data:       net.ParseIP("255.255.255.255").To4(),
				Type:       Stub,
				MetricsNum: 0,
				Metric:     0,
			},
			{
				LinkID:     net.ParseIP("10.100.23.254").To4(),
				Data:       net.ParseIP("10.100.23.254").To4(),
				Type:       Transit,
				MetricsNum: 0,
				Metric:     1,
			},
		},
	}

	p22LS2 = LS2{
		LSA:     p22LSA,
		Netmask: net.ParseIP("255.255.255.0").To4(),
		AttachedRouters: []net.IP{
			net.ParseIP("10.255.0.23").To4(),
			net.ParseIP("10.250.0.1").To4(),
			net.ParseIP("10.252.29.4").To4(),
		},
	}
)