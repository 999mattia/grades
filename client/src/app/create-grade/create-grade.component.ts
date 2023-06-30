import { Component } from '@angular/core';
import { ActivatedRoute, Route, Router } from '@angular/router';
import { GradeService } from 'src/shared/grade.service';
import { CreateGradeModel } from 'src/shared/models';

@Component({
	selector: 'app-create-grade',
	templateUrl: './create-grade.component.html',
})
export class CreateGradeComponent {
	grade: CreateGradeModel = {
		name: '',
		grade: null,
		moduleId: 0,
	};

	constructor(
		private route: ActivatedRoute,
		private router: Router,
		private gradeService: GradeService
	) {}

	ngOnInit() {
		this.grade.moduleId = +this.route.snapshot.params['id'];
	}

	onSubmit() {
		this.gradeService.createGrade(this.grade).subscribe(() => {
			this.router.navigate(['']);
		});
	}
}
