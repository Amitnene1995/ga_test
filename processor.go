// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"math/rand"
        //"strconv"
	"strings"
)


// MutationRate is the rate of mutation
var MutationRate = 0.25

// PopSize is the size of the population
var PopSize = 200

// PoolSize is the max size of the pool
var PoolSize = 30

// FitnessLimit is the fitness of the evolved image we are satisfied with
//var FitnessLimit int64 = 750

var chromosome [200][14][3]int

var msreq [14]float32

var res [14]float32

var thr [14]float32

var fitness [200]float32

var index [200]int

var ms [14]int

var processorLock = &sync.Mutex{}


func initialise(){
for i:=0;i<14;i++{
	for j:=0;j<3;j++{
		l:=rand.Intn(3)
		chromosome[0][i][j]=l
}
}
fmt.Println("Randomly generated 1st chromosome : ",chromosome[0])
fmt.Println("Number of containers of 1st microservice : ",len(chromosome[0][0]))
for i:=1;i<200;i++{
    for j:=0;j<14;j++{
	k:=rand.Intn(14)
	l:=rand.Intn(3)
	chromosome[i][k][l],chromosome[i-1][k][l]=chromosome[i-1][k][l],chromosome[i][k][l]
	}
	}
}

func calcfitness(x int) float32{
var d float32
d = 0
for i:=0;i<14;i++{
    var a float32 = msreq[i]
    var b float32 = res[i]
    var c float32 = thr[i]
    d+=(a*b)/3-c
    
}
fitness[x]=d
return d
}

func naturalselection() (int,int){
for i:=0;i<200;i++{
    for j:=0;j<200;j++{
	if fitness[i]>fitness[j]{
	    fitness[i],fitness[j]=fitness[j],fitness[i]
            index[i],index[j]=index[j],index[i]
}
}
}
return index[199],index[198]
}

func crossover(i,j int){
a:=rand.Intn(3)
for x:=0;x<14;x++{
	for y:=a;y<3;y++{
		chromosome[i][x][y],chromosome[j][x][y]=chromosome[j][x][y],chromosome[i][x][y]
}
}
}

func mutate(x int){
for y:=0;y<14;y++{
	i:=rand.Intn(3)
	j:=rand.Intn(3)
	chromosome[x][y][i]=chromosome[x][y][j]
	}
}


func reconcileUnscheduledPods(interval int, done chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Second):
			err := schedulePod()
			if err != nil {
				log.Println(err)
			}
		case <-done:
			wg.Done()
			log.Println("Stopped reconciliation loop.")
			return
		}
	}
}

func monitorUnscheduledPods(done chan struct{}, wg *sync.WaitGroup) {
	/*pods, errc := watchUnscheduledPods()

	for {
		select {
		case err := <-errc:
			log.Println(err)
		case pod := <-pods:
			processorLock.Lock()
			time.Sleep(2 * time.Second)
			err := schedulePod()
			if err != nil {
				log.Println(err)
			}
			processorLock.Unlock()
		case <-done:
			wg.Done()
			log.Println("Stopped scheduler.")
			return
		}
	}*/
}

func schedulePod() error {

	initialise()
	index:=0
	msreq = [...]float32{3.2,1.8,3.2,1.4,2.3,0.8,15.1,15.1,12.0,3.2,0.1,3.2,3.2,3.2}
	res = [...]float32{0.1,11.7,20.0,0.1,27.1,2.8,3.8,0.5,0.2,41.3,45.1,26.3,4.0,13.2}
	thr = [...]float32{1.0,25.0,200.0,10.0,80.0,30.0,50.0,10.0,3.0,100.0,100.0,80.0,40.0,100.0}
	for i:=0;i<300;i++{
		
		
			a,b:=naturalselection()
			crossover(a,b)
			mutate(a)
			mutate(b)
	
	
	}
	var a float32
	a=0
	
	for i:=0;i<200;i++{
		if fitness[i]>a{
			a=fitness[i]
			index=i;	
	}




	
	podList, err := getPods()
	if err != nil {
		return err
	}

	nodes,err := getNodes()
	
	/*if len(nodes) == 0 {
		return fmt.Errorf("Unable to schedule pod (%s) failed to fit in any node", pod.Metadata.Name)
	}*/
	var str [14]string 
	str = [...]string{"worker","shipping","queue-master","payment","orders","login","front-end","edge-router","catalogue","cart","accounts","weavedb","rabbitmq","consul"}
	ms = [...]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	
	
	
	
	var nodename []Node
	
	for _, node := range nodes.Items {
		
		nodename = append(nodename,node)	
	}
		for _, p := range podList.Items {
			if p.Spec.NodeName == "" {
			continue
			}
			var i int
			if strings.HasPrefix(p.Metadata.Name,str[0]){
				i=0
			}
			if strings.HasPrefix(p.Metadata.Name,str[1]){
				i=1
			}
			if strings.HasPrefix(p.Metadata.Name,str[2]){
				i=2
			}
			if strings.HasPrefix(p.Metadata.Name,str[3]){
				i=3
			}
			if strings.HasPrefix(p.Metadata.Name,str[4]){
				i=4
			}
			if strings.HasPrefix(p.Metadata.Name,str[5]){
				i=5
			}
			if strings.HasPrefix(p.Metadata.Name,str[6]){
				i=6
			}	
			if strings.HasPrefix(p.Metadata.Name,str[7]){
				i=7
			}
			if strings.HasPrefix(p.Metadata.Name,str[8]){
				i=8
			}
			if strings.HasPrefix(p.Metadata.Name,str[9]){
				i=9
			}
			if strings.HasPrefix(p.Metadata.Name,str[10]){
				i=10
			}
			if strings.HasPrefix(p.Metadata.Name,str[11]){
				i=11
			}
			if strings.HasPrefix(p.Metadata.Name,str[12]){
				i=12
			}
			if strings.HasPrefix(p.Metadata.Name,str[13]){
				i=13
			} 
				err = bind(&p, nodename[chromosome[index][i][ms[i]]])
				if err != nil {
					return err
					ms[i]++
				}	
			}
	/*node, err := bestPrice(nodes)
	if err != nil {
		return err
	}
	err = bind(pod, node)
	if err != nil {
		return err
	}*/
}
	return nil
}


