import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { CreateGradeComponent } from 'src/app/create-grade/create-grade.component';
import { CreateGradeModel } from './models';

@Injectable({
	providedIn: 'root',
})
export class GradeService {
	constructor(private http: HttpClient, private authService: AuthService) {}

	createGrade(grade: CreateGradeModel) {
		const headers = new HttpHeaders({
			Authorization: this.authService.getToken(),
		});

		const options = { headers: headers };

		return this.http.post<CreateGradeComponent>(
			`https://api.grades.mattiag.ch/module/${grade.moduleId}/grade`,
			grade,
			options
		);
	}

	deleteGrade(id: number) {
		const headers = new HttpHeaders({
			Authorization: this.authService.getToken(),
		});

		const options = { headers: headers };

		return this.http.delete(
			`https://api.grades.mattiag.ch/grade/${id}`,
			options
		);
	}
}
