package workload

import (
	"log"

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
		ginkgo.It("Query secrets info with wrong cluster", func() {
			responseInfo, err := client.ListClusterSecrets("errorcluster11", tools.KpandaNameSpace)
			tools.CheckErr(err)
			if err {
				log.Fatal("No data is queried")
			}
			gomega.Expect(len(responseInfo.Payload.Secrets) == 0).To(gomega.Equal(true))
			gomega.Expect(responseInfo.Payload.Pagination.Total == 0).To(gomega.Equal(true))
		})
		ginkgo.It("Query secrets info with wrong cluster", func() {
			log.Fatal("No data is queried")
		})
	})
})
