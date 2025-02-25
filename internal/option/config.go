// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package option

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the main configuration as obtained from command-line arguments,
// environment variables and config files.
var Config = &CertGenConfig{}

const (
	// Debug enables debug messages.
	Debug = "debug"

	// CiliumNamespace is the Kubernetes namespace in which Cilium is
	// installed.
	CiliumNamespace = "cilium-namespace"

	// K8sKubeConfigPath is the path to the kubeconfig If empty, the in-cluster
	// configuration is used.
	K8sKubeConfigPath = "k8s-kubeconfig-path"
	// K8sRequestTimeout specifies the timeout for K8s API requests.
	K8sRequestTimeout = "k8s-request-timeout"

	// CACertFile is the path to the Cilium CA cert PEM (if CAGenerate is
	// false).
	CACertFile = "ca-cert-file"
	// CAKeyFile is the path to the Cilium CA key PEM (if CAGenerate is false).
	CAKeyFile = "ca-key-file"

	// CAGenerate can be set to true to generate a new Cilium CA secret.
	// If CAReuseSecret is true, then a new CA secret only is created if
	// existing one is not found.
	CAGenerate = "ca-generate"
	// CAReuseSecret can be set to true to store and load the Cilium CA from
	// the secret if it exists. Setting to false will delete the old Secret and
	// force regeneration.
	CAReuseSecret = "ca-reuse-secret" //#nosec
	// CACommonName is the Cilium CA x509 certificate CN value.
	CACommonName = "ca-common-name"
	// CAValidityDuration represent how much time the Cilium CA certificate
	// generated by certgen is valid.
	CAValidityDuration = "ca-validity-duration"
	// CASecretName is the Kubernetes Secret in which the Cilium CA certificate
	// is read from and/or written to.
	CASecretName = "ca-secret-name"
	// CASecretNamespace is the Kubernetes Namespace in which the Cilium CA
	// Secret will be stored.
	CASecretNamespace = "ca-secret-namespace"

	// HubbleServerCertGenerate can be set to true to generate and store a
	// Hubble server TLS certificate.
	HubbleServerCertGenerate = "hubble-server-cert-generate"
	// HubbleServerCertCommonName is the Hubble server x509 certificate CN
	// value (also used as DNS SAN).
	HubbleServerCertCommonName = "hubble-server-cert-common-name"
	// HubbleServerCertValidityDuration represent how much time the Hubble
	// server certificate generated by certgen is valid.
	HubbleServerCertValidityDuration = "hubble-server-cert-validity-duration"
	// HubbleServerCertSecretName is the Kubernetes Secret in which the Hubble
	// server certificate is written to.
	HubbleServerCertSecretName = "hubble-server-cert-secret-name" //#nosec
	// HubbleServerCertSecretNamespace is the Kubernetes Namespace in which the
	// Hubble server certificate Secret will be stored.
	HubbleServerCertSecretNamespace = "hubble-server-cert-secret-namespace" //#nosec

	// HubbleMetricsServerCertGenerate can be set to true to generate and store a
	// Hubble metrics server TLS certificate.
	HubbleMetricsServerCertGenerate = "hubble-metrics-server-cert-generate"
	// HubbleMetricsServerCertCommonName is the Hubble metrics server x509 certificate CN
	// value (also used as DNS SAN).
	HubbleMetricsServerCertCommonName = "hubble-metrics-server-cert-common-name"
	// HubbleMetricsServerCertValidityDuration represent how much time the Hubble
	// server certificate generated by certgen is valid.
	HubbleMetricsServerCertValidityDuration = "hubble-metrics-server-cert-validity-duration"
	// HubbleMetricsServerCertSecretName is the Kubernetes Secret in which the Hubble
	// server certificate is written to.
	HubbleMetricsServerCertSecretName = "hubble-metrics-server-cert-secret-name" //#nosec
	// HubbleMetricsServerCertSecretNamespace is the Kubernetes Namespace in which the
	// Hubble metrics server certificate Secret will be stored.
	HubbleMetricsServerCertSecretNamespace = "hubble-metrics-server-cert-secret-namespace" //#nosec

	// HubbleRelayServerCertGenerate can be set to true to generate and store a
	// Hubble Relay server TLS certificate.
	HubbleRelayServerCertGenerate = "hubble-relay-server-cert-generate"
	// HubbleRelayServerCertCommonName is the Hubble Relay server x509
	// certificate CN value (also used as DNS SAN).
	HubbleRelayServerCertCommonName = "hubble-relay-server-cert-common-name"
	// HubbleRelayServerCertValidityDuration represent how much time the Hubble
	// Relay server certificate generated by certgen is valid.
	HubbleRelayServerCertValidityDuration = "hubble-relay-server-cert-validity-duration"
	// HubbleRelayServerCertSecretName is the Kubernetes Secret in which the
	// Hubble Relay server certificate is written to.
	HubbleRelayServerCertSecretName = "hubble-relay-server-cert-secret-name" //#nosec
	// HubbleRelayServerCertSecretNamespace is the Kubernetes Namespace in
	// which the Hubble Relay server certificate Secret will be stored.
	HubbleRelayServerCertSecretNamespace = "hubble-relay-server-cert-secret-namespace" //#nosec

	// HubbleRelayClientCertGenerate can be set to true to generate and store a
	// Hubble Relay client TLS certificate (used for the mTLS handshake with
	// the Hubble servers).
	HubbleRelayClientCertGenerate = "hubble-relay-client-cert-generate"
	// HubbleRelayClientCertCommonName is the Hubble Relay client x509
	// certificate CN value.
	HubbleRelayClientCertCommonName = "hubble-relay-client-cert-common-name"
	// HubbleRelayClientCertValidityDuration represent how much time the Hubble
	// Relay client certificate generated by certgen is valid.
	HubbleRelayClientCertValidityDuration = "hubble-relay-client-cert-validity-duration"
	// HubbleRelayClientCertSecretName is the Kubernetes Secret in which the
	// Hubble Relay client certificate is written to.
	HubbleRelayClientCertSecretName = "hubble-relay-client-cert-secret-name" //#nosec
	// HubbleRelayClientCertSecretNamespace is the Kubernetes Namespace in
	// which the Hubble Relay client certificate Secret will be stored.
	HubbleRelayClientCertSecretNamespace = "hubble-relay-client-cert-secret-namespace" //#nosec

	// ClustermeshApiserverServerCertGenerate can be set to true to generate
	// and store a new Clustermesh API server TLS certificate.
	ClustermeshApiserverServerCertGenerate = "clustermesh-apiserver-server-cert-generate"
	// ClustermeshApiserverServerCertCommonName is the Clustermesh API server
	// x509 certificate CN value (also used as DNS SAN).
	ClustermeshApiserverServerCertCommonName = "clustermesh-apiserver-server-cert-common-name"
	// ClustermeshApiserverServerCertValidityDuration represent how much time
	// Clustermesh API server certificate generated by certgen is valid.
	ClustermeshApiserverServerCertValidityDuration = "clustermesh-apiserver-server-cert-validity-duration"
	// ClustermeshApiserverServerCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API server certificate is written to.
	ClustermeshApiserverServerCertSecretName = "clustermesh-apiserver-server-cert-secret-name"
	// ClustermeshApiserverServerCertSANs is the list of SANs to add to the
	// Clustermesh API server certificate.
	ClustermeshApiserverServerCertSANs = "clustermesh-apiserver-server-cert-sans"

	// ClustermeshApiserverAdminCertGenerate can be set to true to generate and
	// store a new Clustermesh API admin TLS certificate.
	ClustermeshApiserverAdminCertGenerate = "clustermesh-apiserver-admin-cert-generate"
	// ClustermeshApiserverAdminCertCommonName is the Clustermesh API admin
	// x509 certificate CN value.
	ClustermeshApiserverAdminCertCommonName = "clustermesh-apiserver-admin-cert-common-name"
	// ClustermeshApiserverAdminCertValidityDuration represent how much time
	// Clustermesh API admin certificate generated by certgen is valid.
	ClustermeshApiserverAdminCertValidityDuration = "clustermesh-apiserver-admin-cert-validity-duration"
	// ClustermeshApiserverAdminCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API admin certificate is written to.
	ClustermeshApiserverAdminCertSecretName = "clustermesh-apiserver-admin-cert-secret-name"

	// ClustermeshApiserverClientCertGenerate can be set to true to generate and
	// store a new Clustermesh API client TLS certificate.
	ClustermeshApiserverClientCertGenerate = "clustermesh-apiserver-client-cert-generate"
	// ClustermeshApiserverClientCertCommonName is the Clustermesh API client
	// x509 certificate CN value.
	ClustermeshApiserverClientCertCommonName = "clustermesh-apiserver-client-cert-common-name"
	// ClustermeshApiserverClientCertValidityDuration represent how much time
	// Clustermesh API client certificate generated by certgen is valid.
	ClustermeshApiserverClientCertValidityDuration = "clustermesh-apiserver-client-cert-validity-duration"
	// ClustermeshApiserverClientCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API client certificate is written to.
	ClustermeshApiserverClientCertSecretName = "clustermesh-apiserver-client-cert-secret-name"

	// ClustermeshApiserverRemoteCertGenerate can be set to true to generate
	// and store a new ClustermeshApiserver remote secret. If true then any
	// existing secret is overwritten with a new one.
	ClustermeshApiserverRemoteCertGenerate = "clustermesh-apiserver-remote-cert-generate"
	// ClustermeshApiserverRemoteCertCommonName is the Clustermesh API remote
	// x509 certificate CN value.
	ClustermeshApiserverRemoteCertCommonName = "clustermesh-apiserver-remote-cert-common-name"
	// ClustermeshApiserverRemoteCertValidityDuration represent how much time
	// Clustermesh API remote certificate generated by certgen is valid.
	ClustermeshApiserverRemoteCertValidityDuration = "clustermesh-apiserver-remote-cert-validity-duration"
	// ClustermeshApiserverRemoteCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API remote certificate is written to.
	ClustermeshApiserverRemoteCertSecretName = "clustermesh-apiserver-remote-cert-secret-name"
)

// CertGenConfig contains the main configuration options
type CertGenConfig struct {
	// Debug enables debug messages.
	Debug bool

	// CiliumNamespace is the Kubernetes namespace in which Cilium is
	// installed.
	CiliumNamespace string

	// K8sKubeConfigPath is the path to the kubeconfig If empty, the in-cluster
	// configuration is used.
	K8sKubeConfigPath string
	// K8sRequestTimeout specifies the timeout for K8s API requests
	K8sRequestTimeout time.Duration

	// CACertFile is the path to the Cilium CA cert PEM (if CAGenerate is
	// false).
	CACertFile string
	// CAKeyFile is the path to the Cilium CA key PEM (if CAGenerate is false).
	CAKeyFile string

	// CAGenerate can be set to true to generate a new Cilium CA secret.  If
	// CAReuseSecret is true, then a new CA secret only is created if existing
	// one is not found.
	CAGenerate bool
	// CAReuseSecret can be set to true to store and load the Cilium CA from
	// the secret if it exists. Setting to false will delete the old Secret and
	// force regeneration.
	CAReuseSecret bool
	// CACommonName is the Cilium CA x509 certificate CN value.
	CACommonName string
	// CAValidityDuration represent how much time the Cilium CA certificate
	// generated by certgen is valid.
	CAValidityDuration time.Duration
	// CASecretName is the Kubernetes Secret in which the Cilium CA certificate
	// is read from and/or written to.
	CASecretName string
	// CASecretNamespace is the Kubernetes Namespace in which the Cilium CA
	// Secret will be stored.
	CASecretNamespace string

	// HubbleRelayClientCertGenerate can be set to true to generate and store a
	// Hubble Relay client TLS certificate (used for the mTLS handshake with
	// the Hubble servers).
	HubbleRelayClientCertGenerate bool
	// HubbleRelayClientCertCommonName is the Hubble Relay client x509
	// certificate CN value.
	HubbleRelayClientCertCommonName string
	// HubbleRelayClientCertValidityDuration represent how much time the Hubble
	// Relay client certificate generated by certgen is valid.
	HubbleRelayClientCertValidityDuration time.Duration
	// HubbleRelayClientCertSecretName is the Kubernetes Secret in which the
	// Hubble Relay client certificate is written to.
	HubbleRelayClientCertSecretName string
	// HubbleRelayClientCertSecretNamespace is the Kubernetes Namespace in
	// which the Hubble Relay client certificate Secret will be stored.
	HubbleRelayClientCertSecretNamespace string

	// HubbleRelayServerCertGenerate can be set to true to generate and store a
	// Hubble Relay server TLS certificate.
	HubbleRelayServerCertGenerate bool
	// HubbleRelayServerCertCommonName is the Hubble Relay server x509
	// certificate CN value (also used as DNS SAN).
	HubbleRelayServerCertCommonName string
	// HubbleRelayServerCertValidityDuration represent how much time the Hubble
	// Relay server certificate generated by certgen is valid.
	HubbleRelayServerCertValidityDuration time.Duration
	// HubbleRelayServerCertSecretName is the Kubernetes Secret in which the
	// Hubble Relay server certificate is written to.
	HubbleRelayServerCertSecretName string
	// HubbleRelayServerCertSecretNamespace where the Hubble Relay server cert
	// and key will be stored.
	HubbleRelayServerCertSecretNamespace string

	// HubbleServerCertGenerate can be set to true to generate and store a
	// Hubble server TLS certificate.
	HubbleServerCertGenerate bool
	// HubbleServerCertCommonName is the Hubble server x509 certificate CN
	// value (also used as DNS SAN).
	HubbleServerCertCommonName string
	// HubbleServerCertValidityDuration represent how much time the Hubble
	// server certificate generated by certgen is valid.
	HubbleServerCertValidityDuration time.Duration
	// HubbleServerCertSecretName is the Kubernetes Secret in which the Hubble
	// server certificate is written to.
	HubbleServerCertSecretName string
	// HubbleServerCertSecretNamespace is the Kubernetes Namespace in which the
	// Hubble server certificate Secret will be stored.
	HubbleServerCertSecretNamespace string

	// HubbleMetricsServerCertGenerate can be set to true to generate and store a
	// Hubble metrics server TLS certificate.
	HubbleMetricsServerCertGenerate bool
	// HubbleMetricsServerCertCommonName is the Hubble metrics server x509 certificate CN
	// value (also used as DNS SAN).
	HubbleMetricsServerCertCommonName string
	// HubbleMetricsServerCertValidityDuration represent how much time the Hubble
	// server certificate generated by certgen is valid.
	HubbleMetricsServerCertValidityDuration time.Duration
	// HubbleMetricsServerCertSecretName is the Kubernetes Secret in which the Hubble
	// server certificate is written to.
	HubbleMetricsServerCertSecretName string
	// HubbleMetricsServerCertSecretNamespace is the Kubernetes Namespace in which the
	// Hubble metrics server certificate Secret will be stored.
	HubbleMetricsServerCertSecretNamespace string

	// ClustermeshApiserverServerCertGenerate can be set to true to generate
	// and store a new Clustermesh API server TLS certificate.
	ClustermeshApiserverServerCertGenerate bool
	// ClustermeshApiserverServerCertCommonName is the Clustermesh API server
	// x509 certificate CN value (also used as DNS SAN).
	ClustermeshApiserverServerCertCommonName string
	// ClustermeshApiserverServerCertValidityDuration represent how much time
	// Clustermesh API server certificate generated by certgen is valid.
	ClustermeshApiserverServerCertValidityDuration time.Duration
	// ClustermeshApiserverServerCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API server certificate is written to.
	ClustermeshApiserverServerCertSecretName string
	// ClustermeshApiserverServerCertSANs is the list of SANs to add to the
	// Clustermesh API server certificate.
	ClustermeshApiserverServerCertSANs []string

	// ClustermeshApiserverAdminCertGenerate can be set to true to generate and
	// store a new Clustermesh API admin TLS certificate.
	ClustermeshApiserverAdminCertGenerate bool
	// ClustermeshApiserverAdminCertCommonName is the Clustermesh API admin
	// x509 certificate CN value.
	ClustermeshApiserverAdminCertCommonName string
	// ClustermeshApiserverAdminCertValidityDuration represent how much time
	// Clustermesh API admin certificate generated by certgen is valid.
	ClustermeshApiserverAdminCertValidityDuration time.Duration
	// ClustermeshApiserverAdminCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API admin certificate is written to.
	ClustermeshApiserverAdminCertSecretName string

	// ClustermeshApiserverClientCertGenerate can be set to true to generate and
	// store a new Clustermesh API client TLS certificate.
	ClustermeshApiserverClientCertGenerate bool
	// ClustermeshApiserverClientCertCommonName is the Clustermesh API client
	// x509 certificate CN value.
	ClustermeshApiserverClientCertCommonName string
	// ClustermeshApiserverClientCertValidityDuration represent how much time
	// Clustermesh API client certificate generated by certgen is valid.
	ClustermeshApiserverClientCertValidityDuration time.Duration
	// ClustermeshApiserverClientCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API client certificate is written to.
	ClustermeshApiserverClientCertSecretName string

	// ClustermeshApiserverRemoteCertGenerate can be set to true to generate and
	// store a new Clustermesh API remote TLS certificate.
	ClustermeshApiserverRemoteCertGenerate bool
	// ClustermeshApiserverRemoteCertCommonName is the Clustermesh API remote
	// x509 certificate CN value.
	ClustermeshApiserverRemoteCertCommonName string
	// ClustermeshApiserverRemoteCertValidityDuration represent how much time
	// Clustermesh API remote certificate generated by certgen is valid.
	ClustermeshApiserverRemoteCertValidityDuration time.Duration
	// ClustermeshApiserverRemoteCertSecretName is the Kubernetes Secret in
	// which the Clustermesh API remote certificate is written to.
	ClustermeshApiserverRemoteCertSecretName string
}

// getStringWithFallback returns the value associated with the key as a string
// if it is non-empty. If the value is empty, this function returns the value
// associated with fallbackKey
func getStringWithFallback(vp *viper.Viper, key, fallbackKey string) string { //nolint:unparam
	if value := vp.GetString(key); value != "" {
		return value
	}
	return vp.GetString(fallbackKey)
}

// PopulateFrom populates the config struct with the values provided by vp
func (c *CertGenConfig) PopulateFrom(vp *viper.Viper) {
	c.Debug = vp.GetBool(Debug)
	c.K8sKubeConfigPath = vp.GetString(K8sKubeConfigPath)
	c.K8sRequestTimeout = vp.GetDuration(K8sRequestTimeout)

	c.CACertFile = vp.GetString(CACertFile)
	c.CAKeyFile = vp.GetString(CAKeyFile)

	c.CAGenerate = vp.GetBool(CAGenerate)
	c.CAReuseSecret = vp.GetBool(CAReuseSecret)
	c.CACommonName = vp.GetString(CACommonName)
	c.CAValidityDuration = vp.GetDuration(CAValidityDuration)
	c.CASecretName = vp.GetString(CASecretName)
	c.CASecretNamespace = getStringWithFallback(vp, CASecretNamespace, CiliumNamespace)

	c.HubbleRelayClientCertGenerate = vp.GetBool(HubbleRelayClientCertGenerate)
	c.HubbleRelayClientCertCommonName = vp.GetString(HubbleRelayClientCertCommonName)
	c.HubbleRelayClientCertValidityDuration = vp.GetDuration(HubbleRelayClientCertValidityDuration)
	c.HubbleRelayClientCertSecretName = vp.GetString(HubbleRelayClientCertSecretName)
	c.HubbleRelayClientCertSecretNamespace = getStringWithFallback(vp, HubbleRelayClientCertSecretNamespace, CiliumNamespace)

	c.HubbleRelayServerCertGenerate = vp.GetBool(HubbleRelayServerCertGenerate)
	c.HubbleRelayServerCertCommonName = vp.GetString(HubbleRelayServerCertCommonName)
	c.HubbleRelayServerCertValidityDuration = vp.GetDuration(HubbleRelayServerCertValidityDuration)
	c.HubbleRelayServerCertSecretName = vp.GetString(HubbleRelayServerCertSecretName)
	c.HubbleRelayServerCertSecretNamespace = getStringWithFallback(vp, HubbleRelayServerCertSecretNamespace, CiliumNamespace)

	c.HubbleServerCertGenerate = vp.GetBool(HubbleServerCertGenerate)
	c.HubbleServerCertCommonName = vp.GetString(HubbleServerCertCommonName)
	c.HubbleServerCertValidityDuration = vp.GetDuration(HubbleServerCertValidityDuration)
	c.HubbleServerCertSecretName = vp.GetString(HubbleServerCertSecretName)
	c.HubbleServerCertSecretNamespace = getStringWithFallback(vp, HubbleServerCertSecretNamespace, CiliumNamespace)

	c.HubbleMetricsServerCertGenerate = vp.GetBool(HubbleMetricsServerCertGenerate)
	c.HubbleMetricsServerCertCommonName = vp.GetString(HubbleMetricsServerCertCommonName)
	c.HubbleMetricsServerCertValidityDuration = vp.GetDuration(HubbleMetricsServerCertValidityDuration)
	c.HubbleMetricsServerCertSecretName = vp.GetString(HubbleMetricsServerCertSecretName)
	c.HubbleMetricsServerCertSecretNamespace = getStringWithFallback(vp, HubbleMetricsServerCertSecretNamespace, CiliumNamespace)

	c.CiliumNamespace = vp.GetString(CiliumNamespace)

	c.ClustermeshApiserverServerCertGenerate = vp.GetBool(ClustermeshApiserverServerCertGenerate)
	c.ClustermeshApiserverServerCertCommonName = vp.GetString(ClustermeshApiserverServerCertCommonName)
	c.ClustermeshApiserverServerCertValidityDuration = vp.GetDuration(ClustermeshApiserverServerCertValidityDuration)
	c.ClustermeshApiserverServerCertSecretName = vp.GetString(ClustermeshApiserverServerCertSecretName)
	c.ClustermeshApiserverServerCertSANs = vp.GetStringSlice(ClustermeshApiserverServerCertSANs)

	c.ClustermeshApiserverAdminCertGenerate = vp.GetBool(ClustermeshApiserverAdminCertGenerate)
	c.ClustermeshApiserverAdminCertCommonName = vp.GetString(ClustermeshApiserverAdminCertCommonName)
	c.ClustermeshApiserverAdminCertValidityDuration = vp.GetDuration(ClustermeshApiserverAdminCertValidityDuration)
	c.ClustermeshApiserverAdminCertSecretName = vp.GetString(ClustermeshApiserverAdminCertSecretName)

	c.ClustermeshApiserverClientCertGenerate = vp.GetBool(ClustermeshApiserverClientCertGenerate)
	c.ClustermeshApiserverClientCertCommonName = vp.GetString(ClustermeshApiserverClientCertCommonName)
	c.ClustermeshApiserverClientCertValidityDuration = vp.GetDuration(ClustermeshApiserverClientCertValidityDuration)
	c.ClustermeshApiserverClientCertSecretName = vp.GetString(ClustermeshApiserverClientCertSecretName)

	c.ClustermeshApiserverRemoteCertGenerate = vp.GetBool(ClustermeshApiserverRemoteCertGenerate)
	c.ClustermeshApiserverRemoteCertCommonName = vp.GetString(ClustermeshApiserverRemoteCertCommonName)
	c.ClustermeshApiserverRemoteCertValidityDuration = vp.GetDuration(ClustermeshApiserverRemoteCertValidityDuration)
	c.ClustermeshApiserverRemoteCertSecretName = vp.GetString(ClustermeshApiserverRemoteCertSecretName)
}
