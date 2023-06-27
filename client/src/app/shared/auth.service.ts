import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { LoginData } from './models';

@Injectable({
	providedIn: 'root',
})
export class AuthService {
	constructor(private http: HttpClient, public router: Router) {}

	logIn(user: LoginData) {
		return this.http
			.post<any>('https://api.grades.mattiag.ch/login', user)
			.subscribe((res: any) => {
				localStorage.setItem('access_token', res.token);
			});
	}
}
