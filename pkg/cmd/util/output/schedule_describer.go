/*
Copyright 2017 the Heptio Ark contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package output

import (
	"fmt"

	"github.com/heptio/ark/pkg/apis/ark/v1"
)

func DescribeSchedule(schedule *v1.Schedule) string {
	return Describe(func(d *Describer) {
		d.DescribeMetadata(schedule.ObjectMeta)

		d.Println()
		DescribeScheduleSpec(d, schedule.Spec)

		d.Println()
		DescribeScheduleStatus(d, schedule.Status)
	})
}

func DescribeScheduleSpec(d *Describer, spec v1.ScheduleSpec) {
	d.Printf("Schedule:\t%s\n", spec.Schedule)

	d.Println()
	d.Println("Backup Template:")
	d.Prefix = "\t"
	DescribeBackupSpec(d, spec.Template)
	d.Prefix = ""
}

func DescribeScheduleStatus(d *Describer, status v1.ScheduleStatus) {
	phase := status.Phase
	if phase == "" {
		phase = v1.SchedulePhaseNew
	}

	d.Printf("Validation errors:")
	if len(status.ValidationErrors) == 0 {
		d.Printf("\t<none>\n")
	} else {
		for _, ve := range status.ValidationErrors {
			d.Printf("\t%s\n", ve)
		}
	}

	d.Println()
	lastBackup := "<never>"
	if !status.LastBackup.Time.IsZero() {
		lastBackup = fmt.Sprintf("%v", status.LastBackup.Time)
	}
	d.Printf("Last Backup:\t%s\n", lastBackup)
}
