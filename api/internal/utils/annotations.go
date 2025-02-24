package utils

import "sigs.k8s.io/kustomize/api/konfig"

const (
	// build annotations
	BuildAnnotationPreviousKinds      = konfig.ConfigAnnoDomain + "/previousKinds"
	BuildAnnotationPreviousNames      = konfig.ConfigAnnoDomain + "/previousNames"
	BuildAnnotationPrefixes           = konfig.ConfigAnnoDomain + "/prefixes"
	BuildAnnotationSuffixes           = konfig.ConfigAnnoDomain + "/suffixes"
	BuildAnnotationPreviousNamespaces = konfig.ConfigAnnoDomain + "/previousNamespaces"
	BuildAnnotationsRefBy             = konfig.ConfigAnnoDomain + "/refBy"
	BuildAnnotationsGenBehavior       = konfig.ConfigAnnoDomain + "/generatorBehavior"
	BuildAnnotationsGenAddHashSuffix  = konfig.ConfigAnnoDomain + "/needsHashSuffix"

	// the following are only for patches, to specify whether they can change names
	// and kinds of their targets
	BuildAnnotationAllowNameChange = konfig.ConfigAnnoDomain + "/allowNameChange"
	BuildAnnotationAllowKindChange = konfig.ConfigAnnoDomain + "/allowKindChange"

	// for keeping track of origin and transformer data
	OriginAnnotationKey      = "config.kubernetes.io/origin"
	TransformerAnnotationKey = "config.kubernetes.io/transformations"

	Enabled = "enabled"
)
