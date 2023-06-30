import { Component, Input } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Grade } from 'src/shared/models';
import { GradeService } from 'src/shared/grade.service';

@Component({
	selector: 'app-grade-table',
	templateUrl: './grade-table.component.html',
})
export class GradeTableComponent {
	@Input() grades: Grade[] = [];

	displayedColumns: string[] = ['name', 'grade', 'actions'];
	dataSource: MatTableDataSource<Grade>;

	constructor(private gradeService: GradeService) {
		this.dataSource = new MatTableDataSource();
		this.dataSource.data = this.grades;
	}

	handleDeleteGrade(id: number) {
		this.gradeService.deleteGrade(id).subscribe(() => {
			this.grades = this.grades.filter((grade) => grade.id !== id);
			this.dataSource.data = this.grades;
		});
	}
}
