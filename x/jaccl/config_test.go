package jaccl

import (
	"errors"
	"path/filepath"
	"testing"
)

type testTopology map[int][]int

func (t testTopology) Peers(rank int) []int { return t[rank] }

func TestConfigValidate(t *testing.T) {
	valid := Config{
		Rank:        0,
		Size:        2,
		GroupID:     "test",
		Coordinator: "127.0.0.1:9000",
		Port:        1,
		Topology:    Mesh(2),
	}
	tests := []struct {
		name string
		edit func(*Config)
	}{
		{"size zero", func(c *Config) { c.Size = 0 }},
		{"size one", func(c *Config) { c.Size = 1 }},
		{"rank", func(c *Config) { c.Rank = 2 }},
		{"group id", func(c *Config) { c.GroupID = "" }},
		{"coordinator", func(c *Config) { c.Coordinator = "" }},
		{"port", func(c *Config) { c.Port = 0 }},
		{"topology", func(c *Config) { c.Topology = nil }},
		{"self peer", func(c *Config) { c.Topology = testTopology{0: {0}} }},
		{"duplicate peer", func(c *Config) { c.Topology = testTopology{0: {1, 1}} }},
		{"out of range peer", func(c *Config) { c.Topology = testTopology{0: {2}} }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := valid
			test.edit(&cfg)
			_, err := cfg.validate()
			if !errors.Is(err, ErrInvalidConfig) {
				t.Fatalf("validate error = %v, want ErrInvalidConfig", err)
			}
		})
	}
	peers, err := valid.validate()
	if err != nil {
		t.Fatal(err)
	}
	if len(peers) != 1 || peers[0] != 1 {
		t.Fatalf("peers = %v, want [1]", peers)
	}
}

func TestConfigRejectsNonJACCLTopology(t *testing.T) {
	config := Config{
		Rank: 0, Size: 4, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1,
		Topology: testTopology{0: {1}, 1: {0}, 2: {3}, 3: {2}},
	}
	if _, err := config.validate(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("validate error = %v, want ErrInvalidConfig", err)
	}
}

func TestConfigRejectsMeshAboveJACCLLimit(t *testing.T) {
	config := Config{
		Rank: 0, Size: 9, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Mesh(9),
	}
	if _, err := config.validate(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("validate error = %v, want ErrInvalidConfig", err)
	}
}

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("JACCL_IBV_DEVICES", filepath.Join("testdata", "jaccl_devices.json"))
	t.Setenv("JACCL_COORDINATOR", "127.0.0.1:9000")
	t.Setenv("JACCL_RANK", "1")
	config, err := ConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if config.Rank != 1 || config.Size != 3 || config.Device != "rdma0" || config.Port != 1 {
		t.Fatalf("config = %+v", config)
	}
	if peers := config.Topology.Peers(1); len(peers) != 2 || peers[0] != 0 || peers[1] != 2 {
		t.Fatalf("topology peers = %v, want [0 2]", peers)
	}
	t.Setenv("JACCL_RING", "1")
	config, err = ConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if peers := config.Topology.Peers(1); len(peers) != 2 || peers[0] != 0 || peers[1] != 2 {
		t.Fatalf("ring topology peers = %v, want [0 2]", peers)
	}
}

func TestConfigFromEnvMissing(t *testing.T) {
	t.Setenv("JACCL_IBV_DEVICES", "")
	t.Setenv("MLX_IBV_DEVICES", "")
	if _, err := ConfigFromEnv(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("ConfigFromEnv error = %v, want ErrInvalidConfig", err)
	}
}

func TestConfigFromEnvAliasesAndAToIRank(t *testing.T) {
	t.Setenv("JACCL_IBV_DEVICES", "")
	t.Setenv("JACCL_COORDINATOR", "")
	t.Setenv("JACCL_RANK", "")
	t.Setenv("MLX_IBV_DEVICES", filepath.Join("testdata", "jaccl_devices.json"))
	t.Setenv("MLX_JACCL_COORDINATOR", "127.0.0.1:9000")
	t.Setenv("MLX_RANK", "2")
	if _, err := ConfigFromEnv(); err == nil {
		t.Fatal("ConfigFromEnv accepted an empty primary JACCL device path")
	}

	t.Setenv("JACCL_IBV_DEVICES", filepath.Join("testdata", "jaccl_devices.json"))
	t.Setenv("JACCL_COORDINATOR", "127.0.0.1:9000")
	t.Setenv("JACCL_RANK", "1 trailing")
	config, err := ConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if config.Rank != 1 {
		t.Fatalf("rank = %d, want 1", config.Rank)
	}
}

func TestParseJACCLRank(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"", 0},
		{"  +12rest", 12},
		{"-7", -7},
		{"abc", 0},
	}
	for _, test := range tests {
		got, err := parseJACCLRank(test.text)
		if err != nil || got != test.want {
			t.Fatalf("parseJACCLRank(%q) = %d, %v, want %d, nil", test.text, got, err, test.want)
		}
	}
}

func TestLinkPlans(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want []linkPlan
	}{
		{
			name: "mesh selects first device",
			cfg: Config{
				Rank: 0, Size: 3, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Mesh(3),
				Devices: [][][]string{
					{nil, {"rdma0", "rdma1"}, {"rdma2"}},
					{{"rdma0"}, nil, {"rdma2"}},
					{{"rdma0"}, {"rdma2"}, nil},
				},
			},
			want: []linkPlan{{Peer: 1, Device: "rdma0"}, {Peer: 2, Device: "rdma2"}},
		},
		{
			name: "ring preserves each directional wire",
			cfg: Config{
				Rank: 1, Size: 3, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Ring(3), PreferRing: true,
				Devices: [][][]string{
					{nil, {"r01a", "r01b"}, {"r02a", "r02b"}},
					{{"r10a", "r10b"}, nil, {"r12a", "r12b"}},
					{{"r20a", "r20b"}, {"r21a", "r21b"}, nil},
				},
			},
			want: []linkPlan{
				{Peer: 0, Direction: -1, Wire: 0, Device: "r10a"},
				{Peer: 0, Direction: -1, Wire: 1, Device: "r10b"},
				{Peer: 2, Direction: 1, Wire: 0, Device: "r12a"},
				{Peer: 2, Direction: 1, Wire: 1, Device: "r12b"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.cfg.linkPlans()
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(test.want) {
				t.Fatalf("plans = %#v, want %#v", got, test.want)
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Fatalf("plans = %#v, want %#v", got, test.want)
				}
			}
		})
	}
}

func TestLinkPlansRejectUnsupportedDeviceConnectivity(t *testing.T) {
	config := Config{
		Rank: 0, Size: 9, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Mesh(9),
		Devices: make([][][]string, 9),
	}
	for rank := range config.Devices {
		config.Devices[rank] = make([][]string, 9)
		for peer := range config.Devices[rank] {
			if rank != peer {
				config.Devices[rank][peer] = []string{"rdma0"}
			}
		}
	}
	if _, err := config.linkPlans(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("linkPlans error = %v, want ErrInvalidConfig", err)
	}
}

func TestLinkPlansRejectTopologyDeviceMismatch(t *testing.T) {
	config := Config{
		Rank: 0, Size: 4, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Mesh(4),
		Devices: [][][]string{
			{nil, {"r01"}, nil, {"r03"}},
			{{"r10"}, nil, {"r12"}, nil},
			{nil, {"r21"}, nil, {"r23"}},
			{{"r30"}, nil, {"r32"}, nil},
		},
	}
	if _, err := config.linkPlans(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("linkPlans error = %v, want ErrInvalidConfig", err)
	}
}

func TestMesh(t *testing.T) {
	tests := []struct {
		rank int
		want []int
	}{
		{0, []int{1, 2}},
		{1, []int{0, 2}},
		{2, []int{0, 1}},
	}
	for _, test := range tests {
		got := Mesh(3).Peers(test.rank)
		if len(got) != len(test.want) {
			t.Fatalf("rank %d peers = %v, want %v", test.rank, got, test.want)
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Fatalf("rank %d peers = %v, want %v", test.rank, got, test.want)
			}
		}
	}
}

func TestMeshInvalidSizeDoesNotPanic(t *testing.T) {
	if peers := Mesh(-1).Peers(0); peers != nil {
		t.Fatalf("negative mesh peers = %v, want nil", peers)
	}
}

func TestRing(t *testing.T) {
	tests := []struct {
		size int
		rank int
		want []int
	}{
		{2, 0, []int{1}},
		{2, 1, []int{0}},
		{4, 0, []int{3, 1}},
		{4, 2, []int{1, 3}},
	}
	for _, test := range tests {
		got := Ring(test.size).Peers(test.rank)
		if len(got) != len(test.want) {
			t.Fatalf("ring size=%d rank=%d peers=%v, want %v", test.size, test.rank, got, test.want)
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Fatalf("ring size=%d rank=%d peers=%v, want %v", test.size, test.rank, got, test.want)
			}
		}
	}
	if peers := Ring(1).Peers(0); peers != nil {
		t.Fatalf("single-rank ring peers=%v, want nil", peers)
	}
}

func TestDTypeSize(t *testing.T) {
	tests := []struct {
		dtype DType
		want  int
	}{
		{Bool, 1},
		{Int8, 1},
		{Int16, 2},
		{Int32, 4},
		{Int64, 8},
		{UInt8, 1},
		{UInt16, 2},
		{UInt32, 4},
		{UInt64, 8},
		{Float16, 2},
		{BFloat16, 2},
		{Int32, 4},
		{Float32, 4},
		{Float64, 8},
		{Complex64, 8},
	}
	for _, test := range tests {
		got, err := test.dtype.Size()
		if err != nil || got != test.want {
			t.Fatalf("dtype %d size = %d, %v, want %d, nil", test.dtype, got, err, test.want)
		}
	}
	if _, err := DType(99).Size(); !errors.Is(err, ErrProtocol) {
		t.Fatalf("invalid dtype error = %v, want ErrProtocol", err)
	}
}
