package utils

import (
	"fmt"
	"testing"

	"github.com/go-gota/gota/dataframe"
)

type UsageData struct {
	Report_period_start            string `dataframe:"report_period_start,string"`
	Report_period_end              string `dataframe:"report_period_end,string"`
	Interval_start                 string `dataframe:"interval_start,string"`
	Interval_end                   string `dataframe:"interval_end,string"`
	Container_name                 string `dataframe:"container_name,string"`
	Pod                            string `dataframe:"pod,string"`
	Owner_name                     string `dataframe:"owner_name,string"`
	Owner_kind                     string `dataframe:"owner_kind,string"`
	Workload                       string `dataframe:"workload,string"`
	Workload_type                  string `dataframe:"workload_type,string"`
	Namespace                      string `dataframe:"namespace,string"`
	Image_name                     string `dataframe:"image_name,string"`
	Node                           string `dataframe:"node,string"`
	Resource_id                    string `dataframe:"resource_id,string"`
	Cpu_request_container_avg      string `dataframe:"cpu_request_container_avg,float"`
	Cpu_request_container_sum      string `dataframe:"cpu_request_container_sum,float"`
	Cpu_limit_container_avg        string `dataframe:"cpu_limit_container_avg,float"`
	Cpu_limit_container_sum        string `dataframe:"cpu_limit_container_sum,float"`
	Cpu_usage_container_avg        string `dataframe:"cpu_usage_container_avg,float"`
	Cpu_usage_container_min        string `dataframe:"cpu_usage_container_min,float"`
	Cpu_usage_container_max        string `dataframe:"cpu_usage_container_max,float"`
	Cpu_usage_container_sum        string `dataframe:"cpu_usage_container_sum,float"`
	Cpu_throttle_container_avg     string `dataframe:"cpu_throttle_container_avg,float"`
	Cpu_throttle_container_max     string `dataframe:"cpu_throttle_container_max,float"`
	Cpu_throttle_container_sum     string `dataframe:"cpu_throttle_container_sum,float"`
	Memory_request_container_avg   string `dataframe:"memory_request_container_avg,float"`
	Memory_request_container_sum   string `dataframe:"memory_request_container_sum,float"`
	Memory_limit_container_avg     string `dataframe:"memory_limit_container_avg,float"`
	Memory_limit_container_sum     string `dataframe:"memory_limit_container_sum,float"`
	Memory_usage_container_avg     string `dataframe:"memory_usage_container_avg,float"`
	Memory_usage_container_min     string `dataframe:"memory_usage_container_min,float"`
	Memory_usage_container_max     string `dataframe:"memory_usage_container_max,float"`
	Memory_usage_container_sum     string `dataframe:"memory_usage_container_sum,float"`
	Memory_rss_usage_container_avg string `dataframe:"memory_rss_usage_container_avg,float"`
	Memory_rss_usage_container_min string `dataframe:"memory_rss_usage_container_min,float"`
	Memory_rss_usage_container_max string `dataframe:"memory_rss_usage_container_max,float"`
	Memory_rss_usage_container_sum string `dataframe:"memory_rss_usage_container_sum,float"`
}

func Test_filter_valid_k8s_object_types(t *testing.T) {
	// Check valid k8s object type
	usage_data := []UsageData{
		// k8s object type DaemonSet
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "DaemonSet", "testdeploymentconfig", "daemonset", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
		// k8s object type Replicaset
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicaSet", "<none>", "deployment", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
		// k8s object type Deployment
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicaSet", "testdeployment", "deployment", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
		// k8s object type ReplicationController
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicationController", "<none>", "deploymentconfig", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
		// k8s object type Deploymentconfig
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicationController", "testdeploymentconfig", "deploymentconfig", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
		// k8s object type StatefulSet
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "StatefulSet", "testdeploymentconfig", "statefulset", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
	}
	df := dataframe.LoadStructs(usage_data)
	df = determine_k8s_object_type(df)
	result := filter_valid_k8s_object_types(df)
	fmt.Println(result.Nrow())
	if result.Nrow() != 6 {
		t.Error("Data not filtered properly. Some of the valid k8s object type got dropped")
	}

	// check if Invalid k8s object type is dropped
	usage_data = []UsageData{
		// k8s object type Job
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "Job", "testdeploymentconfig", "job", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
		},
	}
	df = dataframe.LoadStructs(usage_data)
	df = determine_k8s_object_type(df)
	result = filter_valid_k8s_object_types(df)
	if result.Nrow() != 0 {
		t.Error("Invalid k8s object type did not get dropped")
	}
}

func Test_filter_valid_csv_records(t *testing.T) {
	usage_data := []UsageData{
		// k8s object with missing data
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicaSet", "testdeployment", "deployment", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "",
		},
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicaSet", "testdeployment", "deployment", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "", "", "", "", "1", "1", "1", "1", "1", "1", "1", "", "", "", "", "", "", "", "",
		},
		// k8s object with 0 CPU, Memory and RSS usage
		{
			"2023-02-01 00:00:00 +0000 UTC", "2023-03-01 00:00:00 +0000 UTC", "2023-06-02 00:00:01 +0000 UTC", "2023-06-02 00:15:00 +0000 UTC",
			"Yuptoo-service", "Yuptoo-app-standalone-1", "Yuptoo-app", "ReplicaSet", "testdeployment", "deployment", "Yuptoo-prod",
			"quay.io/cloudservices/yuptoo", "ip-10-0-176-227.us-east-2.compute.internal", "i-0dfbb3fa4d0e8fc94",
			"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0",
		},
	}
	df := dataframe.LoadStructs(usage_data)
	df = determine_k8s_object_type(df)
	result, no_of_dropped_records := filter_valid_csv_records(df)
	if result.Nrow() != 1 || no_of_dropped_records != 2 {
		t.Error("Invalid k8s object type did not get dropped")
	}

}
