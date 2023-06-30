import { Component, Input } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Grade } from 'src/app/shared/models';

@Component({
	selector: 'app-grade-table',
	templateUrl: './grade-table.component.html',
	styleUrls: ['./grade-table.component.css'],
})
export class GradeTableComponent {
	@Input() grades: Grade[] = [];

	displayedColumns: string[] = ['name', 'grade', 'actions'];
	dataSource: MatTableDataSource<Grade>;

	constructor() {
		this.dataSource = new MatTableDataSource();
		this.dataSource.data = this.grades;
	}
}
