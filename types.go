package main

type OutputTemplate struct {
	Name      string           `json:"Name"`
	Scenarios []ScenarioObject `json:"Scenarios"`
}

type ScenarioObject struct {
	StartTime       string           `json:"StartTime"`
	UnscheduledPods []UnscheduledPod `json:"UnscheduledPods"`
	NodeGroups      []NodeGroup      `json:"NodeGroups"`
}

type UnscheduledPod struct {
	UID               string            `json:"UID"`
	Name              string            `json:"Name"`
	Namespace         string            `json:"Namespace"`
	NodeName          string            `json:"NodeName"`
	NominatedNodeName string            `json:"NominatedNodeName"`
	Labels            map[string]string `json:"Labels"`
	Requests          RequestsObject    `json:"Requests"`
	Spec              SpecObject        `json:"Spec"`
}

type NodeGroup struct {
	RowID        int    `json:"RowID"`
	Name         string `json:"Name"`
	CurrentSize  int    `json:"CurrentSize"`
	TargetSize   int    `json:"TargetSize"`
	MinSize      int    `json:"MinSize"`
	MaxSize      int    `json:"MaxSize"`
	Zone         string `json:"Zone"`
	MachineType  string `json:"MachineType"`
	Architecture string `json:"Architecture"`
	PoolName     string `json:"PoolName"`
	PoolMin      int    `json:"PoolMin"`
	PoolMax      int    `json:"PoolMax"`
	PoolZones    string `json:"PoolZones"`
}

type SpecObject struct {
	Containers []ContainerObject `json:"containers"`
	Priority   int               `json:"priority"`
}

type ContainerObject struct {
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	Resources ResourcesObject `json:"resources"`
}

type ResourcesObject struct {
	Requests RequestsObject `json:"requests"`
}

type RequestsObject struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type InputTemplate struct {
	Name                    string                   `json:"Name"`
	Requests                RequestsObject           `json:"requests"`
	PodNames                []string                 `json:"podNames"`
	Priorities              []int                    `json:"priorities"`
	InputContainerTemplates []InputContainerTemplate `json:"containerTemplate"`
	StartTime               string                   `json:"StartTime"`
	InputNodeGroupTemplates []InputNodeGroupTemplate `json:"nodeGroupTemplate"`
}

type InputContainerTemplate struct {
	Name     string         `json:"name"`
	Requests RequestsObject `json:"requests"`
}
type InputNodeGroupTemplate struct {
	CurrentSize  int    `json:"CurrentSize"`
	TargetSize   int    `json:"TargetSize"`
	MinSize      int    `json:"MinSize"`
	MaxSize      int    `json:"MaxSize"`
	Zone         string `json:"Zone"`
	MachineType  string `json:"MachineType"`
	Architecture string `json:"Architecture"`
	PoolName     string `json:"PoolName"`
	PoolMin      int    `json:"PoolMin"`
	PoolMax      int    `json:"PoolMax"`
}
