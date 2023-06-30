import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserData } from './models';
import { AuthService } from './auth.service';
import { HttpHeaders } from '@angular/common/http';

@Injectable({
	providedIn: 'root',
})
export class UserService {
	constructor(private http: HttpClient, private authService: AuthService) {}

	getUserData(): Observable<UserData> {
		const headers = new HttpHeaders({
			Authorization: this.authService.getToken(),
		});

		const options = { headers: headers };

		return this.http.get<UserData>(
			`https://api.grades.mattiag.ch/user/${this.authService.getCurrentUser()}`,
			options
		);
	}
}
