type ContainerUpdateOptions struct {
	NameOrID                        string
	RestartPolicy                   *string
	RestartRetries                  *uint
	Resources                       *spec.LinuxResources
	Env                            []string
	UnsetEnv                       []string
	ChangedHealthCheckConfiguration *define.UpdateHealthCheckConfig
	DevicesLimits                  *define.UpdateContainerDevicesLimits
	Rlimits                        []spec.POSIXRlimit
} 