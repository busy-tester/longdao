package workload

import (
	"log"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/daocloud/kpanda/test/tools"
)

var _ = ginkgo.Describe("Api e2e test - ListClusterSecrets", func() {
	client := &tools.Client{
		Client: tools.NewClientByKubeConfig(tools.KubeConfigPath),
	}

	ginkgo.Context("ListClusterSecrets - Forward testcase", func() {
		ginkgo.It("Query GlobalCluster kpanda services secrets", func() {
			responseInfo, err := client.ListClusterSecrets(tools.GlobalCluster, tools.KpandaNameSpace)
			tools.CheckErr(err, "Query secrets error")
			if len(responseInfo.Payload.Secrets) >= 1 {
				for _, listInfo := range responseInfo.Payload.Secrets {
					gomega.Expect(listInfo.Metadata.Cluster).To(gomega.Equal(tools.GlobalCluster))
					gomega.Expect(listInfo.Metadata.Namespace).To(gomega.Equal(tools.KpandaNameSpace))
				}
			} else {
				log.Fatal("No data is queried")
			}
		})
	})

	ginkgo.Context("ListClusterSecrets - Reverse testcase", func() {
		ginkgo.It("Query secrets info with wrong NameSpace name", func() {
			responseInfo, err := client.ListClusterSecrets(tools.GlobalCluster, "errorNameSpace")
			tools.CheckErr(err)
			gomega.Expect(len(responseInfo.Payload.Secrets) == 0).To(gomega.Equal(true))
			gomega.Expect(responseInfo.Payload.Pagination.Total == 0).To(gomega.Equal(true))
		})
		ginkgo.It("Query secrets info with wrong", func() {
			responseInfo, err := client.ListClusterSecrets("errorcluster", tools.KpandaNameSpace)
			log.Fatal("No data is queried")
			gomega.Expect(len(responseInfo.Payload.Secrets) == 12).To(gomega.Equal(true))
			gomega.Expect(responseInfo.Payload.Pagination.Total == 0).To(gomega.Equal(true))
		})
		for i:=1;i<5;i++{
			log.Fatal("No data is queried")
			log.Fatal("No data is queried")
		}
	})
})
