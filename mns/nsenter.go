/*
This utility file manages the namespace of the container.

NAMESPACES ACCORDING TO THE OCI SPEC:
The namespace elements of a container according to the OCI spec is listed here:
	https://github.com/opencontainers/runtime-spec/blob/main/schema/config-linux.json#L30
The types for that value is listed here:
	https://github.com/opencontainers/runtime-spec/blob/main/schema/defs-linux.json#L301-L313


GO TYPES:
The types used in this package will be that of the OCI spec Go package.
*/

package mns

import "github.com/opencontainers/runtime-spec/specs-go"

// enter a namespace
func nsEnter(namespace specs.LinuxNamespace) error {

	return nil
}
