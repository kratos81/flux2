## flux get auto image-repository

Get ImageRepository statuses

### Synopsis

The get auto image-repository command prints the status of ImageRepository objects.

```
flux get auto image-repository [flags]
```

### Examples

```
  # List all image repositories and their status
  flux get auto image-repository

 # List image repositories from all namespaces
  flux get auto image-repository --all-namespaces

```

### Options

```
  -h, --help   help for image-repository
```

### Options inherited from parent commands

```
  -A, --all-namespaces      list the requested object(s) across all namespaces
      --context string      kubernetes context to use
      --kubeconfig string   path to the kubeconfig file (default "~/.kube/config")
  -n, --namespace string    the namespace scope for this operation (default "flux-system")
      --timeout duration    timeout for this operation (default 5m0s)
      --verbose             print generated objects
```

### SEE ALSO

* [flux get auto](flux_get_auto.md)	 - Get automation statuses
