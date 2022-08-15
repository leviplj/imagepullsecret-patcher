package podutils

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
)

func IncludeImagePullSecret(pod *corev1.Pod, secretName string) bool {
	for _, imagePullSecret := range pod.Spec.ImagePullSecrets {
		if imagePullSecret.Name == secretName {
			return true
		}
	}
	return false
}

type Spec struct {
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
}
type patch struct {
	Spec Spec `json:"spec,omitempty"`
}

func GetPatchString(pod *corev1.Pod, secretName string) ([]byte, error) {
	patch := patch{
		Spec{
			// copy the slice
			ImagePullSecrets: append([]corev1.LocalObjectReference(nil), pod.Spec.ImagePullSecrets...),
		},
	}
	if !IncludeImagePullSecret(pod, secretName) {
		patch.Spec.ImagePullSecrets = append(patch.Spec.ImagePullSecrets, corev1.LocalObjectReference{Name: secretName})
	}
	return json.Marshal(patch)
}
