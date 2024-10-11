// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.openshift.api.console.v1.ApplicationMenuSpec
  map:
    fields:
    - name: imageURL
      type:
        scalar: string
    - name: section
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.CLIDownloadLink
  map:
    fields:
    - name: href
      type:
        scalar: string
      default: ""
    - name: text
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleCLIDownload
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleCLIDownloadSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleCLIDownloadSpec
  map:
    fields:
    - name: description
      type:
        scalar: string
      default: ""
    - name: displayName
      type:
        scalar: string
      default: ""
    - name: links
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.console.v1.CLIDownloadLink
          elementRelationship: atomic
- name: com.github.openshift.api.console.v1.ConsoleExternalLogLink
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleExternalLogLinkSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleExternalLogLinkSpec
  map:
    fields:
    - name: hrefTemplate
      type:
        scalar: string
      default: ""
    - name: namespaceFilter
      type:
        scalar: string
    - name: text
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleLink
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleLinkSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleLinkSpec
  map:
    fields:
    - name: applicationMenu
      type:
        namedType: com.github.openshift.api.console.v1.ApplicationMenuSpec
    - name: href
      type:
        scalar: string
      default: ""
    - name: location
      type:
        scalar: string
      default: ""
    - name: namespaceDashboard
      type:
        namedType: com.github.openshift.api.console.v1.NamespaceDashboardSpec
    - name: text
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleNotification
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleNotificationSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleNotificationSpec
  map:
    fields:
    - name: backgroundColor
      type:
        scalar: string
    - name: color
      type:
        scalar: string
    - name: link
      type:
        namedType: com.github.openshift.api.console.v1.Link
    - name: location
      type:
        scalar: string
    - name: text
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsolePlugin
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsolePluginBackend
  map:
    fields:
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginService
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: service
        discriminatorValue: Service
- name: com.github.openshift.api.console.v1.ConsolePluginI18n
  map:
    fields:
    - name: loadType
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsolePluginProxy
  map:
    fields:
    - name: alias
      type:
        scalar: string
      default: ""
    - name: authorization
      type:
        scalar: string
    - name: caCertificate
      type:
        scalar: string
    - name: endpoint
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginProxyEndpoint
      default: {}
- name: com.github.openshift.api.console.v1.ConsolePluginProxyEndpoint
  map:
    fields:
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginProxyServiceConfig
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: service
        discriminatorValue: Service
- name: com.github.openshift.api.console.v1.ConsolePluginProxyServiceConfig
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: port
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.console.v1.ConsolePluginService
  map:
    fields:
    - name: basePath
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: port
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.console.v1.ConsolePluginSpec
  map:
    fields:
    - name: backend
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginBackend
      default: {}
    - name: displayName
      type:
        scalar: string
      default: ""
    - name: i18n
      type:
        namedType: com.github.openshift.api.console.v1.ConsolePluginI18n
      default: {}
    - name: proxy
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.console.v1.ConsolePluginProxy
          elementRelationship: atomic
- name: com.github.openshift.api.console.v1.ConsoleQuickStart
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleQuickStartSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleQuickStartSpec
  map:
    fields:
    - name: accessReviewResources
      type:
        list:
          elementType:
            namedType: io.k8s.api.authorization.v1.ResourceAttributes
          elementRelationship: atomic
    - name: conclusion
      type:
        scalar: string
    - name: description
      type:
        scalar: string
      default: ""
    - name: displayName
      type:
        scalar: string
      default: ""
    - name: durationMinutes
      type:
        scalar: numeric
      default: 0
    - name: icon
      type:
        scalar: string
    - name: introduction
      type:
        scalar: string
      default: ""
    - name: nextQuickStart
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: prerequisites
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tasks
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.console.v1.ConsoleQuickStartTask
          elementRelationship: atomic
- name: com.github.openshift.api.console.v1.ConsoleQuickStartTask
  map:
    fields:
    - name: description
      type:
        scalar: string
      default: ""
    - name: review
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleQuickStartTaskReview
    - name: summary
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleQuickStartTaskSummary
    - name: title
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleQuickStartTaskReview
  map:
    fields:
    - name: failedTaskHelp
      type:
        scalar: string
      default: ""
    - name: instructions
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleQuickStartTaskSummary
  map:
    fields:
    - name: failed
      type:
        scalar: string
      default: ""
    - name: success
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleSample
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleSampleContainerImportSource
  map:
    fields:
    - name: image
      type:
        scalar: string
      default: ""
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleContainerImportSourceService
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleSampleContainerImportSourceService
  map:
    fields:
    - name: targetPort
      type:
        scalar: numeric
- name: com.github.openshift.api.console.v1.ConsoleSampleGitImportSource
  map:
    fields:
    - name: repository
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleGitImportSourceRepository
      default: {}
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleGitImportSourceService
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleSampleGitImportSourceRepository
  map:
    fields:
    - name: contextDir
      type:
        scalar: string
      default: ""
    - name: revision
      type:
        scalar: string
      default: ""
    - name: url
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleSampleGitImportSourceService
  map:
    fields:
    - name: targetPort
      type:
        scalar: numeric
- name: com.github.openshift.api.console.v1.ConsoleSampleSource
  map:
    fields:
    - name: containerImport
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleContainerImportSource
    - name: gitImport
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleGitImportSource
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: containerImport
        discriminatorValue: ContainerImport
      - fieldName: gitImport
        discriminatorValue: GitImport
- name: com.github.openshift.api.console.v1.ConsoleSampleSpec
  map:
    fields:
    - name: abstract
      type:
        scalar: string
      default: ""
    - name: description
      type:
        scalar: string
      default: ""
    - name: icon
      type:
        scalar: string
      default: ""
    - name: provider
      type:
        scalar: string
      default: ""
    - name: source
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleSampleSource
      default: {}
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: title
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.ConsoleYAMLSample
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1.ConsoleYAMLSampleSpec
      default: {}
- name: com.github.openshift.api.console.v1.ConsoleYAMLSampleSpec
  map:
    fields:
    - name: description
      type:
        scalar: string
      default: ""
    - name: snippet
      type:
        scalar: boolean
      default: false
    - name: targetResource
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.TypeMeta
      default: {}
    - name: title
      type:
        scalar: string
      default: ""
    - name: yaml
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.Link
  map:
    fields:
    - name: href
      type:
        scalar: string
      default: ""
    - name: text
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1.NamespaceDashboardSpec
  map:
    fields:
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.openshift.api.console.v1alpha1.ConsolePlugin
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.console.v1alpha1.ConsolePluginSpec
      default: {}
- name: com.github.openshift.api.console.v1alpha1.ConsolePluginProxy
  map:
    fields:
    - name: alias
      type:
        scalar: string
      default: ""
    - name: authorize
      type:
        scalar: boolean
    - name: caCertificate
      type:
        scalar: string
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1alpha1.ConsolePluginProxyServiceConfig
      default: {}
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.console.v1alpha1.ConsolePluginProxyServiceConfig
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: port
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.console.v1alpha1.ConsolePluginService
  map:
    fields:
    - name: basePath
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: port
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.console.v1alpha1.ConsolePluginSpec
  map:
    fields:
    - name: displayName
      type:
        scalar: string
    - name: proxy
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.console.v1alpha1.ConsolePluginProxy
          elementRelationship: atomic
    - name: service
      type:
        namedType: com.github.openshift.api.console.v1alpha1.ConsolePluginService
      default: {}
- name: io.k8s.api.authorization.v1.ResourceAttributes
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resource
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: verb
      type:
        scalar: string
    - name: version
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: io.k8s.apimachinery.pkg.apis.meta.v1.TypeMeta
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
