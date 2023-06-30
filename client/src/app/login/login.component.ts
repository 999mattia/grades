import { Component } from '@angular/core';
import { AuthService } from '../../shared/auth.service';
import { Router } from '@angular/router';
import { LoginData } from '../../shared/models';

@Component({
	selector: 'app-login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.css'],
})
export class LoginComponent {
	user: LoginData = { name: '', password: '' };
	constructor(private authService: AuthService, private router: Router) {}

	onSubmit() {
		this.authService.logIn(this.user);
		this.router.navigate(['/']);
	}
}
