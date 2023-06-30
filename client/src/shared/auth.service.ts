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
		this.http
			.post<any>('https://api.grades.mattiag.ch/login', user)
			.subscribe((res: any) => {
				localStorage.setItem('access_token', res.token);
				this.router.navigate(['']);
			});
	}

	logOut() {
		localStorage.removeItem('access_token');
		this.router.navigate(['login']);
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

	getCurrentUsername() {
		if (!this.isLoggedIn()) {
			return;
		}
		return this.http.get<any>(
			`https://api.grades.mattiag.ch/user/${this.getCurrentUser()}/username`
		);
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

	isLoggedIn() {
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
