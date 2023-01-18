


## 상수 
```
const (
    kubeAPIServerManifestFileName = "kube-apiserver.manifest"
    kubeAPIServerConfigScriptName = "configure-kubeapiserver.sh"
    kubeAPIServerStartFuncName    = "start-kube-apiserver"
)
```
```
const etcdTimeout = 2 * time.Second
```



## If
```
if pvSpec.NodeAffinity != nil && pvSpec.NodeAffinity.Required != nil {
```

https://kimjingo.tistory.com/144

```
if rc != nil && rc.Scheduling != nil && rc.Scheduling.NodeSelector != nil {
        // use of deprecated node labels in scheduling's node affinity
        for key := range rc.Scheduling.NodeSelector {
            if msg, deprecated := GetNodeLabelDeprecatedMessage(key); deprecated {
                warnings = append(warnings, fmt.Sprintf("%s: %s", field.NewPath("scheduling", "nodeSelector"), msg))
            }
        }
    }
```


```
// compare values
    if !e.InAddr.Equal(o.InAddr) || !e.In6Addr.Equal(o.In6Addr) {
        return false
    }
```

## 중첩 if
```
case source.Projected != nil:
            for j := range source.Projected.Sources {
                if source.Projected.Sources[j].Secret != nil {
                    if !visitor(source.Projected.Sources[j].Secret.Name) {
                        return false
                    }
                }
            }
```

## Switch
```
// IsExternallyManaged checks if we are in the external CA case (CA certificate provided without the certificate key)
func (rm *Manager) IsExternallyManaged(caBaseName string) (bool, error) {
    switch caBaseName {
    case kubeadmconstants.CACertAndKeyBaseName: //CACertAndKeyBaseName = "ca"

        externallyManaged, err := certsphase.UsingExternalCA(rm.cfg)
        if err != nil {
            return false, errors.Wrapf(err, "Error checking external CA condition for %s certificate authority", caBaseName)
        }
        return externallyManaged, nil
    case kubeadmconstants.FrontProxyCACertAndKeyBaseName: //FrontProxyCACertAndKeyBaseName = "front-proxy-ca"

        externallyManaged, err := certsphase.UsingExternalFrontProxyCA(rm.cfg)
        if err != nil {
            return false, errors.Wrapf(err, "Error checking external CA condition for %s certificate authority", caBaseName)
        }
        return externallyManaged, nil
    case kubeadmconstants.EtcdCACertAndKeyBaseName:
        externallyManaged, err := certsphase.UsingExternalEtcdCA(rm.cfg)
        if err != nil {
            return false, errors.Wrapf(err, "Error checking external CA condition for %s certificate authority", caBaseName)
        }
        return externallyManaged, nil
    default:
        return false, errors.Errorf("unknown certificate authority %s", caBaseName)
    }
}
```
```
/ TryLoadKeyFromDisk tries to load the key from the disk and validates that it is valid
func TryLoadKeyFromDisk(pkiPath, name string) (crypto.Signer, error) {
    privateKeyPath := pathForKey(pkiPath, name)

    // Parse the private key from a file
    privKey, err := keyutil.PrivateKeyFromFile(privateKeyPath)
    if err != nil {
        return nil, errors.Wrapf(err, "couldn't load the private key file %s", privateKeyPath)
    }

    // Allow RSA and ECDSA formats only
    var key crypto.Signer
    switch k := privKey.(type) {
    case *rsa.PrivateKey:
        key = k
    case *ecdsa.PrivateKey:
        key = k
    default:
        return nil, errors.Errorf("the private key file %s is neither in RSA nor ECDSA format", privateKeyPath)
    }

    return key, nil
}
```

## Type 

```
// ContainerType signifies container type
type ContainerType int

const (
    // Containers is for normal containers
    Containers ContainerType = 1 << iota
    // InitContainers is for init containers
    InitContainers
    // EphemeralContainers is for ephemeral containers
    EphemeralContainers
)
```



## Fall through
```
func (s *Service) ServeDatastore(w http.ResponseWriter, r *http.Request) {
    ds, ferr := s.findDatastore(r.URL.Query())
    if ferr != nil {
        log.Printf("failed to locate datastore with query params: %s", r.URL.RawQuery)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    r.URL.Path = strings.TrimPrefix(r.URL.Path, folderPrefix)
    p := path.Join(ds.Info.GetDatastoreInfo().Url, r.URL.Path)

    switch r.Method {
    case http.MethodPost:
        _, err := os.Stat(p)
        if err == nil {
            // File exists
            w.WriteHeader(http.StatusConflict)
            return
        }

        // File does not exist, fallthrough to create via PUT logic
        fallthrough
    case http.MethodPut:
        dir := path.Dir(p)
        _ = os.MkdirAll(dir, 0700)

        f, err := os.Create(p)
        if err != nil {
            log.Printf("failed to %s '%s': %s", r.Method, p, err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        defer f.Close()

        _, _ = io.Copy(f, r.Body)
    default:
        fs := http.FileServer(http.Dir(ds.Info.GetDatastoreInfo().Url))

        fs.ServeHTTP(w, r)
    }
}
```
## For

```
// If Node Allocatable is enforced on a node that has not been drained or is updated on an existing node to a lower value,
    // existing memory usage across pods might be higher than current Node Allocatable Memory Limits.
    // Pod Evictions are expected to bring down memory usage to below Node Allocatable limits.
    // Until evictions happen retry cgroup updates.
    // Update limits on non root cgroup-root to be safe since the default limits for CPU can be too low.
    // Check if cgroupRoot is set to a non-empty value (empty would be the root container)
    if len(cm.cgroupRoot) > 0 {
        go func() {
            for {
                err := cm.cgroupManager.Update(cgroupConfig)
                if err == nil {
                    cm.recorder.Event(nodeRef, v1.EventTypeNormal, events.SuccessfulNodeAllocatableEnforcement, "Updated Node Allocatable limit across pods")
                    return
                }
                message := fmt.Sprintf("Failed to update Node Allocatable Limits %q: %v", cm.cgroupRoot, err)
                cm.recorder.Event(nodeRef, v1.EventTypeWarning, events.FailedNodeAllocatableEnforcement, message)
                time.Sleep(time.Minute)
            }
        }()
    }
```
## 배열 range
```
CA : 트러스트 앵커이는 원래 인증 기관 (CA)입니다.
CERTS 인증서 : https://d1smxttentwwqu.cloudfront.net/wp-content/uploads/2019/07/ca-diagram-b.png

certTree, err := certListFunc().AsMap().CertTree()
    if err != nil {
        return nil, err
    }

    // create a CertificateRenewHandler for each signed certificate in the certificate tree;
    // NB. we are not offering support for renewing CAs; this would cause serious consequences
    for ca, certs := range certTree {
        for _, cert := range certs {
            // create a ReadWriter for certificates stored in the K8s local PKI
            pkiReadWriter := newPKICertificateReadWriter(rm.cfg.CertificatesDir, cert.BaseName)

            // adds the certificateRenewHandler.
            // PKI certificates are indexed by name, that is a well know constant defined
            // in the certsphase package and that can be reused across all the kubeadm codebase
            rm.certificates[cert.Name] = &CertificateRenewHandler{
                Name:       cert.Name,
                LongName:   cert.LongName,
                FileName:   cert.BaseName,
                CAName:     ca.Name,
                CABaseName: ca.BaseName, //Nb. this is a path for etcd certs (they are stored in a subfolder)
                readwriter: pkiReadWriter,
            }
        }

        pkiReadWriter := newPKICertificateReadWriter(rm.cfg.CertificatesDir, ca.BaseName)
        rm.cas[ca.Name] = &CAExpirationHandler{
            Name:       ca.Name,
            LongName:   ca.LongName,
            FileName:   ca.BaseName,
            readwriter: pkiReadWriter,
        }
    }

```
```
사실 tree —> map [k]v
// CertificateTree is represents a one-level-deep tree, mapping a CA to the certs that depend on it.
type CertificateTree map[*KubeadmCert]Certificates
```

## 구조체 
```
// DeploymentSpec specifies the state of a Deployment.
type DeploymentSpec struct {
    // Number of desired pods.
    Replicas int32

    // Label selector for pods. Existing ReplicaSets whose pods are
    // selected by this will be the ones affected by this deployment.
    // +optional
    Selector *metav1.LabelSelector

    // Template describes the pods that will be created.
    Template api.PodTemplateSpec

    // The deployment strategy to use to replace existing pods with new ones.
    // +optional
    Strategy DeploymentStrategy

    // Minimum number of seconds for which a newly created pod should be ready
    // without any of its container crashing, for it to be considered available.
    // Defaults to 0 (pod will be considered available as soon as it is ready)
    // +optional
    MinReadySeconds int32

    // The number of old ReplicaSets to retain to allow rollback.
    // This is a pointer to distinguish between explicit zero and not specified.
    // This is set to the max value of int32 (i.e. 2147483647) by default, which means
    // "retaining all old ReplicaSets".
    // +optional
    RevisionHistoryLimit *int32

    // Indicates that the deployment is paused and will not be processed by the
    // deployment controller.
    // +optional
    Paused bool

    // DEPRECATED.
    // The config this deployment is rolling back to. Will be cleared after rollback is done.
    // +optional
    RollbackTo *RollbackConfig

    // The maximum time in seconds for a deployment to make progress before it
    // is considered to be failed. The deployment controller will continue to
    // process failed deployments and a condition with a ProgressDeadlineExceeded
    // reason will be surfaced in the deployment status. Note that progress will
    // not be estimated during the time a deployment is paused. This is set to
    // the max value of int32 (i.e. 2147483647) by default, which means "no deadline".
    // +optional
    ProgressDeadlineSeconds *int32
}
```

### deploymentspec.yaml 예제 

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: web
spec:
  selector:
    matchLabels:
      app: web
  replicas: 5
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
       —name: nginx
          image: nginx
          ports:
           —containerPort: 80
```

