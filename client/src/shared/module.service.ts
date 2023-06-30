import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CreateModuleModel } from './models';
import { AuthService } from './auth.service';

@Injectable({
	providedIn: 'root',
})
export class ModuleService {
	constructor(private http: HttpClient, private authService: AuthService) {}

	createModule(module: CreateModuleModel) {
		const headers = new HttpHeaders({
			Authorization: this.authService.getToken(),
		});

		const options = { headers: headers };

		return this.http.post<CreateModuleModel>(
			'https://api.grades.mattiag.ch/module',
			module,
			options
		);
	}

	deleteModule(id: number) {
		const headers = new HttpHeaders({
			Authorization: this.authService.getToken(),
		});

		const options = { headers: headers };

		return this.http.delete(
			`https://api.grades.mattiag.ch/module/${id}`,
			options
		);
	}
}
