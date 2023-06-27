import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { DecodedToken, LoginData } from './models';
import jwt_decode from 'jwt-decode';

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

	getToken(): string {
		return localStorage.getItem('access_token') || '';
	}

	getCurrentUser() {
		const decoded: DecodedToken = jwt_decode(
			localStorage.getItem('access_token')!
		);
		return decoded.sub;
	}

	getExpiration() {
		const decoded: DecodedToken = jwt_decode(
			localStorage.getItem('access_token')!
		);
		return decoded.exp;
	}

	getCurrentTimestamp() {
		return Math.floor(Date.now() / 1000);
	}

	public isLoggedIn() {
		if (!localStorage.getItem('access_token')) {
			return false;
		}
		if (this.getCurrentTimestamp() < this.getExpiration()) {
			return true;
		} else {
			return false;
		}
	}
}
