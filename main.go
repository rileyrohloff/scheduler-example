package main

import (
	"context"
	"fmt"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	// where you would read the amount of jobs to spin up per db entry or something similar
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			// idk thought this was interesting
			go createTriggerJob(client, i)
		} else {
			createTriggerJob(client, i)
		}
	}
	// debug inside of container
	time.Sleep(20 * time.Second)
}

func createTriggerJob(client *kubernetes.Clientset, jobNumber int) {
	jobs := client.BatchV1().Jobs("scheduler-example")
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("scheduler-example-trigger-job-%v", jobNumber),
			Namespace: "scheduler-example",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            fmt.Sprintf("scheduler-example-trigger-job-%v", jobNumber),
							Image:           "trigger-job:latest",
							ImagePullPolicy: v1.PullNever,
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}
	triggerJob, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cleaning up job")
	err = jobs.Delete(context.TODO(), triggerJob.Name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}
}
