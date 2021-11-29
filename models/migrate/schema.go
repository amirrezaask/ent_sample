// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "birth_year", Type: field.TypeInt},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
	}
	// StudentCoursesColumns holds the columns for the "student_courses" table.
	StudentCoursesColumns = []*schema.Column{
		{Name: "student_id", Type: field.TypeInt},
		{Name: "course_id", Type: field.TypeInt},
	}
	// StudentCoursesTable holds the schema information for the "student_courses" table.
	StudentCoursesTable = &schema.Table{
		Name:       "student_courses",
		Columns:    StudentCoursesColumns,
		PrimaryKey: []*schema.Column{StudentCoursesColumns[0], StudentCoursesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_courses_student_id",
				Columns:    []*schema.Column{StudentCoursesColumns[0]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "student_courses_course_id",
				Columns:    []*schema.Column{StudentCoursesColumns[1]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoursesTable,
		StudentsTable,
		StudentCoursesTable,
	}
)

func init() {
	StudentCoursesTable.ForeignKeys[0].RefTable = StudentsTable
	StudentCoursesTable.ForeignKeys[1].RefTable = CoursesTable
}
