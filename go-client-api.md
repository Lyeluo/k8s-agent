# kubernetes/client-go
github地址： https://github.com/kubernetes/client-go
## namespace 
```go
type NamespaceInterface interface {
    Create(*v1.Namespace) (*v1.Namespace, error)
    Update(*v1.Namespace) (*v1.Namespace, error)
    UpdateStatus(*v1.Namespace) (*v1.Namespace, error)
    Delete(name string, options *metav1.DeleteOptions) error
    Get(name string, options metav1.GetOptions) (*v1.Namespace, error)
    List(opts metav1.ListOptions) (*v1.NamespaceList, error)
    Watch(opts metav1.ListOptions) (watch.Interface, error)
    Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Namespace, err error)
    NamespaceExpansion
}
```
## deployment
```go
type DeploymentInterface interface {
    Create(*v1.Deployment) (*v1.Deployment, error)
    Update(*v1.Deployment) (*v1.Deployment, error)
    UpdateStatus(*v1.Deployment) (*v1.Deployment, error)
    Delete(name string, options *metav1.DeleteOptions) error
    DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
    Get(name string, options metav1.GetOptions) (*v1.Deployment, error)
    List(opts metav1.ListOptions) (*v1.DeploymentList, error)
    Watch(opts metav1.ListOptions) (watch.Interface, error)
    Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Deployment, err error)
    GetScale(deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
    UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (*autoscalingv1.Scale, error)

    DeploymentExpansion
}
```
## service
```go
type ServiceInterface interface {
    Create(*v1.Service) (*v1.Service, error)
    Update(*v1.Service) (*v1.Service, error)
    UpdateStatus(*v1.Service) (*v1.Service, error)
    Delete(name string, options *metav1.DeleteOptions) error
    Get(name string, options metav1.GetOptions) (*v1.Service, error)
    List(opts metav1.ListOptions) (*v1.ServiceList, error)
    Watch(opts metav1.ListOptions) (watch.Interface, error)
    Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Service, err error)
    ServiceExpansion
}
```
## event
```go
type EventInterface interface {
	Create(*v1.Event) (*v1.Event, error)
	Update(*v1.Event) (*v1.Event, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Event, error)
	List(opts metav1.ListOptions) (*v1.EventList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Event, err error)
	EventExpansion
}
```
## pod
```go
type PodInterface interface {
	Create(*v1.Pod) (*v1.Pod, error)
	Update(*v1.Pod) (*v1.Pod, error)
	UpdateStatus(*v1.Pod) (*v1.Pod, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Pod, error)
	List(opts metav1.ListOptions) (*v1.PodList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Pod, err error)
	GetEphemeralContainers(podName string, options metav1.GetOptions) (*v1.EphemeralContainers, error)
	UpdateEphemeralContainers(podName string, ephemeralContainers *v1.EphemeralContainers) (*v1.EphemeralContainers, error)

	PodExpansion
}
```