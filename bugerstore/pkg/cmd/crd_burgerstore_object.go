package controller

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1"
	newClient "satya.com/burgerstore/pkg/generated/clientset/versioned"
	// newInformerFactory "satya.com/burgerstore/pkg/generated/informers/externalversions"
)

func ListCRDObject(clientset *newClient.Clientset, namespace string) ([]v1alpha1.BurgerStore, error) {
	CRDObjects, err := clientset.BurgerstoreV1alpha1().
		BurgerStores(namespace).
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error listing CRDObjects : %v\n", err)
		os.Exit(1)
	}

	// kubectl get burgerstores -n <namespace> displays 5 columns:

	// | NAME                          | OWNER | CURRENCY  | REPLICAS | AGE |
	// |-------------------------------|-------|-----------|----------|-----|
	// | object1                       | Sam   | Pound     | 1        | 14d |

	fmt.Println("Listing CRDObject objects")
	for _, CRDObject := range CRDObjects.Items {
		name := CRDObject.Name
		owner := CRDObject.Spec.Owner
		currency := CRDObject.Spec.Currency
		investment := CRDObject.Spec.Investment
		address := CRDObject.Spec.Address
		// diff := uint(time.Now().Sub(CRDObject.Status.StartTime.Time).Hours())
		// age := uint(diff / 24)

		fmt.Println(name, owner, currency, investment, address)

	}
	return CRDObjects.Items, nil
}

func GetCRDObject(clientset *newClient.Clientset, namespace string, CRDObjectname string) (*v1alpha1.BurgerStore, error) {

	// kubectl get CRDObject <CRDObject_name> -n <namespace> -o yaml

	CRDObject, err := clientset.BurgerstoreV1alpha1().
		BurgerStores(namespace).
		Get(context.TODO(), CRDObjectname, metav1.GetOptions{})

	if err != nil {
		fmt.Printf("error listing CRDObjects : %v\n", err)
		os.Exit(1)
	}

	name := CRDObject.Name
	owner := CRDObject.Spec.Owner
	currency := CRDObject.Spec.Currency
	investment := CRDObject.Spec.Investment
	address := CRDObject.Spec.Address
	// diff := uint(time.Now().Sub(CRDObject.Status.StartTime.Time).Hours())
	// age := uint(diff / 24)
	fmt.Println("Get Single CRDObject object:")
	fmt.Println(name, owner, currency, investment, address)

	return CRDObject, err
}

func CreateCRDObject(clientset *newClient.Clientset, namespace string, CRDObjectname string) (*v1alpha1.BurgerStore, error) {

	burgerStoreData := v1alpha1.BurgerStoreSpec{
		Owner:      "Satya D",
		Address:    "Delhi IN",
		Currency:   "Rupee",
		Investment: 50000,
	}

	myCRDObject := &v1alpha1.BurgerStore{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BurgerStore",
			APIVersion: "v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      CRDObjectname,
			Namespace: namespace,
		},
		Spec: burgerStoreData,
	}

	crdObj, err := clientset.BurgerstoreV1alpha1().
		BurgerStores(namespace).
		Create(context.TODO(), myCRDObject, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("error creating CRDObject : %v\n", err)
		os.Exit(1)
	}

	crdObjName := crdObj.Name
	fmt.Println("Created Single CRDObject:", crdObjName)

	return crdObj, err

}

func UpdateCRDObject(clientset *newClient.Clientset, namespace string, myCRDObject *v1alpha1.BurgerStore) (*v1alpha1.BurgerStore, error) {

	burgerStoreData := v1alpha1.BurgerStoreSpec{
		Owner:      "Satya Dillikar",
		Address:    "New York, US",
		Currency:   "Yen",
		Investment: 50000,
	}

	myCRDObject.Spec = burgerStoreData

	crdObj, err := clientset.BurgerstoreV1alpha1().
		BurgerStores(namespace).
		Update(context.TODO(), myCRDObject, metav1.UpdateOptions{})

	if err != nil {
		fmt.Printf("error updating CRDObject : %v\n", err)
		os.Exit(1)
	}

	crdObjName := crdObj.Name
	fmt.Println("updated Single CRDObject:", crdObjName)

	return crdObj, err
}

func DeleteCRDObject(clientset *newClient.Clientset, namespace string, CRDObjectname string) error {

	err := clientset.BurgerstoreV1alpha1().
		BurgerStores(namespace).
		Delete(context.TODO(), CRDObjectname, metav1.DeleteOptions{})

	if err != nil {
		fmt.Printf("error Deleting CRDObject : %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Delete Single CRDObject:", CRDObjectname)
	return err
}
