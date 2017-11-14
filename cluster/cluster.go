package cluster

import (
	"fmt"
)

type Cluster struct {
	prototype  *PropertyVector
	vectors    []*PropertyVector
	attention  float64
	similarity float64
}

func NewCluster(prototype *PropertyVector, attention, similarity float64) *Cluster {
	vectors := make([]*PropertyVector, 0)
	return &Cluster{prototype, vectors, attention, similarity}
}

func (cluster *Cluster) GetVectors() []*PropertyVector {
	return cluster.vectors
}

func (cluster *Cluster) AddVector(propertyVector *PropertyVector) bool {
	if cluster.similarityTest(propertyVector) && cluster.attentionTest(propertyVector) {
		cluster.vectors = append(cluster.vectors, propertyVector)
		return true
	}
	return false
}

func (cluster *Cluster) Clean() {
	cluster.vectors = make([]*PropertyVector, 0)
}

func (cluster *Cluster) similarityTest(propertyVector *PropertyVector) bool {
	prototypeValue := float64(cluster.prototype.And(propertyVector).weight()) / (cluster.similarity + float64(cluster.prototype.weight()))
	candidateValue := float64(propertyVector.weight()) / (cluster.similarity + float64(cluster.prototype.Len()))
	return prototypeValue > candidateValue
}

func (cluster *Cluster) attentionTest(propertyVector *PropertyVector) bool {
	return (float64(cluster.prototype.And(propertyVector).weight()))/float64(propertyVector.weight()) >= cluster.attention
}

func (cluster Cluster) String() string {
	prototype := cluster.prototype.String()
	vectors := ""
	for _, v := range cluster.vectors {
		vectors += v.String() + "\n"
	}
	return fmt.Sprintf("[prototype]\n%v\n[vectors]\n%v\n", prototype, vectors)
}
