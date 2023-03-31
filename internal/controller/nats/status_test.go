package nats

import (
	"errors"
	"testing"

	natsv1alpha1 "github.com/kyma-project/nats-manager/api/v1alpha1"
	"github.com/kyma-project/nats-manager/testutils"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_syncNATSStatus(t *testing.T) {
	t.Parallel()

	// define test cases
	testCases := []struct {
		name           string
		givenNATS      *natsv1alpha1.Nats
		wantNATSStatus natsv1alpha1.NatsStatus
		wantResult     bool
	}{
		{
			name: "should update the status",
			givenNATS: testutils.NewSampleNATSCR(
				testutils.WithNATSCRStatusInitialized(),
				testutils.WithNATSStateProcessing(),
			),
			wantNATSStatus: natsv1alpha1.NatsStatus{
				State: natsv1alpha1.StateReady,
				Conditions: []metav1.Condition{
					{
						Type:               string(natsv1alpha1.ConditionStatefulSet),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonProcessing),
						Message:            "",
					},
					{
						Type:               string(natsv1alpha1.ConditionAvailable),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonProcessing),
						Message:            "",
					},
				},
			},
			wantResult: false,
		},
	}

	// run test cases
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// given
			testEnv := NewMockedUnitTestEnvironment(t, tc.givenNATS)
			reconciler := testEnv.Reconciler

			newNATS := tc.givenNATS.DeepCopy()
			newNATS.Status = tc.wantNATSStatus

			// when
			err := reconciler.syncNATSStatus(testEnv.Context, newNATS, testEnv.Logger)

			// then
			require.NoError(t, err)
			gotNats, err := testEnv.GetNATS(tc.givenNATS.GetName(), tc.givenNATS.GetNamespace())
			require.NoError(t, err)
			require.True(t, gotNats.Status.IsEqual(tc.wantNATSStatus))
		})
	}
}

func Test_syncNATSStatusWithErr(t *testing.T) {
	t.Parallel()

	// define test cases
	testCases := []struct {
		name           string
		givenNATS      *natsv1alpha1.Nats
		givenError     error
		wantNATSStatus natsv1alpha1.NatsStatus
		wantResult     bool
	}{
		{
			name: "should update the status with error message",
			givenNATS: testutils.NewSampleNATSCR(
				testutils.WithNATSCRStatusInitialized(),
				testutils.WithNATSStateProcessing(),
			),
			givenError: errors.New("test error"),
			wantNATSStatus: natsv1alpha1.NatsStatus{
				State: natsv1alpha1.StateError,
				Conditions: []metav1.Condition{
					{
						Type:               string(natsv1alpha1.ConditionStatefulSet),
						Status:             metav1.ConditionFalse,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonSyncFailError),
						Message:            "",
					},
					{
						Type:               string(natsv1alpha1.ConditionAvailable),
						Status:             metav1.ConditionFalse,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonProcessingError),
						Message:            "test error",
					},
				},
			},
			wantResult: false,
		},
	}

	// run test cases
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// given
			testEnv := NewMockedUnitTestEnvironment(t, tc.givenNATS)
			reconciler := testEnv.Reconciler

			newNATS := tc.givenNATS.DeepCopy()
			newNATS.Status = tc.wantNATSStatus

			// when
			err := reconciler.syncNATSStatusWithErr(testEnv.Context, newNATS, tc.givenError, testEnv.Logger)

			// then
			require.NoError(t, err)
			gotNats, err := testEnv.GetNATS(tc.givenNATS.GetName(), tc.givenNATS.GetNamespace())
			require.NoError(t, err)
			require.True(t, gotNats.Status.IsEqual(tc.wantNATSStatus))
		})
	}
}

func Test_updateStatus(t *testing.T) {
	t.Parallel()

	// define test cases
	testCases := []struct {
		name           string
		givenNATS      *natsv1alpha1.Nats
		wantNATSStatus natsv1alpha1.NatsStatus
		wantResult     bool
	}{
		{
			name: "should update the status",
			givenNATS: testutils.NewSampleNATSCR(
				testutils.WithNATSCRStatusInitialized(),
				testutils.WithNATSStateProcessing(),
			),
			wantNATSStatus: natsv1alpha1.NatsStatus{
				State: natsv1alpha1.StateReady,
				Conditions: []metav1.Condition{
					{
						Type:               string(natsv1alpha1.ConditionStatefulSet),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonProcessing),
						Message:            "",
					},
					{
						Type:               string(natsv1alpha1.ConditionAvailable),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(natsv1alpha1.ConditionReasonProcessing),
						Message:            "",
					},
				},
			},
			wantResult: false,
		},
	}

	// run test cases
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// given
			testEnv := NewMockedUnitTestEnvironment(t, tc.givenNATS)
			reconciler := testEnv.Reconciler

			oldNATS, err := testEnv.GetNATS(tc.givenNATS.GetName(), tc.givenNATS.GetNamespace())
			require.NoError(t, err)
			newNATS := oldNATS.DeepCopy()
			newNATS.Status = tc.wantNATSStatus

			// when
			err = reconciler.updateStatus(testEnv.Context, &oldNATS, newNATS, testEnv.Logger)

			// then
			require.NoError(t, err)
			gotNats, err := testEnv.GetNATS(tc.givenNATS.GetName(), tc.givenNATS.GetNamespace())
			require.NoError(t, err)
			require.True(t, gotNats.Status.IsEqual(tc.wantNATSStatus))
		})
	}
}
