package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Template file not specified")
		os.Exit(1)
	}

	filename := os.Args[1]
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	i := &InputTemplate{}
	json.Unmarshal(byteValue, i)

	// Prepare output
	o := OutputTemplate{}
	o.Name = i.Name
	s := ScenarioObject{}
	s.StartTime = i.StartTime
	// generate pods
	for index, podName := range i.PodNames {
		u := UnscheduledPod{}
		u.UID = uuid.New().String()
		u.Name = podName
		u.Namespace = "default"
		u.NodeName = ""
		u.NominatedNodeName = ""
		u.Labels = map[string]string{"app.kubernetes.io/component": i.Name, "app.kubernetes.io/name": u.Name}
		u.Requests = i.Requests
		u.Spec.Priority = i.Priorities[index]

		for _, container := range i.InputContainerTemplates {
			u.Spec.Containers = append(u.Spec.Containers, ContainerObject{Name: container.Name, Image: "nginx", Resources: ResourcesObject{Requests: container.Requests}})
		}
		s.UnscheduledPods = append(s.UnscheduledPods, u)
	}
	// generate node groups
	for index, nodeGroup := range i.InputNodeGroupTemplates {
		n := NodeGroup{}
		n.RowID = index
		n.Name = uuid.New().String()
		n.Architecture = nodeGroup.Architecture
		n.CurrentSize = nodeGroup.CurrentSize
		n.TargetSize = nodeGroup.TargetSize
		n.MachineType = nodeGroup.MachineType
		n.MaxSize = nodeGroup.MaxSize
		n.MinSize = nodeGroup.MinSize
		n.Zone = nodeGroup.Zone
		n.PoolName = nodeGroup.PoolName
		n.PoolZones = nodeGroup.Zone
		n.PoolMin = nodeGroup.PoolMin
		n.PoolMax = nodeGroup.PoolMax
		s.NodeGroups = append(s.NodeGroups, n)

	}
	o.Scenarios = []ScenarioObject{s}

	out, err := json.Marshal(o)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	outputFilename := strings.Split(filename, ".")[0] + "_out.json"

	if err := os.WriteFile(outputFilename, out, 0644); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
