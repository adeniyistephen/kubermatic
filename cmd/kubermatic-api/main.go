package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kubermatic/api/handler"
	"github.com/kubermatic/api/provider"
	"github.com/kubermatic/api/provider/cloud"
	"github.com/kubermatic/api/provider/kubernetes"
	"golang.org/x/net/context"

	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
)

func main() {
	// parse flags
	flag.Parse()

	// create CloudProviders
	cps := map[string]provider.CloudProvider{
		provider.FakeCloudProvider:         cloud.NewFakeCloudProvider(),
		provider.DigitaloceanCloudProvider: nil,
		// provider.LinodeCloudProvider: nil,
	}

	// create KubernetesProvider for each context in the kubeconfig
	kps := map[string]provider.KubernetesProvider{
		"fake-1": kubernetes.NewKubernetesFakeProvider("fake-1", cps),
		"fake-2": kubernetes.NewKubernetesFakeProvider("fake-2", cps),
	}
	clientcmdConfig, err := clientcmd.LoadFromFile(".kubeconfig")
	if err != nil {
		log.Fatal(err)
	}
	for ctx := range clientcmdConfig.Contexts {
		clientconfig := clientcmd.NewNonInteractiveClientConfig(
			*clientcmdConfig,
			ctx,
			&clientcmd.ConfigOverrides{},
		)
		cfg, err := clientconfig.ClientConfig()
		if err != nil {
			log.Fatal(err)
		}

		kps[ctx] = kubernetes.NewKubernetesProvider(cfg, cps)
	}

	// start server
	ctx := context.Background()
	mux := mux.NewRouter()

	mux.
		Methods("GET").
		Path("/").
		HandlerFunc(handler.StatusOK)

	mux.
		Methods("POST").
		Path("/api/v1/dc/{dc}/cluster/{cluster}").
		Handler(handler.NewCluster(ctx, kps, cps))

	mux.
		Methods("GET").
		Path("/api/v1/dc/{dc}/cluster").
		Handler(handler.Clusters(ctx, kps, cps))

	mux.
		Methods("GET").
		Path("/api/v1/dc/{dc}/cluster/{cluster}").
		Handler(handler.Cluster(ctx, kps, cps))

	mux.
		Methods("GET").
		Path("/api/v1/dc/{dc}/cluster/{cluster}/node").
		Handler(handler.Nodes(ctx, kps, cps))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(":8080", ghandlers.CombinedLoggingHandler(os.Stdout, mux)))
}
