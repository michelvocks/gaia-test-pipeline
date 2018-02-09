package main

import (
	"github.com/gaia-pipeline/gaia/proto"
	sdk "github.com/gaia-pipeline/gosdk"
)

var jobs = []sdk.JobsWrapper{
	sdk.JobsWrapper{
		FuncPointer: CreateUser,
		Job: proto.Job{
			Title:       "Create DB User",
			Description: "Creates a database user with least privileged permissions.",
			Priority:    0,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: MigrateDB,
		Job: proto.Job{
			Title:       "DB Migration",
			Description: "Imports newest test data dump and migrates to newest version.",
			Priority:    10,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: CreateNamespace,
		Job: proto.Job{
			Title:       "Create K8S Namespace",
			Description: "Creates a new Kubernetes namespace for the new test environment.",
			Priority:    20,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: CreateDeployment,
		Job: proto.Job{
			Title:       "Create K8S Deployment",
			Description: "Creates a new Kubernetes deployment for the new test environment.",
			Priority:    30,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: CreateService,
		Job: proto.Job{
			Title:       "Create K8S Service",
			Description: "Creates a new Kubernetes service for the new test environment.",
			Priority:    30,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: CreateIngress,
		Job: proto.Job{
			Title:       "Create K8S Ingress",
			Description: "Creates a new Kubernetes ingress for the new test environment.",
			Priority:    30,
		},
	},
	sdk.JobsWrapper{
		FuncPointer: Cleanup,
		Job: proto.Job{
			Title:       "Clean up",
			Description: "Removes all temporary files.",
			Priority:    40,
		},
	},
}
