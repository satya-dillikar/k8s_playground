package dyn_client

import (
	"fmt"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

type contrller struct {
	client   dynamic.Interface
	informer cache.SharedIndexInformer
	// workqueue workqueue.RateLimitingInterface
}

func newController(client dynamic.Interface, dynInformer dynamicinformer.DynamicSharedInformerFactory) *contrller {
	inf := dynInformer.ForResource(schema.GroupVersionResource{
		Group:    "burgerstore.dev",
		Version:  "v1alpha1",
		Resource: "burgerstore",
	}).Informer()

	inf.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    handleAdd,
			UpdateFunc: handleUpdate,
			DeleteFunc: handleDelete,
		},
	)

	return &contrller{
		client:   client,
		informer: inf,
		// workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "BurgerStores"),
	}

}

func handleAdd(obj interface{}) {
	log.Println("handleAdd was called")
}
func handleDelete(obj interface{}) {
	log.Println("handleDelete was called")
}

func handleUpdate(oldObj interface{}, newObj interface{}) {
	log.Println("handleUpdate was called")
}

func (c *contrller) run(ch <-chan struct{}) {
	// defer c.workqueue.ShutDown()
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.informer.HasSynced) {
		fmt.Print("waiting for cache to be synced\n")
	}

	go wait.Until(c.worker, 1*time.Second, ch)

	<-ch
}

func (c *contrller) worker() {
	for c.processItem() {

	}
}

func (c *contrller) processItem() bool {
	return true
}

// // processNextWorkItem will read a single work item off the workqueue and
// // attempt to process it, by calling the syncHandler.
// func (c *contrller) processItem() bool {
// 	obj, shutdown := c.workqueue.Get()

// 	if shutdown {
// 		return false
// 	}

// 	// We wrap this block in a func so we can defer c.workqueue.Done.
// 	err := func(obj interface{}) error {
// 		defer c.workqueue.Done(obj)
// 		var key string
// 		var ok bool
// 		if key, ok = obj.(string); !ok {
// 			c.workqueue.Forget(obj)
// 			return nil
// 		}
// 		// if err := c.syncHandler(key); err != nil {
// 		// 	c.workqueue.AddRateLimited(key)
// 		// 	return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
// 		// }
// 		c.workqueue.Forget(obj)
// 		klog.Infof("Successfully synced '%s'", key)
// 		return nil
// 	}(obj)

// 	if err != nil {
// 		utilruntime.HandleError(err)
// 		return true
// 	}

// 	return true
// }
