package main

import (
	"github.com/ajoshi-nuwm/ai-lab-5-art1/util"
	strings2 "strings"
	"github.com/ajoshi-nuwm/ai-lab-5-art1/cluster"
	"fmt"
)

const (
	attention  = 0.1
	similarity = 0.1
)

var (
	vectors  []*cluster.PropertyVector
	clusters []*cluster.Cluster
)

func main() {
	strings, _ := util.ReadFromFile("C:\\workspace\\bin\\art1.txt")
	subjects := strings2.Split(strings[0], " ")

	for i := 1; i < len(strings); i++ {
		vector := make(map[string]bool)
		vectorString := strings2.Split(strings[i], " ")
		for j := 0; j < len(subjects); j++ {
			vector[subjects[j]] = Stob(vectorString[j])
		}
		vectors = append(vectors, cluster.NewPropertyVector(vector))
	}

	clusters := append(clusters, cluster.NewCluster(vectors[0], attention, similarity))
	Print(vectors, clusters)

	clusters = Calculate(vectors, clusters)
	Print(vectors, clusters)
	vectors = GetNewVectors(clusters)

	Calculate(vectors, clusters)
	Print(vectors, clusters)
}

func Stob(i string) bool {
	if i == "0" {
		return false
	}
	return true
}

func Print(vectors []*cluster.PropertyVector, clusters []*cluster.Cluster) {
	//fmt.Println("----------vectors------------")
	//for _, v := range vectors {
	//	fmt.Println(v)
	//}
	fmt.Println("----------clusters------------")
	for _, c := range clusters {
		fmt.Println(c)
	}
}

func Calculate(vectors []*cluster.PropertyVector, clusters []*cluster.Cluster) []*cluster.Cluster {
	for _, v := range vectors {
		var success = false
		for _, c := range clusters {
			if c.AddVector(v) {
				success = true
				break
			}
		}
		if !success {
			clusters = append(clusters, cluster.NewCluster(v, attention, similarity))
		}
	}
	return clusters
}

func GetNewVectors(clusters []*cluster.Cluster) []*cluster.PropertyVector {
	vectors := make([]*cluster.PropertyVector, 0)
	for _, c := range clusters {
		for _, v := range c.GetVectors() {
			vectors = append(vectors, v)
		}
		c.Clean()
	}
	return vectors
}
